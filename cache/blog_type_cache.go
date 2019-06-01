package cache

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/stores"
	"github.com/labstack/gommon/log"
	"sync"
)

var (
	blogTypeCache []*models.BlogType
	blogTypeMap   map[int]string
	sOnce         sync.Once
	lock          = &sync.RWMutex{}
)

func Cache() []*models.BlogType {
	sOnce.Do(cacheInit)
	lock.RLock()
	defer lock.RUnlock()
	return blogTypeCache
}

func Map() map[int]string {
	sOnce.Do(cacheInit)
	lock.RLock()
	defer lock.RUnlock()
	return blogTypeMap
}

func cacheInit() {
	types, err := stores.GetAllBlogType(nil)
	if err != nil {
		log.Error("博客类型缓存失败")
	}
	typeMap := map[int]string{}
	for _, val := range types {
		typeMap[val.ID] = val.TypeName
	}
	lock.Lock()
	defer lock.Unlock()
	blogTypeCache = types
	blogTypeMap = typeMap
}

func UpdateCache(blogType *models.BlogType){
	lock.Lock()
	defer lock.Unlock()
	blogTypeCache = append(blogTypeCache, blogType)
	blogTypeMap[blogType.ID] = blogType.TypeName
}