package ustilss

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/config"
	oardomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/analyze_query"
	oadomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/average"
	ocdomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/count"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/observation"
	odomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/observation"
	osdomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/summation"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/mutex"
	timep "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/time"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/uuid"
)

type Usecase interface {
	Insert(ctx context.Context, e *observation.Observation) (ee *observation.Observation, err error)
	CheckIfExistsByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) (bool, error)
	ListAndCountByFilter(ctx context.Context, ef *observation.ObservationFilter) ([]*observation.Observation, uint64, error)
	DeleteByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) error
}

type observationUsecase struct {
	HasAnalyzer                    bool
	Logger                         *zerolog.Logger
	Time                           timep.Provider
	UUID                           uuid.Provider
	KMutex                         mutex.Provider
	ObservationRepo                odomain.Repository
	ObservationCountRepo           ocdomain.Repository
	ObservationSummationRepo       osdomain.Repository
	ObservationAverageRepo         oadomain.Repository
	ObservationAnalyzerRequestRepo oardomain.Repository
}

// NewObservationUsecase Constructor function for the `UserUsecase` implementation.
func NewObservationUsecase(
	appConf *config.Config,
	logger *zerolog.Logger,
	uuidp uuid.Provider,
	tp timep.Provider,
	kmutexp mutex.Provider,
	o odomain.Repository,
	oc ocdomain.Repository,
	os osdomain.Repository,
	oa oadomain.Repository,
	oar oardomain.Repository,

) *observationUsecase {
	return &observationUsecase{
		HasAnalyzer:                    appConf.Setting.HasAnlyzer,
		Logger:                         logger,
		Time:                           tp,
		UUID:                           uuidp,
		KMutex:                         kmutexp,
		ObservationRepo:                o,
		ObservationCountRepo:           oc,
		ObservationSummationRepo:       os,
		ObservationAverageRepo:         oa,
		ObservationAnalyzerRequestRepo: oar,
	}
}
