package app

import (
	"fmt"
	"github.com/Mindyu/blog_system/config"
	"github.com/Mindyu/blog_system/middleware"
	"github.com/Mindyu/blog_system/middleware/jwt"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/Mindyu/blog_system/routers"
	"github.com/Mindyu/blog_system/utils/trie"
	trie2 "github.com/Mindyu/blog_system/persistence/trie"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
)

//博客应用服务器
type App struct {
	DB      *gorm.DB
	RPool   *redis.Pool
	KeyTrie *trie.Trie
	Conf    *config.BlogConfig
	Server  *gin.Engine
}

func NewApp() *App {
	return &App{}
}

//启动服务器
func (app *App) Launch() error {
	app.Conf = config.Config()
	app.initDB()
	// app.initRedis()
	app.initServer()
	app.initImgServer()
	app.initKeyTrie()
	app.initRouter()
	return app.Server.Run(fmt.Sprintf(":%d", app.Conf.ServerPort))
}

//关闭操作
func (app *App) Destory() {
	if app.DB != nil {
		app.DB.Close()
	}
	if app.RPool != nil {
		app.RPool.Close()
	}
}

//根据配置文件初始化数据库
func (app *App) initDB() {
	app.DB = persistence.GetOrm()
}

//根据配置文件初始化Redis
func (app *App) initRedis() {
	app.RPool = persistence.GetRedisPool()
}

//根据配置初始化服务器
func (app *App) initServer() {
	app.Server = gin.Default()
}

//根据博客名称、分类、标签、作者生成前缀树
func (app *App) initKeyTrie()  {
	app.KeyTrie = trie2.GetKeyTrie()
}

//初始化路由配置
func (app *App) initRouter() {
	//使用中间件
	app.Server.Use(middleware.Cors())  // 跨域请求解决
	routers.NewFrontRouter(app.Server) // 博客前台请求
	app.Server.Use(jwt.JWTAuth())      // Jwt认证，除登陆外所有请求都需要携带tokenren认证
	routers.NewAdminRouter(app.Server) // 博客后端请求
}

//配置图片文件服务器
func (app *App) initImgServer() {
	if _, err := os.Stat(app.Conf.ImgPath); err != nil {
		if err = os.MkdirAll(app.Conf.ImgPath, os.ModePerm); err != nil {
			panic("Create ImagePath Error!")
		}
	}
	//开启图片文件服务器访问,否则无法访问
	app.Server.StaticFS("/file", http.Dir(app.Conf.ImgPath))
}
