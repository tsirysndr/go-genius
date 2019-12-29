package genius

type AnnotationsService service

type Annotation struct {
	ApiPath             string      `json:"api_path,omitempty"`
	Body                interface{} `json:"body,omitempty"`
	CommentCount        int         `json:"comment_count,omitempty"`
	Community           bool        `json:"community,omitempty"`
	CustomPreview       bool        `json:"custom_preview,omitempty"`
	HasVoters           bool        `json:"has_voters,omitempty"`
	ID                  int         `json:"id,omitempty"`
	Pinned              bool        `json:"pinned,omitempty"`
	ShareURL            string      `json:"share_url,omitempty"`
	Source              string      `json:"source,omitempty"`
	State               string      `json:"state,omitempty"`
	URL                 string      `json:"url,omitempty"`
	Verified            bool        `json:"verified,omitempty"`
	VotesTotal          int         `json:"votes_total,omitempty"`
	CurrentUserMetadata *struct {
		Permissions         []string `json:"permissions,omitempty"`
		ExcludedPermissions []string `json:"excluded_permissions,omitempty"`
		Interactions        *struct {
			Following bool `json:"following,omitempty"`
		} `json:"interactions,omitempty"`
	} `json:"current_user_metadata,omitempty"`
	Authors          []Author      `json:"authors,omitempty"`
	CosignedBy       []interface{} `json:"cosigned_by,omitempty"`
	RejectionComment string        `json:"rejection_comment,omitempty"`
	VerifiedBy       *User         `json:"verified_by,omitempty"`
	Referent         *Referent     `json:"referent,omitempty"`
}

type Author struct {
	Attribution int    `json:"attribution,omitempty"`
	PinnedRole  string `json:"pinned_role,omitempty"`
	User        *User  `json:"user,omitempty"`
}

type AnnotationsParams struct {
	TextFormat string `url:"text_format,omitempty"`
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
