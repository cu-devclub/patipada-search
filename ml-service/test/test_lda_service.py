import unittest

from app.services.lda_service import LDAServerice

class TestLDAServerice(unittest.TestCase):
    def test_perform_lda(self):
        lda_vector = LDAServerice.perform_lda(["สวัสดี ไทย", "ไทย สวัสดี"])
        # Check that the output is not empty
        self.assertTrue(lda_vector, "LDA vector is empty")

if __name__ == '__main__':
    unittest.main()