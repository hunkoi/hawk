package hawk

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"
)

func Nonce(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func parseHawkHeader(headerVal string) map[string]string {
	attrs := make(map[string]string)

	hv := strings.Split(strings.Split(headerVal, "Hawk ")[1], ", ")

	for _, v := range hv {
		r := regexp.MustCompile(`(\w+)="([^"\\]*)"\s*(?:,\s*|$)`)
		group := r.FindSubmatch([]byte(v))
		attrs[string(group[1])] = string(group[2])
	}

	return attrs
}
