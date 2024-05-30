package ustilss

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/config"
	oardomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/analyze_query"
	oadomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/average"
	ocdomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/count"
	osdomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/summation"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/mutex"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/time"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/uuid"
)

// Usecase Provides interface for the observation use cases.
type Usecase interface {
	Insert(ctx context.Context, o *oardomain.ObservationAnalyzerRequest) (*oardomain.ObservationAnalyzerRequest, error)
	ListAll(ctx context.Context) ([]*oardomain.ObservationAnalyzerRequest, error)
	DeleteByPrimaryKey(ctx context.Context, entityID uint64, uuid string) error
	RunAnalyzer(ctx context.Context) error
}

type observationAnalyzerRequestUsecase struct {
	HasAnalyzer                    bool
	Logger                         *zerolog.Logger
	Time                           time.Provider
	UUID                           uuid.Provider
	KMutex                         mutex.Provider
	ObservationAnalyzerRequestRepo oardomain.Repository
	ObservationRepo                oadomain.Repository
	ObservationCountRepo           ocdomain.Repository
	ObservationSummationRepo       osdomain.Repository
	ObservationAverageRepo         oadomain.Repository
}

// NewObservationAnalyzerRequestUsecase Constructor function for the `ObservationAnalyzerRequest` implementation.
func NewObservationAnalyzerRequestUsecase(
	appConf *config.Config,
	logger *zerolog.Logger,
	uuidp uuid.Provider,
	tp time.Provider,
	kmutexp mutex.Provider,
	o oadomain.Repository,
	oc ocdomain.Repository,
	os osdomain.Repository,
	oa oadomain.Repository,
	oar oardomain.Repository,

) *observationAnalyzerRequestUsecase {
	return &observationAnalyzerRequestUsecase{
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
