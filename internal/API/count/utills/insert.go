package utills

import (
	"context"

	osum "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/count"
)

func (uc observationCountUsecase) Insert(ctx context.Context, os *osum.ObservationCount) (*osum.ObservationCount, error) {
	if err := uc.ObservationCountRepo.Insert(ctx, os); err != nil {
		return nil, err
	}
	os, err := uc.ObservationCountRepo.GetByPrimaryKey(ctx, os.EntityID, os.Frequency, os.Start, os.Finish)
	if err != nil {
		return nil, err
	}
	return os, nil
}
