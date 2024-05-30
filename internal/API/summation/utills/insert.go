package utills

import (
	"context"

	osum "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/summation"
)

func (uc observationSummationUsecase) Insert(ctx context.Context, os *osum.ObservationSummation) (*osum.ObservationSummation, error) {
	if err := uc.ObservationSummationRepo.Insert(ctx, os); err != nil {
		return nil, err
	}
	os, err := uc.ObservationSummationRepo.GetByPrimaryKey(ctx, os.EntityID, os.Frequency, os.Start, os.Finish)
	if err != nil {
		return nil, err
	}
	return os, nil
}
