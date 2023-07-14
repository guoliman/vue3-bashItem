package userApi

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
	"vue3-bashItem/model"
	"vue3-bashItem/model/accountAdminModel/deptModel"
	"vue3-bashItem/model/accountAdminModel/menuModel"
	"vue3-bashItem/model/accountAdminModel/roleModel"
	"vue3-bashItem/model/accountAdminModel/userModel"
	"vue3-bashItem/pkg/aesEncryption"
	"vue3-bashItem/pkg/jwt"
	"vue3-bashItem/pkg/logger"
	redisGo "vue3-bashItem/pkg/redis"
	"vue3-bashItem/pkg/request"
	"vue3-bashItem/pkg/response"
	"vue3-bashItem/pkg/settings"
	"vue3-bashItem/pkg/utils"
	"vue3-bashItem/services/accountAdminMethod/deptMethod"
	"vue3-bashItem/services/accountAdminMethod/roleMethod"
	"vue3-bashItem/services/accountAdminMethod/userMethod"
)

// Me 获取当前登录用户信息
func Me(c *gin.Context) {
	// 角色获取菜单列表
	userResult, roleCodeList, menuIdList, menuIdListErr := roleMethod.RoleGetMenu(c)
	if menuIdListErr != nil {
		response.BaseError(c, fmt.Sprintf("角色获取菜单异常: %v ", menuIdListErr.Error()))
		return
	}

	// 根据菜单id查询按钮
	buttonInfo := make([]menuModel.SysMenu, 0)
	roleIdErr := model.Db.Where("type = ? AND id In ?", 4, menuIdList,
	).Where("perm NOT IN ?", []string{"NULL", ""},
	).Find(&buttonInfo).Error
	if roleIdErr != nil {
		response.BaseError(c, fmt.Sprintf("按钮查询异常 %v ", roleIdErr.Error()))
		return
	}
	var buttonList []string
	for _, buttonV := range buttonInfo {
		buttonList = append(buttonList, buttonV.Perm)
	}
	buttonList = utils.RemoveListStr(buttonList) // 去重
	//logger.FileLogger.Debug("去重后的按钮==", buttonList)

	userResult["roles"] = roleCodeList // []string
	userResult["perms"] = buttonList   // []string

	response.Success(c, userResult)

	//	var userResult userMethod.MeData
	//	jsonData := `
	//{"userId":2,"username":"admin","nickname":"系统管理员","avatar":"https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif","roles":["ADMIN"],"perms":["sys:menu:delete","sys:dict_type:add","sys:dept:edit","sys:dict:edit","sys:dict:delete","sys:dict_type:edit","sys:menu:add","sys:user:add","sys:dept:delete","sys:role:edit","sys:user:edit","sys:user:reset_pwd","sys:user:delete","sys:dept:add","sys:dict_type:delete","sys:role:delete","sys:menu:edit","sys:dict:add","sys:role:add"]}
	//`
	//
	//	err := json.Unmarshal([]byte(jsonData), &userResult) // 序列化
	//	if err != nil {
	//		response.BaseError(c, fmt.Sprintln("JSON parsing error: ", err))
	//		return
	//	}
	//response.Vue3Response(c, userResult)
}

