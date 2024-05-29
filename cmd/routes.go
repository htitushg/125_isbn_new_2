package main

import (
	"125_isbn_new/ui"
	"github.com/justinas/alice"
	"net/http"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	app.logger.Info("Entrée dans routes.go")
	mux := http.NewServeMux()

	// Utilisez la fonction http.FileServerFS() pour créer un gestionnaire HTTP qui
	// sert les fichiers intégrés dans ui.Files. Il est important de noter que nos
	// fichiers statiques sont contenus dans le dossier système de fichiers intégré
	//"static" de ui.Files. Ainsi, par exemple, notre feuille de style CSS se trouve à l'adresse
	// "static/css/main.css". Cela signifie que nous n'avons plus besoin d'indiquer le
	// préfixe de l'URL de la requête -- toutes les requêtes commençant par /static/ peuvent
	// transmis directement au serveur de fichiers et au fichier statique correspondant
	// le fichier sera servi (tant qu'il existe).
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	// Sending the assets to the clients: remplacé par la ligne au dessus
	// fs := http.FileServer(http.Dir(models.Path + "assets"))
	// mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
	// Add a new GET /ping route.
	mux.HandleFunc("GET /ping", ping)
	mux.HandleFunc("GET /images/couverture/{name}", app.CouvertureLivreGet)
	// Handling MethodNotAllowed error on /
	//mux.HandleFunc("/{$}", app.indexHandlerNoMeth)
	// Handling StatusNotFound error everywhere else
	//mux.HandleFunc("/", app.indexHandlerOther)

	// Add the authenticate() middleware to the chain.
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.index))
	mux.Handle("GET /index", dynamic.ThenFunc(app.IndexHandlerGet))
	mux.Handle("GET /affichelivres", dynamic.ThenFunc(app.LivresHandlerGet))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/view", dynamic.ThenFunc(app.home))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))
	mux.Handle("GET /user/modif", dynamic.ThenFunc(app.userModifGet))
	mux.Handle("POST /user/modif", dynamic.ThenFunc(app.userModifPost))

	protected := dynamic.Append(app.requireAuthentication)
	mux.Handle("GET /home", protected.ThenFunc(app.HomeHandlerGet))
	mux.Handle("GET /livre", protected.ThenFunc(app.LivreHandlerGet))
	mux.Handle("POST /affichelivre", protected.ThenFunc(app.LivreHandlerPost))
	mux.Handle("GET /afficheauteurs", protected.ThenFunc(app.ListAuteursHandlerGet))
	mux.Handle("GET /afficheediteurs", protected.ThenFunc(app.ListEditeursHandlerGet))
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))
	mux.Handle("GET /user/logout", protected.ThenFunc(app.userLogoutGet))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)

	//return mux
}
