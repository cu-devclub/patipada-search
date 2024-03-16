from flask import Blueprint, jsonify,request
from app.services.tokenize_service import TokenizeService

bp = Blueprint('tokenize', __name__)

@bp.route('/tokenize', methods=['GET'])
def tokenize():
    """
    Tokenize a Thai text string into words.

    This endpoint accepts a single query parameter:
    - 'query': The Thai text string to tokenize.

    The endpoint returns a JSON response with a 'result' key. The value of the 'result' key is
    the tokenized version of the 'query' parameter.

    Args:
        query (str): The Thai text string to tokenize.

    Returns:
        dict: A dictionary with a 'result' key and the tokenized version of the 'query' parameter as the value.
    """
    query = request.args.get('query', default = "", type = str)
    result = TokenizeService.tokenize_from_string(query)
    return jsonify({"result": result})