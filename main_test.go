package main

import (
	"fyne.io/fyne/v2"
	"reflect"
	"testing"
)

func Test_addChatBubble(t *testing.T) {
	type args struct {
		box     *fyne.Container
		message string
		isUser  bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addChatBubble(tt.args.box, tt.args.message, tt.args.isUser)
		})
	}
}

func Test_addMessage(t *testing.T) {
	type args struct {
		sender  string
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addMessage(tt.args.sender, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("addMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createDatabase(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("createDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getMessages(t *testing.T) {
	tests := []struct {
		name    string
		want    []Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMessages()
			if (err != nil) != tt.wantErr {
				t.Errorf("getMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMessages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeApiCall(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeApiCall(); got != tt.want {
				t.Errorf("makeApiCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
