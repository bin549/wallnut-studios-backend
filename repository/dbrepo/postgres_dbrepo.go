package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"wallnut-studios.com/rest-api/models"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) GetSkills() ([]*models.Skill, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := fmt.Sprintf(`
		select 
			id, name, tag_id
		from
			t_skill
	`)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var skills []*models.Skill
	for rows.Next() {
		var skill models.Skill
		err := rows.Scan(
			&skill.ID,
			&skill.Name,
			&skill.TagId,
		)
		if err != nil {
			return nil, err
		}
		skills = append(skills, &skill)
	}
	return skills, err
}
