// healthCheck is an HTTP handler function that responds with a simple "ok" message.
// It is used to check the health status of the application.
//
// Parameters:
//   - w: An http.ResponseWriter to write the response.
//   - r: An http.Request representing the incoming request.
package main

import "net/http"

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
