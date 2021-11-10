package graphql

import (
	"besrabasant/fiber/app/repositories"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users *repositories.UserRepository
}
