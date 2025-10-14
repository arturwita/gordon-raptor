package migrations

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RunMigrations(db *mongo.Database) error {
	collection := db.Collection("migrations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for name, migration := range MigrationsMap {
		var record MigrationRecord
		err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&record)

		if err == nil {
			fmt.Printf("Migration %s already applied, skipping\n", name)
			continue
		} else if err != mongo.ErrNoDocuments {
			return fmt.Errorf("failed to check migration %s: %w", name, err)
		}

		if err := migration(db); err != nil {
			return fmt.Errorf("migration %s failed: %w", name, err)
		}

		_, err = collection.InsertOne(ctx, MigrationRecord{
			Name:      name,
			AppliedAt: time.Now(),
		})

		if err != nil {
			return fmt.Errorf("failed to record migration %s: %w", name, err)
		}
	}

	fmt.Printf("All migrations applied successfully\n")

	return nil
}
