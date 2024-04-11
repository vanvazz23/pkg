package pkg

import (
	"net/url"
	"regexp"
	"strings"
)

// ByteSliceToString converts a byte slice to a string.
func ByteSliceToString(b []byte) string {
	return string(b)
}

// StringToByteSlice converts a string to a byte slice.
func StringToByteSlice(s string) []byte {
	return []byte(s)
}

// IsSameDomain checks if two URLs belong to the same domain.
func IsSameDomain(u1, u2 string) bool {
	parsedURL1, err := url.Parse(u1)
	if err != nil {
		return false
	}

	parsedURL2, err := url.Parse(u2)
	if err != nil {
		return false
	}

	return parsedURL1.Host == parsedURL2.Host
}

// URLDepth calculates the depth of a URL relative to a reference URL.
func URLDepth(u, referenceURL string) int {
	refURL, err := url.Parse(referenceURL)
	if err != nil {
		return -1 // Error parsing reference URL
	}

	parsedURL, err := url.Parse(u)
	if err != nil {
		return -1 // Error parsing URL
	}

	refPath := strings.TrimSuffix(refURL.Path, "/")
	parsedPath := strings.TrimSuffix(parsedURL.Path, "/")

	if !strings.HasPrefix(parsedPath, refPath) {
		return 0 // No common path prefix, depth is 0
	}

	relPath := strings.TrimPrefix(parsedPath, refPath)
	depth := strings.Count(relPath, "/")

	if relPath == "" {
		return 0 // URL is the same as the reference URL, depth is 0
	}

	return depth
}

// RemoveAnyQueryParam removes any query parameters from a URL.
func RemoveAnyQueryParam(u string) string {
	if strings.Contains(u, "?") {
		return strings.Split(u, "?")[0]
	}
	return u
}

// RemoveAnyAnchors removes any anchors from a URL.
func RemoveAnyAnchors(u string) string {
	if strings.Contains(u, "#") {
		return strings.Split(u, "#")[0]
	}
	return u
}

// GetBaseURL extracts the base URL from a full URL.
func GetBaseURL(u string) string {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return ""
	}
	return parsedURL.Scheme + "://" + parsedURL.Host
}

// ExtractEmailsFromText extracts email addresses from a text using regular expressions.
func ExtractEmailsFromText(text string) []string {
	// Regular expression to match email addresses
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Find all email addresses in the text
	emails := re.FindAllString(text, -1)

	return emails
}

// RelativeToAbsoluteURL converts a relative URL to an absolute URL based on the current URL and base URL.
func RelativeToAbsoluteURL(href, currentURL, baseURL string) string {
	if strings.HasPrefix(href, "http") {
		return href
	}

	if strings.HasPrefix(href, "/") {
		return baseURL + href
	}
	if strings.HasPrefix(href, "./") {
		return currentURL + href[2:]
	}

	return currentURL + href
}
