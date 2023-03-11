package APIS

import (
	"deps/Server/DB"
	"deps/Server/Models"
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func CreateMainReserves(c *gin.Context) {
	var reqbody struct {
		Name        string
		Email       string
		Age         int
		Nationality string
	}
	c.Bind(&reqbody)

	MainForm := Models.Postmain{Name: reqbody.Name, Email: reqbody.Email, Age: reqbody.Age, Nationality: reqbody.Nationality}
	DB.DBconn.Create(&MainForm)
//os.Getenv("SENDGRID_API_KEY")
	// Send email
	from := mail.NewEmail("Example User", "jose.naranjo.martinez@udla.edu.ec")
	subject := "New reservation"
	to := mail.NewEmail("Example User", "naranjojose256@gmail.com")
	plainTextContent := fmt.Sprintf("Name: %s\nEmail: %s\nAge: %d\nNationality: %s", reqbody.Name, reqbody.Email, reqbody.Age, reqbody.Nationality)
	htmlContent := fmt.Sprintf("<strong>Name:</strong> %s<br><strong>Email:</strong> %s<br><strong>Age:</strong> %d<br><strong>Nationality:</strong> %s", reqbody.Name, reqbody.Email, reqbody.Age, reqbody.Nationality)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println("Error sending email:", err)
		c.JSON(500, gin.H{
			"message": "Error sending email",
		})
		return
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		log.Println("Email sent. Response:", response.StatusCode)
		c.JSON(200, gin.H{
			"message": "Reservation created and email sent",
		})
	} else {
		log.Println("Error sending email. Response:", response.StatusCode)
		c.JSON(500, gin.H{
			"message": "Error sending email",
		})
	}
}




