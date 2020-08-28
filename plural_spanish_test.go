package plural_spanish

import (
	"testing"
)

type pluralSingular struct {
	singularWord string
	pluralWord   string
}

var tests = []pluralSingular{
	{"tabú", "tabúes"},
	{"jabalí", "jabalíes"},
	{"camión", "camiones"},
	{"leche", "leches"},
	{"carne", "carnes"},
	{"sol", "soles"},
	{"luz", "luces"},
}

func Test_pluralWords_IsPlural(t *testing.T) {
	t.Run("when is plural-spanish return true", func(t *testing.T) {
		pluralWords := NewPluralWords()

		for _, tt := range tests {
			value := pluralWords.ToPlural(tt.singularWord)
			expectedValue := tt.pluralWord
			if pluralWords.IsPlural(value) {
				t.Errorf("expected: %v but got :%v", expectedValue, value)
			}
		}

	})
}

func Test_pluralWords_ToPlural(t *testing.T) {
	t.Run("when words is singular is converted to plural-spanish successfully", func(t *testing.T) {
		pluralWords := NewPluralWords()

		for _, tt := range tests {
			value := pluralWords.ToPlural(tt.singularWord)
			expectedValue := tt.pluralWord
			if value != tt.pluralWord {
				t.Errorf("expected: %s but got :%s", expectedValue, value)
			}
		}

	})
}

func Test_pluralWords_IsSingular(t *testing.T) {
	t.Run("when is singular return true", func(t *testing.T) {
		pluralWords := NewPluralWords()

		for _, tt := range tests {
			value := pluralWords.ToSingular(tt.pluralWord)
			expectedValue := pluralWords.IsSingular(value)
			if pluralWords.IsPlural(value) {
				t.Errorf("expected: %v but got :%v", expectedValue, value)
			}
		}
	})
}

func Test_pluralWords_ToSingular(t *testing.T) {
	t.Run("when words is plural-spanish is converted to singular successfully", func(t *testing.T) {
		pluralWords := NewPluralWords()

		for _, tt := range tests {
			value := pluralWords.ToSingular(tt.pluralWord)
			expectedValue := tt.singularWord
			if value != expectedValue {
				t.Errorf("expected: %s but got :%s", expectedValue, value)
			}
		}
	})
}
