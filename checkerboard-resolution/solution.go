package kata

func Checkerboard(width, height, resolution int64) int64 {
	if width == 0 || height == 0 {
		return 0
	}

	full_w := width / resolution
	full_h := height / resolution
	rem_w := width % resolution
	rem_h := height % resolution

	var count int64 = 0

	// full grid of resolution×resolution cells
	if full_w > 0 && full_h > 0 {
		full_grid_black := (full_w * full_h) / 2
		count += full_grid_black * resolution * resolution
	}

	// partial column (if rem_w > 0)
	if rem_w > 0 && full_h > 0 {
		col_black := full_h / 2
		if full_w%2 == 1 {
			col_black = (full_h + 1) / 2
		}
		count += col_black * rem_w * resolution
	}

	// partial row (if rem_h > 0)
	if rem_h > 0 && full_w > 0 {
		row_black := full_w / 2
		if full_h%2 == 1 {
			row_black = (full_w + 1) / 2
		}
		count += row_black * resolution * rem_h
	}

	// partial corner cell (if both rem_w > 0 and rem_h > 0)
	if rem_w > 0 && rem_h > 0 {
		if (full_w+full_h)%2 == 1 {
			count += rem_w * rem_h
		}
	}

	return count
}