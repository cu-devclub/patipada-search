import unittest
from app.services.stopWords_service import StopWordsService

class TestStopWordService(unittest.TestCase):
    def test_remove_stop_words(self):
        tokens = ["สวัสดี", "ครับ", "ผม", "ชื่อ", "สปาย"]
        result = StopWordsService.remove_stop_words_from_list(tokens)
        self.assertEqual(result, ["สวัสดี", "ผม", "ชื่อ", "สปาย"])
    
