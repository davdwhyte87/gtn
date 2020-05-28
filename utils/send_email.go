package utils


import (
	"html/template"
	"bytes"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
	"log"
)

// EmailData ... Data format for emails
type EmailData struct {
	Title string
	ContentData interface{}
	EmailTo string
	Template string
}

// SendEmail ... This sends emails to users from the app
func SendEmail(data EmailData) (bool) {
	// "./utils/html_templates/hello.html"
	template, errTemp := template.ParseFiles("utils/html_templates/"+data.Template)
	if errTemp != nil {
	panic(errTemp)
	}
	var buf bytes.Buffer
	// data := map[string]interface{}{"Msg": "Holla!"}
	errExc := template.Execute(&buf, data.ContentData)
	
	if errExc != nil {
	panic(errExc)
	}

	from := mail.NewEmail("Travelfy", os.Getenv("EMAIL"))
	subject := data.Title
	to := mail.NewEmail("",data.EmailTo)
	plainTextContent := "Hello"
	htmlContent := buf.String()
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, errMail := client.Send(message)
	if errMail != nil {
		log.Println(errMail)
		return false
	}
	return true
}