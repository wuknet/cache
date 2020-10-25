package fun

import (
	"fmt"
	"strconv"
	"time"
	//"github.com/tidwall/gjson"
)

//////////////////////////////////////////////////////
type Datajg struct {
	title   string
	content interface{}
	expire  int
}

var (
	CacheVal    []Datajg
	SinCacheval Datajg
)

//////////////////////////////////////////////////////
//检测缓存是否存在
func CheckCache(title string) (int, bool) {
	i := 0
	Isexist := false
	for i = 0; i < len(CacheVal); i++ {
		if CacheVal[i].title == title {
			Isexist = true
			break
		}
	}
	return i, Isexist
}

//添加或设置缓存
func SetCache(title string, content interface{}, expire int) {
	SinCacheval.title = title
	SinCacheval.content = content
	SinCacheval.expire = expire

	index, Isexist := CheckCache(title) //检测缓存是否存在，如果存在返回索引数据
	if Isexist == false {               //如果不存在
		SinCacheval.title = title
		SinCacheval.content = content
		if expire != 0 { //如果是0不过期
			expire_unix, _ := strconv.Atoi(fmt.Sprintf("%v", time.Now().Unix()))
			expire = expire_unix + expire
		}
		SinCacheval.expire = expire
		CacheVal = append(CacheVal, SinCacheval)
	} else { //存在则编辑
		CacheVal[index].content = content
		CacheVal[index].expire = expire
	}
}

//获取缓存内容
func GetCache(title string) interface{} {
	DelExpireCache() //删除过期的缓存
	index, Isexist := CheckCache(title)
	//fmt.Println(len(CacheVal))
	if Isexist == false {
		return ""
	} else {
		return CacheVal[index].content
	}
}

//删除缓存项目
func DelCache(title string) {
	for i := 0; i < len(CacheVal); i++ {
		if CacheVal[i].title == title {
			CacheVal = append(CacheVal[:i], CacheVal[i+1:]...)
			break
		}
	}
}

//删除所有过期缓存项目
func DelExpireCache() {
	NowTime, _ := strconv.Atoi(fmt.Sprintf("%v", time.Now().Unix()))
	for i := 0; i < len(CacheVal); i++ {
		if NowTime-CacheVal[i].expire >= 0 && CacheVal[i].expire != 0 {
			CacheVal = append(CacheVal[:i], CacheVal[i+1:]...)
			i--
		}
	}
}
