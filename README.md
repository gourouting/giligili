# Golang API 范例项目

## 目的

本项目采用了一系列 Golang 中比较流行的组件，可以以本项目为基础快速搭建 Restful Web API 。

## 特色

本项目已经整合了许多开发 API 所必要的组件：

1. Gin: 轻量级 Web 框架，自称路由速度是 Golang 最快的
2. GORM: ORM 工具。本项目需要配合Mysql使用
3. Gin-Session: Gin 框架提供的 Session 操作工具
4. Go-Redis: Golang Redis 客户端
5. godotenv: 开发环境下的环境变量工具，方便使用环境变量
6. Gin-Cors: Gin 框架提供的跨域中间件
7. 自行实现了国际化 i18n 的一些基本功能
8. 本项目是使用基于 Cookie 实现的 Session 来保存登录状态的，如果需要可以自行修改为 token 验证


本项目已经预先实现了一些常用的代码方便参考和复用:

1. 创建了用户模型
2. 实现了 `/api/v1/user/register` 用户注册接口
3. 实现了 `/api/v1/user/login` 用户登录接口
4. 实现了 `/api/v1/user/me` 用户资料接口(需要登录后获取 Session)
5. 实现了 `/api/v1/user/logout` 用户登出接口(需要登录后获取 Session)

本项目已经预先创建了一系列文件夹划分出下列模块:

1. `api` 文件夹就是MVC框架的 Controller，负责协调各部件完成任务
2. `model` 文件夹负责存储数据库模型和数据库操作相关的代码
3. `service` 负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. `serializer` 储存通用的 JSON 模型，把 `model` 得到的数据库模型转换成 API 需要的 JSON 对象
5. `cache` 负责 Redis 缓存相关的代码
6. `auth` 权限控制文件夹
7. `util` 一些通用的小工具
8. `conf` 放一些静态存放的配置文件，其中 `locales` 内放置翻译相关的配置文件

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建 `.env` 文件设置环境变量便于使用（建议开发环境使用）

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRE="" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

## Go 依赖

本项目使用 `go mod` 管理依赖。

```shell
export GOPROXY=https://goproxy.io
go mod tidy
go mod vendor
```

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)
