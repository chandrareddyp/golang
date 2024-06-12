package main

import (
	"fmt"
	"math/rand"
)


func main() {
	rand.Seed(0)
	test(
		[]string{
			"hi friend",
			"What's going on?",
			"Welcome to the business",
			"I'll pay you to be my friend",
		},
		[]string{
			"Will you make your appointment?",
			"Let's be friends",
			"What are you doing?",
			"I can't believe you've done this.",
		},
	)
 
}

func logMessages(chEmails, chSms chan string) {
	x, y := false, false
 	for{
		select{
			case  email, ok := <- chEmails:
				if !ok{
					x = true
				}else{
					logEmail(email)
				}
				 
			case sms, ok := <- chSms:
				if !ok{
					y = true
				}else{
					logSms(sms)
				}
				
		}
		if x && y	{
			return
		}
	}
}

// TEST SUITE - Don't touch below this line

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

func test(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chEmails, chSms)

	fmt.Println("===============================")
}


func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func(){
		defer close(chSms)
		for _, s:= range sms{
			chSms <- s
		}
	}()
	go func(){
		defer close(chEmails)
		for _, e:= range emails{
			chEmails <- e
		}
	}()
	return chSms, chEmails
}
