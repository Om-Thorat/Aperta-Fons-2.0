package run

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Aperta",
	})
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
		Name:   name,
		Prs:    []pr{},
		Points: 0,
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
		Title:  newpr.Title,
		Link:   newpr.Link,
		ID:     newpr.ID,
		Points: 10,
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

// func static(c *gin.Context) {
// 	// read from file
// 	data, err := os.ReadFile(fmt.Sprintf("./build/assets/%s", c.Param("p")))
// 	if err != nil {
// 		c.String(http.StatusTeapot, "hi")
// 	}
// 	switch path.Ext(c.Request.URL.Path) {
// 	case ".html":
// 		c.Header("Content-Type", "text/html")
// 	case ".css":
// 		c.Header("Content-Type", "text/css")
// 	case ".js":
// 		c.Header("Content-Type", "application/javascript")
// 		// ...
// 	}
// 	_, _ = c.Writer.Write(data)
// }
