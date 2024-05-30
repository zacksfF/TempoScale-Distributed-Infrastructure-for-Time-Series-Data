package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/adapters"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/config"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/mutex"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/time"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/uuid"
)

func init() {
	rootCmd.AddCommand(rpcServeCmd)
}

var rpcServeCmd = &cobra.Command{
	Use:   "rpc_serve",
	Short: "Run the RPC API for the server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Load up all the environment variables.
		appConf := config.AppConfig()

		uuidProvider := uuid.NewUUIDProvider()
		timeProvider := time.NewTimeProvider()
		kmutexProvider := mutex.NewKeyMutexProvider()

		// Load up all the interface adapters.
		adapters, err := adapters.NewServices(appConf)
		if err != nil {
			log.Fatal().
				Err(err).
				Msgf("Cannot start interfaceadapters %s", err)
		}

		// Load up all the app services.
		appServices := app.NewAppServices(appConf, uuidProvider, timeProvider, kmutexProvider, adapters)

		// Load up our HTTP server and connect it with the rest of our application.
		inputPortsServices := ports.NewServices(appConf, uuidProvider, timeProvider, kmutexProvider, appServices)

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		go inputPortsServices.CronServer.ListenAndServe()
		go inputPortsServices.RPCServer.ListenAndServe()

		addr := fmt.Sprintf("%s:%s", appConf.Server.IP, appConf.Server.Port)

		log.Info().Msgf("rpc server started listening on UDP via %s", addr)

		// Run the main loop blocking code.
		<-done

		stopMainRuntimeLoop(inputPortsServices)
	},
}

func stopMainRuntimeLoop(services inputports.Services) {
	log.Info().Msg("Starting graceful shutdown now...")

	// DEVELOPERS NOTE:
	// Write your closing code here.
	// . . .

	log.Info().Msg("Graceful shutdown finished.")
	log.Info().Msg("Server Exited")
}
