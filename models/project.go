package models

import (
	"wallnut-studios.com/rest-api/cmd"
)

type Project struct {
	ID        int64  `binding:"required"`
	Name      string `binding:"required"`
	TypeId    int64  `binding:"required"`
	ProfileId int64  `binding:"required"`
}

var projects = []Project{}

func (p Project) Save() error {
	query := `
	INSERT INTO events(name, type_id, profile_id) 
	VALUES (?, ?, ?)`
	stmt, err := cmd.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(p.Name, p.TypeId, p.ProfileId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	p.ID = id
	return err
}

func GetAllProjects() ([]Project, error) {
	query := "SELECT * FROM t_project"
	rows, err := cmd.DB.Query(query)
	if err != nil {
		return nil, err
	}
	rows.Close()
	for rows.Next() {
		var project Project
		err := rows.Scan(&project.ID, &project.Name, &project.TypeId, &project.ProfileId)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func GetProjectById(id int64) (*Project, error) {
	query := "SELECT * FROM t_project WHERE id = ?"
	row := cmd.DB.QueryRow(query, id)
	var project Project
	err := row.Scan(&project.ID, &project.Name, &project.TypeId, &project.ProfileId)
	if err != nil {
		return nil, err
	}
	return &project, err
}

func (project Project) Update() error {
	query := `
		UPDATE t_events
		SET name = ?, type_id = ?, profile_id = ?
		WHERE id = ?
	`
	stmt, err := cmd.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(project.Name, project.TypeId, project.ProfileId, project.ID)
	return err
}

func (project Project) Delete() error {
	query := "DELETE FROM t_project WHERE id = ?"
	stmt, err := cmd.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(project.ID)
	return err
}
