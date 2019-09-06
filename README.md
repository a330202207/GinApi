# GinApi
```
目录
├── config  系统配置
│    └── config.ini 配置文件
│
├── controller      控制器
│     
├── doc             文档存放 
│     
├── middleware      中间件
│    ├── err          
│    ├── jwt                
│    └── loger          
├── models          模型
├── doc             文档存放
├── pkg             第三方包 
│    └── setting    系统设置        
├── routes          路由
├── service         服务
├── static          静态文件
├── storage         缓存日志文件
├── views           视图文件
└── README.md  
```

安装依赖包：

日志：
github.com/sirupsen/logrus
github.com/lestrrat-go/file-rotatelogs
github.com/rifflock/lfshook

数据加密:
golang.org/x/crypto/bcrypt

jwt:
github.com/dgrijalva/jwt-go

数据验证：
github.com/go-ozzo/ozzo-validation

https://github.com/gin-contrib/cors
