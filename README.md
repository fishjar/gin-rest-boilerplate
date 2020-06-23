# GIN+GROM 的 REST 项目模板

## 支持特性

- [GIN](https://github.com/gin-gonic/gin)+[GROM](https://github.com/jinzhu/gorm) 开箱即用
- 简易登录及`JWT`验证
- 模型中[`validator.v8`](https://godoc.org/gopkg.in/go-playground/validator.v8)数据校验

## 使用指引

### 开发

```sh
# 确保已安装go，及$GOPATH环境变量已配置
# Go 1.13 and above
go version
echo $GOPATH

# 设置代理
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

# 创建并进入目录
mkdir -p $GOPATH/src/github.com/fishjar/gin-rest-boilerplate && cd "$_"

# 克隆项目
git clone https://github.com/fishjar/gin-rest-boilerplate.git .

# 确认配置文件(尤其数据库相关配置)
vi config/config.go
vi db/db.go

# 如有需要，运行下列命令启动一个mysql数据库服务
docker-compose -f docker-compose-mysql.yml up -d

# 启动redis
docker-compose -f docker-compose-redis.yml up -d

# 安装依赖
go get

# 生成Swagger文档
swag init

# 开发启动
go run main.go

# 测试：登录
curl -X POST http://localhost:4000/admin/account/login \
-H "Content-Type: application/json" \
-d '{"username":"gabe","password":"123456"}'

# 测试：查询记录，注意替换<token>为实际值
curl http://localhost:4000/admin/users \
-H "Authorization: Bearer <token>"
```

### 部署

```sh
# build
GOOS=linux GOARCH=amd64 go build

# alpine build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
# docker启动
docker-compose -f docker-compose-alpine.yml up
```

## TODO

- 文件上传/下载，从 reader 读取数据
- 静态文件服务
- 原始 SQL 查询
- 模型的空字段
- 内部、外部重定向
- 单元测试
- 批量更新时 Hooks 不会运行
- 自动记录 createdBy/updatedBy/DeletedBy
  - https://github.com/qor/audited
- swagger 结构体校验
- 任务队列，定时任务
- 秒杀抢购
- 日志搜集
- ES
- 缓存

## 模型问题

- gorm tags 参考：https://gorm.io/docs/models.html
- binding tags 参考：https://godoc.org/gopkg.in/go-playground/validator.v8
- 时间格式比较严格，参考：https://golang.org/pkg/time/#pkg-constants
- 模型定义中全部使用指针类型，是为了可以插入 null 值到数据库，但这样会造成一些使用的麻烦
- 也可以使用"database/sql"或"github.com/guregu/null"包中封装的类型
- 但是这样会造成 binding 验证失效，目前没有更好的实现办法，所以暂时全部使用指针类型

- readonly clomn
- 时间字段格式化 Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
- 自定义验证器，结构体级别的验证器
- 如果想忽略某个字段 Password bool `json:"-"`
- https://zhuanlan.zhihu.com/p/91312616
- https://www.tizi365.com/archives/343.html
- JSON type https://github.com/jinzhu/gorm/issues/1935
- https://github.com/gin-gonic/gin/issues/961
- 结构体验证:https://github.com/go-playground/validator/issues/546
- 如果你想更新或忽略某些字段，你可以使用 Select，Omit
- https://colobu.com/2017/06/21/json-tricks-in-Go/
- omitempty 不会忽略某个字段，而是忽略空的字段，当字段的值为空值的时候，它不会出现在 JSON 数据中
- https://github.com/asaskevich/govalidator
