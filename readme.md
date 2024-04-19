```bash
myapp/
├── cmd/
│   └── myapp/          # 主应用程序入口
│       └── main.go
├── pkg/
│   ├── config/         # 配置相关代码
│   │   └── config.go
│   ├── logger/         # 日志库封装
│   │   └── logger.go
│   ├── api/            # API 控制器层
│   │   └── handlers.go
│   ├── service/        # 业务逻辑层
│   │   └── service.go
│   ├── repository/     # 数据访问层
│   │   └── repository.go
│   └── model/          # 数据模型定义
│       └── models.go
├── internal/
│   └── util/           # 内部工具和帮助函数
│       └── util.go
├── config/
│   └── app.yaml        # 外部配置文件
├── migrations/         # 数据库迁移文件
│   └── 01_init_schema.sql
├── tests/              # 测试代码
│   ├── api_test.go
│   └── service_test.go
├── go.mod              # Go 模块定义文件
└── go.sum              # Go 模块的依赖项锁定文件

```