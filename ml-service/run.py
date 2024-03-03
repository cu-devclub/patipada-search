from app import create_app
import os

# Load environment variables from app.env
from dotenv import load_dotenv
load_dotenv('app.env')

# Get the environment name from the 'FLASK_ENV' environment variable, default to 'development'
env = os.getenv('FLASK_ENV', 'development')
# Create an instance of the Flask app with the corresponding configuration
app = create_app(env)

if __name__ == "__main__":
    # Run the Flask app with the specified port
    app.run(port=app.config['PORT'],debug=app.config['DEBUG'])

# TODO : tokenized the text apis => Done
# TODO : remove stop words from text apis => Done
# TODO : LDA => Done