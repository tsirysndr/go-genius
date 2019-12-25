package genius

type SongsService service

type Song struct {
	AnnotationCount          int         `json:"annotation_count"`
	ApiPath                  string      `json:"api_path"`
	AppleMusicID             string      `json:"apple_music_id"`
	AppleMusicPlayerURL      string      `json:"apple_music_player_url"`
	Description              interface{} `json:"description"`
	EmbedContent             string      `json:"embed_content"`
	FeaturedVideo            bool        `json:"featured_video"`
	FullTitle                string      `json:"full_title"`
	HeaderImageThumbnailURL  string      `json:"header_image_thumbnail_url"`
	HeaderImageURL           string      `json:"header_image_url"`
	ID                       int         `json:"id"`
	LyricsOwnerID            int         `json:"lyrics_owner_id"`
	LyricsState              string      `json:"lyrics_state"`
	Path                     string      `json:"path"`
	PyongsCount              int         `json:"pyongs_count"`
	SongArtImageThumbnailURL string      `json:"song_art_image_thumbnail_url"`
	SongArtImageURL          string      `json:"song_art_image_url"`
	Stats                    *struct {
		UnreviewedAnnotations int  `json:"unreviewed_annotations"`
		Concurrents           int  `json:"concurrents"`
		Hot                   bool `json:"hot"`
		PageViews             int  `json:"page_views"`
	} `json:"stats"`
	Title               string `json:"title"`
	TitleFeatured       string `json:"title_features"`
	URL                 string `json:"url"`
	CurrentUserMetadata *struct {
		Permissions         []string `json:"permissions"`
		ExcludedPermissions []string `json:"excluded_permissions"`
		Interactions        *struct {
			Following bool `json:"following"`
		} `json:"interactions"`
	} `json:"current_user_metadata"`
	Album                 *Album        `json:"album"`
	CustomPerformances    []interface{} `json:"cutom_performances"`
	DescriptionAnnotation *struct {
		Type           string `json:"_type"`
		AnnotatorID    int    `json:"annotator_id"`
		AnnotatorLogin string `json:"annotator_login"`
		ApiPath        string `json:"api_path"`
		Classification string `json:"classification"`
		Fragment       string `json:"fragment"`
		ID             int    `json:"id"`
		IsDescription  bool   `json:"is_description"`
		Path           string `json:"path"`
		Range          *struct {
			Content string `json:"content"`
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
	} `json:"description_annotation"`
	FeaturedArtists        []interface{}      `json:"featured_artists"`
	LyricsMarkedCompleteBy interface{}        `json:"lyrics_marked_complete_by"`
	Media                  []Media            `json:"media"`
	PrimaryArtist          *Artist            `json:"primary_artist"`
	ProducerArtists        []Artist           `json:"producer_artists"`
	SongRelationships      []SongRelationship `json:"song_relationships"`
	VerifiedAnnotationsBy  []interface{}      `json:"verified_annotations_by"`
	VerifiedContributors   []interface{}      `json:"verified_contributors"`
	VerifiedLyricsBy       []interface{}      `json:"verified_lyrics_by"`
	WriterArtists          []Artist           `json:"writer_artists"`
}

type Album struct {
	ApiPath     string  `json:"api_path"`
	CoverArtURL string  `json:"cover_art_url"`
	FullTitle   string  `json:"full_title"`
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	URL         string  `json:"url"`
	Artist      *Artist `json:"artist"`
}

type Media struct {
	Provider string `json:"provider"`
	Start    int    `json:"start"`
	Type     string `json:"type"`
	URL      string `json:"url"`
}

type SongRelationship struct {
	Type  string `json:"type"`
	Songs []Song `json:"songs"`
}

type SongParams struct {
	TextFormat string `url:"text_format"`
}

func (s *SongsService) Get(ID string) (*Song, error) {
	var err error
	params := &SongParams{"plain"}
	res := new(ApiResponse)
	s.client.base.Path("songs/").Get(ID).QueryStruct(params).Receive(res, err)
	if err != nil {
		return nil, err
	}
	return res.Response.Song, nil
}
