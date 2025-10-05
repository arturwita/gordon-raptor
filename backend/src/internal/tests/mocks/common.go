package tests_mocks

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MockTimestamp = primitive.NewDateTimeFromTime(time.Date(2025, 5, 13, 12, 0, 0, 0, time.UTC)) // 2025-05-13T12:00:00Z
