class LDAServerice:
    @staticmethod
    def perform_lda():
        # Dummy implementation for LDA
        lda_vector = [0.1, 0.2, 0.3, 0.4]  # Dummy LDA vector
        return lda_vector

    @staticmethod
    def perform_bulk_lda():
        # Dummy implementation for bulk LDA
        bulk_lda_vectors = [
            {'document_id': 1, 'lda_vector': [0.1, 0.2, 0.3, 0.4]},
            {'document_id': 2, 'lda_vector': [0.2, 0.3, 0.4, 0.5]},
            {'document_id': 3, 'lda_vector': [0.3, 0.4, 0.5, 0.6]}
        ]  # Dummy bulk LDA vectors
        return bulk_lda_vectors
