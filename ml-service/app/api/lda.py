from flask import Blueprint,jsonify
from app.services.lda_service import LDAServerice

bp = Blueprint('lda', __name__)

@bp.route('/lda', methods=['POST'])
def perform_lda():
    lda_vector = LDAServerice.perform_lda()
    return jsonify({'lda_vector': lda_vector})
