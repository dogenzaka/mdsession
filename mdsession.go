package mdsession

import (
	"github.com/gin-gonic/gin"
	"github.com/liverpool-music/mds"
)

type (
	Options struct {
		Sets []*Option
	}

	Option struct {
		Dn string
		Name string // access name
		Make bool
	}

)

func MiddlewareMongoDB(options *Options) (gin.HandlerFunc) {

	// pre-valid
	for _, option := range options.Sets {

		// DataStore
		ds, err := mds.GetDataStoreMongoDB(option.Dn)
		if err != nil {
			panic(err)
		}

		// Connect
		if !ds.Connected {
			err = ds.Connect()
			if err != nil {
				panic(err)
			}
		}

		// Session
		_, err = ds.GetSession(false)
		if err != nil {
			panic(err)

		}

	}
	return func (c *gin.Context) {
		for _, option := range options.Sets {
			ds, _ := mds.GetDataStoreMongoDB(option.Dn)
			s, _ := ds.GetSession(option.Make)
			defer s.Close()
			c.Set(option.Name, s)
		}

		c.Next()

	}
}
