package utilss

import (
	"context"

	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/entity"
	edomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/entity"
)

type Usecase interface {
	Insert(ctx context.Context, e *entity.Entity) (ee *entity.Entity, err error)
	ListAndCountByFilter(ctx context.Context, ef *entity.EntityFilter) ([]*entity.Entity, uint64, error)
	Delete(ctx context.Context, entityID uint64) (err error)
}

type entityUsecase struct {
	Time       timep.Provider
	UUID       uuid.Provider
	EntityRepo edomain.Repository
}

// NewEntityUsecase Constructor function for the `UserUsecase` implementation.
func NewEntityUsecase(
	uuidp uuid.Provider,
	tp timep.Provider,
	e edomain.Repository,

) *entityUsecase {
	return &entityUsecase{
		Time:       tp,
		UUID:       uuidp,
		EntityRepo: e,
	}
}
