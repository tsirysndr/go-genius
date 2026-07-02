package genius

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const geniusBrowserUA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

// GetLyrics fetches and parses lyrics from a Genius song page URL.
// It scrapes plain text from div[data-lyrics-container="true"] elements.
// This is a public scraping helper — it does not require a Client or API token.
func GetLyrics(songURL string) (string, error) {
	return GetLyricsWithClient(songURL, &http.Client{})
}

// GetLyricsWithClient is like GetLyrics but accepts a custom http.Client.
func GetLyricsWithClient(songURL string, httpClient *http.Client) (string, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	req, err := http.NewRequest("GET", songURL, nil)
	if err != nil {
		return "", fmt.Errorf("genius lyrics: create request: %w", err)
	}
	req.Header.Set("User-Agent", geniusBrowserUA)

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("genius lyrics: request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("genius lyrics: page returned %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("genius lyrics: parse html: %w", err)
	}

	var sb strings.Builder
	extractLyricsContainers(doc, &sb)
	lyrics := strings.TrimSpace(sb.String())
	if lyrics == "" {
		return "", fmt.Errorf("genius lyrics: no lyrics found at %s", songURL)
	}
	return lyrics, nil
}

func extractLyricsContainers(n *html.Node, sb *strings.Builder) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "data-lyrics-container" && a.Val == "true" {
				extractNodeText(n, sb)
				sb.WriteByte('\n')
				return
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractLyricsContainers(c, sb)
	}
}

func extractNodeText(n *html.Node, sb *strings.Builder) {
	if n.Type == html.TextNode {
		sb.WriteString(n.Data)
		return
	}
	if n.Type == html.ElementNode {
		if n.Data == "br" {
			sb.WriteByte('\n')
			return
		}
		for _, a := range n.Attr {
			if a.Key == "data-exclude-from-selection" && a.Val == "true" {
				return
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractNodeText(c, sb)
	}
}
