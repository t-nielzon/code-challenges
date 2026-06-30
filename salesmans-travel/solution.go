package kata

import "strings"

func travel(r string, zipcode string) string {
	var streets []string
	var numbers []string

	if r != "" {
		for _, addr := range strings.Split(r, ",") {
			if !strings.HasSuffix(addr, zipcode) {
				continue
			}

			body := strings.TrimSuffix(addr, zipcode)
			body = strings.TrimRight(body, " ")

			parts := strings.SplitN(body, " ", 2)
			if len(parts) < 2 {
				continue
			}

			numbers = append(numbers, parts[0])
			streets = append(streets, parts[1])
		}
	}

	return zipcode + ":" + strings.Join(streets, ",") + "/" + strings.Join(numbers, ",")
}