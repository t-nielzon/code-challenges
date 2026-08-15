package main

func Dirreduc(arr []string) []string {
	result := []string{}
	
	for _, dir := range arr {
		if len(result) > 0 {
			last := result[len(result)-1]
			
			if (last == "NORTH" && dir == "SOUTH") ||
				(last == "SOUTH" && dir == "NORTH") ||
				(last == "EAST" && dir == "WEST") ||
				(last == "WEST" && dir == "EAST") {
				result = result[:len(result)-1]
			} else {
				result = append(result, dir)
			}
		} else {
			result = append(result, dir)
		}
	}
	
	return result
}