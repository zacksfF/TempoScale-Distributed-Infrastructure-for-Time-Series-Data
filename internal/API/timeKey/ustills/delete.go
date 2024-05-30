package ustills

import (
	"context"

	timekey "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/timeKey"
)

func (uc timekeyUsecase) DeleteByFilter(ctx context.Context, tkf *timekey.TimeKeyFilter) error {
	return uc.TimeKeyRepo.DeleteByFilter(ctx, tkf)
}
