package ustills

import (
	"context"

	timekey "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/timeKey"
)

func (uc timekeyUsecase) Insert(ctx context.Context, tk *timekey.TimeKey) (*timekey.TimeKey, error) {
	if err := uc.TimeKeyRepo.Insert(ctx, tk); err != nil {
		return nil, err
	}
	tk, err := uc.TimeKeyRepo.GetByPrimaryKey(ctx, tk.EntityID, tk.Timestamp)
	if err != nil {
		return nil, err
	}
	return tk, nil
}
