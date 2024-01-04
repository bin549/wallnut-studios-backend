import databases
import sqlalchemy
from app.config import config

metadata = sqlalchemy.MetaData()

profile_table = sqlalchemy.Table(
    "t_profile",
    metadata,
    sqlalchemy.Column("id", sqlalchemy.Integer, primary_key=True, autoincrement=True),
    sqlalchemy.Column("name", sqlalchemy.String),
    sqlalchemy.Column("birthday", sqlalchemy.Date)
)

skill_tag_table = sqlalchemy.Table(
    "t_skill_tag",
    metadata,
    sqlalchemy.Column("id", sqlalchemy.Integer, primary_key=True, autoincrement=True),
    sqlalchemy.Column("name", sqlalchemy.String)
)

skill_table = sqlalchemy.Table(
    "t_skill",
    metadata,
    sqlalchemy.Column("id", sqlalchemy.Integer, primary_key=True, autoincrement=True),
    sqlalchemy.Column("name", sqlalchemy.String),
    sqlalchemy.Column("tag_id", sqlalchemy.ForeignKey("t_skill_tag.id"), nullable=False)
)

project_type_table = sqlalchemy.Table(
    "t_project_type",
    metadata,
    sqlalchemy.Column("id", sqlalchemy.Integer, primary_key=True, autoincrement=True),
    sqlalchemy.Column("name", sqlalchemy.String)
)

project_table = sqlalchemy.Table(
    "t_project",
    metadata,
    sqlalchemy.Column("id", sqlalchemy.Integer, primary_key=True, autoincrement=True),
    sqlalchemy.Column("name", sqlalchemy.String),
    sqlalchemy.Column("type_id", sqlalchemy.ForeignKey("t_project_type.id"), nullable=False),
    sqlalchemy.Column("profile_id", sqlalchemy.ForeignKey("t_profile.id"), nullable=False)
)

music_table = sqlalchemy.Table(
    "t_music",
    metadata,
    sqlalchemy.Column("id", sqlalchemy.Integer, primary_key=True, autoincrement=True),
    sqlalchemy.Column("name", sqlalchemy.String),
    sqlalchemy.Column("profile_id", sqlalchemy.ForeignKey("t_profile.id"), nullable=False)
)

write_table = sqlalchemy.Table(
    "t_write",
    metadata,
    sqlalchemy.Column("id", sqlalchemy.Integer, primary_key=True, autoincrement=True),
    sqlalchemy.Column("name", sqlalchemy.String),
    sqlalchemy.Column("profile_id", sqlalchemy.ForeignKey("t_profile.id"), nullable=False)
)

music_skill_table = sqlalchemy.Table(
    "t_music_skill",
    metadata,
    sqlalchemy.Column("id", sqlalchemy.Integer, primary_key=True, autoincrement=True),
    sqlalchemy.Column("music_id", sqlalchemy.ForeignKey("t_music.id"), nullable=False),
    sqlalchemy.Column("skill_id", sqlalchemy.ForeignKey("t_skill.id"), nullable=False)
)

project_skill_table = sqlalchemy.Table(
    "t_project_skill",
    metadata,
    sqlalchemy.Column("id", sqlalchemy.Integer, primary_key=True, autoincrement=True),
    sqlalchemy.Column("project_id", sqlalchemy.ForeignKey("t_project.id"), nullable=False),
    sqlalchemy.Column("skill_id", sqlalchemy.ForeignKey("t_skill.id"), nullable=False)
)

engine = sqlalchemy.create_engine(
    config.DATABASE_URL
)

# metadata.drop_all(engine)
metadata.create_all(engine)
database = databases.Database(
    config.DATABASE_URL, force_rollback=config.DB_FORCE_ROLL_BACK
)
