package fileServer

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testapp/config"
)
func writeIndex(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Something"))
}

func StartServer(config config.Config) {

	// This could definitely be done better
	index := "<html>\n<pre>\n"

	router := mux.NewRouter()
	for dir := range config.Directories {
		relative := config.Directories[dir].UrlPath
		file := config.Directories[dir].FilePath
		// Setup fileserver
		handler := http.StripPrefix(relative, http.FileServer(http.Dir(file)))
		router.PathPrefix(relative).Handler(handler)
		// Debug log
		fmt.Printf("Serving %s at %s\n", file, relative)
		// Add to directoryList so we can build an index
		index += fmt.Sprintf("<a href=\"%s\">%s</a>\n", relative, relative)
	}
	index += "</pre>\n</html>"
	router.HandleFunc("/",
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "text/html")
			_, err := writer.Write([]byte(index))
			if err != nil {
				log.Fatalln("guess I'll die")
			}
		},
	)
	httpServer := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:5001",
	}
	log.Fatal(httpServer.ListenAndServe())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil))
	log.Fatalln("idk")
}