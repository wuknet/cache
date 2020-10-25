# cache
golang实现的内存缓存

用golang实现的一个非常简单用内存实现缓存的应用

详细了解可查看 https://887d.com

创建或更改缓存
SetCache("变量名","内容",超时时间/秒)

获取缓存
content:=GetCache("变量名")

删除缓存
DelCache("变量名")

怎么样，简单吧
