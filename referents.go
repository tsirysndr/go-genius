package genius

type ReferentsService service

type Referent struct {
	Type           string `json:"_type,omitempty"`
	AnnotatorID    int    `json:"annotator_id,omitempty"`
	AnnotatorLogin string `json:"annotator_login,omitempty"`
	ApiPath        string `json:"api_path,omitempty"`
	Classification string `json:"classification,omitempty"`
	Featured       bool   `json:"featured,omitempty"`
	Fragment       string `json:"fragment,omitempty"`
	ID             int    `json:"id,omitempty"`
	IsDescription  bool   `json:"is_description,omitempty"`
	Path           string `json:"path,omitempty"`
	Range          *struct {
		Start       string `json:"start,omitempty"`
		StartOffset string `json:"startOffset,omitempty"`
		End         string `json:"end,omitempty"`
		EndOffset   string `json:"endOffset,omitempty"`
		Before      string `json:"before,omitempty"`
		After       string `json:"after,omitempty"`
		Content     string `json:"content,omitempty"`
	} `json:"range,omitempty"`
	SongID               int    `json:"song_id,omitempty"`
	URL                  string `json:"url,omitempty"`
	VerifiedAnnotatorIDs []int  `json:"verified_annotator_ids,omitempty"`
	Annotatable          *struct {
		ApiPath   string `json:"api_path,omitempty"`
		Context   string `json:"context,omitempty"`
		ID        int    `json:"id,omitempty"`
		ImageURL  string `json:"image_url,omitempty"`
		LinkTitle string `json:"link_title,omitempty"`
		Title     string `json:"title,omitempty"`
		Type      string `json:"type,omitempty"`
		URL       string `json:"url,omitempty"`
	} `json:"annotatable,omitempty"`
	Annotations []Annotation `json:"annotations,omitempty"`
}

type ReferentsParams struct {
	WebPageID   int    `url:"web_page_id,omitempty"`
	CreatedByID int    `url:"create_by_id,omitempty"`
	SongID      int    `url:"song_id,omitempty"`
	PerPage     int    `url:"per_page,omitempty"`
	Page        int    `url:"page,omitempty"`
	TextFormat  string `url:"text_format,omitempty"`
}

func (s *ReferentsService) GetBySongID(ID int) ([]Referent, error) {
	_, refs, err := s.GetBySongIDPages(ID, 0, 0)
	return refs, err
}

func (s *ReferentsService) GetByUserID(ID int) ([]Referent, error) {
	_, refs, err := s.GetByUserIDPages(ID, 0, 0)
	return refs, err
}

func (s *ReferentsService) GetByWebPageID(ID int) ([]Referent, error) {
	_, refs, err := s.GetByWebPageIDPages(ID, 0, 0)
	return refs, err
}

// GetBySongIDPages returns referents for a song with pagination.
// page=0 and perPage=0 means no pagination (API defaults).
// Returns (nextPage, referents, error). nextPage is nil on the last page.
func (s *ReferentsService) GetBySongIDPages(ID int, page, perPage int) (*int, []Referent, error) {
	params := &ReferentsParams{
		SongID:     ID,
		PerPage:    perPage,
		Page:       page,
		TextFormat: "plain",
	}
	res := new(ApiResponse)
	_, err := s.client.base.Get("referents").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, nil, err
	}
	if res.Response == nil {
		return nil, nil, nil
	}
	return res.Response.NextPage, res.Response.Referents, nil
}

// GetByUserIDPages returns referents by user with pagination.
func (s *ReferentsService) GetByUserIDPages(ID int, page, perPage int) (*int, []Referent, error) {
	params := &ReferentsParams{
		CreatedByID: ID,
		PerPage:     perPage,
		Page:        page,
		TextFormat:  "plain",
	}
	res := new(ApiResponse)
	_, err := s.client.base.Get("referents").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, nil, err
	}
	if res.Response == nil {
		return nil, nil, nil
	}
	return res.Response.NextPage, res.Response.Referents, nil
}

// GetByWebPageIDPages returns referents for a web page with pagination.
func (s *ReferentsService) GetByWebPageIDPages(ID int, page, perPage int) (*int, []Referent, error) {
	params := &ReferentsParams{
		WebPageID:  ID,
		PerPage:    perPage,
		Page:       page,
		TextFormat: "plain",
	}
	res := new(ApiResponse)
	_, err := s.client.base.Get("referents").QueryStruct(params).Receive(res, nil)
	if err != nil {
		return nil, nil, err
	}
	if res.Response == nil {
		return nil, nil, nil
	}
	return res.Response.NextPage, res.Response.Referents, nil
}
