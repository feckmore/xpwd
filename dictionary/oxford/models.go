package oxford

import "strings"

// ProficiencyLevel is an enum
type ProficiencyLevel int

const (
	A1 ProficiencyLevel = iota
	A2
	B1
	B2
	C1
	C2
)

func (p ProficiencyLevel) String() string {
	levels := []string{"A1", "A2", "B1", "B2", "C1", "C2"}
	return levels[p]
}

// SpeechPart is an enum
type SpeechPart int

const (
	Adjective SpeechPart = iota
	Adverb
	Article
	AuxiliaryVerb
	Conjunction
	Determiner
	Exclamation
	ModalVerb
	Noun
	Number
	Preposition
	Pronoun
	Verb
)

func (s SpeechPart) String() string {
	parts := []string{"Adjective", "Adverb", "Article", "Auxiliary Verb", "Conjunction", "Determiner", "Exclamation", "Modal Verb", "Noun", "Number", "Preposition", "Pronoun", "Verb"}
	return parts[s]
}

type SpeechParts []SpeechPart

func (s SpeechParts) String() string {
	parts := []string{}
	for _, part := range s {
		parts = append(parts, part.String())
	}
	return strings.Join(parts, ", ")
}

type Word struct {
	Word             string
	ProficiencyLevel ProficiencyLevel
	SpeechParts      SpeechParts
}
