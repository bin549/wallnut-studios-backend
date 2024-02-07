package repository

import (
	"database/sql"
	"wallnut-studios.com/rest-api/pojo"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetSkills() ([]*pojo.Skill, error)
	GetSkillTags() ([]*pojo.SkillTag, error)
}
