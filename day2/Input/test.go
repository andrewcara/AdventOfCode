package main

func main() {

	row := make([]int, 0)
	// row = append(row, 5)
	row = append(row, 72)
	row = append(row, 74)
	row = append(row, 75)
	row = append(row, 77)
	row = append(row, 80)
	row = append(row, 81)
	row = append(row, 81)
	//72 74 75 77 80 81 81
	//45 47 49 53 54
	increasing := row[0] < row[len(row)-1] || row[1] < row[len(row)-1]
	prev_val := -1
	incorrect := 0
	total := 0
	// nums_reset := 0
	for i, _ := range row {

		if increasing {
			//condition is not met

			for index, row_val := range row {
				if i == index {
					//println(i)
				} else if prev_val == -1 {
					prev_val = row_val
				} else if !(row_val > prev_val) || !(row_val-prev_val < 4) {
					incorrect++
				} else {
					prev_val = row_val
				}
			}

			if incorrect > 0 {
				total++
			}
			incorrect = 0
			prev_val = -1
		}

	}
	println(total)
}
