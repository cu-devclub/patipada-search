import os

class Config:
    # Define default configurations here
    DEBUG = False
    PORT = int(os.getenv('FLASK_RUN_PORT', 8084))  # Get port number from environment variable, default to 8084

class DevelopmentConfig(Config):
    DEBUG = True

class ProductionConfig(Config):
    DEBUG = False
    # Additional production configurations can be added here

# Define dictionary to map configuration names to configuration classes
config_map = {
    'development': DevelopmentConfig,
    'production': ProductionConfig,
}
