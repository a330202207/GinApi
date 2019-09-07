# GinApi
```
目录
├── app             系统配置
│ 
├── config          系统配置
│    ├── config.ini     配置文件
│    └── LoadConfig     加载配置文件
│
├── controller      控制层
│     
├── doc             文档存放 
│     
├── middleware      中间件
│    ├── casbin         权限          
│    ├── jwt            jwt        
│    └── loger          日志    
│ 
├── model           模型层
│ 
├── doc             文档存放
│ 
├── package         第三方包 
│    ├── error          系统错误  
│    └── loger          日志  
│        
├── routes          路由
│ 
├── service         服务层
│ 
├── static          静态文件
│ 
├── storage         缓存日志文件
│ 
├── util            工具
│ 
├── views           视图文件
│ 
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
