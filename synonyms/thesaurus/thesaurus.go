package thesaurus

type Thesaurus interface {
	Synonyms(true string) ([]string, error)
}
