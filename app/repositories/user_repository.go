package repositories

import "besrabasant/fiber/graphql/model"

type UserRepository struct{}

func (*UserRepository) GetAll() ([]*model.User, error) {
	return nil, nil
}

func (*UserRepository) Add(user *model.NewUser) (*model.User, error) {
	return nil, nil
}
