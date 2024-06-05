### README 说明参考 gorse 项目

这里的调试修改 gorse-in-one :
```
go run cmd/gorse-in-one/main.go -c config/config.toml

➜  gorse_for_me git:(main) cat config/config.toml
[database]
data_store = "mongodb://root:password@127.0.0.1:27017/gorse?authSource=admin&connect=direct"
cache_store = "redis://127.0.0.1:6379/0"

[master]
port = 8086
host = "0.0.0.0"
http_port = 8088
http_host = "0.0.0.0"

[recommend]
cache_size = 5000
cache_expire = "168h"
fit_jobs = 4
n_neighbors = 10

```

数据库连接配置，是docker搭起来的MongoDB，和redis。

具体配置参考 deploy/gorse_inone_mongo

这里搭建器服务，解决panic问题，另外可能需要实现推荐结果外发。


