package main

import (
	"deps/Server/DB"
	"net/http"
	"github.com/gin-gonic/gin"
	"deps/Server/API_Routes"
	"github.com/gin-contrib/cors"
	//"errors"
	//"fmt"
	"os"
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

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func main() {
	router := gin.Default();
	// middlewares 
	router.Use(cors.Default())

	// routes 
    router.GET("/",hello);
	router.GET("/reserves", GetReserves);
	router.POST("/reservar",APIS.CreateReserves);
	router.GET("/reservas",APIS.GetReservas);
	router.POST("/api/sendform",APIS.CreateMainReserves);
	router.POST("/register",APIS.RegisterUser);
	router.POST("/login",APIS.LoginPost);
	router.POST("api/mainform",APIS.SaveMainForm);
port := os.Getenv("PORT")
	if(port == ""){
port = "8080"
	}
	router.Run("0.0.0.0:"+port);


	//fmt.println("Server Running port 5000");
	
}
