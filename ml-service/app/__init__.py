from flask import Flask
from config import config_map

def create_app(env):
    # Create Flask app instance
    app = Flask(__name__)
   # Load configuration based on environment
    app.config.from_object(config_map.get(env))

    # Register blueprints (API endpoints)
    from app.api import lda, stop_word, tokenize
    app.register_blueprint(lda.bp)
    app.register_blueprint(stop_word.bp)
    app.register_blueprint(tokenize.bp)
    return app
