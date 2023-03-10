package main

import (
	"deps/Server/DB"
	"net/http"
	"github.com/gin-gonic/gin"
	"deps/Server/API_Routes"
	//"errors"
	//"fmt"
)

func init(){
	DB.DBconnect()
}
 
type reserve struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type reserveList struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var reserves = []reserve{
	{ID: "1", Name: "Juan", Email: "juan123@juan.com", Phone: "12345"},
}

func GetReserves(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, reserves)
} 

func main() {
	router := gin.Default();
	router.GET("/reserves", GetReserves);
	router.POST("/reservar",APIS.CreateReserves);
	router.Run("localhost:5000");

	//fmt.println("Server Running port 5000");
	
}
 