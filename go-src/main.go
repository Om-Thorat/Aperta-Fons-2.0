package main

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pr struct {
	Link  string
	Title string
}

type inpr struct {
	Name  string
	Link  string
	Title string
}
type participant struct {
	Name string
	Prs  []pr
}

func conn() *mongo.Collection {
	uri := "mongodb+srv://om:fFm8snrFI1ZrXI2W@track01.debfrj4.mongodb.net/?retryWrites=true&w=majority"
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

func insertParticipant(c *gin.Context) {
	coll := conn()
	name := c.Param("name")
	npar := participant{
		Name: name,
	}
	_, err := coll.InsertOne(context.TODO(), npar)
	if err != nil {
		panic(err)
	}
}

func insertPr(c *gin.Context) {
	coll := conn()
	var newpr inpr
	c.BindJSON(&newpr)
	npr := pr{
		Title: newpr.Title,
		Link:  newpr.Link,
	}
	cpa := coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: newpr.Name}})
	var res participant
	cpa.Decode(&res)
	res.Prs = append(res.Prs, npr)
	_, err := coll.ReplaceOne(context.TODO(), bson.D{{Key: "name", Value: newpr.Name}}, res)
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
	router.POST("/inpr", insertPr)
	router.Run("localhost:8080")
}
