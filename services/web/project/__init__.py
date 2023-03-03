from flask import Flask
from .api.auth.views import auth_namespace
from .api.orders.views import order_namespace
from .config.config import Config
from .models.users import User
from .models.orders import Order
from .utils.db import db
from flask_restx import Api
from flask_jwt_extended import JWTManager
from flask_migrate import Migrate
from flask_cors import CORS


def create_app():
    app = Flask(__name__)
    app.config.from_object(Config)
    authorizations = {
        "Bearer Auth": {
            'type': "apiKey",
            'in': 'header',
            'name': "Authorization",
            'description': "Add a JWT with ** Bearer &lt;JWT&gt; to authorize"
        }
    }
    api = Api(app,
              title="Pizza Delivery API",
              description="A REST API for a Pizza Delievry service",
              authorizations=authorizations,
              security="Bearer Auth"
              )
    api.add_namespace(order_namespace)
    api.add_namespace(auth_namespace, path='/auth')
    db.init_app(app)
    jwt = JWTManager(app)
    migrate = Migrate(app, db)
    CORS(app)

    @app.shell_context_processor
    def make_shell_context():
        return {
            'db': db,
            'User': User,
            'Order': Order,
        }

    return app
