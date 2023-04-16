package services

import "gobitly/models"

type Gobitly interface {
	CreateGobitly(*models.Gobitly) error
	GetGobitly(*string) (*models.Gobitly, error)
	GetAllGobitlies() ([]*models.Gobitly, error)
	UpdateGobitly(*string) error
	DeleteGobitly(*string) error
}
