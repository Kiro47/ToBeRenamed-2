package fileServer

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"testapp/config"
)

var Index string = "<html>\nIndex:\n<pre>\n"
var IndexData []string

func writeIndex(writer http.ResponseWriter, request *http.Request) {
	var response = ""
	if request.Header.Get("Content-type") == "application/json" {
		writer.Header().Set("Content-type", "application/json")
		response = "{\"data\":["
		for _, dir := range IndexData {
			// Will do this differently once moved to object based with marshalling
			response += fmt.Sprintf("{\"name\": \"%s\", \"size\": 0, \"type\": \"href\"},", dir)
		}
		// Strip extra comma
		response = strings.TrimRight(response, ",")
		response += "]}"
	} else {
		writer.Header().Set("Content-type", "text/html; charset=utf-8")
		response = "<html><body>\nIndex:\n<pre>\n"
		for _, dir := range IndexData {
			response += fmt.Sprintf("<a href=\"%s\">%s</a>\n", dir, dir)
		}
		response += "</pre>\n</body></html>"
	}
	_, err := writer.Write([]byte(response))
	if err != nil {
		// Can't really do much, do debug logging later
		return
	}
}

func buildRouter(config config.Config) *mux.Router{
	router := mux.NewRouter()
	for dir := range config.Directories {
		relative := config.Directories[dir].UrlPath
		file := config.Directories[dir].FilePath
		// Setup fileservers
		handler := http.StripPrefix(relative, http.FileServer(http.Dir(file)))
		router.PathPrefix(relative).Handler(handler)
		// Debug log
		fmt.Printf("Serving %s at %s\n", file, relative)
		// Append to hacky directory index
		IndexData = append(IndexData, relative)
	}
	return router
}

func StartServer(config config.Config) {
	router := buildRouter(config)

	router.HandleFunc("/", writeIndex)
	httpServer := &http.Server{
		Handler: router,
		Addr: fmt.Sprintf("%s:%d", config.Server.ListenAddr, config.Server.Port),
	}
	log.Fatal(httpServer.ListenAndServe())
}