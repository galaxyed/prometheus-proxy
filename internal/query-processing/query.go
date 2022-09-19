package queryprocessing

import (
	"fmt"
	"regexp"
)

func UpdateQuery(query string, limitFlag string) string {
	var r *regexp.Regexp
	// Case 1: Single Match
	r = regexp.MustCompile("^\\s*([a-zA-Z\\_]+)\\s*$")

	if r.MatchString(query) {
		// fmt.Println(r.MatchString(query))
		finalString := r.ReplaceAllString(query, fmt.Sprintf("%v{%v}", query, limitFlag))
		fmt.Println(finalString)
		return finalString
	}

	// Case 2: Match ...
	r = regexp.MustCompile("({\\s*})")
	if r.MatchString(query) {
		// fmt.Println(r.MatchString(query))
		finalString := r.ReplaceAllString(query, "{"+limitFlag+"}")
		fmt.Println(finalString)
		return finalString
	}

	// Case 3: Match ...
	r = regexp.MustCompile("({\\s*)")
	if r.MatchString(query) {
		// fmt.Println("case 3")
		finalString := r.ReplaceAllString(query, "{"+limitFlag+",")
		fmt.Println(finalString)
		return finalString
	}
	return query
}
