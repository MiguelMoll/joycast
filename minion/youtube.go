package minion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

type youtubeHandler struct{}

func YoutubeHandler() *youtubeHandler {
	return nil
}

func (y *youtubeHandler) Handle(msg string) {
	fmt.Println(msg)

	service, err := newClient()
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
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
}

func newClient() (*youtube.Service, error) {
	oauth := []byte(`{"expiry": "2019-02-28T21:37:12.260700309-05:00", "token_type": "Bearer", "access_token": "ya29.Glu_BqRabMrE9ZxvnKarqr27XFM8L8uEYEh2w42vdUDET1nxQqmhCY4r1_WUMNuNG4UyfEXbeqXT7QZDaNElkMulXcr1Ik-uQ7a-crGWQ06KxAo6hJC7GUpDSUiJ", "refresh_token": "1/AtDMHsIN7q3yBfo4PEMe5uo6mH3hkl2WjKmub2H1t8OBOHW3FaS9vglVQSuVbRAk"}`)

	var token oauth2.Token
	if err := json.Unmarshal(oauth, &token); err != nil {
		return nil, err
	}
	config, err := newConfig()
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background(), &token)
	return youtube.New(client)
}

func newConfig() (*oauth2.Config, error) {
	secret := []byte(`{"web":{"client_id":"1003910230744-gqvdgs38d2kvllslba3jvddsgvo5pckq.apps.googleusercontent.com","project_id":"projectcast","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"4WqhzHGa9_eJnVovcHQsS_CY","redirect_uris":["http://localhost:8000/youtube/authorized"],"javascript_origins":["http://www.sire.ninja"]}}`)
	return google.ConfigFromJSON(secret, youtube.YoutubeReadonlyScope, youtube.YoutubeUploadScope)
}
