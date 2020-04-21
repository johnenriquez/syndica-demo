package main

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	SGApiKey = "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"
)

func sendToEmail(name, email, subject, textbody, htmlbody string) error {
	from := mail.NewEmail("Syndica", "peter@syndica.net")
	to := mail.NewEmail(name, email)
	message := mail.NewSingleEmail(from, subject, to, textbody, htmlbody)
	client := sendgrid.NewSendClient(SGApiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println("Error sending email: ", err.Error())
		return err
	}
	if response.StatusCode != 202 {
		log.Println("Error sending email: ", response.StatusCode)
		log.Println(response.Headers)
		log.Println(response.Body)
		return fmt.Errorf("error sending email: %d", response.StatusCode)
	}
	return nil
}

func SendContactUsEmail(name, email, textbody string) error {
	from := mail.NewEmail(name, email)
	to := mail.NewEmail("Contact Us", "peter@syndica.net")
	subject := "Contact Us Form"
	message := mail.NewSingleEmail(from, subject, to, textbody, textbody)
	client := sendgrid.NewSendClient(SGApiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println("Error sending email: ", err.Error())
		return err
	}
	if response.StatusCode != 202 {
		log.Println("Error sending email: ", response.StatusCode)
		log.Println(response.Headers)
		log.Println(response.Body)
		return fmt.Errorf("error sending email: %d", response.StatusCode)
	}
	return nil
}

func SendWelcomeEmailMBA(email, hash string) {
	subject := "Syndica Signup"
	textbody := fmt.Sprintf("Hello,\n\nThank you for signing up as an MBA Advisor on the Syndica platform.\n\nIn order to verify your email address and continue your registration process, please visit this URL:\nhttps://platform.syndica.net/verify?e=%s&q=%s\n\nAfter you register, please log in and complete your user profile.\nIf you have any problems signing up please contact me.\n\nPeter | Founder\npeter@syndica.net", email, hash)
	htmlbody := fmt.Sprintf("Hello,<br><br>Thank you for signing up as an MBA Advisor on the Syndica platform.<br><br>In order to verify your email address and continue your registration process, please visit this URL:<br>https://platform.syndica.net/verify?e=%s&q=%s<br><br>After you register, please log in and complete your user profile.<br>If you have any problems signing up please contact me.<br><br>Peter | Founder<br>peter@syndica.net", email, hash)
	sendToEmail("", email, subject, textbody, htmlbody)
}

func SendWelcomeEmailClient(email, hash string) {
	subject := "Syndica Signup"
	textbody := fmt.Sprintf("Hello,\n\nThank you for signing up on the Syndica platform.\n\nIn order to verify your email address and continue the submission process, please visit this URL:\nhttps://platform.syndica.net/verify?e=%s&q=%s\n\nAfter you submit your startup information, we will post your startup onto the Syndica platform.\nIf your startup receives responses from our MBA network, we will provide the responses to you via email.\n\nThank you again for your interest.\n\nPeter | Founder\npeter@syndica.net", email, hash)
	htmlbody := fmt.Sprintf("Hello,<br><br>Thank you for signing up on the Syndica platform.<br><br>In order to verify your email address and continue the submission process, please visit this URL:<br>https://platform.syndica.net/verify?e=%s&q=%s<br><br>After you submit your startup information, we will post your startup onto the Syndica platform.<br>If your startup receives responses from our MBA network, we will provide the responses to you via email.<br><br>Thank you again for your interest.<br><br>Peter | Founder<br>peter@syndica.net", email, hash)
	sendToEmail("", email, subject, textbody, htmlbody)
}

func SendPeterEmail(subject, body string) {
	sendToEmail("Syndica", "peter@syndica.net", subject, body, body)
}
