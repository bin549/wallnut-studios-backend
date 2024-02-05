package repository

import (
	"database/sql"
	"wallnut-studios.com/rest-api/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetSkills() ([]*models.Skill, error)
}
