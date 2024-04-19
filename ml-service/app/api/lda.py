from flask import Blueprint, jsonify, request
from app.services.lda_service import LDAServerice

bp = Blueprint('lda', __name__)

@bp.route('/lda', methods=['GET'])
def perform_lda():
    # Get the list of strings from the request body
    query = request.args.get('text',default = "", type = str) 
   
    # Perform LDA on the documents
    lda_vector = LDAServerice.perform_lda(query)
    return_vector = [float(x) for x in lda_vector]

    return jsonify({'name': 'lda','embedding': return_vector})