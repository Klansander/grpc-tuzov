package postgresql

import (
	"context"
	"fmt"
	"sfera-server/internal/config"
	"sfera-server/pkg/errors"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	DB *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context) (*Postgres, error) {

	var err error

	pgOnce.Do(func() {

		var (
			pgxCfg *pgxpool.Config
			db     *pgxpool.Pool
		)

		cfg := config.GetConfig(ctx).PSQL
		dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

		if pgxCfg, err = pgxpool.ParseConfig(dsn); err != nil {
			err = errors.Wrap(err)
			return
		}

		// TODO: обязательно разобраться с логером запросов к базе
		// pgxCfg.ConnConfig.Logger = logrusadapter.NewLogger(logger)

		if db, err = pgxpool.ConnectConfig(ctx, pgxCfg); err != nil {
			err = errors.Wrap(err)
			return
		}

		pgInstance = &Postgres{db}

	})

	if err != nil {
		return nil, err
	}

	err = pgInstance.DB.Ping(ctx)

	return pgInstance, errors.Wrap(err)

}

func (pg *Postgres) Close() {

	pgInstance.DB.Close()

}
