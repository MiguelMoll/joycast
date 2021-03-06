package minion

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"

	"github.com/MiguelMoll/joycast/realm"
)

type youtubeOption func(y *youtubeHandler) error

type youtubeHandler struct {
	users *realm.UserService
}

func YoutubeHandler(opts ...youtubeOption) (*youtubeHandler, error) {
	y := &youtubeHandler{}

	for _, opt := range opts {
		if err := opt(y); err != nil {
			return nil, err
		}
	}
	return y, nil
}

func (y *youtubeHandler) Handle(msg string) {
	fmt.Println(msg)

	service, err := y.newClient()
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       "Test Title",
			Description: "Test Description", // can not use non-alpha-numeric characters
			CategoryId:  "22",
			Tags:        []string{"test", "upload", "api"},
		},
		Status: &youtube.VideoStatus{PrivacyStatus: "unlisted"},
	}

	call := service.Videos.Insert("snippet,status", upload)

	filename := "ocean.mp4"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening %v: %v", filename, err)
	}
	defer file.Close()

	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
	}
	fmt.Printf("%#v\n", response)
	// TODO: Double check if upload is successful? Uploading a duplicate video
	// returns successful via API but youtube UI shows error on duplicate videos
	// Double check this assumption
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
}

func YoutubeUsers(u *realm.UserService) youtubeOption {
	return func(y *youtubeHandler) error {
		if u == nil {
			return errors.New("user service not initialized")
		}
		y.users = u
		return nil
	}
}

func (y *youtubeHandler) newClient() (*youtube.Service, error) {
	user, err := y.users.Get(1)

	config, err := newConfig()
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background(), user.OauthToken)
	return youtube.New(client)
}

func newConfig() (*oauth2.Config, error) {
	secret := []byte(`{"web":{"client_id":"1003910230744-gqvdgs38d2kvllslba3jvddsgvo5pckq.apps.googleusercontent.com","project_id":"projectcast","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"4WqhzHGa9_eJnVovcHQsS_CY","redirect_uris":["http://localhost:8000/youtube/authorized"],"javascript_origins":["http://www.sire.ninja"]}}`)
	return google.ConfigFromJSON(secret, youtube.YoutubeReadonlyScope, youtube.YoutubeUploadScope)
}
