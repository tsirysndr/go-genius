package genius

import (
	"github.com/dghubble/sling"
	"golang.org/x/oauth2"
)

type Client struct {
	base        *sling.Sling
	common      service
	Account     *AccountService
	Annotations *AnnotationsService
	Artists     *ArtistsService
	Referents   *ReferentsService
	Search      *SearchService
	Songs       *SongsService
	WebPages    *WebPagesService
}

type service struct {
	client *Client
}

type ApiResponse struct {
	Meta *struct {
		Status int `json:"status"`
	} `json:"meta"`
	Response *struct {
		Annotation *Annotation `json:"annotation"`
		Referents  []Referent  `json:"referents"`
		Song       *Song       `json:"song"`
		Artist     *Artist     `json:"artist"`
		Hits       []Hit       `json:"hits"`
	} `json:"response"`
}

func NewClient(accessToken string) *Client {
	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: accessToken}
	httpClient := config.Client(oauth2.NoContext, token)
	base := sling.New().Base("https://api.genius.com/").Client(httpClient)
	c := &Client{}
	c.base = base
	c.common.client = c
	c.Account = (*AccountService)(&c.common)
	c.Annotations = (*AnnotationsService)(&c.common)
	c.Artists = (*ArtistsService)(&c.common)
	c.Referents = (*ReferentsService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.Songs = (*SongsService)(&c.common)
	c.WebPages = (*WebPagesService)(&c.common)
	return c
}
