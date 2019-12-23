package genius

type ArtistsService service

type Artist struct {
	AlternateNames      []string    `json:"alternate_names"`
	ApiPath             string      `json:"api_path"`
	Description         interface{} `json:"description"`
	FacebookName        string      `json:"facebook_name"`
	FollowersCount      int         `json:"followers_count"`
	HeaderImageURL      string      `json:"header_image_url"`
	ID                  string      `json:"id"`
	ImageURL            string      `json:"image_url"`
	InstagramName       string      `json:"instagram_name"`
	IsMemeVerified      bool        `json:"is_meme_verified"`
	IsVerified          bool        `json:"is_verified"`
	Name                string      `json:"name"`
	TranslationArtist   bool        `json:"translation_artist"`
	TwitterName         string      `json:"twitter_name"`
	URL                 string      `json:"url"`
	CurrentUserMetadata *struct {
		Permissions         []string `json:"permissions"`
		ExcludedPermissions []string `json:"excluded_permissions"`
		Interactions        *struct {
			Following bool `json:"following"`
		} `json:"interactions"`
	} `json:"current_user_metadata"`
	IQ                    int `json:"iq"`
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
	User *User `json:"user"`
}

type User struct {
	ApiPath string `json:"api_path"`
	Avatar  *struct {
		Tiny *struct {
			URL         string `json:"url"`
			BoundingBox *struct {
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"bounding_box"`
		} `json:"tiny"`
		Thumb *struct {
			URL         string `json:"url"`
			BoundingBox *struct {
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"bounding_box"`
		} `json:"thumb"`
		Small *struct {
			URL         string `json:"url"`
			BoundingBox *struct {
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"bounding_box"`
		} `json:"small"`
		Medium *struct {
			URL         string `json:"url"`
			BoundingBox *struct {
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"bounding_box"`
		} `json:"medium"`
	} `json:"avatar"`
	HeaderImageURL              string `json:"header_image_url"`
	HumanReadableRoleForDisplay string `json:"human_readable_role_for_display"`
	ID                          int    `json:"id"`
	IQ                          int    `json:"iq"`
	Login                       string `json:"login"`
	Name                        string `json:"string"`
	RoleForDisplay              string `json:"role_for_display"`
	URL                         string `json:"url"`
	CurrentUserMetadata         *struct {
		Permissions         []string `json:"permissions"`
		ExcludedPermissions []string `json:"excluded_permissions"`
		Interactions        *struct {
			Following bool `json:"following"`
		} `json:"interactions"`
	} `json:"current_user_metadata"`
}

func (s *ArtistsService) Get(ID int) (*Artist, error) {
	return nil, nil
}
