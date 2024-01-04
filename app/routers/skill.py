from fastapi import APIRouter
from app.database import (
    database,
    skill_table
)

router = APIRouter()

@router.get("/skill/list")
async def list_skills():
    query = skill_table.select()
    result = await database.fetch_all(query)
    return result
