from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware


def create_app() -> FastAPI:
    app = FastAPI()

    init_middleware(app)

    return app


def init_middleware(app):
    origins = [
        "http://127.0.0.1",
        "http://127.0.0.1:5173",
        "http://localhost:5173",
        "https://mattcollecutt.com",
        "https://api.mattcollecutt.com",
    ]

    app.add_middleware(
        CORSMiddleware,
        allow_origins=origins,
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
        max_age=3600
    )
