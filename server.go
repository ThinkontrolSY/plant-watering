package main

import (
	"log"
	"plant-watering/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
	"periph.io/x/host/v3"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewSchema())

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}
	// Initialize periph.io host
	if _, err := host.Init(); err != nil {
		log.Fatalf("Failed to initialize periph.io: %v", err)
		return
	}

	// Setting up Gin
	r := gin.Default()
	r.POST("/graph", graphqlHandler())
	// r.GET("/playground", playgroundHandler())

	// 服务静态文件，将 dist 目录设置为根路径
	r.Static("/", "./frontend/dist")

	// 重定向所有未匹配的路由到 index.html
	r.NoRoute(func(c *gin.Context) {
		log.Println(c.Request.URL.Path)
		if c.Request.URL.Path == "" || c.Request.URL.Path == "/" {
			c.File("./frontend/dist/index.html")
		} else {
			c.File("./frontend/dist" + c.Request.URL.Path)
		}
	})

	r.Run()
}
