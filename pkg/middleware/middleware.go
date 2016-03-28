/*
中间件
*/

package middleware

import (
	"errors"
	"log"
	"time"
	//"mime"
	//"io/ioutil"
	//"net/url"

	"app-server/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/gogap/logrus"
	"github.com/gogap/logrus/hooks/graylog"
)

// 输出请求参数的中间件
func InspectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if gin.IsDebugging() {
			// ct1 := c.Request.Header.Get("Content-Type")
			// ct, _, _ := mime.ParseMediaType(ct1)
			// log.Println(ct1, ct, "cao", ct == "application/x-www-form-urlencoded")

			// log.Println(c.Request.Form == nil, c.Request.PostForm == nil)
			// c.Request.ParseForm()
			// log.Printf("\nRequest params: %+v  \n", c.Request.PostForm)

			//log.Println(c.Request, "\n")

			// b, e := ioutil.ReadAll(c.Request.Body)
			// log.Println("fuckyou-------->", string(b), e)
			// log.Println(url.ParseQuery(string(b)))
			//log.Printf("\nRequest params: %+v , PostForm=%+v \n", c.Request.Form, c.Request.PostForm)
			//log.Println(c.Request.Form.Encode())
		}
	}
}

// jwt 认证中间件
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		backend := token.GetJwtBackend()

		err, accessToken := backend.TokenAuthentication(c.Request)
		if err == nil {
			userId := backend.GetUserIdFromToken(accessToken)
			if len(userId) == 0 {
				c.AbortWithError(401, errors.New("[Error]can't get userid from token"))
			} else {
				c.Set("userId", userId)
				log.Println("token中的userid = ", userId)
			}

			c.Next()
		} else {
			c.AbortWithError(401, err)
		}
	}
}

// graylog 中间件，可以保存到文件内
// Requests with errors are logged using logrus.Error().
// Requests without errors are logged using logrus.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func Graylog(utc bool) gin.HandlerFunc {
	//api_v1.Use(middleware.Graylog(time.ANSIC, false))
	glog, err := graylog.NewHook("192.168.1.112:12201", "app-server", nil)
	if err != nil {
		return nil
	}

	logrus.SetLevel(logrus.WarnLevel)
	logrus.AddHook(glog)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		entry := logrus.StandardLogger().WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
			"time":       end.Format("2006-01-02 15:04:05"),
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		}
		// else {
		// 	if gin.Mode() == "debug" {
		// 		entry.Info(time.Now().Format("2006-01-02 15:04:05"))
		// 	}
		// }
	}
}
