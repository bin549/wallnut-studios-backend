package pojo

type Music struct {
}

type Write struct {
}

type SkillTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Skill struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	TagId int    `json:"tag_id"`
}

type ProjectType struct {
}

type ProjectSkill struct {
}

type Project struct {
	ID        int64  `binding:"required"`
	Name      string `binding:"required"`
	TypeId    int64  `binding:"required"`
	ProfileId int64  `binding:"required"`
}

var projects = []Project{}

type Profile struct {
}

type MusicSkill struct {
}
