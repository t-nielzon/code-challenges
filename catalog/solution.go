package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type item struct {
	Name     string `xml:"name"`
	Price    string `xml:"prx"`
	Quantity string `xml:"qty"`
}

type catalog struct {
	Items []item `xml:"prod"`
}

func Catalog(s, article string) string {
	var cat catalog
	
	xmlStr := "<catalog>" + s + "</catalog>"
	
	err := xml.Unmarshal([]byte(xmlStr), &cat)
	if err != nil {
		return "Nothing"
	}
	
	var results []string
	
	for _, product := range cat.Items {
		if strings.Contains(product.Name, article) {
			line := fmt.Sprintf("%s > prx: $%s qty: %s", product.Name, product.Price, product.Quantity)
			results = append(results, line)
		}
	}
	
	if len(results) == 0 {
		return "Nothing"
	}
	
	return strings.Join(results, "\n") + "\n"
}