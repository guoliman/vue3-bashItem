package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"vue3-bashItem/pkg/jwt"
	redisGo "vue3-bashItem/pkg/redis"
	"vue3-bashItem/pkg/response"
	"vue3-bashItem/pkg/settings"
	"vue3-bashItem/services/accountAdminMethod/roleMethod"
)

type SessionUserInfo struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	UpdatedAt   time.Time `json:"updated_at"`
	LastLoginAt time.Time `json:"last_login_at"`
}

type APIResponse struct {
	Code int              `json:"code"`
	Data *SessionUserInfo `json:"data"`
	Msg  string           `json:"msg"`
}

type PermData struct {
	RoleKey string
	Perms   []string
}

// config.yaml取默认路由 存在则为true 不存在为false
func splitIn(c *gin.Context) (string, bool) {
	routeList := settings.DefaultRouteSetting                          //取配置
	urlAddr := strings.Split(fmt.Sprintf("%v", c.Request.URL), "?")[0] // 切割去url 不带get参数

	for i := 0; i < len(routeList); i++ {
		if c.Request.Method == routeList[i]["UrlType"] && urlAddr == routeList[i]["Path"] {
			return urlAddr, true
		}
	}
	return urlAddr, false
}

// JwtMiddleware jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 免验证  例：get请求可以直接访问 无需登录
		//if c.Request.Method == "GET" || c.Request.Method == "POST" || c.Request.Method == "PUT" ||
		//	c.Request.Method == "PATCH" || c.Request.Method == "DELETE" {
		//	c.Next() // 继续执行 会执行下一个中间件 一直下去 最终执行url接口
		//	return
		//}

		//从请求头中获取token
		tokenStr := c.Request.Header.Get("Authorization")
		//token不存在
		if tokenStr == "" {
			response.AuthError(c, fmt.Sprintf("token 不存在"))
			c.Abort() //阻止执行
			return
		}
		//token格式错误  // key名是Authorization value是 Bearer+空格+token字符串
		tokenSlice := strings.SplitN(tokenStr, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			response.AuthError(c, fmt.Sprintf("token传参 格式错误"))
			c.Abort() //阻止执行
			return
		}
		//验证token 值不对或已过期
		tokenStruck, ok := jwt.CheckToken(tokenSlice[1]) // 解析token
		if !ok {
			response.AuthError(c, fmt.Sprintf("token值不正确 或 token已过期"))
			c.Abort() //阻止执行
			return
		}

		// redis缓存中验证token
		redisUser, redisUserErr := redisGo.GetString(tokenStruck.Username)
		if redisUserErr != nil { // redis缓存中不存在
			response.AuthError(c, fmt.Sprintf("%v token已过期 不存在缓存中", tokenStruck.Username))
			c.Abort() //阻止执行
			return
		}

		if redisUser != strings.Split(tokenStr, " ")[1] { // 多设备登录 token是最新访问的那个
			response.AuthError(c, fmt.Sprintf("%v token失效 在其他设备登录 ", tokenStruck.Username))
			c.Abort() //阻止执行
			return
		}

		//logger.FileLogger.Debug("token内第三段信息==", tokenStruck)
		c.Set("username", tokenStruck.Username) //  c.Set向接口里插入数据 tokenStruck是个人用户全部信息
		//sessionUserInfoRaw, _ := c.Get("userInfo")  //   c.Get在接口里获取数据

		// 后端路由权限认证url
		if urlAddr, urlStatus := splitIn(c); urlStatus != true { // 默认路由不验证
			aErr := roleMethod.UrlAuth(c, c.Request.Method, urlAddr) // 验证权限
			if aErr != nil {
				response.AuthError(c, fmt.Sprintf("无权限访问 %v", aErr))
				c.Abort() //阻止执行
				return
			}
		}

		c.Next()
	}
}
