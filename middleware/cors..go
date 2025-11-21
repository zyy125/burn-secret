package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin")


func CorsMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		// 1. 允许任何来源访问 (生产环境要改成具体的域名，比如 "http://localhost:3000")
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

        // 2. 允许的请求方法
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

        // 3. 允许的请求头 (前端发送 JSON 时会带上 Content-Type)
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

        // 4. 处理 "预检" 请求 (Preflight)
        // 浏览器在发 POST 之前，经常会先发一个 OPTIONS 请求问服务器：“我能连你吗？”
        // 如果是 OPTIONS 请求，直接返回 204 (No Content) 并终止后续处理
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        // 5. 不是 OPTIONS 请求，放行，去执行后面的路由代码
        c.Next()
	}
}