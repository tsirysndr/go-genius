package genius

type SearchService service

type Hit struct {
	HighLights interface{} `json:"highlights,omitempty"`
	Index      string      `json:"string,omitempty"`
	Type       string      `json:"type,omitempty"`
	Result     *struct {
		AnnotationCount          int    `json:"annotation_count,omitempty"`
		ApiPath                  string `json:"api_path,omitempty"`
		FullTitle                string `json:"full_title,omitempty"`
		HeaderImageThumbnailURL  string `json:"header_image_thumbnail_url,omitempty"`
		ID                       int    `json:"id,omitempty"`
		LyricsOwnerID            int    `json:"lyrics_owner_id,omitempty"`
		LyricsState              string `json:"lyrics_state,omitempty"`
		Path                     string `json:"path,omitempty"`
		PyongsCount              int    `json:"pyongs_count,omitempty"`
		SongArtImageThumbnailURL string `json:"song_art_image_thumbnail_url,omitempty"`
		SongArtImageURL          string `json:"song_art_image_url,omitempty"`
		Stats                    *struct {
			UnreviewedAnnotations int  `json:"unreviewed_annotations,omitempty"`
			Concurrents           int  `json:"concurrents,omitempty"`
			Hot                   bool `json:"hot,omitempty"`
			PageViews             int  `json:"page_views,omitempty"`
		} `json:"stats,omitempty"`
		Title             string  `json:"title,omitempty"`
		TitleWithFeatured string  `json:"title_with_featured,omitempty"`
		URL               string  `json:"url,omitempty"`
		PrimaryArtist     *Artist `json:"primary_artist,omitempty"`
	} `json:"result,omitempty"`
}

type SearchParams struct {
	Q string `url:"q,omitempty"`
}

func (s *SearchService) Get(q string) ([]Hit, error) {
	var err error
	params := &SearchParams{q}
	res := new(ApiResponse)
	s.client.base.Get("search").QueryStruct(params).Receive(res, err)
	if err != nil {
		return nil, err
	}
	hits := res.Response.Hits
	return hits, nil
}
