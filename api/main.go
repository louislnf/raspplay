package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/louislnf/raspplay/piplayer"

	"github.com/gorilla/mux"
)

var piPlayer = piplayer.CreatePiPlayer()

func main() {
	piPlayer.SetMediaSource("/home/pi/Videos/sample_video.avi")

	router := mux.NewRouter()
	router.HandleFunc("/player/command/play", PutPiPlayerPlay).Methods("PUT")
	router.HandleFunc("/player/command/pause", PutPiPlayerPause).Methods("PUT")
	router.HandleFunc("/player/command/quit", PutPiPlayerQuit).Methods("PUT")
	router.HandleFunc("/player/status", GetPiPlayerStatus).Methods("GET")
	router.HandleFunc("/player/media", SetPiPlayerMediaSource).Methods("POST")
	router.HandleFunc("/player/media", GetPiPlayerMediaSources).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func PutPiPlayerPlay(w http.ResponseWriter, r *http.Request) {
	piPlayer.Play()
	fmt.Fprintf(w, "Playing...\n")
}

func PutPiPlayerPause(w http.ResponseWriter, r *http.Request) {
	piPlayer.Pause()
	fmt.Fprintf(w, "Paused...\n")
}

func PutPiPlayerQuit(w http.ResponseWriter, r *http.Request) {
	piPlayer.Quit()
	fmt.Fprintf(w, "Terminated the player\n")
}

func GetPiPlayerStatus(w http.ResponseWriter, r *http.Request) {

}

func SetPiPlayerMediaSource(w http.ResponseWriter, r *http.Request) {
	piPlayer.Quit()
	piPlayer.SetMediaSource(r.PostFormValue("source"))
	piPlayer.Play()
	fmt.Fprintf(w, fmt.Sprintf("Set media source to %s\n", r.PostFormValue("source")))
}

func GetPiPlayerMediaSources(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
