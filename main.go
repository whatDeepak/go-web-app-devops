package main

import (
	"log"
	"net/http"
)

// homePage handles requests to the home page and serves the home.html file
func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

// coursePage handles requests to the course page and serves the courses.html file
func coursePage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/courses.html")
}

// aboutPage handles requests to the about page and serves the about.html file
func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

// contactPage handles requests to the contact page and serves the contact.html file
func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	// Route definitions
	http.HandleFunc("/home", homePage)      // When the route is /home, serve home.html
	http.HandleFunc("/courses", coursePage) // When the route is /courses, serve courses.html
	http.HandleFunc("/about", aboutPage)    // When the route is /about, serve about.html
	http.HandleFunc("/contact", contactPage) // When the route is /contact, serve contact.html

	// Start the server
	err := http.ListenAndServe("0.0.0.0:8080", nil) // Listen on port 8080
	if err != nil {
		log.Fatal(err) // Log any errors that occur
	}
}
