package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pr struct {
	Link  string
	Title string
	ID    string
}

type inpr struct {
	Name  string `json:"name" binding:"required"`
	Link  string `json:"link" binding:"required"`
	Title string `json:"title" binding:"required"`
	ID    string `json:"id" binding:"required"`
}

type participant struct {
	Name string
	Prs  []pr
}

func conn() *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	uri := os.Getenv("URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	coll := client.Database("aperta").Collection("participants")
	return coll
}

func getAll(c *gin.Context) {
	coll := conn()
	curr, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		panic(err)
	}
	var results []participant
	if err = curr.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	var outputs []participant
	for _, result := range results {
		curr.Decode(&result)
		if err != nil {
			panic(err)
		}
		outputs = append(outputs, result)
	}
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, outputs)
}

func getUser(c *gin.Context) {
	coll := conn()
	var output participant
	result := coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: c.Param("name")}})
	result.Decode(&output)
	c.IndentedJSON(http.StatusOK, output)
}

func insertParticipant(c *gin.Context) {
	coll := conn()
	name := c.Param("name")
	npar := participant{
		Name: name,
		Prs:  []pr{},
	}
	_, err := coll.InsertOne(context.TODO(), npar)
	if err != nil {
		panic(err)
	}
}

func insertPr(c *gin.Context) {
	coll := conn()
	var newpr inpr
	err := c.BindJSON(&newpr)
	if err != nil {
		panic(err)
	}
	npr := pr{
		Title: newpr.Title,
		Link:  newpr.Link,
		ID:    newpr.ID,
	}
	cpa := coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: newpr.Name}})
	var res participant
	cpa.Decode(&res)
	if res.Prs != nil {
		res.Prs = append(res.Prs, npr)
	} else {
		res.Prs = []pr{npr}
	}
	_, err = coll.ReplaceOne(context.TODO(), bson.D{{Key: "name", Value: newpr.Name}}, res)
	if err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))
	router.GET("/all", getAll)
	router.POST("/inuser/:name", insertParticipant)
	router.GET("/getuser/:name", getUser)
	router.POST("/inpr", insertPr)
	router.Run("localhost:8080")
}
