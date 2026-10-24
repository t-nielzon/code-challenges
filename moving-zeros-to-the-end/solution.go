package kata

func MoveZeros(arr []interface{}) []interface{} {
	zeroCount := 0
	for _, v := range arr {
		switch val := v.(type) {
		case int:
			if val == 0 {
				zeroCount++
			}
		case int32:
			if val == 0 {
				zeroCount++
			}
		case int64:
			if val == 0 {
				zeroCount++
			}
		case float32:
			if val == 0 {
				zeroCount++
			}
		case float64:
			if val == 0 {
				zeroCount++
			}
		}
	}
	
	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		isZero := false
		switch val := v.(type) {
		case int:
			if val == 0 {
				isZero = true
			}
		case int32:
			if val == 0 {
				isZero = true
			}
		case int64:
			if val == 0 {
				isZero = true
			}
		case float32:
			if val == 0 {
				isZero = true
			}
		case float64:
			if val == 0 {
				isZero = true
			}
		}
		
		if !isZero {
			result = append(result, v)
		}
	}
	
	for i := 0; i < zeroCount; i++ {
		result = append(result, 0)
	}
	
	return result
}