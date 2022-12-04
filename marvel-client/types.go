package main

type character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Modified    string `json:"modified"`
}

type CharacterResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int                 `json:"offset"`
		Limit   int                 `json:"limit"`
		Total   int                 `json:"total"`
		Count   int                 `json:"count"`
		Results []CharacterResponse `json:"data"`
	}
}
