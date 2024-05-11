package util

import (
	"testing"
)

func TestUtil(t *testing.T) {

	t.Run("Test Success DecodeHTMLText", func(t *testing.T) {
		encodedString := "Hello, &quot;World&quot;"
		expected := "Hello, \"World\""
		result := DecodeHTMLText(encodedString)
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})

	t.Run("Test Success ExtractRawStringFromHTMLTags", func(t *testing.T) {
		encodedString := "<p>Hello, <b>World</b></p>"
		expected := "Hello, World"
		result := ExtractRawStringFromHTMLTags(encodedString)
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})

	t.Run("Test success : decode HTML then extract raw string (usecase)", func(t *testing.T) {
		encodedString := "&lt;p&gt;ทดสอบ พระขอเงิน แล้วเราให้พระ เราจะบาปไหม&lt;/p&gt;"
		expected := "ทดสอบ พระขอเงิน แล้วเราให้พระ เราจะบาปไหม"
		result := ExtractRawStringFromHTMLTags(DecodeHTMLText(encodedString))
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})
}
