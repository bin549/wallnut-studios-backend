package models

import "wallnut-studios.com/rest-api/db"

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
	stmt, err := db.DB.Prepare(query)
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
	query := "SELECT * FROM projects"
	rows, err := db.DB.Query(query)
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
