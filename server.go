package main

import (
	"context"
	"log"
	"plant-watering/graph"

	"plant-watering/ent"
	"plant-watering/ent/migrate"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

// Defining the Graphql handler
func graphqlHandler(entClient *ent.Client) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewSchema(entClient))

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
	databaseFile := "file:./plant_warting.db?_fk=1"
	drv, err := sql.Open(dialect.SQLite, databaseFile)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer drv.Close()

	// 创建 Ent 客户端
	entClient := ent.NewClient(ent.Driver(drv))

	if err := entClient.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Setting up Gin
	r := gin.Default()
	r.POST("/graph", graphqlHandler(entClient))
	// r.GET("/playground", playgroundHandler())

	// 服务静态文件，将 dist 目录设置为根路径
	r.Static("/", "./frontend/dist")

	// 重定向所有未匹配的路由到 index.html
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	r.Run()
}
