package adapters

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/adapters/migrations"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/adapters/storage/postgres"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/config"
)

// Services contains the exposed services of interface adapters
type Services struct {
	Logger                         *zerolog.Logger
	EntityRepo                     entity_d.Repository
	TimeKeyRepo                    timekey_d.Repository
	ObservationAnalyzerRequestRepo observation_analyz_d.Repository
	ObservationRepo                observation_d.Repository
	ObservationSummationRepo       observation_sum_d.Repository
	ObservationCountRepo           observation_cnt_d.Repository
	ObservationAverageRepo         observation_avg_d.Repository
}

// NewServices Instantiates the interface adapter services
func NewServices(appConf *config.Config) (*Services, error) {
	// Step 2: Connect to database.
	db, err := postgres.ConnectDB(appConf)
	if err != nil {
		return nil, err
	}

	// Step 2: Perform our automatic database migrations (if enabled)
	if appConf.DB.HasAutoMigarations {
		if err := migrations.RunOnDB(db); err != nil {
			return nil, err
		}
	} else {
		log.Warn().Msg("No migrations occured - you must do this manually.")
	}

	// Step 3: Default level for this example is info, unless debug flag is present
	var logger zerolog.Logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if appConf.Setting.HasDebugging {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// The following line of code adds a pretty output to the console.
	logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().
		Timestamp(). // Add timestamp to every call.
		Logger()

	log.Logger = logger

	return &Services{
		Logger:                         &logger,
		EntityRepo:                     entity_r.NewEntityRepoImpl(db, &logger),
		ObservationAnalyzerRequestRepo: observation_analyz_r.NewObservationAnalyzerRequestRepoImpl(db, &logger),
		ObservationRepo:                observation_r.NewObservationRepoImpl(db, &logger),
		ObservationSummationRepo:       observation_sum_r.NewObservationSummationRepoImpl(db, &logger),
		ObservationCountRepo:           observation_cnt_r.NewObservationCountRepoImpl(db, &logger),
		ObservationAverageRepo:         observation_avg_r.NewObservationAverageRepoImpl(db, &logger),
		TimeKeyRepo:                    timekey_r.NewTimeKeyRepoImpl(db, &logger),
	}, nil
}
