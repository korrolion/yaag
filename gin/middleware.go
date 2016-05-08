package gin

import (
	"log"
	"net/http/httptest"
	"strings"

	"github.com/korrolion/yaag/middleware"
	"github.com/korrolion/yaag/yaag"
	"github.com/korrolion/yaag/yaag/models"
	"github.com/gin-gonic/gin"
	"regexp"
)

func Document() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !yaag.IsOn() {
			return
		}
		writer := httptest.NewRecorder()
		apiCall := models.ApiCall{}
		middleware.Before(&apiCall, c.Request)
		c.Next()
		if writer.Code != 404 {
			validPath := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
			currentPath := strings.Split(c.Request.RequestURI, "?")[0]
			if !validPath.MatchString(currentPath) {
				return
			}

			apiCall.MethodType = c.Request.Method
			apiCall.CurrentPath = currentPath
			apiCall.ResponseBody = ""
			apiCall.ResponseCode = c.Writer.Status()
			headers := map[string]string{}
			for k, v := range c.Writer.Header() {
				log.Println(k, v)
				headers[k] = strings.Join(v, " ")
			}
			apiCall.ResponseHeader = headers
			go yaag.GenerateHtml(&apiCall)
		}
	}
}
