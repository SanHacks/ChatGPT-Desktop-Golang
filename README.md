# ChatGPT-Desktop-Golang
Simple Attempt To Have Chat With API Enpoints Via Chat, Desktop

At this point it is only talking to the Ron Swanson, Quotes API.

You can change endpoints by modifying the `apiCall` functions

<img width="1440" alt="Screenshot 2023-03-13 at 17 06 21" src="https://user-images.githubusercontent.com/13138647/224782941-fbbff41b-3bb3-4044-a9c3-60f0a58a9a16.png">


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
- Golang
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

