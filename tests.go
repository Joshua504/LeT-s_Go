package main

import "net/http"

func headerManipulating(w http.ResponseWriter, r *http.Request) {

	//? Set a new cache-control header. If an existing "Cache-Control" header exists
	//? it will be overwritten.
	w.Header().Set("Cache-Control", "public, max-age=31536000")

	// ? In contrast, the Add() method appends a new "Cache-Control" header and can
	// ? be called multiple times.
	w.Header().Add("Cache-Control", "Public")
	w.Header().Add("Cache-Control", "max-age=315360000")

	// ? Delete all values for the "Cache-Control" header.
	w.Header().Del("Cache-Control")

	// ? Retrieve the first value for the "Cache-Control" header.
	w.Header().Get("Cache-Conttrol")

	// ? Retrieve a slice of all values for the "Cache-Control" header.
	w.Header().Values("Cache-Control")
	
}
