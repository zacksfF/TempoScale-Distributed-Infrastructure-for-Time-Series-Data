package ustills

import (
	"context"

	timekey "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/timeKey"
	tkdomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/timeKey"
	timep "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/time"
	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/pkg/uuid"
)

// Usecase Provides interface for the timekey use cases.
type Usecase interface {
	Insert(ctx context.Context, tk *timekey.TimeKey) (ee *timekey.TimeKey, err error)
	ListAndCountByFilter(ctx context.Context, tkf *timekey.TimeKeyFilter) ([]*timekey.TimeKey, uint64, error)
	DeleteByFilter(ctx context.Context, tkf *timekey.TimeKeyFilter) error
}

type timekeyUsecase struct {
	Time        timep.Provider
	UUID        uuid.Provider
	TimeKeyRepo tkdomain.Repository
}

// NewTimeKeyUsecase Constructor function for the `UserUsecase` implementation.
func NewTimeKeyUsecase(
	uuidp uuid.Provider,
	tp timep.Provider,
	o tkdomain.Repository,

) *timekeyUsecase {
	return &timekeyUsecase{
		Time:        tp,
		UUID:        uuidp,
		TimeKeyRepo: o,
	}
}
