from flask import Flask
from config import config_map

def create_app(env='development'):
    # Create Flask app instance
    app = Flask(__name__)
   # Load configuration based on environment
    app.config.from_object(config_map.get(env, 'development'))

    # Register blueprints (API endpoints)
    from app.api import bulk_lda, lda
    app.register_blueprint(bulk_lda.bp)
    app.register_blueprint(lda.bp)

    return app
