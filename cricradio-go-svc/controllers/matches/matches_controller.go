package matches

import (
	"cricradio-go-svc/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListMatches(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	matches, err := services.MatchesService.List()
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, matches.Marshall())
}