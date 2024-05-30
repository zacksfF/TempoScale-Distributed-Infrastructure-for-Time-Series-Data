package utills

import (
	"context"

	osum "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/count"
)

func (uc observationCountUsecase) ListAndCountByFilter(ctx context.Context, ef *osum.ObservationCountFilter) ([]*osum.ObservationCount, uint64, error) {
	arrCh := make(chan []*osum.ObservationCount)
	countCh := make(chan uint64)

	go func() {
		arr, err := uc.ObservationCountRepo.ListByFilter(ctx, ef)
		if err != nil {
			arrCh <- nil
			return
		}
		arrCh <- arr[:]
	}()

	go func() {
		count, err := uc.ObservationCountRepo.CountByFilter(ctx, ef)
		if err != nil {
			countCh <- 0
			return
		}
		countCh <- count
	}()

	arr, count := <-arrCh, <-countCh

	return arr, count, nil
}
