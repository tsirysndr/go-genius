package genius

type SongsService service

type Song struct {
	AnnotationCount          int         `json:"annotation_count,omitempty"`
	ApiPath                  string      `json:"api_path,omitempty"`
	AppleMusicID             string      `json:"apple_music_id,omitempty"`
	AppleMusicPlayerURL      string      `json:"apple_music_player_url,omitempty"`
	Description              interface{} `json:"description,omitempty"`
	EmbedContent             string      `json:"embed_content,omitempty"`
	FeaturedVideo            bool        `json:"featured_video,omitempty"`
	FullTitle                string      `json:"full_title,omitempty"`
	HeaderImageThumbnailURL  string      `json:"header_image_thumbnail_url,omitempty"`
	HeaderImageURL           string      `json:"header_image_url,omitempty"`
	ID                       int         `json:"id,omitempty"`
	LyricsOwnerID            int         `json:"lyrics_owner_id,omitempty"`
	LyricsState              string      `json:"lyrics_state,omitempty"`
	Path                     string      `json:"path,omitempty"`
	PyongsCount              int         `json:"pyongs_count,omitempty"`
	SongArtImageThumbnailURL string      `json:"song_art_image_thumbnail_url,omitempty"`
	SongArtImageURL          string      `json:"song_art_image_url,omitempty"`
	Stats                    *struct {
		UnreviewedAnnotations int  `json:"unreviewed_annotations,omitempty"`
		Concurrents           int  `json:"concurrents,omitempty"`
		Hot                   bool `json:"hot,omitempty"`
		PageViews             int  `json:"page_views,omitempty"`
	} `json:"stats,omitempty"`
	Title               string `json:"title,omitempty"`
	TitleFeatured       string `json:"title_features,omitempty"`
	URL                 string `json:"url,omitempty"`
	CurrentUserMetadata *struct {
		Permissions         []string `json:"permissions,omitempty"`
		ExcludedPermissions []string `json:"excluded_permissions,omitempty"`
		Interactions        *struct {
			Following bool `json:"following,omitempty"`
		} `json:"interactions,omitempty"`
	} `json:"current_user_metadata,omitempty"`
	Album                 *Album        `json:"album,omitempty"`
	CustomPerformances    []interface{} `json:"cutom_performances,omitempty"`
	DescriptionAnnotation *struct {
		Type           string `json:"_type,omitempty"`
		AnnotatorID    int    `json:"annotator_id,omitempty"`
		AnnotatorLogin string `json:"annotator_login,omitempty"`
		ApiPath        string `json:"api_path,omitempty"`
		Classification string `json:"classification,omitempty"`
		Fragment       string `json:"fragment,omitempty"`
		ID             int    `json:"id,omitempty"`
		IsDescription  bool   `json:"is_description,omitempty"`
		Path           string `json:"path,omitempty"`
		Range          *struct {
			Content string `json:"content,omitempty"`
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
	} `json:"description_annotation,omitempty"`
	FeaturedArtists        []interface{}      `json:"featured_artists,omitempty"`
	LyricsMarkedCompleteBy interface{}        `json:"lyrics_marked_complete_by,omitempty"`
	Media                  []Media            `json:"media,omitempty"`
	PrimaryArtist          *Artist            `json:"primary_artist,omitempty"`
	ProducerArtists        []Artist           `json:"producer_artists,omitempty"`
	SongRelationships      []SongRelationship `json:"song_relationships,omitempty"`
	VerifiedAnnotationsBy  []interface{}      `json:"verified_annotations_by,omitempty"`
	VerifiedContributors   []interface{}      `json:"verified_contributors,omitempty"`
	VerifiedLyricsBy       []interface{}      `json:"verified_lyrics_by,omitempty"`
	WriterArtists          []Artist           `json:"writer_artists,omitempty"`
}

type Album struct {
	ApiPath     string  `json:"api_path,omitempty"`
	CoverArtURL string  `json:"cover_art_url,omitempty"`
	FullTitle   string  `json:"full_title,omitempty"`
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	URL         string  `json:"url,omitempty"`
	Artist      *Artist `json:"artist,omitempty"`
}

type Media struct {
	Provider string `json:"provider,omitempty"`
	Start    int    `json:"start,omitempty"`
	Type     string `json:"type,omitempty"`
	URL      string `json:"url,omitempty"`
}

type SongRelationship struct {
	Type  string `json:"type,omitempty"`
	Songs []Song `json:"songs,omitempty"`
}

type SongParams struct {
	TextFormat string `url:"text_format,omitempty"`
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
