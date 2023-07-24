package model

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"vue3-bashItem/pkg/aesEncryption"
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/settings"
	"vue3-bashItem/pkg/utils"
)

type BaseModel struct {
	//gorm.Model  // 包含了Id CreatedAt UpdatedAt DeletedAt
	CreatedAt time.Time      `gorm:"comment:创建时间"` // 在创建时，如果该字段值为零值，则使用当前时间填充
	UpdatedAt time.Time      `gorm:"comment:更新时间"` // 在创建时该字段值为零值或者在更新时，使用当前时间填充
	DeletedAt gorm.DeletedAt `gorm:"comment:软删除"`
}
type GormModel struct {
	gorm.Model
}

var Db *gorm.DB

func Setup() {
	mysqlPass, deErr := aesEncryption.DePwdCode(settings.DatabaseSetting.Password)
	if deErr != nil {
		errData := fmt.Sprintf("mysqlPass解密报错: %v", zap.Error(deErr))
		logger.FileLogger.Error(errData)
		panic(errData)
		//os.Exit(1)
	}

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.DatabaseSetting.User,
		mysqlPass,
		//settings.DatabaseSetting.Password,
		settings.DatabaseSetting.Host,
		settings.DatabaseSetting.Port,
		settings.DatabaseSetting.Name)

	// db log
	loggerLevel := gormLogger.Error
	if settings.ServerSetting.RunMode == "debug" {
		loggerLevel = gormLogger.Info
	}
	dbLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  loggerLevel,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   settings.DatabaseSetting.TablePrefix, // 设置默认表头
			SingularTable: true,
		},
		Logger: dbLogger,
	})

	if err != nil {
		//logger.Logger.Error("DB Connection Error: ", zap.Error(err))
		//os.Exit(1)
		errData := fmt.Sprintf("DB Connection Error: %v", zap.Error(deErr))
		logger.FileLogger.Error(errData)
		panic(errData)
	}

	sqlDB, err := Db.DB()
	if err != nil {
		//logger.Logger.Error("DB Connection Error: ", zap.Error(err))
		//os.Exit(1)
		errData := fmt.Sprintf("DB Connection Error: %v", zap.Error(deErr))
		logger.FileLogger.Error(errData)
		panic(errData)
	}
	sqlDB.SetMaxIdleConns(10)           // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 设置最大开放连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接池里连接最大存活时长
}

