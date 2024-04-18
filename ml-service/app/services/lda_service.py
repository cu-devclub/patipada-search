from typing import List
from gensim import corpora, models
from app.services.stopWords_service import StopWordsService
from app.services.tokenize_service import TokenizeService
import pickle

class LDAServerice:
    id2word = pickle.load(open(f'model\id2word.pkl', 'rb'))
    lda_model = pickle.load(open(f'model\LDA_model.pkl', 'rb'))

    @staticmethod

    def perform_lda(document: str) -> List[float]:
        """
        Perform LDA on a list of documents.

        Args:
            documents (List[str]): The documents to perform LDA on.

        Returns:
            List[float]: The LDA vector for the documents.
        """
        if len(document) == 0:
            print("No data provided")
            return []
        # Tokenize the documents
        text = StopWordsService.remove_stop_words_from_list(TokenizeService.tokenize_from_string(document))

        # word2vec
        vector = LDAServerice.id2word.doc2bow(text)

        # Perform LDA
        sparse_lda_vector = LDAServerice.lda_model.get_document_topics(vector)

        # Convert the sparse LDA vector to a dense vector
        dense_lda_vector = [0.0] * 30
        for index, value in sparse_lda_vector:
            dense_lda_vector[index] = value

        return dense_lda_vector
    