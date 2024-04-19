from typing import List

class StopWordsService:
    """
    A service class for removing Thai stop words from a list of tokens.
    """

    with(open('stopword/pythainlp-corpus-stopwords_th.txt', 'rb')) as f:
        stopWords = set(f.read().decode('utf-8').splitlines())

    @staticmethod
    def remove_stop_words_from_list(tokens: List[str]) -> List[str]:
        """
        Remove Thai stop words from a list of tokens.

        Args:
            tokens (list): A list of tokens from which to remove stop words.

        Returns:
            list: A new list of tokens with stop words removed.
        """
        return list(filter(lambda x: x not in StopWordsService.stopWords, tokens))