import unittest
from app.services.tokenize_service import TokenizeService

class TestTokenizeService(unittest.TestCase):
    def test_tokenize(self):
        text = "สวัสดีครับ ผมชื่อสปาย"
        tokens = TokenizeService.tokenize_from_string(text)
        self.assertEqual(tokens, ["สวัสดี", "ครับ", "ผม", "ชื่อ", "สปาย"])

if __name__ == '__main__':
    unittest.main()