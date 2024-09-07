package spotify

import (
	"context"
	"fmt"
)

func getPlaylist() {
	spotifyClient := GetClient()
	ctx := context.Background()
	playlist, err := spotifyClient.GetPlaylistItems(ctx, "id")
	if err != nil {
		fmt.Println("Error getting user playlist")
	}
	fmt.Println(playlist)
}
