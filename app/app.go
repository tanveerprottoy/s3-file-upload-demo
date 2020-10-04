package app

import (
	"log"
	"net/http"
)

// App struct
type App struct {
	Mux *http.ServeMux
}

const BucketName = "com.tanveershafeeprottoy.test"

func (app *App) initS3() {
	InitS3()
	err := GetBucket(BucketName)
	if err != nil {
		// bucket does not exist create it
		log.Println("bucket does not exist create it")
		err = CreateBucket(BucketName)
		if err != nil {
			panic("Failed to create bucket")
		}
	}
	log.Println("bucket exists")
}

// Init app
func (app *App) Init() {
	app.Mux = http.NewServeMux()
	app.Mux.HandleFunc("/", PostFile)
	app.initS3()
}

// Run app
func (app *App) Run() {
	log.Fatal(
		http.ListenAndServe(
			":3000",
			app.Mux,
		),
	)
}
