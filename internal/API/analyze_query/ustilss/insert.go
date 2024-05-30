package ustilss

import (
	"context"

	oardomain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/analyze_query"
)

func (uc observationAnalyzerRequestUsecase) Insert(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (*oardomain.ObservationAnalyzerRequest, error) {
	if err := uc.ObservationAnalyzerRequestRepo.Insert(ctx, req); err != nil {
		return nil, err
	}
	o, err := uc.ObservationAnalyzerRequestRepo.GetByPrimaryKey(ctx, req.EntityID, req.UUID)
	if err != nil {
		return nil, err
	}
	return o, nil
}
