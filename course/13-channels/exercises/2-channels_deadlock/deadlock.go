package main

import "fmt"


func main() {
	test([]email{
		{ 
			header:"hi",
			body:"how are you doing",
		},
		{
			header:"hello",
			body:"what are you doing",
		},
	})
}

type email struct {
	header string
	body string
}

func test(emails []email) {
	sendEmail(emails)
}

func sendEmail(emails []email) {
	body := make(chan string)
	for i, _ := range emails {
		fmt.Println("Sending email: ", emails[i].header)
		go func(header string) {
			fmt.Println("Email sent: ",header)
			body <- emails[i].body
		}( emails[i].header)
	}
	for range emails {
		fmt.Println("Email received: ", <-body)
	}
	/* Deadlock code:::
	fmt.Println("Email received: ", <-body)
	fmt.Println("Email received: ", <-body)
	fmt.Println("Email received: ", <-body)
	*/
	 
}

 