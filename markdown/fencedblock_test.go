package markdown

import "testing"

func TestFencedBlock(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string // Description of the test case
		fence       string // Fenced block string
		expected    string // Expected fenced block string
	}{
		{
			description: "Empty",
			fence:       FencedBlock("").String(),
			expected:    "---\n\n---",
		},
		{
			description: "With content",
			fence:       FencedBlock("Content").String(),
			expected:    "---\nContent\n---",
		},
		{
			description: "With multiple lines of content",
			fence:       FencedBlock("Line 1\nLine 2\nLine 3").String(),
			expected:    "---\nLine 1\nLine 2\nLine 3\n---",
		},
		{
			description: "With metadata",
			fence:       FencedBlock("Content").WithMetadata([]string{"key1", "key2"}).String(),
			expected:    "--- key1 key2\nContent\n---",
		},
		{
			description: "With multiple lines of content and metadata",
			fence:       FencedBlock("Line 1\nLine 2\nLine 3").WithMetadata([]string{"key1", "key2"}).String(),
			expected:    "--- key1 key2\nLine 1\nLine 2\nLine 3\n---",
		},
		{
			description: "With custom fence",
			fence:       FencedBlock("Content").SetFence("~~~").String(),
			expected:    "~~~\nContent\n~~~",
		},
		{
			description: "With custom fence and metadata",
			fence:       FencedBlock("Content").SetFence("~~~").WithMetadata([]string{"key1", "key2"}).String(),
			expected:    "~~~ key1 key2\nContent\n~~~",
		},
		{
			description: "With custom fence and multiple lines of content",
			fence:       FencedBlock("Line 1\nLine 2\nLine 3").SetFence("~~~").String(),
			expected:    "~~~\nLine 1\nLine 2\nLine 3\n~~~",
		},
		{
			description: "With custom fence, multiple lines of content and metadata",
			fence:       FencedBlock("Line 1\nLine 2\nLine 3").SetFence("~~~").WithMetadata([]string{"key1", "key2"}).String(),
			expected:    "~~~ key1 key2\nLine 1\nLine 2\nLine 3\n~~~",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.fence != testCase.expected {
			t.Errorf("Expected:\n%s\n===\ngot:\n%s", testCase.expected, testCase.fence)
		}
	}

}
