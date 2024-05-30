package repositopry

import (
	"context"

	domain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/count"
)

func (r *ObservationCountRepoImpl) InsertOrGetByPrimaryKey(ctx context.Context, oc *domain.ObservationCount) (*domain.ObservationCount, error) {
	doesExist, err := r.CheckIfExistsByPrimaryKey(
		ctx, oc.EntityID, oc.Frequency, oc.Start, oc.Finish,
	)
	if err != nil {
		return nil, err
	}

	if doesExist == true {
		return r.GetByPrimaryKey(ctx, oc.EntityID, oc.Frequency, oc.Start, oc.Finish)
	}
	if err := r.Insert(ctx, oc); err != nil {
		return nil, err
	}
	return r.GetByPrimaryKey(ctx, oc.EntityID, oc.Frequency, oc.Start, oc.Finish)
}
