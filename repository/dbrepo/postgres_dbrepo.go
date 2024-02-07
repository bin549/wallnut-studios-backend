package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"wallnut-studios.com/rest-api/pojo"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) GetSkillTags() ([]*pojo.SkillTag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := fmt.Sprintf(`
		select 
			id, name
		from 
			t_skill_tag
	`)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var skill_tags []*pojo.SkillTag
	for rows.Next() {
		var skill_tag pojo.SkillTag
		err := rows.Scan(
			&skill_tag.ID,
			&skill_tag.Name,
		)
		if err != nil {
			return nil, err
		}
		skill_tags = append(skill_tags, &skill_tag)
	}
	return skill_tags, err
}

func (m *PostgresDBRepo) GetSkills() ([]*pojo.Skill, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := fmt.Sprintf(`
		select 
			id, name, tag_id
		from
			t_skill
		order by
		    id asc
	`)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var skills []*pojo.Skill
	for rows.Next() {
		var skill pojo.Skill
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
