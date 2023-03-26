package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gastrodon/popplio/ifunny"
)

func main() {
	bearer := os.Getenv("IFUNNY_BEARER")
	if bearer == "" {
		panic("IFUNNY_BEARER must be set")
	}

	userAgent := os.Getenv("IFUNNY_USER_AGENT")
	if userAgent == "" {
		panic("IFUNNY_USER_AGENT must be set")
	}

	client := ifunny.MakeClient("bearer "+bearer, userAgent)
	client.Request("GET", "/v4", nil)

	chat, err := client.Connect(bearer)
	if err != nil {
		panic(err)
	}

	<-time.After(4 * time.Second)

	userID := "641f57b56f4a823e897e6f36"
	wsInvite := chat.Invites(userID)
	wsChats := chat.Chats(userID)

	for {
		select {
		case invite := <-wsInvite:
			fmt.Printf("invite: %+v\n", invite)
		case chats := <-wsChats:
			fmt.Printf("chats: %+v\n", chats)
		}
	}
}
