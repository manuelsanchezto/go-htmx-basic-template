package dto

type SectionsToEdit struct {
	Query  string         `json:"query"`
	Module string         `json:"module"`
	Result ResultSections `json:"result"`
}

type ResultSections struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
