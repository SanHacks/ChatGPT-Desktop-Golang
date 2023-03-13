# ChatGPT-Desktop-Golang
Simple Attempt To Have Chat With API Endpoints Via Chat, Desktop

At this point it is only talking to the Ron Swanson, Quotes API.

You can change endpoints by modifying the `apiCall` functions

<img width="1440" alt="Screenshot 2023-03-13 at 17 06 21" src="https://user-images.githubusercontent.com/13138647/224782941-fbbff41b-3bb3-4044-a9c3-60f0a58a9a16.png">

This tool is made as an experiment and is not intended for production use. 
## Features
- [x] Chat with API
- [x] Save Responses to SQLite3 Database
- [ ] Save Responses to JSON File
- [ ] Save Responses to CSV File
- [ ] Save Responses to XML File
- [ ] Save Response and Convert to Audio
- [ ] Show images in chat
- [ ] Get image from API

## How To Install
### Pre-reqs
- Golang </br>
`brew install go`</br>
This will install the latest version of Golang on your machine </br>
`go version`</br>
This will show you the version of Golang installed on your machine </br>
- SQLite3

## How To Run

`go run main.go`
or simply
`go build .`

## How To Use

- Type in the text box
- Press Enter
- The API will be called
- The response will be displayed in the chat box
- The response is stored in SQLite3 Database

## How To Build
You can build the app for your OS by running the following command in the root directory of the project
first install `fyne-cross` </br>
`go get -u fyne.io/fyne/cmd/fyne-cross`</br>
fyne-cross <command> [options]

The commands are:

	darwin        Build and package a fyne application for the darwin OS
	linux         Build and package a fyne application for the linux OS
	windows       Build and package a fyne application for the windows OS
	android       Build and package a fyne application for the android OS
	ios           Build and package a fyne application for the iOS OS
	freebsd       Build and package a fyne application for the freebsd OS
	version       Print the fyne-cross version information

Use "fyne-cross <command> -help" for more information about a command.
