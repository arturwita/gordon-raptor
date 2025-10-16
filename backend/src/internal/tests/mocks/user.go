package tests_mocks

import (
	"gordon-raptor/src/internal/domains/users"
	"gordon-raptor/src/pkg/db"
	"gordon-raptor/src/pkg/utils"
)

var MockUserId1 = "68dc4708c9b3f71d63510091"
var MockUserId2 = "68dc4708c9b3f71d63510092"
var MockUserId3 = "68dc4708c9b3f71d63510093"
var MockUserId4 = "68dc4708c9b3f71d63510094"
var MockUserId5 = "68dc4708c9b3f71d63510095"

var DefaultUserMock = users.UserModel{
	Id:        db.EnsureMongoId(MockUserId1),
	CreatedAt: MockTimestamp,
	UpdatedAt: MockTimestamp,
	Role:      users.UserRole,
	Email:     "john.doe@example.com",
	FirstName: utils.ToStrPointer("John"),
	LastName:  utils.ToStrPointer("Doe"),
	Picture:   utils.ToStrPointer("https://lh3.googleusercontent.com/a/ACg8ocL0h6eKQMyIpMdmarcTK78y2z_YyE8SJKxfLz3rmdyoOD0TFg=s96-c"),
}
