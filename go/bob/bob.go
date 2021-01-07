package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if strings.ToUpper(remark) == remark && strings.Contains(remark, "?") {
		return "Calm down, I know what I'm doing!"
	} else if strings.ToUpper(remark) == remark {
		return "Whoa, chill out!"
	} else if strings.Contains(remark, "?") {
		return "Sure."
	} else if len(remark) == 0 {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
