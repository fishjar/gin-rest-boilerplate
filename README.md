# GIN+GROM 的 REST 项目模板

## 支持特性

- [GIN](https://github.com/gin-gonic/gin)+[GROM](https://github.com/jinzhu/gorm) 开箱即用
- 简易登录及`JWT`验证、续签
- 模型中[`validator.v8`](https://godoc.org/gopkg.in/go-playground/validator.v8)数据校验

## 缺陷（待改进）

- 配置文件未区分开发、测试、生产环境
- 日志工具未使用协程，且未按级别及日期分开多文件
- 中间件获取`body`数据流后，又写入`body`中，可能对性能有影响
- 实例中缺少批量增、删、改的功能
- 缺少文件上传、下载功能

## 目录结构

```sh
├── config
│   └── config.go       #配置文件
├── db
│   └── db.go           #数据库实例
├── handler
│   ├── foo.go          #示例handler
│   └── login.go        #登录handler
├── log                 #日志目录
│   └── gin.log
├── main.go             #主程
├── middleware          #中间件包
│   ├── bodyLogger.go   #日志中间件
│   └── jwtAuth.go      #JWT验证中间件
├── model
│   ├── base.go         #共用model
│   ├── foo.go          #示例model
│   └── user.go         #用户model
├── README.md
├── router              #路由配置
│   └── router.go
└── utils               #工具包
    ├── jwt.go
    ├── logger.go
    └── md5.go
```

## 使用指引

```sh
# 确保已安装go，及$GOPATH环境变量已配置
go version
echo $GOPATH

# 创建并进入目录
mkdir -p $GOPATH/src/github.com/fishjar/gin-rest-boilerplate && cd "$_"

# 克隆项目
git clone https://github.com/fishjar/gin-rest-boilerplate.git .

# 确认配置文件(尤其数据库相关配置)
vi config/config.go
vi db/db.go

# 如有需要，运行下列命令启动一个mysql数据库服务
# 否则跳过此行
sudo docker-compose -f db/docker-compose.mysql.yml up -d

# 安装依赖
go get

# 启动
go run main.go

# 测试：登录
curl -X POST http://localhost:4000/login/account \
-H "Content-Type: application/json" \
-d '{"userName":"gabe","userType":"admin","password":"123456"}'

# 测试：创建记录，注意替换<token>为实际值
curl -X POST http://localhost:8000/foos \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <token>" \
-d '{"name":"gabe","good_time":"2019-06-06T00:00:00Z","status":1}'

# 测试：查询记录，注意替换<token>为实际值
curl http://localhost:8000/foos \
-H "Authorization: Bearer <token>"
```

## TODO

- 文件上传/下载，从 reader 读取数据
- 静态文件服务
- req，res 记录
- 原始 SQL 查询
- 部署
- 嵌套路由组
- 在中间件中使用 Goroutine
- 定义路由日志的格式
- 模型的空字段
- ShouldBindUri
- Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
- 批量创建、更新
- 绑定表单数据至自定义结构体
- c.MustGet("example").(string)
- 自定义验证器，结构体级别的验证器
- 内部、外部重定向
- 测试
- Official CORS gin's middleware
- 事物
- https://colobu.com/2017/06/21/json-tricks-in-Go/
- omitempty 不会忽略某个字段，而是忽略空的字段，当字段的值为空值的时候，它不会出现在 JSON 数据中
- 如果想忽略某个字段 Password bool `json:"-"`
- https://zhuanlan.zhihu.com/p/91312616
- https://www.tizi365.com/archives/343.html
- JSON type https://github.com/jinzhu/gorm/issues/1935
- https://github.com/gin-gonic/gin/issues/961
- 结构体验证:https://github.com/go-playground/validator/issues/546

## 问题

// gorm tags 参考：https://gorm.io/docs/models.html
// binding tags 参考：https://godoc.org/gopkg.in/go-playground/validator.v8
// 时间格式比较严格，参考：https://golang.org/pkg/time/#pkg-constants
// 模型定义中全部使用指针类型，是为了可以插入 null 值到数据库，但这样会造成一些使用的麻烦
// 也可以使用"database/sql"或"github.com/guregu/null"包中封装的类型
// 但是这样会造成 binding 验证失效，目前没有更好的实现办法，所以暂时全部使用指针类型
