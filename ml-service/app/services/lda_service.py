import numpy as np
from typing import List
from app.services.stopWords_service import StopWordsService
from app.services.tokenize_service import TokenizeService
import pickle

class LDAServerice:
    id2word = pickle.load(open('model/id2word.pkl', 'rb'))
    lda_model = pickle.load(open('model/LDA_model.pkl', 'rb'))

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
        vector = [0]*len(LDAServerice.id2word)
        for word in text:
            vector[LDAServerice.id2word[word]] += 1
        vector = np.array(vector)

        # Perform LDA
        lda_vector = LDAServerice.lda_model.transform(vector).tolist()

        return lda_vector[0]
    