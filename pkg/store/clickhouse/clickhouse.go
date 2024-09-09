package clickhouse

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/pkg/errors"
)

func init() {
	if err := initDefaultEnv(); err != nil {
		panic(err)
	}
}

func initDefaultEnv() error {
	if len(os.Getenv("CLICKHOUSE_HOST")) == 0 {
		if err := os.Setenv("CLICKHOUSE_HOST", "localhost"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("CLICKHOUSE_PORT")) == 0 {
		if err := os.Setenv("CLICKHOUSE_PORT", "9000"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("CLICKHOUSE_DATABASE")) == 0 {
		if err := os.Setenv("CLICKHOUSE_DATABASE", "default"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("CLICKHOUSE_USER")) == 0 {
		if err := os.Setenv("CLICKHOUSE_USER", "default"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("CLICKHOUSE_PASSWORD")) == 0 {
		if err := os.Setenv("CLICKHOUSE_PASSWORD", ""); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("CLICKHOUSE_SSLMODE")) == 0 {
		if err := os.Setenv("CLICKHOUSE_SSLMODE", "disable"); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

type Store struct {
	Conn clickhouse.Conn
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
	args := make([]string, 0)

	if len(s.Host) > 0 {
		args = append(args, fmt.Sprintf("host=%s", s.Host))
	}
	if s.Port > 0 {
		args = append(args, fmt.Sprintf("port=%d", s.Port))
	}
	if len(s.Database) > 0 {
		args = append(args, fmt.Sprintf("database=%s", s.Database))
	}
	if len(s.User) > 0 {
		args = append(args, fmt.Sprintf("username=%s", s.User))
	}
	if len(s.Password) > 0 {
		args = append(args, fmt.Sprintf("password=%s", s.Password))
	}
	if len(s.SSLMode) > 0 {
		args = append(args, fmt.Sprintf("sslmode=%s", s.SSLMode))
	}

	return strings.Join(args, "&")
}

func New(settings Settings) (*Store, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", settings.Host, settings.Port)},
		Auth: clickhouse.Auth{
			Database: settings.Database,
			Username: settings.User,
			Password: settings.Password,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = conn.Ping(ctx); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Store{Conn: conn}, nil
}