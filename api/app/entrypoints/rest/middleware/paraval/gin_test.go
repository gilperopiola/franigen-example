package paraval

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupGinServer() (*httptest.ResponseRecorder, *gin.Context, *gin.Engine) {

	handler := func(c *gin.Context) {
		c.JSON(http.StatusOK, c.MustGet("validated"))
	}

	rec := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(rec)
	r.Use(GinValidate())
	r.GET("/no-params", handler)
	r.GET("/ignore-param/:aaaa", handler)
	r.GET("/validate-param/:aID", handler)
	r.GET("/validate-params/:aID/a/:bID", handler)
	return rec, c, r
}

func errBody(msg string) interface{} {
	return struct {
		Description string `json:"description"`
	}{
		Description: msg,
	}
}

func TestGinValidate(t *testing.T) {

	const invalid = "invalid url param"
	const missing = "missing url param"

	cases := []struct {
		url  string
		code int
		body interface{}
	}{
		{"/no-params", http.StatusOK, Validated{}},
		{"/ignore-param/123", http.StatusOK, Validated{}},
		{"/ignore-param/asd", http.StatusOK, Validated{}},
		{"/validate-param/asd", http.StatusBadRequest, errBody(invalid)},
		{"/validate-param/2", http.StatusOK, Validated{"aID": 2}},
		{"/validate-params/a/a/a", http.StatusBadRequest, errBody(invalid)},
		{"/validate-params/1/a/a", http.StatusBadRequest, errBody(invalid)},
		{"/validate-params/a/a/2", http.StatusBadRequest, errBody(invalid)},
		{"/validate-params/1/a/2", http.StatusOK, Validated{"aID": 1, "bID": 2}},
		{"/validate-params//a/2", http.StatusBadRequest, errBody(missing)},
	}

	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("url:'%s'", c.url), func(t *testing.T) {
			t.Parallel()

			rec, context, r := setupGinServer()

			context.Request, _ = http.NewRequest(http.MethodGet, c.url, nil)
			r.ServeHTTP(rec, context.Request)
			resp := rec.Result()
			body, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			if resp.StatusCode != c.code {
				t.Errorf("got %d, want %d", resp.StatusCode, c.code)
			}

			want, _ := json.Marshal(c.body)
			if string(body) != string(want) {
				t.Errorf("got %s, want %s", body, want)
			}
		})
	}

}
