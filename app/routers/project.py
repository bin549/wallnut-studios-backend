from fastapi import APIRouter
from app.database import (
    database,
    project_table
)


router = APIRouter()


@router.get("/project/list")
async def list_project():
    query = project_table.select()
    result = await database.fetch_all(query)
    return result