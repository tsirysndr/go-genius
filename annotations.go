package genius

type AnnotationsService service

type Annotation struct {
	ApiPath             string      `json:"api_path"`
	Body                interface{} `json:"body"`
	CommentCount        int         `json:"comment_count"`
	Community           bool        `json:"community"`
	CustomPreview       bool        `json:"custom_preview"`
	HasVoters           bool        `json:"has_voters"`
	ID                  int         `json:"id"`
	Pinned              bool        `json:"pinned"`
	ShareURL            string      `json:"share_url"`
	Source              string      `json:"source"`
	State               string      `json:"state"`
	URL                 string      `json:"url"`
	Verified            bool        `json:"verified"`
	VotesTotal          int         `json:"votes_total"`
	CurrentUserMetadata *struct {
		Permissions         []string `json:"permissions"`
		ExcludedPermissions []string `json:"excluded_permissions"`
		Interactions        *struct {
			Following bool `json:"following"`
		} `json:"interactions"`
	} `json:"current_user_metadata"`
	Authors          []Author      `json:"authors"`
	CosignedBy       []interface{} `json:"cosigned_by"`
	RejectionComment string        `json:"rejection_comment"`
	VerifiedBy       *User         `json:"verified_by"`
	Referent         *Referent     `json:"referent"`
}

type Author struct {
	Attribution int    `json:"attribution"`
	PinnedRole  string `json:"pinned_role"`
	User        *User  `json:"user"`
}

type AnnotationsParams struct {
	TextFormat string `url:"text_format"`
}

func (s *AnnotationsService) Get(ID string) (*Annotation, error) {
	var err error
	params := &AnnotationsParams{"plain"}
	res := new(ApiResponse)
	s.client.base.Path("annotations/").Get(ID).QueryStruct(params).Receive(res, err)
	if err != nil {
		return nil, err
	}
	return res.Response.Annotation, nil
}
