import unittest

from app.services.lda_service import LDAServerice

class TestLDAServerice(unittest.TestCase):
    def test_perform_lda(self):
        num_topics = 5
        lda_vector = LDAServerice.perform_lda(["สวัสดี ไทย", "ไทย สวัสดี"], num_topics)
        # Check that the output is not empty
        self.assertTrue(lda_vector, "LDA vector is empty")
