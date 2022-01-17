package router

import (
	"go_gin/common"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)

func InitRouter(r *gin.Engine) {
	r.GET("/s", First)
}
	func First(c *gin.Context) {
		ts := strconv.FormatInt(common.GetTimeUnix(), 10)
		res := map[string]interface{}{}
		params := url.Values{
			"name":  []string{"a"},
			"price": []string{"10"},
			"ts":    []string{ts},
		}
		res["sn"] = common.CreateSign(params)
		res["ts"] = ts
		common.RetJson("hello word")
	}
