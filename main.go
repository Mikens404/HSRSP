package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/Mikens404/HSRSP/internal/application"
	"github.com/Mikens404/HSRSP/internal/infrastructure"
	"github.com/Mikens404/HSRSP/internal/presentation"
)

var server http.Handler

func BuildServer() (*presentation.Server, error) {
	//reservation_epository
	trainRepository := infrastructure.NewTrainRepository()
	reservationRepository := infrastructure.NewReservationRepository()

	trainService := application.NewTrainService(trainRepository)
	reservationService := application.NewReservationService(reservationRepository)

	handler := presentation.NewHandler(reservationService, trainService)
	return presentation.NewServer(handler)
}

func init() {
	s, err := BuildServer()
	if err != nil {
		log.Fatal(err)
	}
	server = s
}

func main() {
	slog.Info("hello.")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target internal/presentation --clean --package presentation api/openapi.yaml
