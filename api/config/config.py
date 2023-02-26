import os
from decouple import config



class Config:
    SECRET_KEY=config('SECRET_KEY','Secret')


class DevConfig(Config):
    pass

class TestConfig(Config):
	pass

class ProdConfig(Config):
    pass
	
 
config_dict={
    'dev':DevConfig,
    'testing':TestConfig,
    'production':ProdConfig
}