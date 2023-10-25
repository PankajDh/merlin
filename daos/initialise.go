package daos

import (
	"context"
	"database/sql"
	"merlin/utils"
	"net/url"
	"os"
	"path"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
)

type Daos struct {
	DB *sql.DB
}

func Initialise(ctx context.Context) *Daos {
	logger := utils.GetLogger(ctx)
	db, err := configureDB(ctx)
	if err != nil {
		logger.Warnln("Error  while configuring the DB")
		db = &sql.DB{}
	}

	return &Daos{
		DB: db,
	}
}

func configureDB(ctx context.Context) (*sql.DB, error) {
	logger := utils.GetLogger(ctx)

	connectionString := ""
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logger.WithError(err).Errorln("Error while connecting to the DB")
		return nil, err
	}

	pingError := db.Ping()
	if pingError != nil {
		logger.WithError(err).Errorln("Error while pinging to the DB")
		return nil, err
	}

	logger.Infoln("DB connected")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.WithError(err).Errorln("Error while finding db driver for migration")
		return nil, err
	}

	pwd, err := os.Getwd()
	if err != nil {
		logger.WithError(err).Errorln("Error while getting pwd")
		return nil, err
	}

	base, err := url.Parse(pwd)
	if err != nil {
		logger.WithError(err).Errorln("Error while creating base path for migrations folder")
		return nil, err
	}
	base.Path = path.Join(base.Path, "migrations")

	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+base.String(),
		"postgres",
		driver,
	)
	if err != nil {
		logger.WithError(err).Errorln("Error while initialising the migration tool")
		return nil, err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		logger.WithError(err).Errorln("Error while running the migration")
		return nil, err
	}

	_, err = db.Exec("\\c merlin")
	if err != nil {
		logger.WithError(err).Errorln("Error while connecting to the database inside the main DB")
		return nil, err
	}

	return db, nil
}
