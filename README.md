# 毒鸡汤 dujitang  
go 简单学习例子

#### 软件架构
网页效果和数据库数据由：https://github.com/egotong/nows  提供

网站：https://juxiaoba.com

程序访问地址：http://dujitang.juxiaoba.com

把资源打包
```cassandraql
go get -u github.com/shuLhan/go-bindata/...
执行go-bindata ./assets/...会出现一个bindata.go文件
再行main.go
```