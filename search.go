package genius

type SearchService service

type Hit struct {
	HighLights interface{} `json:"highlights"`
	Index      string      `json:"string"`
	Type       string      `json:"type"`
	Result     *struct {
		AnnotationCount          int    `json:"annotation_count"`
		ApiPath                  string `json:"api_path"`
		FullTitle                string `json:"full_title"`
		HeaderImageThumbnailURL  string `json:"header_image_thumbnail_url"`
		ID                       int    `json:"id"`
		LyricsOwnerID            int    `json:"lyrics_owner_id"`
		LyricsState              string `json:"lyrics_state"`
		Path                     string `json:"path"`
		PyongsCount              int    `json:"pyongs_count"`
		SongArtImageThumbnailURL string `json:"song_art_image_thumbnail_url"`
		SongArtImageURL          string `json:"song_art_image_url"`
		Stats                    *struct {
			UnreviewedAnnotations int  `json:"unreviewed_annotations"`
			Concurrents           int  `json:"concurrents"`
			Hot                   bool `json:"hot"`
			PageViews             int  `json:"page_views"`
		} `json:"stats"`
		Title             string  `json:"title"`
		TitleWithFeatured string  `json:"title_with_featured"`
		URL               string  `json:"url"`
		PrimaryArtist     *Artist `json:"primary_artist"`
	} `json:"result"`
}

type SearchParams struct {
	Q string `url:"q"`
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
