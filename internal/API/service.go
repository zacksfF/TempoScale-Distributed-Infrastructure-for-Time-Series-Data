package services

import (
	"github.com/rs/zerolog"
	observation_avg_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/analyze_query/ustilss"
	average "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/average/repository"
	observation_cnt_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/count/utills"
	entity_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/entity/utilss"
	observation_anlyz_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/observation/ustilss"
	observation_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/observation/ustilss"
	observation_sum_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/summation/utills"
	timekey_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/timeKey/ustills"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/adapters"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/config"
	analyzequery "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/analyze_query"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/count"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/summation"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/mutex"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/time"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/uuid"
)

// Services contains all exposed services of the application layer
type Services struct {
	Logger                            *zerolog.Logger
	EntityUsecase                     entity_usecase.Usecase
	ObservationUsecase                observation_usecase.Usecase
	ObservationAnalyzerRequestUsecase observation_anlyz_usecase.Usecase
	ObservationSummationUsecase       observation_sum_usecase.Usecase
	ObservationCountUsecase           observation_cnt_usecase.Usecase
	ObservationAverageUsecase         observation_avg_usecase.Usecase
	TimeKeyUsecase                    timekey_usecase.Usecase
}

func NewObservationAnalyzerRequestUsecase(
	appConf *config.Config,
	logger *zerolog.Logger,
	uuidProvider uuid.Provider,
	timeProvider time.Provider,
	kmutexProvider mutex.Provider,
	avgRepo average.ObservationAverageRepoImpl,
	cntRepo count.Repository,
	sumRepo summation.Repository,
	avgRepo2 average.ObservationAverageRepoImpl, // Possibly a typo or duplicate in the function definition.
	anlyzQueryRepo analyzequery.Repository,
) observation_avg_usecase.Usecase

// NewAppServices Bootstraps Application Layer dependencies
func NewAppServices(
	appConf *config.Config,
	uuidProvider uuid.Provider,
	timeProvider time.Provider,
	kmutexProvider mutex.Provider,
	adapters *adapters.Services,
) Services {
	return Services{
		Logger: adapters.Logger,
		EntityUsecase: entity_usecase.NewEntityUsecase(
			uuidProvider,
			timeProvider,
			adapters.EntityRepo,
		),
		ObservationUsecase: observation_usecase.NewObservationUsecase(
			appConf,
			adapters.Logger,
			uuidProvider,
			timeProvider,
			kmutexProvider,
			adapters.ObservationRepo,
			adapters.ObservationCountRepo,
			adapters.ObservationSummationRepo,
			adapters.ObservationAverageRepo,
			adapters.ObservationAnalyzerRequestRepo,
		),
		ObservationAnalyzerRequestUsecase: observation_anlyz_usecase.NewObservationUsecase(
			appConf,
			adapters.Logger,
			uuidProvider,
			timeProvider,
			kmutexProvider,
			adapters.ObservationRepo,
			adapters.ObservationCountRepo,
			adapters.ObservationSummationRepo,
			adapters.ObservationAverageRepo,
			adapters.ObservationAnalyzerRequestRepo,
		),
		ObservationSummationUsecase: observation_sum_usecase.NewObservationSummationUsecase(
			uuidProvider,
			timeProvider,
			adapters.ObservationSummationRepo,
		),
		ObservationCountUsecase: observation_cnt_usecase.NewObservationCountUsecase(
			uuidProvider,
			timeProvider,
			adapters.ObservationCountRepo,
		),
		ObservationAverageUsecase: observation_avg_usecase.NewObservationAnalyzerRequestUsecase(
			appConf,
			adapters.Logger,
			uuidProvider,
			timeProvider,
			kmutexProvider,
			adapters.ObservationAverageRepo,
			adapters.ObservationCountRepo,
			adapters.ObservationSummationRepo,
			adapters.ObservationAnalyzerRequestRepo,
		),
		TimeKeyUsecase: timekey_usecase.NewTimeKeyUsecase(
			uuidProvider,
			timeProvider,
			adapters.TimeKeyRepo,
		),
	}
}
