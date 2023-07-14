package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
	menuAdmin "vue3-bashItem/controller/accountAdminApi/menuApi"
	roleAdmin "vue3-bashItem/controller/accountAdminApi/roleApi"
	"vue3-bashItem/controller/example/testExample"
	"vue3-bashItem/middleware"
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/settings"

	"vue3-bashItem/controller/accountAdminApi/deptApi"
	"vue3-bashItem/controller/accountAdminApi/dictApi"
	userAdmin "vue3-bashItem/controller/accountAdminApi/userApi"
)

func InitRouter() *gin.Engine {
	gin.SetMode(settings.ServerSetting.RunMode) // 设置运行模式
	r := gin.New()                              // 生成引擎

	//  ========中间件 生效在路由前==========
	r.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, true)) // log中间件  终端输出信息
	r.Use(middleware.Cors())                                // 跨域
	r.Use(middleware.Recovery())                            // recovery中间件 异常报错，服务不会崩溃
	r.Use(middleware.RequestId())                           // request id中间件   + 日志进入请求显示
	//r.Use(middleware.JwtMiddleware())  	// 登录认证校验

	//=========== route example==============
	//r.GET("/healthz", healthz.HealthZ)						   // 一层路径 不加组
	//r.GET("/authTest", middleware.AuthCheck(), healthz.AuthTest) // 一层路径 不加组 加中间件

	//openAPI := r.Group("/api/open", middleware.aa(), middleware.bb())  //  一级路由 可以写多个
	//{
	//	openAPI.GET("/accounts", middleware.cc()，open.ListAccounts)
	//}

	bashGroup := r.Group("/api") // 一级路由 可以写多个
	{
		// swagger 二级路由
		swagGroup := bashGroup.Group("/swag")
		{
			// swagger web访问http://127.0.0.1:8095/api/swag/swagger/index.html
			swagGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
			// swage example
			swagGroup.GET("/aa/bb", testExample.Bb)
			swagGroup.GET("/cc", testExample.Helloworld)
			swagGroup.DELETE("/loginAA", testExample.LoginBB) // swagger也是post类型

		}

		// 权限管理 二级路由
		authAdmin := bashGroup.Group("authAdmin")
		{

			// 字典管理
			dictGroup := authAdmin.Group("/dict", middleware.JwtMiddleware())
			{
				// 字段类型 也叫字段组
				dictGroup.GET("/dictTypes", dictApi.GetDictType)     // 获取字典类型
				dictGroup.POST("/dictTypes", dictApi.CreateDictType) // 新增字典类型
				dictGroup.PUT("/dictTypes", dictApi.PutDictType)     // 更改字典类型
				dictGroup.DELETE("/dictTypes", dictApi.DelDictType)  // 删除字典类型
				// 字段类型子级
				dictGroup.GET("/dictSon", dictApi.GetDictSon)     // 获取字典子级
				dictGroup.POST("/dictSon", dictApi.CreateDictSon) // 创建字典子级
				dictGroup.PUT("/dictSon", dictApi.PutDictSon)     // 更改字典子级
				dictGroup.DELETE("/dictSon", dictApi.DelDictSon)  // 删除字典子级
			}
			// 部门管理
			deptGroup := authAdmin.Group("/dept", middleware.JwtMiddleware())
			{
				deptGroup.GET("/deptOptions", deptApi.SelectDept) // 获取部门下拉菜单 递归
				deptGroup.GET("/deptTree", deptApi.GetDept)       // 获取部门 递归
				deptGroup.POST("/deptTree", deptApi.CreateDept)   // 新增部门
				deptGroup.PUT("/deptTree", deptApi.PutDept)       // 更改部门
				deptGroup.DELETE("/deptTree", deptApi.DelDept)    // 删除部门数据
			}
			// 菜单管理
			menuGroup := authAdmin.Group("/menus", middleware.JwtMiddleware())
			{
				menuGroup.GET("/menuOperate", menuAdmin.GetMenu)          // 获取全部菜单 递归
				menuGroup.POST("/menuOperate", menuAdmin.CreateMenu)      // 新增菜单
				menuGroup.PUT("/menuOperate", menuAdmin.PutMenu)          // 更改菜单
				menuGroup.DELETE("/menuOperate", menuAdmin.DelMenu)       // 删除菜单
				menuGroup.GET("/dirMenuS", menuAdmin.DirAndMenuSelect)    // 获取目录和菜单(用于新增或修改菜单时使用) 递归
				menuGroup.GET("/menuSelect", menuAdmin.RoleGetMenuSelect) // 角色分配权限时获取下拉菜单 递归

				//menuGroup.GET("/routes", menuAdmin.MenuRoutes)       // 左侧动态路由 假数据
				menuGroup.GET("/routes", menuAdmin.GetMoveRoute) // 左侧动态路由 真数据
			}
			// 角色管理
			roleGroup := authAdmin.Group("/roles", middleware.JwtMiddleware())
			{
				roleGroup.GET("/roleOperate", roleAdmin.GetRole)         // 获取角色
				roleGroup.POST("/roleOperate", roleAdmin.CreateRole)     // 创建角色
				roleGroup.PUT("/roleOperate", roleAdmin.UpdateRole)      // 更改角色
				roleGroup.DELETE("/roleOperate", roleAdmin.DeleteRole)   // 删除角色
				roleGroup.POST("/roleSonMenu", roleAdmin.GetRoleSonMenu) // 获取角色的菜单权限
				roleGroup.PUT("/roleSonMenu", roleAdmin.PostRoleSonMenu) // 更新角色的菜单权限
				roleGroup.GET("/roleOptions", roleAdmin.SelectRole)      // 获取角色下拉列表
			}
			// 用户管理
			userGroup := authAdmin.Group("/users")
			{
				// 用户管理
				userGroup.GET("/userOperate", middleware.JwtMiddleware(), userAdmin.GetUser)                 // 获取用户
				userGroup.POST("/userOperate", middleware.JwtMiddleware(), userAdmin.CreateUser)             // 创建用户
				userGroup.PUT("/userOperate", middleware.JwtMiddleware(), userAdmin.UpdateUser)              // 更改用户
				userGroup.DELETE("/userOperate", middleware.JwtMiddleware(), userAdmin.DeleteUser)           // 删除用户
				userGroup.PATCH("/userStatusChange", middleware.JwtMiddleware(), userAdmin.UserStatusChange) // 用户状态更改
				userGroup.PATCH("/userPassChange", middleware.JwtMiddleware(), userAdmin.UserPassChange)     // 用户密码更改
				userGroup.GET("/getDeptRole", middleware.JwtMiddleware(), userAdmin.GetDeptRole)             // 获取全部部门和全部角色

				// 登录管理
				userGroup.POST("/login", userAdmin.Login)                                // 登录
				userGroup.DELETE("/logout", middleware.JwtMiddleware(), userAdmin.LoOut) // 注销
				//userGroup.GET("/captcha", userAdmin.Captcha) // 获取图片验证
				userGroup.GET("/me", middleware.JwtMiddleware(), userAdmin.Me) // 获取当前登录用户信息
			}

		}

	}

	return r
}
