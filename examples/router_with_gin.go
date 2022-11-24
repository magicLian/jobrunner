package examples

import (
	"github.com/gin-gonic/gin"
	"github.com/magicLian/jobrunner"
)

func main() {
	routes := gin.Default()

	// Resource to return the JSON data
	routes.GET("/jobrunner/json", JobJson)

	routes.Run(":8080")
}

func JobJson(c *gin.Context) {
	// returns a map[string]interface{} that can be marshalled as JSON
	c.JSON(200, jobrunner.StatusJson())
}

func JobHtml(c *gin.Context) {
	// Returns the template data pre-parsed
	c.HTML(200, "", jobrunner.StatusPage())

}
