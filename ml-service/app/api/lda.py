from flask import Blueprint, jsonify, request
from app.services.lda_service import LDAServerice

bp = Blueprint('lda', __name__)

@bp.route('/lda', methods=['GET'])
def perform_lda():
    # Get the list of strings from the request body
    query = request.args.get('query',default = "", type = str) 
   
    # Perform LDA on the documents
    lda_vector = LDAServerice.perform_lda(query)
    return_vector = [float(x) for x in lda_vector]

    return jsonify({'result': return_vector})

# Pre-data : run bulkLDA.py (called text2vec api inside) to generate the lda csv data
# Search : api route lda:  query => vector 

# All move external 
# 1. Tokenize -- include in text2vec
# 2. Remove stop words -- include in text2vec
# 3. Perform LDA // Get text2vec

# Still have stopword file to use for stopword removal in TF-IDF

# text2vec : external service called api
# multiple ways (LDA,BERT,GPT) and determine weight
# define in config

# user set weight in config file for each text2vec models
# same location as api URL for that model