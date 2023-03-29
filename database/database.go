package database

import (
	_ "embed"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"maunium.net/go/maulogger/v2"

	"maunium.net/go/mautrix/util/dbutil"

	"go.mau.fi/mautrix-signal/database/upgrades"
)

type Database struct {
	*dbutil.Database

	User     *UserQuery
	Portal   *PortalQuery
}

func New(baseDB *dbutil.Database, log maulogger.Logger) *Database {
	db := &Database{Database: baseDB}
	db.UpgradeTable = upgrades.Table
	db.User = &UserQuery{
		db:  db,
		log: log.Sub("User"),
	}
	db.Portal = &PortalQuery{
		db:  db,
		log: log.Sub("Portal"),
	}
	return db
}

func strPtr(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}