// Captcha 图片验证
func Captcha(c *gin.Context) {
	var captchaData userMethod.CaptchaResponse
	jsonData := `{"verifyCodeKey":"77f25f7b80de47b4a9b74f34ece72023","verifyCodeBase64":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAHgAAAAkCAIAAADNSmkJAAAEiUlEQVR4Xu2Yy09UVxzH2bX/QxdNmjQ2abti5cK4MCYa42tD2Bh1UaP4Ij7ic4gvEl3AQmp9BG1aGFDQgvEd0HEg0cgrqTNTIIUCDcM4+IAZGOyA+Ov3cHRy5nfv3GGKnjH0fHIWc3733LnnfOZ3f/fcySGDFnJ4wPBxMKI1YURrwojWhBGtCSNaE0a0JoxoTRjRmjCiNWFEayKN6PB4uPRxaX5t/uKfFy8sX7i6enVxU/FgZJCP08amTTYtSwwMkNtNR4/Sjh20fTuVllJ3Nx+TwEm0L+xbdGlR7vlc1hD0h/18tB6slrMn2jqRzZspEODDJE6iN9RtgNb1devbh9pjk7Hx+Dg+rPttHYIb6zfy0XrIqlnG8ePk8dCLFzQ1RcGgyGhMrbiYD5M4iUatgNNgNKgG0UUQh9RgpuScb0Pj0dnwKYlmjIyIqRUU8LjESfSq6lWpRK+pXqMGM2VeikZSY2oHDvC4xEn0tT+uydLREepA6UBLlI66zrp3g3p76cIF2rtX1KcjR+jsWWppoenppC+yMFfRx46JzNm/n2pqaGKCj8kGQ0N0+LCY2t27/JDESTS48+edvNo89UmI7r2ee+8ONzbSvn3U3CxuGxCJUFsblZTQiROidKVmrqLVBunxOB+WDnVFDo2floK+PiosFHMpK0uZY2lEV/mqlvyyRL02uu6nbnGsp0dYlooZ9+/TwYMUi/H4e/676DNnqKtLZPGrV+Jn3rZNrO/GDT4sHVanto2fZsfz57Rrl5jFuXPiqZgKJ9H1nfW42Ar3Ck+fJ/JPBA0f0EXwetd18cVNTfycBNhhXr2a6Emzzk05edY8fCiWiK1s9oAGTOHiRXr7lh9ScRIti0ZbMEkBurkzBYT27LFPZwmKFpL6PVat1uY4zxSgWGGVW7fyuEZk0Rgd5XGGk2i5vcMzUA2imyu3d1u2ON0qk5O26//B21/46G+WwiPxN9/XBsr8YWXg7JCiUUAyxFolbBs/zQ75pHBOZ3IWvfTXpbgYdhpqsCXYgiAOiYwOhdRDSYTDtjudLyp/7xx5ba0VjYPRb65k/raJhwFWiTeHDLE6tW38NDuk6LQ4iXY9cOFiK6tWevu90XhU1ujllcsRLPIUiUdsays/J8GtW1RZyYNEn5W3x6amraKR1J+Xd6gRG06epCdPxE+I2wXPoNu3xU2DVTY08JGfHk6iQ2OhZRXLrL8zgs/Gnok1YydnCyzs3k3DwzxOtOCKv3vUJqObQmNfX/apERtk8rB2+nTKLZUWPkBGg5cTL0selay9vBZFGQ0f0EVQHMPysF/GzcuAZZeLvF4enyGvofenwLBVdH7jX67WpFdQG/r7qaKCiopEIu/cSadOiaukrY4fmQ8jOg14K0EhRokYnPnjFCX75k2xq0xhGTwIRr90P0WhUIN4DH5V7YtOJgXnGXMTDfBWgv2yfP08dEhIt6sYKgXNA9/VBh6Hx1+/mUbFQC7DMuoJHze/mLPozMGt/qN/+NuaAKoHdhqoGJHkBJ+XZEH0/xMjWhNGtCaMaE0Y0ZowojXxLxH/W3QVz3OZAAAAAElFTkSuQmCC"}`

	if err := json.Unmarshal([]byte(jsonData), &captchaData); err != nil {
		response.BaseError(c, fmt.Sprintf("json转译失败 %v", err))
		return
	}

	//msg := "一切ok"
	response.Vue3Response(c, captchaData)

}

