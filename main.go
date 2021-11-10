package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"webserver/webRoutes"
)

const DiscordToken string = "OTA4MDkyNTE4OTg4NjQ0Mzkz.YYwtEA.at7ZiRqox9msQl0doDCMVyYS6Pc"

func main() {
	// Setup Discord client
	client := setupDiscord()
	// Setup WebServer
	go setupWeb()

	// Graceful shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT, os.Interrupt, os.Kill)
	<-sc
	_ = client.Close()
}

func setupDiscord() *discordgo.Session {
	client, err := discordgo.New("Bot " + DiscordToken)

	if err != nil {
		log.Fatalf("Invalid token: %v", err)
	}

	// Add event handlers
	client.AddHandler(onReady)
	client.AddHandler(onMessage)

	// Connect to Discord
	if err = client.Open(); err != nil {
		log.Fatalf("Could not connect to discord %v", err)
	}

	log.Printf("Connected")

	return client
}
func setupWeb() {
	r := mux.NewRouter()
	r.HandleFunc("/", webRoutes.Index)
	// We can get {name} later when specifying it her
	r.HandleFunc("/hello/{name}", webRoutes.HelloName)
	r.HandleFunc("/discord", webRoutes.Discord)

	log.Fatal(http.ListenAndServe(":8080", r))
}
