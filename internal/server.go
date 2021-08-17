package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/2beens/anas-api/internal/middleware"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
}

func NewServer() *Server {
	s := &Server{}
	return s
}

func (s *Server) Serve(port int) {
	router := s.routerSetup()

	ipAndPort := fmt.Sprintf("%s:%d", "localhost", port)

	httpServer := &http.Server{
		Handler:      router,
		Addr:         ipAndPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	chOsInterrupt := make(chan os.Signal, 1)
	signal.Notify(chOsInterrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Infof(" > server listening on: [%s]", ipAndPort)
		log.Fatal(httpServer.ListenAndServe())
	}()

	receivedSig := <-chOsInterrupt

	log.Warnf("signal [%s] received ...", receivedSig)

	// go to sleep 🥱
	s.gracefulShutdown(httpServer)
}

func (s *Server) routerSetup() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("I'm OK"))
	})

	r.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("pong"))
	})

	r.Use(middleware.PanicRecovery())
	r.Use(middleware.LogRequest())
	r.Use(middleware.Cors())
	r.Use(middleware.DrainAndCloseRequest())

	return r
}

func (s *Server) gracefulShutdown(httpServer *http.Server) {
	log.Debug("graceful shutdown initiated ...")

	maxWaitDuration := time.Second * 15
	ctx, timeoutCancel := context.WithTimeout(context.Background(), maxWaitDuration)
	defer timeoutCancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error(" >>> failed to gracefully shutdown http server")
	}
	log.Warn("server shut down")
}
