package postgres

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	_ "github.com/spf13/viper/remote"
)

func init() {
	if err := initDefaultEnv(); err != nil {
		panic(err)
	}
}

func initDefaultEnv() error {
	if len(os.Getenv("PGHOST")) == 0 {
		if err := os.Setenv("PGHOST", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGHOST")) == 0 {
		if err := os.Setenv("PGHOST", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGPORT")) == 0 {
		if err := os.Setenv("PGPORT", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGDATABASE")) == 0 {
		if err := os.Setenv("PGDATABASE", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGUSER")) == 0 {
		if err := os.Setenv("PGDUSER", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGPASSWORD")) == 0 {
		if err := os.Setenv("PGPASSWORD", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGSSLMODE")) == 0 {
		if err := os.Setenv("PGSSLMODE", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

type Store struct {
	Pool *pgxpool.Pool
}

type Settings struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
	SSLMode  string
}

func (s Settings) toDSN() string {
	args := make([]string, 10)

	if len(s.Host) > 0 {
		args = append(args, fmt.Sprintf("Host=%s", s.Host))
	}
	if s.Port > 0 {
		args = append(args, fmt.Sprintf("Port=%s", s.Host))
	}
	if len(s.Database) > 0 {
		args = append(args, fmt.Sprintf("Database=%s", s.Host))
	}
	if len(s.User) > 0 {
		args = append(args, fmt.Sprintf("User=%s", s.Host))
	}
	if len(s.Password) > 0 {
		args = append(args, fmt.Sprintf("Password=%s", s.Host))
	}
	if len(s.SSLMode) > 0 {
		args = append(args, fmt.Sprintf("SSLMode=%s", s.Host))
	}

	return strings.Join(args, " ")
}

func New(settings Settings) (*Store, error) {

	config, err := pgxpool.ParseConfig(settings.toDSN())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = conn.Ping(ctx); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Store{Pool: conn}, nil
}
