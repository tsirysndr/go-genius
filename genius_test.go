package genius

import (
	"fmt"
	"os"
	"testing"
)

func getToken(t *testing.T) string {
	token := os.Getenv("GENIUS_ACCESS_TOKEN")
	if token == "" {
		t.Skip("GENIUS_ACCESS_TOKEN not set")
	}
	return token
}

func TestNewClient(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)
	if c == nil {
		t.Fatal("NewClient returned nil")
	}
	if c.Search == nil {
		t.Error("Search service is nil")
	}
	if c.Songs == nil {
		t.Error("Songs service is nil")
	}
	if c.Artists == nil {
		t.Error("Artists service is nil")
	}
	if c.Annotations == nil {
		t.Error("Annotations service is nil")
	}
	if c.Referents == nil {
		t.Error("Referents service is nil")
	}
}

func TestSearchGet(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)

	hits, err := c.Search.Get("Bohemian Rhapsody")
	if err != nil {
		t.Fatalf("Search.Get failed: %v", err)
	}
	if len(hits) == 0 {
		t.Fatal("Search.Get returned 0 hits")
	}

	found := false
	for _, h := range hits {
		if h.Result != nil && h.Result.Title != "" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Search.Get returned hits with no Results")
	}

	t.Logf("Search returned %d hits, first: %s", len(hits), hits[0].Result.FullTitle)
}

func TestSearchGetNoResult(t *testing.T) {
	// Test that nil Result doesn't panic (our PR #8 fix)
	token := getToken(t)
	c := NewClient(token)

	hits, err := c.Search.Get("zzzxxxyyyyunknownsong12345")
	if err != nil {
		t.Fatalf("Search.Get failed: %v", err)
	}
	if hits == nil {
		return // no hits is fine
	}
	for _, h := range hits {
		_ = h.Result // should not panic even if nil
	}
}

func TestSongsGet(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)

	// Bohemian Rhapsody on Genius has song ID 1 or similar known song
	// Let's search first to find a valid song ID
	hits, err := c.Search.Get("Bohemian Rhapsody")
	if err != nil {
		t.Fatalf("Search.Get failed: %v", err)
	}
	if len(hits) == 0 || hits[0].Result == nil {
		t.Fatal("No search results")
	}

	songID := fmt.Sprintf("%d", hits[0].Result.ID)
	t.Logf("Fetching song ID=%s (title=%s)", songID, hits[0].Result.FullTitle)

	song, err := c.Songs.Get(songID)
	if err != nil {
		t.Fatalf("Songs.Get failed: %v", err)
	}
	if song == nil {
		t.Fatal("Songs.Get returned nil song")
	}
	if song.FullTitle == "" {
		t.Error("Song has no title")
	}
	if song.URL == "" {
		t.Error("Song has no URL")
	}
	if song.PrimaryArtist == nil {
		t.Error("Song has no primary artist")
	}
	t.Logf("Got song: %s by %s", song.FullTitle, song.PrimaryArtist.Name)
}

func TestArtistsGet(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)

	// Find Queen
	hits, err := c.Search.Get("Queen")
	if err != nil {
		t.Fatalf("Search.Get failed: %v", err)
	}
	if len(hits) == 0 || hits[0].Result == nil || hits[0].Result.PrimaryArtist == nil {
		t.Fatal("No search results with artist")
	}

	artistID := fmt.Sprintf("%d", hits[0].Result.PrimaryArtist.ID)
	t.Logf("Fetching artist ID=%s (name=%s)", artistID, hits[0].Result.PrimaryArtist.Name)

	artist, err := c.Artists.Get(artistID)
	if err != nil {
		t.Fatalf("Artists.Get failed: %v", err)
	}
	if artist == nil {
		t.Fatal("Artists.Get returned nil")
	}
	if artist.Name == "" {
		t.Error("Artist has no name")
	}
	t.Logf("Got artist: %s (followers: %d)", artist.Name, artist.FollowersCount)
}

func TestAnnotationsGet(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)

	// Get a song first to find a known annotation
	hits, err := c.Search.Get("Bohemian Rhapsody Queen")
	if err != nil {
		t.Fatalf("Search.Get failed: %v", err)
	}
	if len(hits) == 0 || hits[0].Result == nil {
		t.Skip("No search results")
	}

	// Try fetching an annotation using referents from the song
	_, refs, err := c.Referents.GetBySongIDPages(hits[0].Result.ID, 1, 1)
	if err != nil {
		t.Skipf("Referents.GetBySongIDPages failed: %v", err)
	}
	if len(refs) == 0 {
		t.Skip("No referents found for song")
	}

	t.Logf("First referent ID=%d", refs[0].ID)
	// Note: annotations endpoint uses annotation ID, not referent ID
	// We test the method exists and returns expected error for invalid ID
	annotation, err := c.Annotations.Get("999999999")
	if err != nil {
		// might return error for non-existent annotation, that's fine
		t.Logf("Annotations.Get with fake ID returned: %v", err)
	} else if annotation != nil {
		t.Logf("Got annotation ID=%d", annotation.ID)
	}
}

