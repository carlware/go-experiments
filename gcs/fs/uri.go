package fs

import "crypto/rand"

const (
	// StdLen is a standard length of uniuri string to achive ~95 bits of entropy.
	StdLen = 16
)

// StdChars is a set of standard characters allowed in uniuri string.
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// New returns a new random string of the standard length, consisting of
// standard characters.
func NewURI() string {
	return NewLenChars(StdLen, StdChars)
}

// NewLen returns a new random string of the provided length, consisting of
// standard characters.
func NewLen(length int) string {
	return NewLenChars(length, StdChars)
}

// NewLenChars returns a new random string of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
func NewLenChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("uniuri: wrong charset length for NewLenChars")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("uniuri: error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}
