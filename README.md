## LittlePrince

一个简单方便的 go 后端脚手架  

### 第三方依赖：

- web 框架 - gin
- 数据库 - gorm.v2
- 日志库 - logrus
- 配置库 - go-ini

*可以随意更换*

### 文件布局

conf - 存放配置文件

data - 连接数据源

db - 数据库操作

doc - 存放文档

moddleware - 中间件实现

model - 数据模型，包括请求、响应、数据库表

nosql - 非关系型数据库操作

pkg - 工具包

router - 路由的注册与实现

runtime - 存放程序产生依赖的数据

service - 业务实现层，当业务代码**比较复杂**时，router 中应该只做参数校验功能，具体的业务处理代码在 service 实现