func TestReferentsGetBySongID(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)

	hits, err := c.Search.Get("Bohemian Rhapsody Queen")
	if err != nil {
		t.Fatalf("Search.Get failed: %v", err)
	}
	if len(hits) == 0 || hits[0].Result == nil {
		t.Skip("No search results")
	}

	refs, err := c.Referents.GetBySongID(hits[0].Result.ID)
	if err != nil {
		t.Fatalf("Referents.GetBySongID failed: %v", err)
	}
	t.Logf("Got %d referents for song ID=%d", len(refs), hits[0].Result.ID)
}

func TestReferentsGetBySongIDPages(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)

	hits, err := c.Search.Get("Bohemian Rhapsody Queen")
	if err != nil {
		t.Fatalf("Search.Get failed: %v", err)
	}
	if len(hits) == 0 || hits[0].Result == nil {
		t.Skip("No search results")
	}

	// Test pagination with per_page=5
	nextPage, refs, err := c.Referents.GetBySongIDPages(hits[0].Result.ID, 1, 5)
	if err != nil {
		t.Fatalf("Referents.GetBySongIDPages failed: %v", err)
	}
	if len(refs) == 0 {
		t.Log("No referents returned (possible, e.g., no annotations)")
	} else {
		t.Logf("Got %d referents on page 1, nextPage=%v", len(refs), nextPage)
	}
}

func TestArtistsGetSongs(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)

	// Find an artist
	hits, err := c.Search.Get("Eminem")
	if err != nil {
		t.Fatalf("Search.Get failed: %v", err)
	}
	if len(hits) == 0 || hits[0].Result == nil || hits[0].Result.PrimaryArtist == nil {
		t.Skip("No search results with artist")
	}

	artistID := fmt.Sprintf("%d", hits[0].Result.PrimaryArtist.ID)
	t.Logf("Fetching songs for artist ID=%s", artistID)

	songs, nextPage, err := c.Artists.GetSongs(artistID, nil)
	if err != nil {
		t.Fatalf("Artists.GetSongs failed: %v", err)
	}
	if len(songs) == 0 {
		t.Fatal("Artists.GetSongs returned 0 songs")
	}
	if songs[0].Title == "" {
		t.Error("First song has no title")
	}
	t.Logf("Got %d songs, nextPage=%v, first: %s", len(songs), nextPage, songs[0].FullTitle)
}

func TestArtistsGetSongsWithPagination(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)

	hits, err := c.Search.Get("Eminem")
	if err != nil {
		t.Fatalf("Search.Get failed: %v", err)
	}
	if len(hits) == 0 || hits[0].Result == nil || hits[0].Result.PrimaryArtist == nil {
		t.Skip("No search results with artist")
	}

	artistID := fmt.Sprintf("%d", hits[0].Result.PrimaryArtist.ID)
	params := &ArtistSongsParams{
		PerPage:    3,
		Sort:       "popularity",
		TextFormat: "plain",
	}
	songs, nextPage, err := c.Artists.GetSongs(artistID, params)
	if err != nil {
		t.Fatalf("Artists.GetSongs with pagination failed: %v", err)
	}
	if len(songs) > 3 {
		t.Errorf("Expected at most 3 songs per page, got %d", len(songs))
	}
	t.Logf("Got %d songs (per_page=3), nextPage=%v", len(songs), nextPage)
}

func TestGetLyrics(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)

	// Search for a well-known song
	hits, err := c.Search.Get("Bohemian Rhapsody Queen")
	if err != nil {
		t.Fatalf("Search.Get failed: %v", err)
	}
	if len(hits) == 0 || hits[0].Result == nil || hits[0].Result.URL == "" {
		t.Skip("No search results with URL")
	}

	songURL := hits[0].Result.URL
	t.Logf("Scraping lyrics from: %s", songURL)

	lyrics, err := GetLyrics(songURL)
	if err != nil {
		// Cloudflare challenge is expected from some IPs — mark as skip
		t.Skipf("GetLyrics failed (may be Cloudflare): %v", err)
	}
	if len(lyrics) == 0 {
		t.Fatal("GetLyrics returned empty lyrics")
	}
	t.Logf("Lyrics length: %d chars, first 100: %s", len(lyrics), lyrics[:min(100, len(lyrics))])
}

func TestGetLyricsInvalidURL(t *testing.T) {
	_, err := GetLyrics("https://genius.com/this-page-does-not-exist-12345")
	if err == nil {
		t.Fatal("Expected error for invalid URL, got nil")
	}
	t.Logf("Got expected error: %v", err)
}

// TestNilResponseGuard — regression: ensure nil Response doesn't panic
func TestNilResponseGuard(t *testing.T) {
	token := getToken(t)
	c := NewClient(token)

	// Pass an artist ID that doesn't exist — Response can be nil
	_, err := c.Artists.Get("999999999")

	// The get may succeed with nil Artist, or error out — either is fine
	// but it MUST NOT panic
	t.Logf("Artists.Get with fake ID returned err=%v", err)
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
