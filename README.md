## YKlili-Blog

[![Build Status](https://travis-ci.org/sinmahod/yklili.svg?branch=master)](https://travis-ci.org/sinmahod/yklili)

简易个人博客系统，使用`Golang 1.7.4`+[`Beego 1.7.2`](https://github.com/astaxie/beego)，后台使用[`ace admin`](https://github.com/bopoda/ace)。

初学Go语言练手，目前只实现了比较简单的一些功能。

[Demo](https://blog.yklili.com)

配置数据库（仅支持mysql）`github.com/sinmahod/conf/app.conf`
```
#应用名称
appname = YKlili
#端口号
httpport = 8080
#运行模式（开发模式打印SQL语句）
runmode = dev
#默认站点ID，为数据库主键ID，这个设计有点失败（设计之初的构想是多站点模式）
siteid = 1

#数据库设置
DB.Type = mysql
DB.IP = localhost
DB.Port = 3306
DB.Name = yklili
DB.UserName = root
DB.Password = 123456

#缓存设置
#实际业务中并没有用到,只是为了熟悉了Golang操作redis的库(github.com/garyburd/redigo/redis）
Cache = true
Cache.Type = redis
Redis.IP = 127.0.0.1
Redis.Port = 6379
Redis.Password = qweqwe
Redis.DBNum = 0
```

######<font color=blue>代码中调用了结巴分词的Go语言版本，包含部分C++代码编译时需要gcc</font>
