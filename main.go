package main

import(
	"fmt"
	"net/http"
	"log"
	"time"
	"strconv"
	"github.com/getsentry/sentry-go"

)

func main(){
	err:= sentry.Init(sentry.ClientOptions{
		Dsn: "you key",
		EnableTracing: true,
		// Specify a fixed sample rate:
		// We recommend adjusting this value in production
		TracesSampleRate: 1.0,
	})
	if err != nil{
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // If an error occurs during the execution of the program, report it to Sentry
        defer func() {
            if err := recover(); err != nil {
                sentry.CaptureException(fmt.Errorf("Panic: %v", err))
                http.Error(w, "Something went wrong", http.StatusInternalServerError)
            }
        }()

        // Retrieve the user's name and age from the HTTP request
        if err != nil {
            http.Error(w, "Invalid age", http.StatusBadRequest)
            return
        }

        if r.Method == "POST" {
			ageStr := r.FormValue("age")
			age, err := strconv.Atoi(ageStr)
			if err != nil {
				fmt.Fprintln(w, "invalid age")
				return
			}
			if age < 18 {
				fmt.Fprintln(w, "you are not old enough")
			} else {
				fmt.Fprintln(w, "you are old enough")
			}
		} else {
			http.ServeFile(w, r, "index.html")
		}
    })

    // Start the web server on port 8080
    fmt.Println("Listening on port 8085")
    http.ListenAndServe(":8085", nil)
}