package repository

import (
	"context"
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"
	domain "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/domain/analyze_query"
)

func (s *ObservationAnalyzerRequestRepoImpl) ListAll(ctx context.Context) ([]*domain.ObservationAnalyzerRequest, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(s.dbCache).PlaceholderFormat(sq.Dollar)
	rds := psql.Select(
		"entity_id",
		"uuid",
		"timestamp",
		"type",
		"observation",
	).From("observation_analyzer_requests")

	rds = rds.OrderBy("uuid" + " " + "ASC")

	// Note:
	// (1) https://ivopereira.net/efficient-pagination-dont-use-offset-limit
	// (2) https://github.com/Masterminds/squirrel/blob/def598cbb358368fbfc3f6a9a914699a36846992/select_test.go#L41

	// rds = rds.Offset(f.Offset).Suffix("FETCH FIRST ? ROWS ONLY", f.Limit)

	// Build the SQL statement and the accomponing arguments.
	sql, args, err := rds.ToSql()

	// // For debugging purposes only.
	// log.Println("sql:", sql)
	// log.Println("args:", args)
	// log.Println("err:", err)

	stmt, err := s.db.Prepare(sql)
	if err != nil {
		s.logger.Error().Err(err).Caller().Msg("failed prepare")
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		s.logger.Error().Err(err).Caller().Msg("failed query context")
		return nil, err
	}

	var arr []*domain.ObservationAnalyzerRequest
	defer rows.Close()
	for rows.Next() {
		var obin []byte
		m := new(domain.ObservationAnalyzerRequest)
		err := rows.Scan(
			&m.EntityID,
			&m.UUID,
			&m.Timestamp,
			&m.Type,
			&obin,
		)
		if err != nil {
			s.logger.Error().Err(err).Caller().Msg("database scan error")
			return nil, err
		}
		if err := json.Unmarshal(obin, &m.Observation); err != nil {
			return nil, err
		}
		arr = append(arr, m)
	}
	err = rows.Err()
	if err != nil {
		s.logger.Error().Err(err).Caller().Msgf("database")
		return nil, err
	}

	if arr == nil {
		return []*domain.ObservationAnalyzerRequest{}, nil
	}
	return arr, err
}
