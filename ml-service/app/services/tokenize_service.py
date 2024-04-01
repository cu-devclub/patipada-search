from typing import List
from pythainlp.tokenize import word_tokenize

class TokenizeService:
    """
    A service class for tokenizing Thai text into words.
    """

    @staticmethod
    def tokenize_from_string(text: str) -> List[str]:
        """
        Tokenize Thai text into words.

        Args:
            text (str): The Thai text to tokenize.

        Returns:
            List[str]: A list of tokens (words) from the text.
        """
        tokens = word_tokenize(text, engine='newmm')
        # Remove any ' ' (space) tokens
        tokens = list(filter(lambda x: x != ' ', tokens))
        return tokens