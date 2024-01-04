from fastapi import FastAPI, HTTPException
from contextlib import asynccontextmanager
from app.database import database


@asynccontextmanager
async def lifespan(app: FastAPI):
    await database.connect()
    yield
    await database.disconnect()



app = FastAPI(lifespan=lifespan)