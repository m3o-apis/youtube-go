package example

import (
	"fmt"
	"os"

	"github.com/micro/services/clients/go/youtube"
)

// Search for videos on YouTube
func SearchForVideos() {
	youtubeService := youtube.NewYoutubeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := youtubeService.Search(&youtube.SearchRequest{
		Query: "donuts",
	})
	fmt.Println(rsp, err)
}
