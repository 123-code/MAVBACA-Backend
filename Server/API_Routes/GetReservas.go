package APIS

import (
	"deps/Server/DB"
	"deps/Server/Models"

	"github.com/gin-gonic/gin"
)

func GetReservas(c *gin.Context){
	var reserves []Models.Reserve
	DB.DBconn.Find(&reserves);
   c.JSON(200,gin.H{
	   "reserves":reserves,
   })
}