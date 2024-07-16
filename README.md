# polaris
**北极星资产统计平台**

## 版本

```version
    Go 1.19.1+
    
    gin 1.8.1
    
    mysql  5.7+
    
    redis 6.2.6+
```

## 目录结构

```目录
  main.py             入口文件
  
  go.mod              依赖文件
  
  configs             配置文件
  
  controller          控制器(主逻辑)
  
  crond               定时任务
  
  dosc                swag生成器
  
  middleware          中间件
  
  model               数据库
  
  pkg                 模块工具
  
  router              路由
  
  scripts             初始化数据目录
  
  services            服务调用函数
  
  log                 日志目录
```



## 基础功能

```bash
  路由管理     restful
  
  数据库管理   gorm        
  
  跨域        cors        
  
  配置管理     yaml
  
  探活模块     关联服务探活
  
  加解密模块   aes  
  
  用户认证     jwt
  
  防撞库      防撞库模块
  
  登录        单点登录功能
  
  权限管理     用户/角色/菜单
  
  日志模块     logger
  
  小功能       分页模块/搜索模块
  
  异常处理     服务返回异常处理
  
  request模块 统一返回状态
  
  crond      定时任务
```

## 安装依赖

```
  go mod init vue3-bashItem # 初始化
  go mod tidy          # 安装依赖
```

## swagger初始化

```
  swag init
```

## 初始化数据库
```
本地 默认配置文件  
go run main.go -mode=migrate #=可去掉

线上 指定配置文件  
./vue3-bashItem -c configs/prodConfig.yaml -mode=migrate
  
```

## 本地启动
```
本地
go run main.go
  
线上 
cd /backend/ && ./vue3-bashItem -c ./configs/prodConfig.yaml
```

## 编译+上线部署
```
# sh scripts/build-linux.bat  # windows编译到Linux 
# sh scripts/build-linux.sh   # mac编译到Linux     
    
# 测试环境自动部署  
sh scripts/stage.sh
 
# 线上环境自动部署  
sh scripts//prod.sh.sh
```
