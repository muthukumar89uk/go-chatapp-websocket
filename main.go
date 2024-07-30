package main

import (
	"chatApp/drivers"
	"chatApp/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	drivers.DbConnection()

	drivers.Migration()

	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	port := viper.GetString(`listen.port`)

	router := mux.NewRouter()

	router.HandleFunc("/user", handlers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/ws", handlers.ChatHandler)

	log.Fatal(http.ListenAndServe(port, router))
}
