package postgres

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"vendetta/internal/app/config"
	"vendetta/pkg/utils"
)

type Database = gorm.DB
type Connection = sql.DB

type Service struct {
	c *config.Config
	l *utils.Logger

	conn     *Connection
	database *Database

	traceName string
}

func (s *Service) Init() error {
	s.l.DebugT(s.traceName, "Initializing postgres service")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	url := utils.NewURLConnectionString("postgres", fmt.Sprintf("%s:%d", s.c.PostgresHost, s.c.PostgresPort), "", s.c.PostgresDatabase, s.c.PostgresUser, s.c.PostgresPassword)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{TranslateError: true, Logger: newLogger})
	if err != nil {
		s.l.ErrorT(s.traceName, "Failed to connect to postgres database", err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		s.l.ErrorT(s.traceName, "Failed to connect to postgres database", err)
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		s.l.ErrorT(s.traceName, "Failed to ping postgres database", err)
		return err
	}

	s.database = db
	s.conn = sqlDB

	s.l.DebugT(s.traceName, "Postgres service initialized")

	return nil
}

func (s *Service) Disconnect() error {
	if err := s.conn.Close(); err != nil {
		s.l.ErrorT(s.traceName, "Failed to close connection", err)
	}

	return nil
}

func (s *Service) GetConnect() *Connection {
	return s.conn
}

func (s *Service) GetDatabase() *Database {
	return s.database
}

func NewPostgresService(c *config.Config, l *utils.Logger) *Service {
	return &Service{
		c: c,
		l: l,

		traceName: "[PostgresService]",
	}
}
