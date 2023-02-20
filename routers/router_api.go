package routers

import (
	"MobileIDAPI/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func IdentityAPI(router *mux.Router) *mux.Router {
	router.Handle("/mobileid/v1/identify",
		negroni.New(
			negroni.HandlerFunc(controllers.IdentityByISDNAndIPDestination),
		)).Methods("POST")
	return router

}
