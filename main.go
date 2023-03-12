package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//err := createDatabase()
	//if err != nil {
	//	log.Printf("Error creating database: %v", err)
	//}
	myApp := app.New()
	myWindow := myApp.NewWindow("Sage Chatbot")
	myWindow.Resize(fyne.NewSize(1200, 1200))

	// create a scrollable chat bubble for displaying messages
	messageBox := container.NewVBox()

	// create a text input field for users to enter their messages
	inputBox := widget.NewMultiLineEntry()
	inputBox.Wrapping = fyne.TextWrapWord
	inputBox.PlaceHolder = "Enter your message here..."

	// add chat bubbles to the message box
	messages, err := getMessages()
	if err != nil {
		log.Printf("Error getting messages: %v", err)
	}

	for _, message := range messages {
		addChatBubble(messageBox, message.Sender+": "+message.Content, message.Sender == "Bot")
	}

	messageCall := makeApiCall()
	addChatBubble(messageBox, "YOU: I am looking for a quote", false)
	addChatBubble(messageBox, "Bot: "+messageCall, true)

	// create a send button for sending messages
	sendButton := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		message := inputBox.Text
		//increase width of input box
		fmt.Println(message)
		if message != "" {
			// send message
			addUserMessage := addMessage("YOU", message)
			addChatBubble(messageBox, "YOU: "+message, false)
			inputBox.SetText("")
			messageCall := makeApiCall()
			addChatBubble(messageBox, "Bot: "+messageCall, true)
			addMessage := addMessage("Bot", messageCall)
			if addUserMessage != nil {
				log.Printf("Error adding user message: %v", addUserMessage)
			}
			if addMessage != nil {
				log.Printf("Error adding bot message: %v", addMessage)

			}

		}
	})

	// create a horizontal box for the input field and send button
	inputBoxContainer := container.NewVSplit(inputBox, sendButton)
	inputBoxContainer.Size()
	// create a vertical box for the message box and input field/send button
	content := container.NewVBox(messageBox, layout.NewSpacer(), inputBoxContainer)

	// set the content of the window and show it
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}

func addChatBubble(box *fyne.Container, message string, isUser bool) {
	// create a new label with the message
	label := widget.NewLabel(message)

	// create a new chat bubble with the label
	bubble := container.NewHBox(label)
	bubble.Resize(fyne.NewSize(1000, 500))

	// add the chat bubble to the message box
	box.Add(bubble)
}

func makeApiCall() string {

	url := "https://ron-swanson-quotes.herokuapp.com/v2/quotes"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body[1 : len(body)-1])

}
