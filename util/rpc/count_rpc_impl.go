package rpc

import "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/util/dtos"

func (s *StockYardRPCClient) ListObservationCounts(dto *dtos.ObservationCountFilterRequestDTO) (*dtos.ObservationCountListResponseDTO, error) {
	s.Logger.Info().Msg("calling remote list observation count function...")

	// Create the response payload that will be filled out by the server.
	var reply dtos.ObservationCountListResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.ListObservationCounts", dto, &reply)
	if err != nil {
		s.Logger.Error().Err(err).Caller().Str("RemoteAddress", s.serverAddress).Msg("failed making remote proceedure call")
		return nil, err
	}

	s.Logger.Info().
		Msg("succesfully called remote proceedure")

	return &reply, nil
}
