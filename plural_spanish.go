package plural_spanish

import (
	"github.com/gertd/go-pluralize"
)

type pluralWords struct {
	client *pluralize.Client
}

func NewPluralWords() *pluralWords {
	client := pluralize.NewClient()
	p := &pluralWords{client: client}
	p.registerPluralRules()
	p.registerSingularRules()
	return p
}

func (p *pluralWords) registerPluralRules() {
	p.client.AddPluralRule("(?i)y$", "yes")
	p.client.AddPluralRule("(?i)í$", "íes")
	p.client.AddPluralRule("(?i)ú$", "úes")
	p.client.AddPluralRule("(?i)a$", "as")
	p.client.AddPluralRule("(?i)e$", "es")
	p.client.AddPluralRule("(?i)i$", "is")
	p.client.AddPluralRule("(?i)o$", "os")
	p.client.AddPluralRule("(?i)u$", "us")
	p.client.AddPluralRule("(?i)d$", "des")
	p.client.AddPluralRule("(?i)s$", "ses")
	p.client.AddPluralRule("(?i)j$", "jes")
	p.client.AddPluralRule("(?i)l$", "les")
	p.client.AddPluralRule("(?i)r$", "res")
	p.client.AddPluralRule("(?i)z$", "ces")
	p.client.AddPluralRule("(?i)ón$", "ones")
}

func (p *pluralWords) registerSingularRules() {
	p.client.AddSingularRule(`(?i)yes$`, `y`)
	p.client.AddSingularRule("(?i)íes$", "í")
	p.client.AddSingularRule("(?i)úes$", "ú")
	p.client.AddSingularRule("(?i)as$", "a")
	p.client.AddSingularRule("(?i)nes$", "ne")
	p.client.AddSingularRule("(?i)ches$", "che")
	p.client.AddSingularRule("(?i)is$", "i")
	p.client.AddSingularRule("(?i)os$", "o")
	p.client.AddSingularRule("(?i)us$", "u")
	p.client.AddSingularRule("(?i)des$", "d")
	p.client.AddSingularRule("(?i)ses$", "s")
	p.client.AddSingularRule("(?i)jes$", "j")
	p.client.AddSingularRule("(?i)les$", "l")
	p.client.AddSingularRule("(?i)res$", "r")
	p.client.AddSingularRule("(?i)ces$", "z")
	p.client.AddSingularRule("(?i)ones$", "ón")
}

func (p *pluralWords) IsPlural(word string) bool {
	return p.client.IsPlural(word)
}

func (p *pluralWords) ToPlural(word string) string {
	value := p.client.Plural(word)
	return value
}

func (p *pluralWords) IsSingular(word string) bool {
	return p.client.IsSingular(word)
}

func (p *pluralWords) ToSingular(word string) string {
	value := p.client.Singular(word)
	return value
}
