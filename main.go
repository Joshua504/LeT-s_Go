package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// * Define a home handler function which writes a byte slice containing
// * "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	//? Use the Header().Add() method to add a 'Server: Go' header to the
	//? response header map. The first parameter is the header name, and
	//? the second parameter is the header value.
	w.Header().Add("Server", "GO")

	//? Initialize a slice containing the paths to the two files. It's important
	//? to note that the file containing our base template must be the *first*
	//? file in the slice.
	//? Include the navigation partial in the template files.
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	//? Use the template.ParseFiles() function to read the template file into a
	//? template set. If there's an error, we log the detailed error message, use
	//? the http.Error() function to send an Internal Server Error response to the
	//? user, and then return from the handler so no subsequent code is executed.
	//! Use the template.ParseFiles() function to read the files and store the
	//! templates in a template set. Notice that we use ... to pass the contents
	//! of the files slice as variadic arguments.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//? Then we use the Execute() method on the template set to write the
	//? template content as the response body. The last parameter to Execute()
	//? represents any dynamic data that we want to pass in, which for now we'll
	//? leave as nil.
	//! Use the ExecuteTemplate() method to write the content of the "base"
	//! template as the response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	//* Extract the value of the id wildcard from the request using r.PathValue()
	//* and try to convert it to an integer using the strconv.Atoi() function. If
	//* it can't be converted to an integer, or the value is less than 1, we
	//* return a 404 page not found response.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	//* Use the fmt.Sprintf() function to interpolate the id value with a
	//* message, then write it as the HTTP response.
	// msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	// w.Write([]byte(msg))

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	//* Use the w.WriteHeader() method to send a 201 status code.
	w.WriteHeader(http.StatusCreated)

	//* Then w.Write() method to write the response body as normal.
	w.Write([]byte("Saving a new snippet..."))
}

func main() {
	//* Use the http.NewServeMux() function to initialize a new servemux, then
	//* register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home) //? Restrict this route to exact matches on / only.

	//! Create a file server which serves files out of the "./ui/static" directory.
	//! Note that the path given to the http.Dir function is relative to the project
	//! directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//? Use the mux.Handle() function to register the file server as the handler for
	//? all URL paths that start with "/static/". For matching paths, we strip the
	//? "/static" prefix before the request reaches the file server.
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	//* Register the two new handler functions and corresponding route patterns with
	//* the servemux, in exactly the same way that we did before.

	//* Prefix the route patterns with the required HTTP method (for now, we will
	//* restrict all three routes to acting on GET requests).
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) //? Add the {id} wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost) //? route which is restricted to POST requests only.

	//* Print a log message to say that the server is starting.
	log.Print("Starting server on :4000")

	//* Use the http.ListenAndServe() function to start a new web server. We pass in
	//* two parameters: the TCP network address to listen on (in this case ":4000")
	//* and the servemux we just created. If http.ListenAndServe() returns an error
	//* we use the log.Fatal() function to log the error message and exit. Note
	//* that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
