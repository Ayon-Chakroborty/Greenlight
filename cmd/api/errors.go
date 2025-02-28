package main

import (
	"fmt"
	"net/http"
)

// log and error to the console
func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

// log an error as a response to the client
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any){
	env := envelope{"error":message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil{
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// Error when there is a unexpected problem at runtime
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error){
	app.logError(r, err)

	message := "the server envountered a problem and could not process you request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Error when the resource is not found
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request){
	message := "the resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// method will be used to send a 405 Method Not Allowed
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request){
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func(app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string){
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}





