package repository

import (
	"database/sql"
	"wallnut-studios.com/rest-api/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetSkills() ([]*models.Skill, error)
	GetSkillTags() ([]*models.SkillTag, error)
	GetProjects(typeId ...int) ([]*models.Project, error)
}
