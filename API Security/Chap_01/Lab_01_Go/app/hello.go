package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!\n", r.URL.Path[1:])
	fmt.Fprintf(w, "\nMethod: %s and the type is %T\n", r.Method, r.Method)
	fmt.Fprintf(w, "\nURLPath: %s and the type is %T\n", r.URL.Path, r.URL.Path)
	fmt.Fprintf(w, "\nRemoteAddress: %s and the type is %T\n", r.RemoteAddr, r.RemoteAddr)

	// ✅ These must start with 'w' as the writer
	fmt.Fprintf(w, "\nFull URL object: %+v\n", r.URL)
	fmt.Fprintf(w, "\nFull URL (Go syntax): %#v\n", r.URL)
	fmt.Fprintf(w, "\nScheme: %s\n", r.URL.Scheme)
	fmt.Fprintf(w, "\nHost: %s\n", r.URL.Host)
	fmt.Fprintf(w, "\nPath: %s\n", r.URL.Path)
	fmt.Fprintf(w, "\nRawQuery: %s\n", r.URL.RawQuery)
	fmt.Fprintf(w, "\nFragment: %s\n", r.URL.Fragment)

	// ✅ Proper JSON marshal for headers
	headersJSON, err := json.MarshalIndent(r.Header, "", "  ")
	if err != nil {
		http.Error(w, "Failed to marshal headers", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "\nHeaders (JSON):\n%s\n", headersJSON)
}

func main() {
	http.HandleFunc("/josh", handler)
	http.ListenAndServe(":8080", nil)
}
