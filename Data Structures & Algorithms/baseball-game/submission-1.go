func calPoints(operations []string) int {
	// x == push
	// + == peek+x push
	// D == peek*2 & push
	// C == pop
	record := []int{}
	sum := 0
	for _, op := range operations {
		if val, err := strconv.Atoi(op); err == nil {
			record = append(record, val)
			sum += val
		} else {
			switch op {
			case "D":
				val := 2 * record[len(record)-1]
				record = append(record, val)
				sum += val
			case "C":
				sum -= record[len(record)-1]
				record = record[:len(record)-1]
			case "+":
				val := record[len(record)-1]+record[len(record)-2]
				record = append(record,val)
				sum += val
			}
		}
	}

	return sum
}
