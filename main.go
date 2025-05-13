package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type Config struct {
	Mux *http.ServeMux
}

func (app *Config) InitRouter() {
	mux := app.Mux
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	mux.HandleFunc("/", app.HelloHandler)
}

func (app *Config) Run(port int) {
	/// Setting up Logging middleware
	handler := app.LoggingMiddleware(app.Mux)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
	fmt.Printf("Server listening on http://localhost%s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func main() {
	/// New App config ...my style of doing this =)
	app := Config{
		Mux: http.DefaultServeMux,
	}
	/// initialize router and add routes to internal default servemux
	app.InitRouter()
	/// run our app on port 3000
	app.Run(3000)
}

func (app *Config) HelloHandler(w http.ResponseWriter, r *http.Request) {

	posts := []BlogPost{
		{
			Title:       "Building a Blog with Go",
			Date:        "May 13, 2025",
			Tags:        []string{"golang", "docker", "templ"},
			ImageSrc:    "/assets/img/card_placeholder.jpeg",
			Description: "After getting fed up with React, SPAs, and Javascript in 2023 I decided to re-write my personal webpage in Go from Sveltekit.It was one of my best decisions simply because Go is fun to write and blazingly fast.",
		},
		{
			Title:       "Concurrency in Go: Channels and WaitGroups",
			Date:        "May 10, 2025",
			Tags:        []string{"golang", "concurrency"},
			ImageSrc:    "/assets/img/channels.webp",
			Description: "Concurrency is a powerful feature in Go (Golang) that allows developers to write efficient and scalable applications.",
		},
		{
			Title:       "Why Rust is a great fit for embedded software",
			Date:        "May 10, 2025",
			Tags:        []string{"rust", "Embedded"},
			ImageSrc:    "/assets/img/embeddedlovesrust.png",
		    Description: "If you've heard of Rust, you've heard about it's memory safety features. Those features ensure that Rust eliminates a whole classes of bugs from your application. And it can do so at compile time, greatly shortening the feedback loop of the developer.",
		},
		// Add more posts here
	}
	component := page(posts)
	templ.Handler(component).ServeHTTP(w, r)
}

func (app *Config) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(r.Method + " http://" + r.Host + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
