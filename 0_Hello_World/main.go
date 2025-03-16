package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// create a struct that holds information in our html file
type Welcome struct {
	Name string
	Time string
}

// GO Application - Entry point
func main() {

	// Instantiate a welcome struct object and pass in some information
	// We should get the name along with the timestamp as a query parameter from the URL.
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	// we will tell GO exactly where we can find our HTML file, we ask GO to parse
	// the HTML file (Notice the relative path), we wrap it in a call to
	// template.Must() which handles any errors and halts if there are any fatal errors.

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	//Our HTML comes with CSS that GO needs to provide when we run the app.
	// Here we tell GO to create a handle that looks in the static directory,
	// GO then uses the "/static/" as a URL that our HTML can refer to
	// when looking for our CSS and other files.

	http.Handle("/static/", // final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	//Go looks in the relative static directory first, then matches it to a
	//url of our choice as shown in http.Handle("/static/").
	//This url is what we need when referencing our css files
	// once the server begins. Our html code would therefore
	// be <link rel="stylesheets"  href="/static/stylesheets/...">
	//It is important to note the final url can be whatever we like,
	// so long as we are consistent.

	// This method takes in the URL path "/" and a function that takes in a
	// Response Writer and a HTTP Request.
	http.HandleFunc("/", func(write http.ResponseWriter, req *http.Request) {
		// Takes the name from the URL query, e.g. ?name=Brijesh,
		// will set welcome.Name = "Brijesh".
		if name := req.FormValue("name"); name != "" {
			welcome.Name = name
		}

		// If errors show an internal server error message
		// We also pass the welcome struct to the welcome-template.html file.
		if err := templates.ExecuteTemplate(write, "welcome-template.html", welcome); err != nil {
			http.Error(write, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the web server, set the port to listen to 8080.
	// Without a path it assumes localhost
	// Print any errors from starting the webserver using fmt
	fmt.Println(http.ListenAndServe(":8080", nil))
}
