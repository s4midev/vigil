package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gotify/go-api-client/auth"
	"github.com/gotify/go-api-client/client/message"
	"github.com/gotify/go-api-client/gotify"
	"github.com/gotify/server/model"
	"github.com/typical-developers/discord-webhooks-go/webhooks"
)

func isNotified(id string) bool {
	fileData, _ := os.ReadFile(listPath)

	return strings.Contains(string(fileData), id)
}

func notify(id string) {
	fileData, _ := os.ReadFile(listPath)
	os.WriteFile(listPath, []byte(string(fileData)+"\n"+id), 0755)
}

func createTitleString(mainAttraction Attraction) string {
	return "There is a new `" + mainAttraction.Name + "` event!"
}

func createBodyString(venue Venue) string {
	return "It will be at `" + venue.Name + "` in " + venue.City.Name
}

func getMainAttraction(event Event) Attraction {
	var mainAttraction Attraction

	for _, a := range event.Embedded.Attractions {
		if _, exists := config.Watches[a.ID]; exists {
			mainAttraction = a
			break
		}
	}

	return mainAttraction
}

func discordNotify(event Event, target NotificationTarget) {
	webhook := webhooks.NewWebhookClientFromURL(target.DiscordURL)

	payload := webhooks.WebhookPayload{}

	embed := payload.AddEmbed()

	mainAttraction := getMainAttraction(event)

	embed.SetTitle(createTitleString(mainAttraction))
	embed.SetDescription(createBodyString(event.Embedded.Venues[0]))

	pricefield := embed.AddField()

	pricefield.SetName("Price/PP")

	// TODO: implement this
	pricefield.SetValue("money")
	pricefield.SetInline(true)

	embed.SetURL(event.URL)
	embed.SetImage(event.Embedded.Attractions[0].Images[len(event.Embedded.Attractions[0].Images)-1].URL)

	webhook.SendMessage(&payload)

	notify(event.ID)
}

func gotifyNotify(event Event, target NotificationTarget) {
	url, _ := url.Parse(target.GotifyURL)
	client := gotify.NewClient(url, &http.Client{})

	params := message.NewCreateMessageParams()

	mainAttraction := getMainAttraction(event)

	params.Body = &model.Message{
		Title:         createTitleString(mainAttraction),
		Message:       createBodyString(event.Embedded.Venues[0]),
		ApplicationID: target.GotifyAppID,
	}

	_, err := client.Message.CreateMessage(params, auth.TokenAuth(target.GotifyToken))

	if err != nil {
		log.Fatalf("Could not send message %v", err)
		return
	}
}

func notifyUser(event Event) {
	for i, t := range config.Targets {
		switch t.Type {
		case "discord":
			discordNotify(event, t)

		case "gotify":
			gotifyNotify(event, t)

		default:
			fmt.Println("Target " + strconv.Itoa(i) + "has an invalid type!")
		}
	}
}
