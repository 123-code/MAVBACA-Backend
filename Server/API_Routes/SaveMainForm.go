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

func SaveMainForm(c *gin.Context) {
	var reqbody struct {
		FirstName   string    
		LastName    string   
		Phone       string   
		Email       string    
		Contact     string   
		Residence   string    
		Adults      int        
		Children    int        
		Comments	string   
	}
	c.Bind(&reqbody)

	MainpageForm := Models.MainForm{FirstName: reqbody.FirstName,
		LastName:reqbody.LastName,
		Phone:reqbody.Phone,
		Email: reqbody.Email, 
		Contact: reqbody.Contact, 
		Residence: reqbody.Residence,
		Adults: reqbody.Adults,
		Children: reqbody.Children,
		Comments: reqbody.Comments,
	}
	DB.DBconn.Create(&MainpageForm)
//os.Getenv("SENDGRID_API_KEY")
	// Send email
	from := mail.NewEmail("Example User", "jose.naranjo.martinez@udla.edu.ec")
	subject := "New reservation"
	to := mail.NewEmail("Example User", "naranjojose256@gmail.com")
	                                plainTextContent := fmt.Sprintf("First Name: %s\nLast Name: %s\nPhone: %s\nEmail: %s\nContact: %s\nResidence: %s\nAdults: %d\nChildren: %v\nComments: %s",
	                                reqbody.FirstName, reqbody.LastName, reqbody.Phone, reqbody.Email, reqbody.Contact, reqbody.Residence, reqbody.Adults, reqbody.Children, reqbody.Comments)
									 htmlContent := fmt.Sprintf("<strong>First Name:</strong> %s<br><strong>Last Name:</strong> %s<br><strong>Phone:</strong> %s<br><strong>Email:</strong> %s<br><strong>Contact:</strong> %s<br><strong>Residence:</strong> %s<br><strong>Adults:</strong> %d<br><strong>Children:</strong> %v<br><strong>Comments:</strong> %s",
									 reqbody.FirstName, reqbody.LastName, reqbody.Phone, reqbody.Email, reqbody.Contact, reqbody.Residence, reqbody.Adults, reqbody.Children, reqbody.Comments)
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




