package cache

import (
	"encoding/json"
	"fmt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/models/common"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/Mindyu/blog_system/utils"
	"github.com/bitly/go-simplejson"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/gommon/log"
	"strconv"
	"strings"
	"time"
)

func SetBlogListToRedis(req *common.BlogPageRequest, blogs []*models.Blog, total int) {
	keyArr := []string{req.SearchWords, strconv.Itoa(req.BlogTypeId), strconv.Itoa(req.SortType),
		strconv.Itoa(req.CurrentPage), strconv.Itoa(req.PageSize)}
	key := strings.Join(keyArr, "_")
	log.Info(key)

	b, _ := json.Marshal(blogs)
	_, err := persistence.GetR().Do("SET", key, string(b), "EX", "600") // 过期时间10分钟
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	key = key[:strings.LastIndex(key, "_")]
	key = key[:strings.LastIndex(key, "_")]
	log.Info(key)
	_, err = persistence.GetR().Do("SET", key, total, "EX", "600") // 过期时间10分钟
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	fmt.Println("redis set success")
}

func GetBlogListToRedis(req *common.BlogPageRequest) (blogs []*models.Blog, total int, err error) {
	keyArr := []string{req.SearchWords, strconv.Itoa(req.BlogTypeId), strconv.Itoa(req.SortType),
		strconv.Itoa(req.CurrentPage), strconv.Itoa(req.PageSize)}
	key := strings.Join(keyArr, "_")

	str, err := redis.String(persistence.GetR().Do("GET", key))
	if err != nil {
		fmt.Println("redis get string failed:", err)
		return
	}
	key = key[:strings.LastIndex(key, "_")]
	key = key[:strings.LastIndex(key, "_")]
	total, err = redis.Int(persistence.GetR().Do("GET", key))
	if err != nil {
		fmt.Println("redis get int failed:", err)
		return
	}

	fmt.Println(str)
	/*var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	_ = json_iterator.Unmarshal(utils.Str2Bytes(str), &blogs)*/
	jsonObject, _ := simplejson.NewJson(utils.Str2Bytes(str))
	rows, _ := jsonObject.Array()
	for _, row := range rows {
		tMap := row.(map[string]interface{})
		blog := &models.Blog{}
		blog.ID, _ = strconv.Atoi(string(tMap["id"].(json.Number)))
		blog.BlogTitle = tMap["blog_title"].(string)
		blog.Keywords = tMap["keywords"].(string)
		blog.Author = tMap["author"].(string)
		blog.ReadCount, _ = strconv.Atoi(string(tMap["read_count"].(json.Number)))
		blog.ReplyCount, _ = strconv.Atoi(string(tMap["reply_count"].(json.Number)))
		blog.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", tMap["created_at"].(string))
		blogs = append(blogs, blog)
	}
	// _ = json.Unmarshal(utils.Str2Bytes(str), &blogs) // 只能解析一条记录，而且日期格式有误
	fmt.Println("redis get success")
	return
}
