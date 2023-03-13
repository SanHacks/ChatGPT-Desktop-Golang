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
	//Uncomment to create DB on startup
	//err := createDatabase()
	//if err != nil {
	//	log.Printf("Error creating database: %v", err)
	//}
	myApp := app.New()
	platform := myApp.NewWindow("Sage Chatbot")
	myApp.Settings().SetTheme(theme.DarkTheme())
	platform.Resize(fyne.Size{
		Width:  500,
		Height: 500,
	})
	platform.SetFixedSize(false)
	platform.MainMenu()

	platform.SetMaster()
	platform.SetOnClosed(func() {
		fmt.Println("Closed")
	})
	messageBox := container.NewVBox()
	messageBox.Resize(fyne.Size{
		Width:  0,
		Height: 0,
	})

	inputBox := widget.NewMultiLineEntry()
	inputBox.Wrapping = fyne.TextWrapWord
	inputBox.PlaceHolder = "Enter your message here..."

	// Add chat bubbles to the message box
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

	// Create a send button for sending messages
	sendButton := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		message := inputBox.Text
		//Increase width of input box
		fmt.Println(message)
		if message != "" {
			// Send message
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

	// Create a horizontal box for the input field and send button
	inputBoxContainer := container.NewVSplit(inputBox, sendButton)
	inputBoxContainer.Size()
	// Create a vertical box for the message box and input field/send button
	content := container.NewVBox(messageBox, layout.NewSpacer(), inputBoxContainer)

	// Set the content of the window and show it
	platform.SetContent(content)
	platform.ShowAndRun()
}

func addChatBubble(box *fyne.Container, message string, isUser bool) {
	// Create a new label with the message
	label := widget.NewLabel(message)
	// Create a new chat bubble with the label
	bubble := container.NewHBox(label)
	// Add the chat bubble to the message box
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
