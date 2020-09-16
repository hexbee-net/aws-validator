package validator

import (
	"strings"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

const dnsDelimiter = "."
const slashSeparator = "/"

// Bad path components to be rejected by the path validity handler.
const (
	dotDotComponent = ".."
	dotComponent    = "."
)

// IsValidBucketName verifies that a bucket name is in accordance with
// Amazon's requirements (i.e. DNS naming conventions). It must be 3-63
// characters long, and it must be a sequence of one or more labels
// separated by periods. Each label can contain lowercase ascii
// letters, decimal digits and hyphens, but must not begin or end with
// a hyphen. See:
// http://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html
func IsValidBucketName(fl validator.FieldLevel) bool {
	bucket := fl.Field().String()

	if len(bucket) < 3 || len(bucket) > 63 {
		return false
	}

	// Split on dot and check each piece conforms to rules.
	pieces := strings.Split(bucket, dnsDelimiter)

	// Does the bucket name look like an IP address?
	return !(len(pieces) == 4 && isAllNumbers(pieces))
}

func isAllNumbers(pieces []string) bool {
	allNumbers := true

	for _, piece := range pieces {
		if piece == "" || piece[0] == '-' ||
			piece[len(piece)-1] == '-' {
			// Current piece has 0-length or starts or ends with a hyphen.
			return false
		}

		// Now only need to check if each piece is a valid 'label' in AWS terminology and if the bucket looks
		// like an IP address.
		isNotNumber := false

		for i := 0; i < len(piece); i++ {
			switch {
			case piece[i] >= 'a' && piece[i] <= 'z' || piece[i] == '-':
				// Found a non-digit character, so this piece is not a number.
				isNotNumber = true
			case piece[i] >= '0' && piece[i] <= '9':
				// Nothing to do.
			default:
				// Found invalid character.
				return false
			}
		}

		allNumbers = allNumbers && !isNotNumber
	}

	return allNumbers
}

// IsValidObjectName verifies an object name in accordance with Amazon's
// requirements. It cannot exceed 1024 characters and must be a valid UTF8
// string.
//
// See:
// http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html
func IsValidObjectName(fl validator.FieldLevel) bool {
	object := fl.Field().String()

	if object == "" {
		return false
	}

	return isValidObjectPrefix(object)
}

// isValidObjectPrefix verifies whether the prefix is a valid object name.
// Its valid to have a empty prefix.
func isValidObjectPrefix(object string) bool {
	if hasBadPathComponent(object) {
		return false
	}

	if !utf8.ValidString(object) {
		return false
	}

	if strings.Contains(object, `//`) {
		return false
	}

	return true
}

// Check if the incoming path has bad path components, such as ".." and ".".
func hasBadPathComponent(path string) bool {
	path = strings.TrimSpace(path)
	for _, p := range strings.Split(path, slashSeparator) {
		switch strings.TrimSpace(p) {
		case dotDotComponent:
			return true
		case dotComponent:
			return true
		}
	}

	return false
}
