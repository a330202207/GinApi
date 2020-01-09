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

### 项目设置

```
//设置代理
vim /etc/profile

export GOPROXY=https://goproxy.cn

//进入项目，执行
go mod download

//后台运行并输出日志
nohup go run main.go > log.out 2>&1 &

//后台运行
nohup go run main.go >/dev/null 2>&1 &
```

### 权限管理配置
配置文件：middleware/casbin/casbin.go


https://casbin.org/docs/zh-CN/function