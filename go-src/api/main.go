package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func conn() *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	uri := os.Getenv("URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	coll := client.Database("aperta").Collection("participants")
	return coll
}

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))
	router.LoadHTMLFiles("build/index.html")
	router.StaticFS("/assets", http.Dir("build/assets"))
	// router.GET("/assets/:p", static)
	router.GET("/", home)
	router.GET("/all", getAll)
	router.POST("/inuser/:name", insertParticipant)
	router.GET("/getuser/:name", getUser)
	router.POST("/inpr", insertPr)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

func main() {
	router.Run("localhost:8080")
}
