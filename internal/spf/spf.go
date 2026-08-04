package spf

import (
	"strings"
)

func parseRecord(txtRecord []string) []string {
	txt := strings.Join(txtRecord, "")
	
	parsed := make([]string, 0)

	var builder strings.Builder
	for i := 0; i<len(txt); i++ {
			
		if (txt[i] == ' ') {
			if (builder.Len() > 0) {
				parsed = append(parsed, builder.String())
				builder.Reset()
			}
		} else {
			builder.WriteString(string(txt[i]))
		}

	}

	if builder.Len() > 0 {
		parsed = append(parsed, builder.String())
	}

	return parsed
}

func checkHost(txtRecords []string) string {
	return "none"
}

