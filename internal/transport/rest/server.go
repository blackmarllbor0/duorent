package rest

import (
	"fmt"
	"net/http"
	"time"
)

func RunNewHTTPServer(port uint, routes http.Handler) error {
	server := http.Server{
		Addr:           fmt.Sprintf(":%v", port),
		Handler:        routes,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	return server.ListenAndServe()
}
