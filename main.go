package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/robfig/cron/v3"
	"github.com/typical-developers/discord-webhooks-go/webhooks"
)

type Watch struct {
	CountryCode string `json:"CountryCode"`
}

type Config struct {
	Watches    map[string]Watch `json:"watches"`
	WebhookURL string           `json:"webhookURL"`
	APIKey     string           `json:"apiKey"`
}

var config Config

var dataPath string

var listPath string = ""

func isNotified(id string) bool {
	fileData, _ := os.ReadFile(listPath)

	return strings.Contains(string(fileData), id)
}

func notify(id string) {
	fileData, _ := os.ReadFile(listPath)
	os.WriteFile(listPath, []byte(string(fileData)+"\n"+id), 0755)
}

func notifyUser(event Event) {
	webhook := webhooks.NewWebhookClientFromURL(config.WebhookURL)

	payload := webhooks.WebhookPayload{}

	embed := payload.AddEmbed()

	var mainAttraction Attraction

	for _, a := range event.Embedded.Attractions {
		if _, exists := config.Watches[a.ID]; exists {
			mainAttraction = a
			break
		}
	}

	embed.SetTitle("There is a new `" + mainAttraction.Name + "` event!")
	embed.SetDescription("It will be at `" + event.Embedded.Venues[0].Name + "` in " + event.Embedded.Venues[0].City.Name)

	pricefield := embed.AddField()

	pricefield.SetName("Price/PP")
	pricefield.SetValue("money")
	pricefield.SetInline(true)

	embed.SetURL(event.URL)
	embed.SetImage(event.Embedded.Attractions[0].Images[len(event.Embedded.Attractions[0].Images)-1].URL)

	webhook.SendMessage(&payload)

	notify(event.ID)
}

func checkWatches() {
	fmt.Println("Checking watches")
	for i, w := range config.Watches {
		endpoint := "https://app.ticketmaster.com/discovery/v2/events?apikey=" + config.APIKey + "&attractionId=" + i + "&CountryCode=" + w.CountryCode

		resp, err := http.Get(endpoint)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var response = APIResponse{}

		json.Unmarshal(body, &response)

		if len(response.Embedded.Events) == 0 {
			fmt.Println("No new events for artist ID " + i + " in " + w.CountryCode)
			continue
		}

		for _, e := range response.Embedded.Events {
			if isNotified(e.ID) {
				continue
			}

			fmt.Println(e.Embedded.Venues[0].Name)
			notifyUser(e)
			fmt.Println(e.ID)
		}
	}
}

// this assumes the path exists, as checked in main
func loadConfig(path string) error {
	configPath := path + "/config.json"

	_, err := os.Stat(configPath)

	if os.IsNotExist(err) {
		defaultValue := Config{
			Watches:    map[string]Watch{},
			WebhookURL: "discord webhook url here",
			APIKey:     "ticketmaster api key here",
		}

		marshal, _ := json.Marshal(defaultValue)

		os.WriteFile(configPath, marshal, 0755)

		return errors.New("You do not have a config.json file. An example has been written, please edit it before re-running vigil")
	} else {
		fileData, _ := os.ReadFile(configPath)

		config = Config{}

		err := json.Unmarshal(fileData, &config)

		if err != nil {
			return errors.New("Your config.json file is invalid")
		}
	}

	return nil
}

func main() {
	dataPath = os.Getenv("VIGIL_DATAPATH")

	if dataPath == "" {
		fmt.Println("Please set the VIGIL_DATAPATH env var! :(")
		return
	}

	_, err := os.Stat(dataPath)

	if os.IsNotExist(err) {
		fmt.Println("VIGIL_DATAPATH env var is not a vaild path :(")
		return
	}

	err = loadConfig(dataPath)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	listPath = dataPath + "/notifiedList"

	fmt.Println("Successfully booted and loaded config :3")

	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0,30 * * * *", checkWatches)
	c.Start()

	select {}
}
