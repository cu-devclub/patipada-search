from flask import Blueprint, jsonify,request
from app.services.stopWords_service import StopWordsService
from app.services.tokenize_service import TokenizeService
bp = Blueprint('stop_word', __name__)

@bp.route('/remove-stopWords-from-text', methods=['GET'])
def remove_stopWords_from_text():
    """
    Remove stop words from a query string 

    This endpoint accepts two query parameters:
    - 'query': The string or list of tokens from which to remove stop words.

    The endpoint returns a JSON response with a 'result' key. The value of the 'result' key is
    the 'query' parameter with stop words removed.

    Args:
        query (str): The string or list of tokens from which to remove stop words.
        from (str): The type of the 'query' parameter. It can be 'list' or 'string'.

    Returns:
        dict: A dictionary with a 'result' key and the 'query' parameter with stop words removed as the value.
    """
    query = request.args.get('query', default = "", type = str)
    tokens = TokenizeService.tokenize_from_string(query)
    result = StopWordsService.remove_stop_words_from_list(tokens)

    return jsonify({"result": result})

@bp.route('/remove-stopWords-from-list', methods=['GET'])
def remove_stopWords_from_list():
    """
    Remove stop words from a list of tokens.

    This endpoint accepts a single query parameter:
    - 'query': The list of tokens from which to remove stop words.

    The endpoint returns a JSON response with a 'result' key. The value of the 'result' key is
    the list of tokens with stop words removed.

    Args:
        query (list): The list of tokens from which to remove stop words.

    Returns:
        dict: A dictionary with a 'result' key and the list of tokens with stop words removed as the value.
    """
    query = request.args.get('query', default = "", type = str)
    tokens = query.split(',')
    result = StopWordsService.remove_stop_words_from_list(tokens)

    return jsonify({"result": result})