// Login 		登录
// @Summary 	login接口
// @Description login接口详情介绍
// @Accept      json
// @Produce     json
// @Param dataInfo body string true "post传参"
// @Success 200 {object}  loginMethod.SwagLogin
// @Failure 404 {object}  loginMethod.AuthError
// @Tags        accountGroup
// @Router      /account/login [post]
func Login(c *gin.Context) {
	var userInfo jwt.MyClaims // 也可以写 var userInfo = &jwt.MyClaims{}
	if userInfoErr := c.ShouldBindJSON(&userInfo); userInfoErr != nil {
		response.BaseError(c, fmt.Sprintf("获取用户或密码异常 %v", userInfoErr))
		return
	}
	//logger.FileLogger.Info(fmt.Sprintf("user:%v, pass:%v", userInfo.Username, userInfo.Password))

	retryName := fmt.Sprintf("%vToken", userInfo.Username)
	// 防撞库 登录失败次数大于等于3次 禁止登陆
	retryInt, retryErr := redisGo.GetInt(retryName)
	if retryErr == nil && retryInt >= 3 {
		response.AuthError(c, fmt.Sprintf("%v 登录失败次数过多 请30分钟后再次尝试", userInfo.Username))
		return
	}

	// 密码加密
	passJm, passErr := aesEncryption.EnPwdCode([]byte(userInfo.Password))
	if passErr != nil {
		logger.Logger.Error(fmt.Sprintf("用户密码加密报错：%s", passErr))
	}
	//logger.FileLogger.Debug(fmt.Sprintf("密码：%v   加密后密码%v", userInfo.Password, passJm))

	// 密码验证  用户不存在或密码错误验证
	userData := userModel.SysUser{}
	userErr := model.GetOneFirst(&userData, "username = ? AND password = ? ", userInfo.Username, passJm) //可多条件
	if userErr != nil {
		retryInt += 1
		response.AuthError(c, fmt.Sprintf("用户或密码错误 请重新登录"))
		// 用户登录失败次数 写入
		retrySetErr := redisGo.SetAndEx(retryName, retryInt, fmt.Sprint(60*30)) //秒 30分钟超时
		if retrySetErr != nil {
			response.BaseError(c, fmt.Sprintf("%v 用户密码验证失败 写入缓存失败: %v", userInfo.Username, retrySetErr))
		}
		return
	}

	// 创建用户token
	token, err := jwt.CreateToken(userInfo.Username, userInfo.Password)
	if err != nil {
		response.BaseError(c, fmt.Sprintf("%v创建token失败 %v", userInfo.Username, err))
		return
	}
	logger.FileLogger.Info(fmt.Sprintf("1111%v", token))

	// 用户token写入缓存 # 作用就是用户在新浏览器登录时 生成新的登录token 在写入redis时 会覆盖之前redis存在的token 用户在之前的老浏览器方式时 redis则验证不通过 实现不能多点登录 只能唯一点登录
	redisErr := redisGo.SetAndEx(userInfo.Username, token, fmt.Sprint(60*60*settings.JwtConfSetting.JwtTime)) //JwtTime超时失效时间 单位小时
	if redisErr != nil {
		response.BaseError(c, fmt.Sprintf("用户token写入缓存失败: %v", redisErr))
		return
	}
	// 用户登录失败次数 清空
	redisCountErr := redisGo.SetAndEx(retryName, 0, fmt.Sprint(10)) //10秒超时
	if redisCountErr != nil {
		response.BaseError(c, fmt.Sprintf("用户登录失败次数 清空缓存失败: %v", redisCountErr))
		return
	}

	// 定义数据类型
	var result = make(map[string]interface{})
	result["accessToken"] = token
	result["tokenType"] = "Bearer"
	response.Success(c, result)
	//c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "token": token})
}

// LoOut		登录退出
// @Summary 	登录退出
// @Description 用户退出接口
// @Accept      json
// @Produce     json
// @Param Authorization header string true "传token 例: Bearer eyJhbGzI1...ciOiJIU"
// @Success 200 {object}  loginMethod.LoginOut
// @Failure 404 {object}  loginMethod.AuthError
// @Tags        accountGroup
// @Router      /account/logout [delete]
func LoOut(c *gin.Context) {
	var userInfo jwt.MyClaims
	if userInfoErr := c.ShouldBindJSON(&userInfo); userInfoErr != nil {
		response.BaseError(c, fmt.Sprintf("获取用户异常 %v", userInfoErr))
		return
	}

	logger.FileLogger.Info(fmt.Sprintf("注销用户:%v", userInfo.Username))
	if _, delErr := redisGo.Delete(fmt.Sprint(userInfo.Username)); delErr != nil {
		response.BaseError(c, fmt.Sprintf("用户注销 清空数据失败: %v", delErr))
		return
	}
	response.Success(c, "logout success")
}

// GetDeptRole 获取全部部门和全部角色
func GetDeptRole(c *gin.Context) {
	// 部门全部数据
	var allDept = make([]*deptModel.SysDept, 0)
	deptErr := model.GetAll(&allDept)
	if deptErr != nil {
		response.BaseError(c, fmt.Sprintf("获取部门数据错误 %v", deptErr.Error()))
		return
	}
	var deptResult = make(map[int64]string)
	for _, deptV := range allDept {
		deptResult[deptV.Id] = deptV.Name
	}

	// 获取全部角色
	var allRole = make([]*roleModel.SysRole, 0)
	roleErr := model.GetAll(&allRole)
	if roleErr != nil {
		response.BaseError(c, fmt.Sprintf("获取角色数据错误 %v", roleErr.Error()))
		return
	}
	var roleResult = make(map[int64]string)
	for _, roleV := range allRole {
		roleResult[roleV.Id] = roleV.Name
	}

	var resultData = map[string]map[int64]string{"dept": deptResult, "role": roleResult}
	response.Success(c, resultData)

}

