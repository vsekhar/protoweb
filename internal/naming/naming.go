package naming

import "strings"

// ProtoFieldName returns a representation of s that is suitable
// to be a field name for a protocol buffer.
//
// E.g. ProtoFieldName("Accept-Encoding") == "accept_encoding"
func ProtoFieldName(s string) string {
	s = strings.ToLower(s)
	replacements := []string{
		"/", "_",
		".", "_",
		"+", "_PLUS_",
		" ", "_",
		"-", "_",
		"(", "_",
		")", "_",
		";", "_",
		":", "_",
		"*", "_STAR_",
		"\n", "_",
		" ", "_",
	}
	repl := strings.NewReplacer(replacements...)
	s = repl.Replace(s)
	s = strings.Trim(s, "_")
	return s
}

// ProtoEnumName returns a representation of s that is suitable
// to be an enum name in a protocol buffer.
//
// E.g. ProtoEnumName("utf-8") == "UTF_8"
func ProtoEnumName(s string) string {
	s = strings.ToUpper(s)
	replacements := []string{
		"/", "_",
		".", "_",
		"+", "_PLUS_",
		" ", "_",
		"-", "_",
		"(", "_",
		")", "_",
		";", "_",
		":", "_",
		"*", "_STAR_",
		"\n", "_",
		" ", "_",
	}
	repl := strings.NewReplacer(replacements...)
	s = repl.Replace(s)
	s = strings.Trim(s, "_")
	return s
}
