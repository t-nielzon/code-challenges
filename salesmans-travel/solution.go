package kata

import "strings"

func Travel(r string, zipcode string) string {
	if r == "" {
		return zipcode + ":/"
	}
	var streets []string
	var numbers []string
	for _, addr := range strings.Split(r, ",") {
		if strings.HasSuffix(addr, zipcode) {
			body := strings.TrimSuffix(addr, " "+zipcode)
			parts := strings.SplitN(body, " ", 2)
			if len(parts) < 2 {
				continue
			}
			numbers = append(numbers, parts[0])
			streets = append(streets, parts[1])
		}
	}
	if len(streets) == 0 {
		return zipcode + ":/"
	}
	return zipcode + ":" + strings.Join(streets, ",") + "/" + strings.Join(numbers, ",")
}