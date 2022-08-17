package main
import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()
	r.GET("/ping",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"hey there",
			"server time":time.Now().String()})
	})
	r.Run(":3333")
}
