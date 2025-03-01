//模块名
module github.com/jonesashbur/go_Learning

//go sdk版本
go 1.22.6

//当前模块（项目）依赖包
//依赖包 版本
//直接依赖
require (
	//dependency latest
)

//间接依赖
require (
	//dependency latest  //indirect
)
//排除的第三方包
exclude (
	//dependency latest
)

//修改依赖包的路径和版本
//迁移，原始包无法访问，二次开发三方包，本地包替换原始包
replace (
	//source latest => target latest
)

//撤回
//当前项目作为其他项目的依赖，如果某个版本出现问题则撤回
retract (
// v1.0.0
)