from typing import List
from gensim import corpora, models

class LDAServerice:
    @staticmethod
    def perform_lda(documents: List[str],num_topics: int) -> List[float]:
        """
        Perform LDA on a list of documents.

        Args:
            documents (List[str]): The documents to perform LDA on.

        Returns:
            List[float]: The LDA vector for the documents.
        """
        if len(documents) == 0:
            print("No data provided")
            return []
        # Tokenize the documents
        texts = [[word for word in document.lower().split()] for document in documents]

        # Create a dictionary from the texts
        dictionary = corpora.Dictionary(texts)

        # Create a corpus from the dictionary
        corpus = [dictionary.doc2bow(text) for text in texts]

        # Perform LDA
        
        lda_model = models.LdaModel(corpus, num_topics=num_topics, id2word=dictionary, passes=15)

        # Get the LDA vector for the first document
        sparse_lda_vector = lda_model[corpus[0]]

        # Convert the sparse LDA vector to a dense vector
        dense_lda_vector = [0.0] * num_topics
        for index, value in sparse_lda_vector:
            dense_lda_vector[index] = value

        return dense_lda_vector