package app

import (
	"net/http"

	"github.com/dzen-it/di/app/handlers/node"
	"github.com/dzen-it/di/common/configs"
	log "github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

func getRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/node/:id", node.Get)
	router.POST("/node", node.Post)
	router.PUT("/node", node.Put)

	return router
}

func withTracing(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%v %v %v", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func Start(address string) {
	router := getRouter()
	log.Infof("Start listen address %s HTTPS is %v", address, configs.Config.HTTPS)

	if configs.Config.HTTPS {
		log.Fatal(http.ListenAndServeTLS(address, configs.Config.Certfile, configs.Config.Keyfile, withTracing(router)))
	} else {
		log.Fatal(http.ListenAndServe(address, withTracing(router)))
	}
}
