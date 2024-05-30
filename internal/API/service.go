package services

import (
	"github.com/rs/zerolog"
	observation_anlyz_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data//Users/zakariasaif/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/analyze_query/ustilss"
	observation_avg_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data//Users/zakariasaif/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/average/utills"
	observation_cnt_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data//Users/zakariasaif/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/count/utills"
	entity_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data//Users/zakariasaif/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/entity/utilss"
	observation_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data//Users/zakariasaif/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/observation/ustilss"
	observation_sum_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data//Users/zakariasaif/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/summation/utills"
	timekey_usecase "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data//Users/zakariasaif/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/API/timeKey/ustills"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/config"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/mutex"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/time"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/uuid"
)

// Services contains all exposed services of the application layer
type Services struct {
	Logger                            *zerolog.Logger
	EntityUsecase                     entity_usecase.Usecase
	ObservationUsecase                observation_usecase.utilss
	ObservationAnalyzerRequestUsecase observation_anlyz_usecase.utilss
	ObservationSummationUsecase       observation_sum_usecase.utilss
	ObservationCountUsecase           observation_cnt_usecase.utilss
	ObservationAverageUsecase         observation_avg_usecase.utilss
	TimeKeyUsecase                    timekey_usecase.utilss
}

// NewAppServices Bootstraps Application Layer dependencies
func NewAppServices(appConf *config.Config, uuidProvider uuid.Provider, timeProvider time.Provider, kmutexProvider mutex.Provider, adapters *interfaceadapters.Services) Services {
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
		ObservationAnalyzerRequestUsecase: observation_anlyz_usecase.NewObservationAnalyzerRequestUsecase(
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
		ObservationAverageUsecase: observation_avg_usecase.NewObservationAverageUsecase(
			uuidProvider,
			timeProvider,
			adapters.ObservationAverageRepo,
		),
		TimeKeyUsecase: timekey_usecase.NewTimeKeyUsecase(
			uuidProvider,
			timeProvider,
			adapters.TimeKeyRepo,
		),
	}
}
