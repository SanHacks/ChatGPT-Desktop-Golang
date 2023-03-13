package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func sendButton(inputBox *widget.Entry, tab1 *fyne.Container) *widget.Button {
	// Create a send button for sending messages
	sendButton := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		message := inputBox.Text
		//Increase width of input box
		fmt.Println(message)
		if message != "" {
			// Send message
			addUserMessage := addMessage("YOU", message)
			addChatBubble(tab1, "YOU: "+message, false)
			// Clear input box
			inputBox.SetText("")
			// Receive message
			messageCall := makeApiCall()

			addChatBubble(tab1, "Bot: "+messageCall, true)
			addMessage := addMessage("Bot", messageCall)
			if addUserMessage != nil {
				log.Printf("Error adding user message: %v", addUserMessage)
			}
			if addMessage != nil {
				log.Printf("Error adding bot message: %v", addMessage)

			}

		}
	})
	return sendButton
}
