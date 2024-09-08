# gee2

golang 版本为1.22.5

本项目使用 Go 语言实现一个简单的 Web 框架. 这个框架中的很多部分实现的功能都很简单，但是尽可能地体现一个框架核心的设计原则。例如Router的设计，虽然支持的动态路由规则有限，但为了性能考虑匹配算法是用Trie树实现的，Router最重要的指标之一便是性能。
其中很多设计参考了Gin. 

* 上下文
* 前缀树路由
* 分组控制
* 中间件
* 模板Template
* 错误恢复
