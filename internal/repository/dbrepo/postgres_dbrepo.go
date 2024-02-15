package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"wallnut-studios.com/rest-api/internal/models"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) GetSkillTags() ([]*models.SkillTag, error) {
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
	var skill_tags []*models.SkillTag
	for rows.Next() {
		var skill_tag models.SkillTag
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

func (m *PostgresDBRepo) GetSkills() ([]*models.Skill, error) {
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

func (m *PostgresDBRepo) GetProjects(typeId ...int) ([]*models.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	where := ""
	if len(typeId) > 0 {
		where = fmt.Sprintf("where type_id = %d ", typeId[0])
	}
	query := fmt.Sprintf(`
		select
			id, name, type_id, profile_id, cover, link
		from
			t_project %s
		order by
		    id asc
	`, where)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []*models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.TypeId,
			&project.ProfileId,
			&project.Cover,
			&project.Link,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}
	return projects, err
}
