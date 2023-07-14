package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
	"vue3-bashItem/pkg/settings"
)

/*
使用文档 https://blog.csdn.net/leo_jk/article/details/123779032
go get github.com/golang-jwt/jwt  已升级 从github.com/dgrijalva/jwt-go升级过来的 不是必须要jwt/v4
*/

//var Jwtkey = []byte("3!gw6zf^3gGYay0nP0Htp@jemneerf5&") // 加盐 秘钥

type MyClaims struct {
	//Username string `form:"username" binding:"required"` // binding:"required"表示值必须存在 binding:"min=6,max=20"表示值最小6位最大20位
	Username string `form:"username"`
	Password string `from:"password"`
	jwt.StandardClaims
}

// 生成token 三段组成
func CreateToken(name string, password string) (string, error) {
	// 过期时间 单位秒   时time.Hour   分time.Minute   秒time.Second
	expireTime := time.Now().Add(time.Duration(settings.JwtConfSetting.JwtTime) * time.Hour)
	nowTime := time.Now() //当前时间
	claims := MyClaims{
		Username: name,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间戳
			IssuedAt:  nowTime.Unix(),    //当前时间戳
			Issuer:    "blogLeo",         //颁发者签名
			Subject:   "userToken",       //签名主题
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenStruct.SignedString([]byte(settings.JwtConfSetting.JwtKey))
}

// 获取 解析并验证token    值 &{11 afsadf { 1664449317  1664445717 blogLeo 0 userToken}}
func CheckToken(token string) (*MyClaims, bool) {
	// 解析成byte数据  [98 108 111 103 95 106 119 116 95 107 101 121]
	tokenObj, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.JwtConfSetting.JwtKey), nil
	})
	//logger.FileLogger.Info(fmt.Sprintf("CheckToken 解析成byte==%v",tokenObj))
	if tokenObj == nil {
		return nil, false
	}
	// 如上值 byte转string
	if key, _ := tokenObj.Claims.(*MyClaims); tokenObj.Valid {
		//logger.FileLogger.Info(fmt.Sprintf("token值 byte转string 正常 ==%v", key))
		return key, true // key是解析后正确的值 &{11 afsadf { 1664449317  1664445717 blogLeo 0 userToken}}
	} else {
		//logger.FileLogger.Info(fmt.Sprintf("CheckToken byte转string 失败"))
		return nil, false // 解析失败 值不正确 或 token过期
	}
}

// redis验证token
// redis里增加token
