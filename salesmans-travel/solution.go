package main

import (
	"strings"
)

func travel(r string, zipcode string) string {
	if r == "" {
		return zipcode + ":/"
	}

	addresses := strings.Split(r, ",")
	var streets []string
	var houses []string

	for _, addr := range addresses {
		addr = strings.TrimSpace(addr)

		if len(addr) < 8 {
			continue
		}

		addrZipcode := addr[len(addr)-8:]
		if addrZipcode != zipcode {
			continue
		}

		parts := strings.Fields(addr)
		if len(parts) < 2 {
			continue
		}

		house := parts[0]
		streetAndTown := addr[len(house)+1 : len(addr)-8]
		streetAndTown = strings.TrimSpace(streetAndTown)

		streets = append(streets, streetAndTown)
		houses = append(houses, house)
	}

	if len(streets) == 0 {
		return zipcode + ":/"
	}

	return zipcode + ":" + strings.Join(streets, ",") + "/" + strings.Join(houses, ",")
}