package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
)

type Watch struct {
	CountryCode string `json:"CountryCode"`
}

type NotificationTarget struct {
	// discord, gotify, curl
	Type string `json:"type"`

	// discord options
	DiscordURL string `json:"discordUrl"`

	// gotify options
	GotifyURL   string `json:"gotifyUrl"`
	GotifyToken string `json:"gotifyToken"`
	GotifyAppID uint   `json:"gotifyAppID"`
}

type Config struct {
	Watches     map[string]Watch     `json:"watches"`
	Targets     []NotificationTarget `json:"targets"`
	APIKey      string               `json:"apiKey"`
	CheckOnBoot bool                 `json:"checkOnBoot"`
}

var config Config

var dataPath string

var listPath string = ""

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
			Watches: map[string]Watch{},
			Targets: []NotificationTarget{
				{
					Type:       "discord",
					DiscordURL: "put your webhook url here",
				},
				{
					Type:        "gotify",
					GotifyURL:   "https://gotify.awesomeperson.com",
					GotifyToken: "gotify token",
					// index of your gotify app, btw
					GotifyAppID: 0,
				},
			},
			APIKey:      "ticketmaster api key here",
			CheckOnBoot: true,
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

func verifyListExists(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		os.WriteFile(path, []byte(""), 0755)
	}
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

	verifyListExists(listPath)
	fmt.Println("Successfully booted and loaded config :3")

	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0,30 * * * *", checkWatches)
	c.Start()

	if config.CheckOnBoot {
		checkWatches()
	}

	select {}
}
