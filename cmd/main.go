package main

import (
	"fmt"

	"github.com/2beens/anas-api/internal"
	"github.com/2beens/anas-api/internal/logging"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("starting ...")

	port := 10222
	logsPath := "./service.log"

	logging.Setup(logsPath, true, "")

	log.Debugf("using port: %d", port)
	log.Debugf("using server logs path: [%s]", logsPath)

	server := internal.NewServer()

	server.Serve(port)
}
