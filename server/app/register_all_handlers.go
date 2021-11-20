package app

import "net/http"

func RegisterAllHandlers() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/ping", Ping)
}
