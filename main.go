package main

func main() {


}

func calculatePrice(books []string) (float64, error) {
	resultingPrice := 0.0
	bookmap := make(map[string]int)

	for _, book := range books {
		bookmap[book] += 1
	}


	for len(bookmap) != 0 {
		diffBookCount := 0
		iterationPrice := 0.0
		for book, count := range bookmap {

			if count == 0 {
				delete(bookmap, book)
				continue
			}
			iterationPrice = iterationPrice + 8.0
			bookmap[book] = bookmap[book] - 1
			diffBookCount = diffBookCount + 1
		}

		discount := 1.0
		switch diffBookCount {
		case 1:
			discount = 0.0
		case 2:
			discount = 0.05
		case 3:
			discount = 0.10
		case 4:
			discount = 0.20
		case 5:
			discount = 0.25
		}

		resultingPrice = resultingPrice + (iterationPrice - iterationPrice * discount)
	}



	return resultingPrice, nil
}
