package ustilss

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/observation"
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
	KMutex                         kmutex.Provider
	ObservationRepo                odomain.Repository
	ObservationCountRepo           ocdomain.Repository
	ObservationSummationRepo       osdomain.Repository
	ObservationAverageRepo         oadomain.Repository
	ObservationAnalyzerRequestRepo oardomain.Repository
}

// NewObservationUsecase Constructor function for the `UserUsecase` implementation.
func NewObservationUsecase(
	appConf *config.Conf,
	logger *zerolog.Logger,
	uuidp uuid.Provider,
	tp timep.Provider,
	kmutexp kmutex.Provider,
	o odomain.Repository,
	oc ocdomain.Repository,
	os osdomain.Repository,
	oa oadomain.Repository,
	oar oardomain.Repository,

) *observationUsecase {
	return &observationUsecase{
		HasAnalyzer:                    appConf.Setting.HasAnalyzer,
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
