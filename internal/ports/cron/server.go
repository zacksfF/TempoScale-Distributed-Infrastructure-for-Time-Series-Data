package cron

import (
	"github.com/go-co-op/gocron"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/config"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/mutex"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/time"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/uuid"
)

// Server Represents the http server running for this service
type Server struct {
	Time        time.Provider
	UUID        uuid.Provider
	KMutext     mutex.Provider
	Scheduler   *gocron.Scheduler
	AServices API.server
}

// NewServer HTTP Server constructor
func NewServer(appConf *config.Config, uuidProvider uuid.Provider, timeProvider time.Provider, kmutexProvider mutex.Provider, appServices api.Services) *Server {

	scheduler := gocron.NewScheduler(t.UTC)

	cronServer := &Server{
		Time:        timeProvider,
		UUID:        uuidProvider,
		KMutext:     kmutexProvider,
		AppServices: appServices,
		Scheduler:   scheduler,
	}
	return cronServer
}

// ListenAndServe Starts listening for requests
func (cronServer *Server) ListenAndServe() {
	cronServer.Scheduler.Cron("*/1 * * * *").Do(cronServer.RunAnalyzer) // Every minute.

	cronServer.AppServices.Logger.Info().Msg("cron service started")
	defer cronServer.AppServices.Logger.Info().Msg("cron service stopped")

	// Execute the analyzer on startup.
	if err := cronServer.RunAnalyzer(); err != nil {
		cronServer.AppServices.Logger.Info().Msg("run analyzer has error")
	}

	// Starts the scheduler and blocks current execution path.
	cronServer.Scheduler.StartBlocking()
}
