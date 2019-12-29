package genius

type ArtistsService service

type Artist struct {
	AlternateNames      []string    `json:"alternate_names,omitempty"`
	ApiPath             string      `json:"api_path,omitempty"`
	Description         interface{} `json:"description,omitempty"`
	FacebookName        string      `json:"facebook_name,omitempty"`
	FollowersCount      int         `json:"followers_count,omitempty"`
	HeaderImageURL      string      `json:"header_image_url,omitempty"`
	ID                  string      `json:"id,omitempty"`
	ImageURL            string      `json:"image_url,omitempty"`
	InstagramName       string      `json:"instagram_name,omitempty"`
	IsMemeVerified      bool        `json:"is_meme_verified,omitempty"`
	IsVerified          bool        `json:"is_verified,omitempty"`
	Name                string      `json:"name,omitempty"`
	TranslationArtist   bool        `json:"translation_artist,omitempty"`
	TwitterName         string      `json:"twitter_name,omitempty"`
	URL                 string      `json:"url,omitempty"`
	CurrentUserMetadata *struct {
		Permissions         []string `json:"permissions,omitempty"`
		ExcludedPermissions []string `json:"excluded_permissions,omitempty"`
		Interactions        *struct {
			Following bool `json:"following,omitempty"`
		} `json:"interactions,omitempty"`
	} `json:"current_user_metadata,omitempty"`
	IQ                    int `json:"iq,omitempty"`
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
	User *User `json:"user,omitempty"`
}

type User struct {
	ApiPath string `json:"api_path,omitempty"`
	Avatar  *struct {
		Tiny *struct {
			URL         string `json:"url,omitempty"`
			BoundingBox *struct {
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"bounding_box,omitempty"`
		} `json:"tiny,omitempty"`
		Thumb *struct {
			URL         string `json:"url,omitempty"`
			BoundingBox *struct {
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"bounding_box,omitempty"`
		} `json:"thumb,omitempty"`
		Small *struct {
			URL         string `json:"url,omitempty"`
			BoundingBox *struct {
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"bounding_box,omitempty"`
		} `json:"small,omitempty"`
		Medium *struct {
			URL         string `json:"url,omitempty"`
			BoundingBox *struct {
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"bounding_box,omitempty"`
		} `json:"medium,omitempty"`
	} `json:"avatar,omitempty"`
	HeaderImageURL              string `json:"header_image_url,omitempty"`
	HumanReadableRoleForDisplay string `json:"human_readable_role_for_display,omitempty"`
	ID                          int    `json:"id,omitempty"`
	IQ                          int    `json:"iq,omitempty"`
	Login                       string `json:"login,omitempty"`
	Name                        string `json:"string,omitempty"`
	RoleForDisplay              string `json:"role_for_display,omitempty"`
	URL                         string `json:"url,omitempty"`
	CurrentUserMetadata         *struct {
		Permissions         []string `json:"permissions,omitempty"`
		ExcludedPermissions []string `json:"excluded_permissions,omitempty"`
		Interactions        *struct {
			Following bool `json:"following,omitempty"`
		} `json:"interactions,omitempty"`
	} `json:"current_user_metadata,omitempty"`
}

type ArtistParams struct {
	TextFormat string `url:"text_format,omitempty"`
}

func (s *ArtistsService) Get(ID string) (*Artist, error) {
	var err error
	params := &ArtistParams{"plain"}
	res := new(ApiResponse)
	s.client.base.Path("artists/").Get(ID).QueryStruct(params).Receive(res, err)
	if err != nil {
		return nil, err
	}
	return res.Response.Artist, nil
}
