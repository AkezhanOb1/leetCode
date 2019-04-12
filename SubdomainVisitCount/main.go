package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}

func subdomainVisits(cpdomains []string) []string {
	counter := make(map[string]int)
	for _, v := range cpdomains {
		s := strings.Split(v, " ")
		amount, err := strconv.Atoi(s[0])
		if err != nil {
			log.Println(err)
		}
		if val, ok := counter[s[1]]; ok {
			counter[s[1]] = val + amount
		} else {
			counter[s[1]] = amount
		}

		in := strings.Index(s[1], ".")
		for in != -1 {
			s[1] = s[1][in+1:]
			if val, ok := counter[s[1]]; ok {
				counter[s[1]] = val + amount
			} else {
				counter[s[1]] = val + amount
			}
			in = strings.Index(s[1], ".")
		}
	}

	result := []string{}
	for key, value := range counter {
		result = append(result, fmt.Sprintf("%d %s", value, key))
	}
	return result
}
