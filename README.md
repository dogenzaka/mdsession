mdsession During the development :)
=========

# What is mdsession?

Middleware which sets up an established database session according to a request at `*gin.Context`

# Features

- Support database
  - [x] MongoDB
  - [ ] Redis
  - [ ] RethinkDB

# Requirements

- [mds](https://github.com/dogenzaka/mds)
  - [mgo](https://github.com/go-mgo/mgo)
  - [gin](github.com/gin-gonic/gin)
  - [mapstructure](https://github.com/mitchellh/mapstructure)


## Test

- [goconvey](https://github.com/smartystreets/goconvey)

# Getgin started

```sh
TODO
```

# Example

```go

var option *Option = &mdsession.Option{
    Dn: DnName, // uniq datastore name
    Name: Name, // Key to the `*gin.Context`
    Make: true, // To connect every time

}

var options *Options = &mdsession.Options{
    Sets: []*Option{option},
}

g := gin.New()

// append gin middleware
g.Use(MiddlewareMongoDB(options))


g.GET("/test", func(c *gin.Context) {
    ret, err := c.Get(Name)
    s, ok := ret.(*mgo.Session)

    if ret != nil && err == nil && ok {
        c.String(200, "OK dbs: " + fmt.Sprint(s.DatabaseNames()))
    } else {
        c.String(500, "NG")
    }
})

```

> Please read the test code.

# Developer

## Test

```sh
$ go get github.com/mattn/gom # package manager
$ make init
$ make test
```

## Coverage

```sh
$ go get github.com/smartystreets/goconvey // Coverage library
$ make cover
```

# License

MIT License
