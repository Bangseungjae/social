package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func New(addr string, maxOpenConns, maxIdleCons int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	// 최대 5초 내에 완료되지 않으면 자동으로 취소되도록 하는 것을 의미
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 5초의 타임아웃을 가짐
	defer cancel()

	// 만약 데이터베이스 서버가 5초 내에 응답하지 않으면, PingContext는 타임아웃 오류를 반환
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleCons)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, fmt.Errorf("invalid maxIdleTime format: %w", err)
	}
	db.SetConnMaxIdleTime(duration)

	return db, nil
}
