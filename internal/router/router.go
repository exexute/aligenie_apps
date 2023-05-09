package router

import (
	"AliGenieServer/internal/news"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSkillRoute)
}

func Init(r *gin.Engine, authKey, authValue string) {
	registerAliGenieServer(r, authKey, authValue)
	registerHeartBeat(r)
	v1 := r.Group("/api/v1")
	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

func registerHeartBeat(r *gin.Engine) {
	r.GET("status", func(context *gin.Context) {
		context.JSON(200, "ok")
	})
}

func registerAliGenieServer(r *gin.Engine, authKey, authValue string) {
	r.GET(fmt.Sprintf("/aligenie/%s.txt", authKey), func(c *gin.Context) {
		c.String(http.StatusOK, authValue)
	})
}

func registerSkillRoute(rg *gin.RouterGroup) {
	r := rg.Group("skill").Use()
	{
		r.POST("/test", news.GetNews)
	}
}
