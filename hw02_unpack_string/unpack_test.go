package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "Case-1: Base case with numbers",
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			desc:     "Case-2: Base case without numbers",
			input:    "abccd",
			expected: "abccd",
		},
		{
			desc:     "Case-3: Empty string",
			input:    "",
			expected: "",
		},
		{
			desc:     "Case-4: Zero number in string",
			input:    "aaa0b",
			expected: "aab",
		},
		{
			desc:     "Case-5: Zero number to empty string",
			input:    "a0",
			expected: "",
		},
		{
			desc:     "Case-6: Complex bytes runes in string without numbers",
			input:    "你好世界",
			expected: "你好世界",
		},
		{
			desc:     "Case-7: Complex bytes runes in string with numbers",
			input:    "你1好2世0界",
			expected: "你好好界",
		},
		{
			desc:     "Case-8: Special characters in string",
			input:    "d\n5abc",
			expected: "d\n\n\n\n\nabc",
		},
		{
			desc:     "Case-9: Symbol characters string",
			input:    "-5$@#3~'1`2",
			expected: "-----$@###~'``",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "0"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
