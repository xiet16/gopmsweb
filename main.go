package main

import (
	"net/url"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"go.xiet16.com/gopmsweb/conf"
	"go.xiet16.com/gopmsweb/modules/cache"
	"go.xiet16.com/gopmsweb/modules/response"
	"go.xiet16.com/gopmsweb/public/common"
	"go.xiet16.com/gopmsweb/models"
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

	//conf.Set(gin.DebugMode)

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
		users := models.

	}
}
