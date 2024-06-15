package main

import (
	"log"
	"plant-watering/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// // Defining the Playground handler
// func playgroundHandler() gin.HandlerFunc {
// 	h := playground.Handler("GraphQL", "/graph")

// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Setting up Gin
	r := gin.Default()
	r.POST("/graph", graphqlHandler())
	// r.GET("/playground", playgroundHandler())

	// 服务静态文件，将 dist 目录设置为根路径
	r.Static("/", "./frontend/dist")

	// 重定向所有未匹配的路由到 index.html
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	r.Run()
}
