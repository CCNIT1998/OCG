package main

import (
	"fmt"
	"math/rand"

	// "sort"
	"bufio"
	"os"
	"strconv"
	"time"
)

type Product struct {
	Name     string
	Category string
	Price    int
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
func main() {
	category := [4]string{"fashion", "electronics", "sport", "food"}
	products := [20]Product{}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(products); i++ {
		products[i] = Product{
			fmt.Sprintf("%s %d", "Product", i),
			category[rand.Intn(len(category))],
			randomInt(100, 200),
		}
	}

	for _, product := range products {
		fmt.Println(product)
	}

	// Gõ vào từ bàn phím tên một sản phẩm, trả về danh sách tìm được
	fmt.Println("_______________ TÌM SẢN PHẨM THEO TÊN _______________")
	scannerName := bufio.NewScanner(os.Stdin)
	fmt.Println("Nhập tên sản phẩm: ")
	scannerName.Scan()
	nameProduct := scannerName.Text()
	for _, product := range products {
		if product.Name == nameProduct {
			fmt.Println(product)
		}
	}

	// Gõ vào category, trả về danh sách tất cả các sản phẩm trong category đó
	fmt.Println("_______________ TÌM SẢN PHẨM THEO CATEGORY _______________")
	scannerCategory := bufio.NewScanner(os.Stdin)
	fmt.Println("Nhập tên category: ")
	scannerCategory.Scan()
	nameCategory := scannerCategory.Text()
	for _, product := range products {
		if product.Category == nameCategory {
			fmt.Println(product)
		}
	}

	// Gõ vào price min, và price max, tìm tất cả các sản phẩm có giá trong dải min đến max
	fmt.Println("_______________ TÌM SẢN PHẨM THEO PRICE _______________")
	scannerPriceMin := bufio.NewScanner(os.Stdin)
	fmt.Println("Price min: ")
	scannerPriceMin.Scan()
	priceMin, _ := strconv.Atoi(scannerPriceMin.Text())

	scannerPriceMax := bufio.NewScanner(os.Stdin)
	fmt.Println("Price max: ")
	scannerPriceMax.Scan()
	priceMax, _ := strconv.Atoi(scannerPriceMax.Text())

	fmt.Println("__ price[", priceMin, " ", priceMax, "] __")
	for _, product := range products {
		if product.Price >= priceMin && product.Price <= priceMax {
			fmt.Println(product)
		}
	}
}
