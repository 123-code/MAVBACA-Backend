package APIS

import (
	"deps/Server/DB"
	"deps/Server/Models"
	"net/http"
	"time"
    "os"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

)
   


func LoginPost(c*gin.Context){
// email and password off req body.
	var reqbody struct{
		UserName string
		Password string
	}
	c.Bind(&reqbody);
// look up requested user in postgres databse
var user Models.User
DB.DBconn.First(&user, "UserName = ? ", reqbody.UserName);

if user.UserName == ""{
	c.JSON(404,gin.H{"message":"User not found"})
	return
} 

// compare entered hash with user hash 

err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(reqbody.Password));
if err != nil{
	c.JSON(http.StatusBadRequest,gin.H{"message":"Wrong password"})
	return
}

//generate jwt
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
"sub":user.ID,
"exp":time.Now().Add(time.Hour * 24 *30).Unix(),
})

tokenString,err := token.SignedString(os.Getenv("JWT_SECRET"));
if err != nil{
	c.JSON(http.StatusBadRequest,gin.H{"message":"Wrong password"})
	return
}

c.JSON(http.StatusOK,gin.H{
	"token":tokenString,
})

login := Models.Login{UserName:reqbody.UserName,Password:reqbody.Password};
DB.DBconn.Create(&login);
}