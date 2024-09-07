package main

import (
	"context"
	"fmt"
	"log"
	"playlist-transfer-go/spotify"
)

func main() {
	fmt.Println("Implement main starting point")
	client := spotify.GetClient()
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.User)
	fmt.Println(user.DisplayName)

}
