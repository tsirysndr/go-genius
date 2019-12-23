package genius

type SongsService service

type Song struct {
	AnnotationCount          int    `json:"annotation_count"`
	ApiPath                  string `json:"api_path"`
	FullTitle                string `json:"full_title"`
	HeaderImageThumbnailURL  string `json:"header_image_thumbnail_url"`
	HeaderImageURL           string `json:"header_image_url"`
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
	Title         string  `json:"title"`
	TitleFeatured string  `json:"title_features"`
	URL           string  `json:"url"`
	PrimaryArtist *Artist `json:"primary_artist"`
}

func (s *SongsService) Get(ID int) (*Song, error) {
	return nil, nil
}
