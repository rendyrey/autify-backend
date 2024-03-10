package cmd

import (
	"strings"
)

func removeFlagsInParams(params []string) []string {
	sanitizedParams := []string{}

	for _, val := range params {
		if strings.HasPrefix(val, "-") || strings.HasPrefix(val, "--") {
			continue
		}

		sanitizedParams = append(sanitizedParams, val)
	}

	return sanitizedParams
}
