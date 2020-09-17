package mimetype

import "testing"

// TestBasics runs some simple tests
func TestBasics(t *testing.T) {
	mimeType, err := ByExtension(".mp4")
	if err != nil {
		t.Errorf(`ByExtension(".mp4"), expected no errors.  Got %s`, err.Error())
	}
	if mimeType != "video/mp4" {
		t.Errorf(`ByExtension(".mp4"), expected "video/mp4".  Got %s`, mimeType)

	}

}
