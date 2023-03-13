package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
	myApp.Settings().PrimaryColor()
	myApp.SendNotification(&fyne.Notification{
		Title:   "Sage Chatbot",
		Content: "Welcome to Sage Chatbot",
	})
	myApp.SetIcon(theme.MailAttachmentIcon())

	tab1 := container.NewVBox()
	tab1.Resize(fyne.Size{
		Width:  0,
		Height: 0,
	})

	inputBox := widget.NewMultiLineEntry()
	inputBox.Wrapping = fyne.TextWrapWord
	inputBox.PlaceHolder = "Enter your message here..."

	// Add chat bubbles to the message box
	messages1, err := getMessages()
	if err != nil {
		log.Printf("Error getting messages: %v", err)
	}

	for _, message := range messages1 {
		addChatBubble(tab1, message.Sender+": "+message.Content, message.Sender == "Bot")
	}

	messageCall := makeApiCall()
	addChatBubble(tab1, "YOU: I am looking for a quote", false)
	addChatBubble(tab1, "Bot: "+messageCall, true)

	// Create a send button for sending messages
	sendButton := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		message := inputBox.Text
		//Increase width of input box
		fmt.Println(message)
		if message != "" {
			// Send message
			addUserMessage := addMessage("YOU", message)
			addChatBubble(tab1, "YOU: "+message, false)
			inputBox.SetText("")
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

	// Create a horizontal box for the input field and send button
	inputBoxContainer := container.NewVSplit(inputBox, sendButton)
	inputBoxContainer.Size()
	inputBoxContainer.Resize(fyne.Size{
		Width:  300,
		Height: 300,
	})

	// Create the chat list
	chatList := widget.NewList(
		func() int {
			return len(messages1)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(messages1[i].Sender + ": " + messages1[i].Content)

		})

	//On click of chat bubble in chat list open chat bubble in chat window
	chatList.OnSelected = func(id widget.ListItemID) {
		myApp.SendNotification(&fyne.Notification{
			Title:   "Sage Chatbot",
			Content: "You clicked on a chat bubble",
		})
		//append message to chat window
		addChatBubble(tab1, messages1[id].Sender+": "+messages1[id].Content, messages1[id].Sender == "Bot")
	}

	// Add the chat list to the tab2 container
	tab2 := container.NewVBox(chatList)
	tab2.Resize(fyne.Size{
		Width:  0,
		Height: 0,
	})
	// create a TabContainer
	tabs := container.NewAppTabs(
		container.NewTabItem("Home", tab1),
		container.NewTabItem("Media", tab2),
	)

	// Set the content of the window and show it
	content := container.NewVBox(tabs)
	platform := myApp.NewWindow("Sage Chatbot")
	platform.SetContent(content)
	platform.SetFixedSize(false)
	platform.MainMenu()
	platform.SetMaster()
	platform.SetOnClosed(func() {
		fmt.Println("Closed")
	})
	//add side menu
	platform.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				fmt.Println("New")
			}),
			fyne.NewMenuItem("Open", func() {
				fmt.Println("Open")
			}),
			fyne.NewMenuItem("Save", func() {
				fmt.Println("Save")
			}),
			fyne.NewMenuItem("Save As", func() {
				fmt.Println("Save As")

			}),
			fyne.NewMenuItem("Quit", func() {
				fmt.Println("Quit")
				platform.Close()
			}),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Cut", func() {
				fmt.Println("Cut")
				if inputBox.SelectedText() != "" {
					platform.Clipboard().SetContent(inputBox.SelectedText()) // copy to clipboard
					inputBox.SetText(inputBox.Text[:inputBox.CursorColumn] + inputBox.Text[inputBox.CursorColumn+len(inputBox.SelectedText()):])
				}

			}),
			fyne.NewMenuItem("Copy", func() {
				fmt.Println("Copy")
				if inputBox.SelectedText() != "" {
					platform.Clipboard().SetContent(inputBox.SelectedText()) // copy to clipboard
				}

			}),
			fyne.NewMenuItem("Paste", func() {
				fmt.Println("Paste")
				if platform.Clipboard().Content() != "" {
					inputBox.SetText(inputBox.Text[:inputBox.CursorColumn] + platform.Clipboard().Content() + inputBox.Text[inputBox.CursorColumn:])
				}

			}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				//Send notification
				myApp.SendNotification(&fyne.Notification{
					Title:   "Sage Chatbot",
					Content: "This is a chatbot that can help you get started with Chatbots",
				})

			}),
		),
	))

	myApp.Settings().SetTheme(theme.DarkTheme())
	platform.Resize(fyne.Size{
		Width:  500,
		Height: 500,
	})
	// Create the side menu
	sideMenu := container.NewVBox(
		widget.NewButton("Sign In", func() {
			fmt.Println("Sign In")
		}),
		widget.NewButton("Sign Up", func() {
			fmt.Println("Sign Up")

		}),
		widget.NewButton("Sign Out", func() {
			fmt.Println("Sign Out")

		}),
		widget.NewButton("Settings", func() {
			fmt.Println("Settings")

		}),
	)

	//add input box
	platform.SetContent(container.NewBorder(nil, inputBoxContainer, sideMenu, nil, content))

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
