package main

import (
	"net/http"

	"snippet-box.kensonjohnson.com/ui"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	// Setup server ping
	mux.HandleFunc("GET /ping", ping)

	// Middleware handler
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	// Define routes
	mux.Handle("GET /{$}", dynamic.ThenFunc((app.home)))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	// Adding authentication middleware to the chain
	protected := dynamic.Append(app.requireAuthentication)

	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	// Middleware that should run on every request
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
