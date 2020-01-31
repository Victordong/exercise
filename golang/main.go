package main

func main() {
	price := [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}

	print(findCheapestPrice(4, price, 0, 3, 1))
}
