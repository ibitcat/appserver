
使用golang开发的app服务器，框架使用gin，数据库采用mongodb，缓存使用redis。

项目集成了swagger ui，使用[yvasiyarov/swagger](https://github.com/yvasiyarov/swagger)生成api文档，类似beego的注解文档。


	git clone https://github.com/shuimu98/domi-appserver.git app-server

注意项目文件夹的名称为 **app-server**,因为有些自己写的package，用的是本地的路径。

golang版本1.6，使用vendor作为包里管理。

如果没有安装 govendor, 需要 `go get -u github.com/kardianos/govendor`

	govendor init
	govendor add +ext
	govendor list
	govendor fetch +m
	go build