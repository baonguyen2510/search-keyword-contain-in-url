package postgresql

import (
	"fmt"
	"search-keyword-service/configs"
	"search-keyword-service/pkg/log"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Connection struct {
	ConnectionName string
	Host           string
	Port           int
	Username       string
	Password       string
	DatabaseName   string
	Schema         string
	MaxIdleConns   int
	MaxOpenConns   int
	MaxLifetime    time.Duration
}

const (
	// defaultMaxIdleConns
	// the default maximum number of connections in the idle connection pool.
	defaultMaxIdleConns = 10
	// defaultMaxOpenConns
	// the default maximum number of open connections to the database.
	defaultMaxOpenConns = 100
	// defaultConnMaxLifetime
	// the default maximum amount of time a connection may be reused.
	defaultConnMaxLifetime = time.Hour
)

// getPostgresSQLDataSourceName return data source name of mysql
// the format of dsn is
// host=[host] port=[port] user=[username] password=[password] dbname=[dbname]
func getPostgresSQLDialector(
	connectionName string, host string, port int,
	username string, password string,
	dbname string,
) gorm.Dialector {
	if connectionName != "" {
		dsn := fmt.Sprintf(
			"host=%s user=%s dbname=%s password=%s sslmode=disable",
			connectionName, username, dbname, password,
		)
		return postgres.New(postgres.Config{
			DriverName: "cloudsqlpostgres",
			DSN:        dsn,
		})
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname,
	)
	return postgres.Open(dsn)
}

func New(conn Connection) (*gorm.DB, error) {
	dialector := getPostgresSQLDialector(
		conn.ConnectionName, conn.Host, conn.Port,
		conn.Username, conn.Password,
		conn.DatabaseName,
	)
	if conn.Schema == "" {
		conn.Schema = "public"
	}
	logger := newLogger(log.Env(), log.Sensitive())
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conn.Schema + ".",
			SingularTable: true,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		return nil, err
	}

	if err := db.Use(
		otelgorm.NewPlugin(otelgorm.WithDBName(configs.Config.DbName)),
	); err != nil {
		return nil, fmt.Errorf("failed to instrument trace: %w", err)
	}

	if conn.MaxIdleConns <= 0 {
		conn.MaxIdleConns = defaultMaxIdleConns
	}
	if conn.MaxOpenConns <= 0 {
		conn.MaxOpenConns = defaultMaxOpenConns
	}
	if conn.MaxLifetime <= 0 {
		conn.MaxLifetime = defaultConnMaxLifetime
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Cannot get postgresql database: %v", err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(conn.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conn.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(conn.MaxLifetime)

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Cannot ping postgresql database")
	}
	return db, err
}
