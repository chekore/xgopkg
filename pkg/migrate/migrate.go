//Package migrate for migrate database
package migrate

import (
	"path/filepath"
	// "database/sql"
	"xgopkg.com/xgopkg/pkg/config"

	"gopkg.in/logger.v1"

	mg "github.com/mattes/migrate"
	"github.com/mattes/migrate/source/go-bindata"
	// "github.com/mattes/migrate/source/go-bindata/examples/migrations"
)

//Migrate start migrations
func Migrate() {
	dbURL := config.GetMySQLURL()
	names, err := AssetDir("migrations")
	if err != nil {
		log.Error(err)
	}
	s := bindata.Resource(names,
		func(name string) ([]byte, error) {
			log.Info(name)
			return Asset(filepath.Join("migrations", name))
		})

	d, err := bindata.WithInstance(s)
	if err != nil {
		log.Error(err)
		//todo
	}

	log.Debug("database url: ", dbURL)

	// migrate.NewWithD
	m, err := mg.NewWithSourceInstance("go-bindata", d, "mysql://"+dbURL)
	if err != nil {
		log.Error(err)
		//todo
	}
	m.Up() // run your migrations and handle the errors above of course
}
