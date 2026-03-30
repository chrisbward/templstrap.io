package anchor

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestShow_AnchorType(t *testing.T) {
	// Pipe the rendered template into goquery.
	r, w := io.Pipe()
	go func() {
		_ = Show(AnchorProps{
			Type: Anchor,
			Text: "my link",
		}).Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}
	// Expect the component to have a class of "".
	if doc.Find(`[class]`).Length() == 0 {
		t.Error("expected element to not have class attribute")
	}
	// Expect the page name to be set correctly.
	expectedLinkContent := "my link"
	if actualLinkContent := doc.Find("a").Text(); actualLinkContent != expectedLinkContent {
		t.Errorf("expected link content %q, got %q", expectedLinkContent, actualLinkContent)
	}
}

func TestShow_ButtonType(t *testing.T) {
	// Pipe the rendered template into goquery.
	r, w := io.Pipe()
	go func() {
		_ = Show(AnchorProps{
			Type: Button,
			Text: "my link",
		}).Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}
	// Expect the component to have a class of "btn".
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// Check if the "class" attribute exists
		actualClassValue, exists := s.Attr("class")

		if !exists {
			t.Errorf("anchor has no class attribute")
		}
		expectedClassValue := "btn"
		if actualClassValue != expectedClassValue {
			t.Errorf("expected rendered class '%q', got '%q'", expectedClassValue, actualClassValue)
		}
	})
	// Expect the page name to be set correctly.
	expectedLinkContent := "my link"
	if actualLinkContent := doc.Find("a").Text(); actualLinkContent != expectedLinkContent {
		t.Errorf("expected link content %q, got %q", expectedLinkContent, actualLinkContent)
	}
}
