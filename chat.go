package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Create a new label with the message and add it to the chat window
// The isUser parameter determines if the message is from the user or the bot
// If the message is from the user, the bubble will be on the right side of the chat window
// If the message is from the bot, the bubble will be on the left side of the chat window
// Create a new label with the message and add it to the chat window
// The isUser parameter determines if the message is from the user or the bot
// If the message is from the user, the bubble will be on the right side of the chat window
// If the message is from the bot, the bubble will be on the left side of the chat window
func addChatBubble(box *fyne.Container, message string, isUser bool) {

	// Create a new label with the message
	label := widget.NewLabel(message)
	label.TextStyle = fyne.TextStyle{Bold: false, Italic: false, Monospace: false}

	// Create a new chat bubble with the label
	bubble := container.NewHBox(label)
	container.NewScroll(bubble)

	// Create a new image widget with the avatar URL
	avatarImg := canvas.NewImageFromFile("source/avatar.jpg")
	avatarImg.SetMinSize(fyne.NewSize(64, 64))

	botAvatarImg := canvas.NewImageFromFile("source/botAvatar.png")
	botAvatarImg.SetMinSize(fyne.NewSize(64, 64))

	// Add the chat bubble to the card
	if isUser {
		// If the message is from the user, add the bubble to the right side of the card
		box.Add(container.NewHBox(
			layout.NewSpacer(),
			widget.NewCard("", "", bubble),
			botAvatarImg,
		))
	} else {
		// If the message is from someone else, add the bubble to the left side of the card
		box.Add(container.NewHBox(
			avatarImg,
			widget.NewCard("", "", bubble),
			layout.NewSpacer(),
		))
	}

	// Create a new audio player widget with the audio file URL
	//audioPlayer := widget.Audio("source/voice.mp3")
	//audioPlayer := container.NewAudioPlayer()
	//audioPlayer.SetPlayer(func() fyne.Resource {
	//	file, err := fyne.LoadResourceFromPath("source/voice.mp3")
	//	if err != nil {
	//		fmt.Println("Error loading audio file:", err)
	//		return nil
	//	}
	//	return file
	//})
	//audioPlayer.SetAutoplay(true)

	// Add the audio player to the chat window
	//box.Add(container.NewHBox(
	//	layout.NewSpacer(),
	//	audioPlayer,
	//))

	// Wrap the container with a ScrollContainer to enable scrolling
	container.NewScroll(box)
}
