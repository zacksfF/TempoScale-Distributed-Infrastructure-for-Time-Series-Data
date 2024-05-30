package ustilss

import (
	"context"

	oardomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/analyze_query"
)

func (uc observationAnalyzerRequestUsecase) ListAll(ctx context.Context) ([]*oardomain.ObservationAnalyzerRequest, error) {
	arrCh := make(chan []*oardomain.ObservationAnalyzerRequest)

	go func() {
		arr, err := uc.ObservationAnalyzerRequestRepo.ListAll(ctx)
		if err != nil {
			arrCh <- nil
			return
		}
		arrCh <- arr[:]
	}()

	arr := <-arrCh

	return arr, nil
}
