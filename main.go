package main

import (
	"fmt"
	"net/http"
)

// // Benefit of this struct type
// // Access to different fields while still having a function that satisfies Handler interface
// type Router struct{}

// func (Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	}
// }

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Check e.g. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:foo@gmail.com\">foo@gmail.com</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1><p>
    <ul>
        <li>Is there a free version? Yes! We offer a free trial for 30 days</li>
    </ul>
    `)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	// URL Path vs RawPath
	// fmt.Fprintln(w, r.URL.Path)
	// fmt.Fprintln(w, r.URL.RawPath) // allows for encoding/decoding
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func main() {
	fmt.Println("Starting the server on :3000")

	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)
	// http.ListenAndServe(":3000", nil)

	// http.Handler - interface with the ServeHTTP method
	// http.HandlerFunc - a function type that accepts same args as ServeHttp method. Also implements http.Handler
	// This looks like a function call but is actually a type conversion to the http.HandlerFunc type
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
