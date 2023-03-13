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
		//Display Conversation with the bot
		displayConvo(message, tab1, inputBox)
	})
	return sendButton
}

func displayConvo(message string, tab1 *fyne.Container, inputBox *widget.Entry) {
	if message != "" {

		userMessages(message, tab1)
		addMessages := addMessage("YOU", message)
		if addMessages != nil {
			log.Printf("Error adding user message: %v", addMessage)
		}

		// Clear input box
		inputBox.SetText("")

		messageCall, err := makeApiCall()
		if err != nil {
			log.Printf("Error making API call: %v", err)
		}

		botMessages(messageCall, err, tab1)
		addMessage := addMessage("Bot", messageCall)
		if addMessage != nil {
			log.Printf("Error adding bot message: %v", addMessage)

		}

	}
}

// botMessages function to display messages from the bot
// This function is used to split messages into multiple chat bubbles if the message is too long
// This function is also used to send voice notes if the message is too long
func botMessages(messageCall string, err error, tab1 *fyne.Container) {
	if len(messageCall) > 90 {
		//Send voice note if message is more than 120 characters
		if len(messageCall) > 90 {
			voiceNote(messageCall, err)
		}
		var messageArray []string
		for i := 0; i < len(messageCall); i += 90 {
			end := i + 90
			if end > len(messageCall) {
				end = len(messageCall)
			}
			messageArray = append(messageArray, messageCall[i:end])
		}
		for _, message := range messageArray {
			addChatBubble(tab1, "Bot: "+message, true)
		}
	} else {
		addChatBubble(tab1, "Bot: "+messageCall, true)
	}
}

func userMessages(message string, tab1 *fyne.Container) {
	if len(message) > 90 {
		var messageArray []string
		for i := 0; i < len(message); i += 90 {
			end := i + 90
			if end > len(message) {
				end = len(message)
			}
			messageArray = append(messageArray, message[i:end])
		}
		for _, message := range messageArray {
			addChatBubble(tab1, "YOU: "+message, false)
		}
	} else {
		addChatBubble(tab1, "YOU: "+message, false)
	}
}
