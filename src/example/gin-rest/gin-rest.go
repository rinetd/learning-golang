// gin restfull sample

package main

import (
	//"os"
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/contrib/static"
	//"tsingson/middleware/logging"
	//log "github.com/cihub/seelog"
)

const dataFile = "./comments.json"

func main() {

	// database end
	r := gin.Default()
	//r.Use(logging.AccessLogger(os.Stdout))

	r.Static("/reflux", "/Users/qinshen/git/project/reactjs_seed/reflux-store/build")
	r.Use(static.Serve("/", static.LocalFile("/Users/qinshen/git/project/refluxweb/src/react_reflux/dest", false)))
	// This handler will match /user/john but will not match neither /user/ or /user

	//comments
	r.GET("/products", func(c *gin.Context) {

		message := []byte(`[
    {"id": 1, "title": "iPad 4 Mini", "price": 500.01, "inventory": 2, "image": "../common/assets/ipad-mini.png"},
    {"id": 2, "title": "H&M T-Shirt White", "price": 10.99, "inventory": 10, "image": "../common/assets/t-shirt.png"},
    {"id": 3, "title": "Charli XCX - Sucker CD", "price": 19.99, "inventory": 5, "image": "../common/assets/sucker.png"}
]`)

		c.String(200, string(message))
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		message := "Hello " + name
		fmt.Println(message)
		c.JSON(200, gin.H{"message": message})
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/join/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Params.ByName("name")
		action := c.Params.ByName("action")
		message := name + " is " + action

		c.JSON(200, gin.H{"user": name, "action": action, "message": message})
	})

	r.GET("/search", func(c *gin.Context) {
		// You need to call ParseForm() on the request to receive url and form params first
		c.Request.ParseForm()

		firstname := c.Request.Form.Get("firstname")
		lastname := c.Request.Form.Get("lastname")

		message := "Hello " + firstname + lastname
		c.String(200, message)
	})

	// session with redis

	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/session/set", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count += 1
		}
		count = 2
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	r.GET("/session/get", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count += 1
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8090")
}
