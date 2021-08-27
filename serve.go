package main

import (
	"testapp/config"
	"testapp/fileServer"
)

func main() {
	config := config.LoadConfig()
	/*
	fmt.Printf("Listening on port %p",onionConfig.GetInt("server.port") )
	http.Handle("/", http.FileServer(http.Dir("/home")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", onionConfig.GetInt("server.port")), nil)
	 */
	fileServer.StartServer(config)
}
