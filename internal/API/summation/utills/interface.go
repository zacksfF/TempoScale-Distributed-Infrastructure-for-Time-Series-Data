package utills

import (
	"context"

	odomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/summation"
	osum "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/summation"
	timep "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/time"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/uuid"
)

// Usecase Provides interface for the observation summation use cases.
type Usecase interface {
	Insert(ctx context.Context, e *osum.ObservationSummation) (ee *osum.ObservationSummation, err error)
	ListAndCountByFilter(ctx context.Context, ef *osum.ObservationSummationFilter) ([]*osum.ObservationSummation, uint64, error)
}

type observationSummationUsecase struct {
	Time                     timep.Provider
	UUID                     uuid.Provider
	ObservationSummationRepo odomain.Repository
}

// NewObservationSummationUsecase Constructor function for the `ObservationSummationUsecase` implementation.
func NewObservationSummationUsecase(
	uuidp uuid.Provider,
	tp timep.Provider,
	o odomain.Repository,

) *observationSummationUsecase {
	return &observationSummationUsecase{
		Time:                     tp,
		UUID:                     uuidp,
		ObservationSummationRepo: o,
	}
}
