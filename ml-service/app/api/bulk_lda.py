from flask import Blueprint,jsonify
from app.services.lda_service import LDAServerice

bp = Blueprint('bulk_lda', __name__)

@bp.route('/bulk_lda', methods=['POST'])
def perform_bulk_lda():
    bulk_lda_vectors = LDAServerice.perform_bulk_lda()
    return jsonify({'bulk_lda_vectors': bulk_lda_vectors})
