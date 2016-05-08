package martiniyaag

import (
	"github.com/korrolion/yaag/middleware"
	"github.com/korrolion/yaag/yaag"
	"github.com/korrolion/yaag/yaag/models"
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httptest"
)

func Document(c martini.Context, w http.ResponseWriter, r *http.Request) {
	if !yaag.IsOn() {
		c.Next()
		return
	}
	apiCall := models.ApiCall{}
	writer := httptest.NewRecorder()
	c.MapTo(writer, (*http.ResponseWriter)(nil))
	middleware.Before(&apiCall, r)
	c.Next()
	middleware.After(&apiCall, writer, w, r)
}
