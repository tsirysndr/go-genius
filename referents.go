package genius

type ReferentsService service

type Referent struct {
	Type           string `json:"_type"`
	AnnotatorID    int    `json:"annotator_id"`
	AnnotatorLogin string `json:"annotator_login"`
	ApiPath        string `json:"api_path"`
	Classification string `json:"classification"`
	Featured       bool   `json:"featured"`
	Fragment       string `json:"fragment"`
	ID             int    `json:"id"`
	IsDescription  bool   `json:"is_description"`
	Path           string `json:"path"`
	Range          *struct {
		Start       string `json:"start"`
		StartOffset string `json:"startOffset"`
		End         string `json:"end"`
		EndOffset   string `json:"endOffset"`
		Before      string `json:"before"`
		After       string `json:"after"`
		Content     string `json:"content"`
	} `json:"range"`
	SongID               int    `json:"song_id"`
	URL                  string `json:"url"`
	VerifiedAnnotatorIDs []int  `json:"verified_annotator_ids"`
	Annotatable          *struct {
		ApiPath   string `json:"api_path"`
		Context   string `json:"context"`
		ID        int    `json:"id"`
		ImageURL  string `json:"image_url"`
		LinkTitle string `json:"link_title"`
		Title     string `json:"title"`
		Type      string `json:"type"`
		URL       string `json:"url"`
	} `json:"annotatable"`
	Annotations []Annotation `json:"annotations"`
}

func (s *ReferentsService) GetBySongID(songID int) ([]Referent, error) {
	return nil, nil
}

func (s *ReferentsService) GetByUserID(songID int) ([]Referent, error) {
	return nil, nil
}

func (s *ReferentsService) GetByWebPageID(songID int) ([]Referent, error) {
	return nil, nil
}
