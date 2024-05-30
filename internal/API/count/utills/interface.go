package utills

import (
	"context"

	odomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/count"
	osum "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/count"
	timep "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/time"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/uuid"
)

// Usecase Provides interface for the observation count use cases.
type Usecase interface {
	Insert(ctx context.Context, e *osum.ObservationCount) (ee *osum.ObservationCount, err error)
	ListAndCountByFilter(ctx context.Context, ef *osum.ObservationCountFilter) ([]*osum.ObservationCount, uint64, error)
}

type observationCountUsecase struct {
	Time                 timep.Provider
	UUID                 uuid.Provider
	ObservationCountRepo odomain.Repository
}

// NewObservationCountUsecase Constructor function for the `ObservationCountUsecase` implementation.
func NewObservationCountUsecase(
	uuidp uuid.Provider,
	tp timep.Provider,
	o odomain.Repository,

) *observationCountUsecase {
	return &observationCountUsecase{
		Time:                 tp,
		UUID:                 uuidp,
		ObservationCountRepo: o,
	}
}