// GetUser 获取用户
func GetUser(c *gin.Context) {

	// 取全部部门
	deptList, deptErr := deptMethod.GetDeptIdList(c)
	if deptErr != nil {
		response.BaseError(c, fmt.Sprintf("获取部门id列表失败 %v", deptErr.Error()))
		return
	}
	logger.FileLogger.Debug("递归得到部门id列表：", deptList)

	// 获取用户的全部的关联角色
	allUserRole := make([]*userModel.SysUserRole, 0)
	userRoleErr := model.GetAll(&allUserRole)
	if userRoleErr != nil {
		response.ParamsError(c, fmt.Sprintf("获取全部用户的管理角色-错误%v", userRoleErr.Error()))
		return
	}
	userRoleList := make(map[int64][]int64)
	for _, userRoleY := range allUserRole {
		if _, existData := userRoleList[userRoleY.UserID]; !existData { // 不存在  新增   !是取反 去掉!是存在
			userRoleList[userRoleY.UserID] = []int64{userRoleY.RoleID}
		} else { // 存在  追加
			userRoleList[userRoleY.UserID] = append(userRoleList[userRoleY.UserID], userRoleY.RoleID)
		}
	}
	logger.FileLogger.Debug("全部用户管理角色数据=====", userRoleList)

	// 分页参数 offset是偏移 limit是条数
	offset, limit, err := request.GetPagerParams(c)
	if err != nil {
		response.ParamsError(c, err.Error())
		return
	}
	// 排序条件 orderExpr=id desc
	orderExpr, err := request.GetOrderParams(c)
	if err != nil {
		response.ParamsError(c, err.Error())
		return
	}

	// 多条件like查询 数据拼接成string
	validFilterFields := []string{"username", "status"}
	likeExpr, likeValues := request.GetLikeFilterQuery(c, validFilterFields, 1)

	// 合并查询 deptList
	keyList := ""
	valueList := make([]interface{}, 0)
	if deptList[0] != 0 { // 条件成立 添加部门搜索  默认是搜索全部用户
		keyList = fmt.Sprintf(" %s In ?", "dept_id")
		valueList = []interface{}{deptList}
	}

	// 添加用户和状态搜索
	if len(likeExpr) != 0 { // 用户有搜索条件
		if len(keyList) != 0 { // 部门无搜索条件
			keyList = fmt.Sprintf("%v and", keyList)
		}
		keyList = fmt.Sprintf("%v %v", keyList, likeExpr)
		valueList = append(valueList, likeValues...)
	}

	logger.FileLogger.Debug("最终查询条件key===", keyList)
	logger.FileLogger.Debug("最终查询条件value===", valueList)

	// 查询db获取数据
	dictData := make([]*userModel.SysUser, 0)
	dataNum, getErr := model.GetListPage(&dictData, offset, limit, orderExpr, keyList, valueList...)
	if getErr != nil {
		response.BaseError(c, fmt.Sprintf("查询db获取数据失败 %v", err.Error()))
		return
	}

	var dictPyteList = make([]map[string]interface{}, 0)
	for _, userY := range dictData {
		oneDictPyte := map[string]interface{}{
			"id":       userY.Id,       // =
			"username": userY.Username, // =
			"nickname": userY.Nickname, // =
			"gender":   userY.Gender,   // my+
			"deptId":   userY.DeptID,   // my+
			"roleIds":  userRoleList[userY.Id],
			//"avatar":      userY.Avatar,                                // 图片
			"mobile":     userY.Mobile,                                   // =
			"status":     userY.Status,                                   //=
			"email":      userY.Email,                                    // =
			"createTime": userY.CreateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 更改时间
			"updateTime": userY.UpdateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 更改时间 my +
		}
		dictPyteList = append(dictPyteList, oneDictPyte)
	}
	var dictPyteResult = make(map[string]interface{})
	dictPyteResult["list"] = dictPyteList
	dictPyteResult["total"] = dataNum
	response.Success(c, dictPyteResult)
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	reqBody := &userMethod.UserData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("创建接口-获取数据失败%v", err.Error()))
		return
	}

	roleData := userModel.SysUser{}
	if roleErr := model.GetOneLast(&roleData, "username = ?", reqBody.Username); roleErr == nil {
		response.ParamsError(c, fmt.Sprintf("用户名%v已存在 请更换", reqBody.Username))
		return
	}

	// 创建用户
	nowTime := time.Now()
	item := &userModel.SysUser{
		Username:   reqBody.Username,
		Nickname:   reqBody.Nickname,
		Gender:     reqBody.Gender,
		DeptID:     reqBody.DeptID,
		Mobile:     reqBody.Mobile,
		Status:     reqBody.Status,
		Email:      reqBody.Email,
		Avatar:     "https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif", // 头像地址
		CreateTime: nowTime,
		UpdateTime: nowTime,
	}
	if err := model.Create(item); err != nil {
		response.BaseError(c, fmt.Sprintf("创建数据错误：%v", err.Error()))
		return
	}
	logger.FileLogger.Debug(fmt.Sprintf("%v 用户创建成功 正在关联角色...", reqBody.Username))

	// 更新用户的角色
	if changeErr := userMethod.ChangeUserRole(item.Id, *reqBody); changeErr != nil {
		response.BaseError(c, fmt.Sprintf("更新管理角色异常: %v", changeErr.Error()))
		return
	}

	response.Success(c, item)
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
	reqBody := &userMethod.UserData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("获取参数错误：%v", err.Error()))
		return
	}
	// 创建时间 string转time 东八区
	local, _ := time.LoadLocation("Asia/Shanghai")
	createTime, _ := time.ParseInLocation("2006-01-02 15:04:05", reqBody.CreateTime, local) // 转成东8区time
	updateTime := time.Now()                                                                // 更新时间为当前时间

	updateColumns := map[string]interface{}{
		"Username":   reqBody.Username,
		"Nickname":   reqBody.Nickname,
		"Gender":     reqBody.Gender,
		"DeptID":     reqBody.DeptID,
		"Mobile":     reqBody.Mobile,
		"Status":     reqBody.Status,
		"Email":      reqBody.Email,
		"CreateTime": createTime,
		"UpdateTime": updateTime,
	}
	// 已软删除的id 不会更改  但会返回更改成功
	if err := model.Update(&userModel.SysUser{}, updateColumns, "id = ?", reqBody.Id); err != nil {
		response.BaseError(c, fmt.Sprintf("更新错误：%v", err.Error()))
		return
	}

	// 更新用户的角色
	if changeErr := userMethod.ChangeUserRole(reqBody.Id, *reqBody); changeErr != nil {
		response.BaseError(c, fmt.Sprintf("更新管理角色异常: %v", changeErr.Error()))
		return
	}

	response.Success(c, "更新成功!")
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	delId := &userMethod.DelId{}
	if err := c.ShouldBindJSON(&delId); err != nil {
		response.ParamsError(c, fmt.Sprintf("删除接口，获取数据失败%v", err.Error()))
		return
	}

	// 批量删除
	idList := strings.Split(delId.Id, ",")
	for iNum := 0; iNum < len(idList); iNum++ {
		idInt, intErr := strconv.ParseInt(idList[iNum], 10, 64) //string转int64
		if intErr != nil {
			response.BaseError(c, fmt.Sprintf("：%v string转int64转译失败 %v", idList[iNum], intErr.Error()))
			return
		}

		// 删除该用户关联的角色
		delErr := model.DeleteByQuery(&userModel.SysUserRole{}, map[string]interface{}{"user_id": idInt})
		if delErr != nil {
			response.ParamsError(c, fmt.Sprintf("删除用户关联角色失败：%v", delErr.Error()))
			return
		}

		// 删除用户
		if err := model.DeleteByQuery(&userModel.SysUser{}, map[string]interface{}{"id": idInt}); err != nil {
			response.BaseError(c, fmt.Sprintf("删除用户数据错误：%v", err.Error()))
			return
		}
	}

	response.Success(c, "删除成功!")
}

