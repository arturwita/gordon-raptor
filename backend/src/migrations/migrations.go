package migrations

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Migration func(db *mongo.Database) error

type MigrationRecord struct {
	Name      string    `bson:"name"`
	AppliedAt time.Time `bson:"appliedAt"`
}

var MigrationsMap = map[string]Migration{
	"001_unique_user_email": CreateUsersEmailUniqueIndex,
}
