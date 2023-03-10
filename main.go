package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Chat App")

	// create a scrollable text area for displaying messages
	messageBox := widget.NewMultiLineEntry()
	messageBox.Wrapping = fyne.TextWrapWord
	messageBox.Resize(fyne.NewSize(1200, 1200))

	// create a text input field for users to enter their messages
	inputBox := widget.NewEntry()
	inputBox.Resize(fyne.NewSize(1200, 50))

	// create a button for sending messages
	sendButton := widget.NewButton("Prompt", func() {
		message := inputBox.Text
		// add the user's message to the message box
		messageBox.SetText(messageBox.Text + "You: " + message + "\n")
		// TODO: send the message to the chatbot and get a response
		// add the chatbot's response to the message box
		response := makeApiCall()

		messageBox.SetText(messageBox.Text + "Chatbot: " + response + "\n")
		// clear the input field
		inputBox.SetText("")
	})

	// create a horizontal box for the input field and send button
	inputContainer := container.NewHBox(inputBox, sendButton)
	inputContainer.Layout = layout.NewMaxLayout()

	// create a vertical box for the message box and input box/button
	chatContainer := container.NewVBox(
		messageBox,
		inputContainer,
	)
	chatContainer.Resize(fyne.NewSize(400, 500))

	// set the chat container as the content of the window and show it
	myWindow.SetContent(chatContainer)
	myWindow.ShowAndRun()
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
	return string(body)
}
