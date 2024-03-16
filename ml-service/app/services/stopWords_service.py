from typing import List
from pythainlp.corpus.common import thai_stopwords

class StopWordsService:
    """
    A service class for removing Thai stop words from a list of tokens.
    """

    @staticmethod
    def remove_stop_words_from_list(tokens: List[str]) -> List[str]:
        """
        Remove Thai stop words from a list of tokens.

        Args:
            tokens (list): A list of tokens from which to remove stop words.

        Returns:
            list: A new list of tokens with stop words removed.
        """
        stop_words = thai_stopwords()
        return list(filter(lambda x: x not in stop_words, tokens))