package main

import (
	"sort"
	"strings"
)

type Guest struct {
	first, last string
}

func meeting(s string) string {
	names := strings.Split(s, ";")
	
	var guests []Guest
	for _, name := range names {
		parts := strings.Split(name, ":")
		if len(parts) == 2 {
			guests = append(guests, Guest{
				first: strings.ToUpper(parts[0]),
				last:  strings.ToUpper(parts[1]),
			})
		}
	}
	
	sort.Slice(guests, func(i, j int) bool {
		if guests[i].last != guests[j].last {
			return guests[i].last < guests[j].last
		}
		return guests[i].first < guests[j].first
	})
	
	var result strings.Builder
	for _, guest := range guests {
		result.WriteString("(")
		result.WriteString(guest.last)
		result.WriteString(", ")
		result.WriteString(guest.first)
		result.WriteString(")")
	}
	
	return result.String()
}