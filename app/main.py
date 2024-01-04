from fastapi import FastAPI, HTTPException
from contextlib import asynccontextmanager
from app.database import database
from app.routers.skill import router as skill_router

@asynccontextmanager
async def lifespan(app: FastAPI):
    await database.connect()
    yield
    await database.disconnect()


app = FastAPI(lifespan=lifespan)

app.router.include_router(skill_router, prefix="/api")