func CheckExisted(instance interface{}, query interface{}, args ...interface{}) (bool, error) {
	count, err := Count(instance, query, args...)
	// 查询报错
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func CheckExistedById(instance interface{}, id uint) (bool, error) {
	return CheckExisted(instance, "`id` = ?", id)
}

func DeleteOne(instance interface{}, query interface{}, args ...interface{}) error {
	stmt := &gorm.Statement{DB: Db}
	if err := stmt.Parse(instance); err != nil {
		return err
	}
	// get instance
	err := GetOne(instance, query, args...)
	if err != nil {
		return err
	}

	valueRef := reflect.ValueOf(instance).Elem()
	updates := make(map[string]interface{}, 0)
	for _, field := range stmt.Schema.Fields {
		// 通过tag检查所有唯一的可能性
		fieldIsUniq := field.Unique
		if _, ok := field.TagSettings["UNIQUEINDEX"]; ok {
			fieldIsUniq = true
		}
		if val, ok := field.TagSettings["INDEX"]; ok {
			indexConfigs := strings.Split(val, ",")
			for _, item := range indexConfigs {
				if item == "unique" {
					fieldIsUniq = true
				}
			}
		}
		// 除了主键，其他是唯一约束的删除时作数据填充操作
		if fieldIsUniq && !field.PrimaryKey {
			switch field.GORMDataType {
			case schema.String:
				newVal := strings.Join([]string{valueRef.FieldByName(field.Name).String(), strconv.Itoa(int(valueRef.FieldByName("ID").Uint()))}, "|||")
				updates[field.DBName] = newVal
			case schema.Uint:
				newVal := uint(valueRef.FieldByName(field.Name).Uint()+valueRef.FieldByName("ID").Uint()) + 50000
				updates[field.DBName] = newVal
			case schema.Int:
				newVal := int(valueRef.FieldByName(field.Name).Uint()+valueRef.FieldByName("ID").Uint()) + 50000
				updates[field.DBName] = newVal
			}
		}
	}
	updates["deleted_at"] = time.Now()
	return Db.Model(instance).Where(query, args...).Updates(updates).Error
}

func baseGetList(results interface{}, fields string, offset, limit int, order string, withCount bool, query interface{}, args ...interface{}) (int64, error) {
	// init count
	var count int64
	// get model schema
	stmt := &gorm.Statement{DB: Db}
	if err := stmt.Parse(results); err != nil {
		return count, err
	}
	// build sql query
	q := Db.Model(results)
	// set query filed if filed is not null
	if len(fields) != 0 {
		fieldList := make([]string, 0)
		for _, item := range strings.Split(fields, ",") {
			trimmedField := strings.TrimSpace(item)
			if len(trimmedField) != 0 {
				if utils.StringInSlice(fieldList, trimmedField) == false && utils.StringInSlice(stmt.Schema.DBNames, trimmedField) == true {
					fieldList = append(fieldList, trimmedField)
				}
			}
		}
		q = q.Select(fieldList)
	}
	// set query null if query is nil
	if query == nil {
		query = struct{}{}
	}

	// ========添加where查询============
	q = q.Where(query, args...)
	// ========添加order排序 条件不存在不添加order========
	if len(order) != 0 {
		orderSlice := strings.Split(order, " ")
		if ix := utils.StringInSlice(stmt.Schema.DBNames, orderSlice[0]); ix == false {
			return count, errors.New("order 参数非法! ")
		}
		orderSlice[0] = fmt.Sprintf("`%s`", orderSlice[0])
		q = q.Order(strings.Join([]string{strings.Join(orderSlice, " "), "id"}, ", "))
	}
	// ========添加offer 条件成立添加limit============
	if offset > 0 {
		q = q.Offset(offset)
	}
	//  ========添加limit 条件成立添加limit============
	if limit > 0 {
		q = q.Limit(limit)
	}
	// 查询值的数量  if need count, query count
	if withCount {
		countErrChan := make(chan error, 1)
		go func() {
			err := Db.Model(results).Where(query, args...).Count(&count).Error
			countErrChan <- err
		}()
		err := q.Find(results).Error
		cErr := <-countErrChan
		if cErr != nil {
			return count, cErr
		}
		return count, err
	}
	return count, q.Find(results).Error
}

func GetListWithFields(results interface{}, fields string, query interface{}, args ...interface{}) error {
	_, err := baseGetList(results, fields, 0, 0, "", false, query, args...)
	return err
}

func GetOrderedList(results interface{}, order string, query interface{}, args ...interface{}) error {
	_, err := baseGetList(results, "", 0, 0, order, false, query, args...)
	return err
}

func GetOrderedListWithFields(results interface{}, fields string, order string, query interface{}, args ...interface{}) error {
	_, err := baseGetList(results, fields, 0, 0, order, false, query, args...)
	return err
}

func GetPageList(results interface{}, offset, limit int, order string, query interface{}, args ...interface{}) (int64, error) {
	return baseGetList(results, "", offset, limit, order, true, query, args...)
}

func GetPageListWithFields(results interface{}, fields string, offset, limit int, order string, query interface{}, args ...interface{}) (int64, error) {
	return baseGetList(results, fields, offset, limit, order, true, query, args...)
}

func GetAllWithFields(results interface{}, fields string) error {
	_, err := baseGetList(results, fields, 0, 0, "", false, nil)
	return err
}

//=============================查询=======================================

// Count 取值的个数
//menuConut,menuConutErr := model.Count(&accountModel.Menu{}, "menu_id = ? AND menu_id IN ?", menuMdl.MenuId, roleInt)
func Count(instance interface{}, query interface{}, args ...interface{}) (int64, error) {
	var count int64
	err := Db.Model(instance).Where(query, args...).Count(&count).Error
	return count, err
}

// model.GetOne(serviceConfObj, "env_id = ? and service_id = ?", req.EnvId, req.ServiceId)
func GetOne(result interface{}, query interface{}, args ...interface{}) error {
	return Db.Model(result).Where(query, args...).First(result).Error
}

func GetOneById(result interface{}, id uint) error {
	return GetOne(result, "`id` = ?", id)
}

// GetOneFirst 获取第一个值 单值返回 {"Id":"aa","Age":18}
//roleData := accountModel.Role{}
//roleErr := model.GetOneFirst(&roleData, "role_id = ?", relationI.RoleId) //可多条件
func GetOneFirst(result interface{}, query interface{}, args ...interface{}) error {
	return Db.Where(query, args...).First(result).Error
}

// GetOneLast 获取最后1个值 单值返回 {"Id":"aa","Age":18}  查1个 不是list  没查询到会报错
//deptSelect := deptModel.SysDept{}
//if deptErr := model.GetOneLast(&deptSelect, "id = ? and type = ?",userInfo.aa,userInfo.bb); deptErr != nil {
//	logger.FileLogger.Error(fmt.Sprintf("查询失败 %v ", deptErr.Error()))
//}
func GetOneLast(result interface{}, query interface{}, args ...interface{}) error {
	return Db.Where(query, args...).Last(result).Error
	//return Db.Model(result).Where(query, args...).Last(result).Error
}

// GetAll 获取全部值 查N个 是list 没查询到 不会报错
//allUser := make([]*accountModel.User, 0)
//userErr := model.GetAll(&allUser)
//if userErr != nil {
//	fmt.Println("获取数据错误", userErr)
//}
func GetAll(results interface{}) error {
	return Db.Find(results).Error
}

// GetMany 获取多个值 可多条件
//relationUserRole :=make([]*accountModel.RelationUserRole,0)
//relationErr:=model.GetMany(&relationUserRole,"user_id = ? or sms_type = ?",userInfo.UserId,userInfo.smsType)
//if relationErr != nil {
//	fmt.Println("获取数据错误", relationErr)
//}
func GetMany(result interface{}, query interface{}, args ...interface{}) error {
	//err :=Db.Model(result).Where(query,args).Find(result).Error
	err := Db.Where(query, args).Find(result).Error
	return err
}

//=============================接口增删改查=======================================

// GetList 多条件查询 用上边的方法
// model.GetList(&deployRecords, "service_conf_id = ?", serviceConfObj.ID)
// model.GetList(&results, "sms_type = ? or sms_type = ? or sms_type = ?","typeCount", "typeMoney", "typeTime")
func GetList(results interface{}, query interface{}, args ...interface{}) error {
	_, err := baseGetList(results, "", 0, 0, "", false, query, args...)
	return err
}

//func GetPageList(results interface{}, offset, limit int, order string, query interface{}, args ...interface{}) (int64, error) {
//	return baseGetList(results, "", offset, limit, order, true, query, args...)
//}

// GetListPage 获取列表
func GetListPage(results interface{}, offset, limit int, order string, query interface{}, args ...interface{}) (int64, error) {
	// ========添加where查询============
	q := Db.Where(query, args...)
	// ========添加order排序 条件不存在不添加order========
	if len(order) != 0 {
		q = q.Order(order)
	}
	// ========添加offer 条件成立添加limit============
	if offset > 0 {
		q = q.Offset(offset)
	}
	//  ========添加limit 条件成立添加limit============
	if limit > 0 {
		q = q.Limit(limit)
	}
	// ========查询数据值=========
	if findErr := q.Find(results).Error; findErr != nil {
		return 0, findErr
	}
	// ========查询数据条数=========
	var countData int64
	if countErr := Db.Model(results).Where(query, args...).Count(&countData).Error; countErr != nil {
		return 0, countErr
	}
	return countData, nil
}

// Create 新增数据
//item := &domainConf.AccountUsage{
//	AliasBigClass:   requestBody.AliasBigClass,
//	AliasName:       requestBody.AliasName,
//}
//if err := model.Create(item); err != nil {
//	response.BaseError(c, fmt.Sprintf("insert 创建AccountUsage失败，DB创建数据错误：%v", err.Error()))
//	return
//}
func Create(instance interface{}) error {
	return Db.Create(instance).Error
}

// UpdateOneById 更新 根据id更新
// &hospitalModel.YchItem{}是model表 updateColumns是自定义interface   ych_id是id uriParams.YchId是值
//if err := model.UpdateOneById(&hospitalModel.YchItem{}, updateColumns, "ych_id", uriParams.YchId); err != nil {
//	response.BaseError(c, fmt.Sprintf("update 更新权限失败，hospitalModel表更新错误：%v", err.Error()))
//	return
//}
func UpdateOneById(model interface{}, updates interface{}, idKey string, idValue uint) error {
	return Update(model, updates, fmt.Sprintf("`%v` = ?", idKey), idValue)
}

// Update 更新 多条件更新
//updateData := map[string]interface{}{"Type": updateMenu.Type, "Icon": updateMenu.Icon,}
//if err := model.Update(&accountModel.Menu{}, updateData,"`menu_id` = ?  AND age =?", u.Id,u.Age); err != nil {
//	response.BaseError(c, fmt.Sprintf("update更新权限失败，DB数据更新错误：%v", err.Error()))
//return
//}
func Update(model interface{}, updates interface{}, query interface{}, args ...interface{}) error {
	return Db.Model(model).Where(query, args...).Updates(updates).Error
}

// DeleteMany 软删除 多条件删除
//model.DeleteMany(&accountModel.RelationUserRole{},"user_id => ? AND email = ?",2,"123@qq.com")
func DeleteMany(result interface{}, query interface{}, args ...interface{}) error {
	return Db.Where(query, args...).Delete(result).Error
}

// SoftGet 软删除 查已经软删除的数据
//getOldRole := make([]accountModel.RelationUserRole,0)  //Kwok是model表
//getErr := model.GetSoftDelete(&getOldRole,"id = ?",updateData.UserId)
func SoftGet(result interface{}, query interface{}, args ...interface{}) error {
	//err :=Db.Unscoped().Where(query,args...).Find(result).Error
	//return err
	return Db.Unscoped().Where(query, args...).Find(result).Error
}

// SoftRecover 软删除 回滚数据
//model.SoftRecover(&accountModel.RelationUserRole{},"id IN (?)",softUserAdd)
func SoftRecover(result interface{}, query interface{}, args interface{}) error {
	var nullDeleteAt gorm.DeletedAt
	err := Db.Unscoped().Model(result).Where(query, args).Updates(
		map[string]interface{}{"deleted_at": nullDeleteAt}).Error
	return err
}

// DeleteByQuery 硬删除 多条件删除 不写key默认key是id  user_id是数据库里的字段 不是model里定义的UserId
//model.DeleteByQuery(&accountModel.RelationUserRole{}, map[string]interface{}{"age": 3})         // 单条件删除
//model.DeleteByQuery(&accountModel.RelationUserRole{}, map[string]interface{}{"age": 3,"id":5})  // 多条件删除
func DeleteByQuery(instance interface{}, query interface{}) error {
	return Db.Unscoped().Delete(instance, query).Error
}

// GetTime 时间查询 查询N天内的数据
//	SELECT * FROM users WHERE created_at BETWEEN '2020-12-11 11:15' AND '2020-12-11 11:16:06.51';
//
//	intNum, _ := strconv.Atoi(dayNum) // string转int
//	today := time.Now()											    // 当前时间 东八区cst
//	lastTime := today.Add(-time.Hour * 24 * time.Duration(intNum))  // 最后时间 例前天
//	ossItemList := make([]*ossModel.OssItem, 0)
//	getErr := model.GetTime(&ossItemList, lastTime, today, "oss_status !='删除'")
//	if getErr != nil {
//		response.BaseError(c, fmt.Sprintf("查询db获取数据失败 %v", getErr.Error()))
//		return
//	}
// GetTime 时间查询 范围数据查询  // SELECT * FROM users WHERE created_at BETWEEN '2020-12-11 11:15' AND '2020-12-11 11:16:06.51';
func GetTime(result interface{}, lastTime, today interface{}, other string) error {
	err := Db.Where("created_at BETWEEN ? AND ?", lastTime, today,
	).Where(other, //).Where("oss_status !='删除'",
	).Find(result).Error
	return err
}

// 时间查询 时间小于等于2006-01-02 15:04:05的一个值 没有就报错 //SELECT * FROM users WHERE updated_at <= '2000-01-01 00:00:00';
//ossItemList := make([]*ossModel.OssItem, 0)
//getErr := model.GetTime(&ossItemList, lastTime, today, "oss_status !='删除'")
//if getErr != nil {
//	response.BaseError(c, fmt.Sprintf("查询db获取数据失败 %v", getErr.Error()))
//return
//}
// 时间查询 查询一条 不存在就异常
//  SELECT * FROM `ych_upgrade` WHERE end_time <= '2023-05-25 09:09:35' AND (hospital_id = 23 and app_server_id = 30 ) ORDER BY `ych_upgrade`.`ych_id` DESC LIMIT 1

func TimeLastLTE(result interface{}, lastTime interface{}, other string) error {
	err := Db.Where("end_time <= ?", lastTime,
	).Where(other, //).Where("oss_status !='删除'",
	).Last(result).Error
	return err
}
