package word

type Lemma struct {
	ID          int    `json:"id"`
	Word        string `json:"word"`
	Definitions string `json:"definitions"`
}

type PartOfSpeech struct {
	ID           int     `json:"id"`
	PartOfSpeech string  `json:"part_of_speech"`
	Lemmas       []Lemma `json:"lemmas"`
}

type Word struct {
	ID            int            `json:"id"`
	Word          string         `json:"word"`
	PartsOfSpeech []PartOfSpeech `json:"parts_of_speech"`
}
