package utils

import "github.com/gosimple/slug"

func GenerateSlug(str string) string {
	return slug.Make(str)
}
