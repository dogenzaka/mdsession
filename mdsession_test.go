package mdsession

import (
	"net/http"
	"net/http/httptest"

	. "github.com/smartystreets/goconvey/convey"
	"testing"

	mgo "gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
	"github.com/liverpool-music/mds"
	"fmt"
)


var DnName string = "DS_MDSESSION"
var DbName string = "DB_MDSESSION"
var Name string = "db" // access


func newMDS() {
	datastores := []map[string]interface{}{
		map[string]interface{}{
			"Use":  true,
			"Dn": DnName,
			"Type": "MongoDB",
			"DialInfo": map[string]interface{}{
				"Addrs":    []string{"localhost:27017"},
				"Database": DbName,
			},
		},
	}

	err := mds.Setup(datastores)
	if err != nil {
		panic(err)
	}

}

func newServer() *gin.Engine {
	var option *Option = &Option{
		Dn: DnName,
		Name: Name,
		Make: true,
	}

	var options *Options = &Options{
		Sets: []*Option{option},
	}

	g := gin.New()

	// append gin middleware
	g.Use(MiddlewareMongoDB(options))

	return g
}


func request(server *gin.Engine, method string, uri string) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, uri, nil)

	server.ServeHTTP(w, req)

	if err != nil {
		panic(err)
	}

	return w
}

func TestMDSession(t *testing.T) {

	newMDS()

	Convey("mdsession operation", t, func() {
		g := newServer()


		g.GET("/test", func(c *gin.Context) {
			ret, err := c.Get(Name)
			s, ok := ret.(*mgo.Session)

			if ret != nil && err == nil && ok {
				c.String(200, "OK dbs: " + fmt.Sprint(s.DatabaseNames()))
			} else {
				c.String(500, "NG")
			}
		})

		r := request(g, "GET", "/test")

		So(r.Code, ShouldEqual, 200)
		fmt.Println(r.Body.String())
	})
}
