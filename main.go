package main

import (
	"net/url"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/wonderivan/logger"
	"go.xiet16.com/gopmsweb/conf"
	"go.xiet16.com/gopmsweb/ctrl"
	"go.xiet16.com/gopmsweb/models"
	"go.xiet16.com/gopmsweb/modules/cache"
	"go.xiet16.com/gopmsweb/modules/response"
	"go.xiet16.com/gopmsweb/public/common"
)

func main() {
	logger.Info("/**********start*********/")
	Load()                     //加载配置
	gin.SetMode(gin.DebugMode) //开发环境
	//gin.SetMode(gin.ReleaseMode) //线上环境

	r := gin.Default()
	store, _ := redis.NewStoreWithPool(cache.RedisClient, []byte("secret"))

	r.Use(sessions.Sessions("gosession", store))
	r.Use(cors.New(GetCorsConfig())) //跨域
	r.Use(Auth())
	//r.Use(cors.Default()) //默认跨域
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", ctrl.Index)
	//r.GET("/upload/image",ctrl.)
	r.GET("/pong",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"pong",
		}):
	})
	r.Run(":8899") 
}

func Load() {
	c := conf.Config{}
	c.Routes = []string{"/pong", "/login", "/role/index", "/info", "/dashboard", "/logout"}
	conf.Set(c)
}

func GetCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://admin.duiniya.com", "http://localhost:9529", "http://localhost:9528", "http://localhost:9527", "http://localhost"}
	config.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"x-requested-with", "Content-Type", "AccessToken", "X-CSRF-Token", "X-Token", "Authorization", "token"}
	return config
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := url.Parse(c.Request.RequestURI)
		if err != nil {
			panic(err)
		}
		if common.InArrayString(u.Path, &conf.Cfg.Routes) {
			c.Next()
			return
		}
		session := sessions.Default(c)
		v := session.Get(conf.Cfg.Token)
		if v == nil {
			c.Abort()
			response.ShowError(c, "nologin")
			return
		}
		uid := session.Get(v)
		users := models.SystemUser{Id: uid.(int), Status: 1}
		has := users.GetRow()
		if !has {
			c.Abort()
			response.ShowError(c, "user_error")
			return
		}
		//特殊账号
		if users.Name == conf.Cfg.Super {
			return
		}
		menuModel := models.SystemMenu{}
		menuMap, err := menuModel.GetRouteByUid(uid)
		if err != nil {
			c.Abort()
			response.ShowError(c, "unauthorized")
			return
		}
		if _, ok = menuMap[u.Path]; !ok {
			c.Abort()
			response.ShowError(c, "unauthorized")
		}
		// access the status we are sending
		//status := c.Writer.Status()
		c.Next()
		//log.Println(status) //状态 200
	}
}