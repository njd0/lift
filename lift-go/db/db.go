package db

import (
	"context"
	"fmt"
	"log/slog"
	config "m/lift/config"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

type DbCreds struct {
	Host     	string 
	Post  		string 
	User    	string 
	Password  int8
	DBname    int8
}

var (
	PgInstance *postgres
	pgOnce     sync.Once
)

func ConnectDB(ctx context.Context) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Secrets.PG.Host,
		config.Secrets.PG.Port, 
		config.Secrets.PG.User, 
		config.Secrets.PG.Password, 
		config.Secrets.PG.DBname,
	)

	NewPG(ctx, psqlInfo)
}

func NewPG(ctx context.Context, connString string) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		
		if err != nil {
			slog.Error("unable to create connection pool", "err", err)
		}

		err = db.Ping(ctx)
		if err != nil {
			panic(err)
		}

		PgInstance = &postgres{db}
		fmt.Println("DB Successfully connected!")
	})
}

// func (pg *postgres) Ping(ctx context.Context) error {
// 	return pg.db.Ping(ctx)
// }

// func (pg *postgres) Close() {
// 	pg.db.Close()
// }

func Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return PgInstance.db.Query(ctx, sql, args...)
}

func QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return PgInstance.db.QueryRow(ctx, sql, args...)
}

func Insert(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return PgInstance.db.Exec(ctx, sql, args...)
}
