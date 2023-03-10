package APIS

import (
	"deps/Server/DB"
	"deps/Server/Models"

	"github.com/gin-gonic/gin"
)



func CreateReserves(c*gin.Context){

	var reqbody struct{
		Name string
		Email string	
		Phone string
	}
	c.Bind(&reqbody);


reserve := Models.Reserve{Name:reqbody.Name,Email:reqbody.Email,Phone:reqbody.Phone};
DB.DBconn.Create(&reserve);
}