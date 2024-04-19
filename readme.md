```bash
my-microservice/
├── cmd/
│   └── myservice/
│       └── main.go       # 程序入口点，启动服务器
├── pkg/
│   ├── config/
│   │   └── config.go     # 配置相关代码，例如读取环境变量
│   ├── api/
│   │   ├── handler.go    # HTTP处理器，定义API端点逻辑
│   │   └── middleware.go # HTTP中间件，如身份验证、日志等
│   ├── service/
│   │   └── service.go    # 业务逻辑层，实现具体的业务需求
│   ├── repository/
│   │   └── repository.go # 数据访问层，与数据库交互
│   └── model/
│       └── model.go      # 数据模型，定义结构体和方法
├── internal/
│   └── util/
│       └── util.go       # 内部使用的实用工具函数
├── migrations/
│   └── 01_init_schema.sql # 数据库迁移脚本
├── tests/
│   ├── api_test.go       # API 测试
│   └── service_test.go   # 服务层测试
├── go.mod                # Go模块文件
└── go.sum                # Go模块的依赖树
```