// UserStatusChange 用户状态更改 1正常 0禁用
func UserStatusChange(c *gin.Context) {

	reqBody := &userMethod.UserData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("获取参数错误：%v", err.Error()))
		return
	}
	updateColumns := map[string]interface{}{"Status": reqBody.Status, "UpdateTime": time.Now()} // 更新
	// 已软删除的id不会更改  但会返回更改成功
	if err := model.Update(&userModel.SysUser{}, updateColumns, "id = ?", reqBody.Id); err != nil {
		response.BaseError(c, fmt.Sprintf("更新错误：%v", err.Error()))
		return
	}

	response.Success(c, "更新成功!")
}

// UserPassChange 用户密码更改
func UserPassChange(c *gin.Context) {
	reqBody := &userMethod.UserPass{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("获取参数错误：%v", err.Error()))
		return
	}

	// 密码加密
	passJm, passErr := aesEncryption.EnPwdCode([]byte(reqBody.Password))
	if passErr != nil {
		response.BaseError(c, fmt.Sprintf("用户密码加密异常: %v", passErr))
		return
	}

	updateColumns := map[string]interface{}{"Password": passJm, "UpdateTime": time.Now()} // 更新
	// 已软删除的id不会更改  但会返回更改成功
	if err := model.Update(&userModel.SysUser{}, updateColumns, "id = ?", reqBody.Id); err != nil {
		response.BaseError(c, fmt.Sprintf("更新错误：%v", err.Error()))
		return
	}

	response.Success(c, "更新成功!")
}
