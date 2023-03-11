package APIS

import (
	"deps/Server/DB"
	"deps/Server/Models"
	"github.com/gin-gonic/gin"
)



func RegisterUser(c*gin.Context){

	var reqbody struct{
		UserName string
		Email string	
		Password string
	}
	c.Bind(&reqbody);


user := Models.User{	UserName:reqbody.UserName,Email:reqbody.Email,Password:reqbody.Password};
DB.DBconn.Create(&user);
}