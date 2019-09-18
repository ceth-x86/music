package dbutils

import (
	"errors"
	"time"

	"github.com/demas/music/internal/models/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/zap"
)

func OpenDbConnection(connectionString string, logSql bool) (*gorm.DB, error) {

	logger := zap.NewExample().Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	database, err := getClient(connectionString)
	if err != nil {
		return nil, err
	}

	database.DB().SetConnMaxLifetime(5 * time.Minute)
	database.LogMode(logSql)
	applyAutoMigrations(database)
	return database, nil
}

func applyAutoMigrations(dbs *gorm.DB) {
	dbs.AutoMigrate(
		&db.Artist{},
		&db.Album{},
		&db.Playlist{},
		&db.Track{},
	)
}

func getClient(connectionString string) (*gorm.DB, error) {

	logger := zap.NewExample().Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	ticker := time.NewTicker(1 * time.Nanosecond)
	timeout := time.After(15 * time.Minute)
	seconds := 1
	for {
		select {
		case <-ticker.C:
			ticker.Stop()

			client, err := gorm.Open("postgres", connectionString)
			if err != nil {
				logger.With(zap.Error(err)).Warn("не удалось установить соединение с PostgreSQL")

				ticker = time.NewTicker(time.Duration(seconds) * time.Second)
				seconds *= 2
				if seconds > 60 {
					seconds = 60
				}

				continue
			}

			logger.Debug("соединение с PostgreSQL успешно установлено")
			return client, nil
		case <-timeout:
			return nil, errors.New("PostgreSQL: connection timeout")
		}
	}
}
