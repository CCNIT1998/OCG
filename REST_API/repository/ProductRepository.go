package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/CCNIT1998/OCG/REST_API/model"
)

type ProductRepository struct {
	products map[int64]*model.Product
	autoID   int64
}

var ProductRepo ProductRepository

func (r *ProductRepository) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

//
// type ReviewRepository struct {
// 	review map[int64]*model.Review
// 	autoID int64
// }

// var ReviewRepo ReviewRepository

// func (r *ProductRepository) getAutoIDReview() int64 {
// 	r.autoID += 1
// 	return r.autoID
// }
// func (r *ProductRepository) CreateNewReviewInProductRepo(product *model.Product) int64 {
// 	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
// 	product.Review = nextID

// 	updateRating(product)
// 	r.products[nextID] = product //tạo mới một phần tử trong map, gán key bằng nextID

// 	return nextID
// }

//

func (r *ProductRepository) InitData(connection string) {
	fmt.Println("Connect to ", connection)
	timeNow := time.Now().Local().Format("01/02/2006")
	r.CreateNewProductRepo(&model.Product{
		Id:         1,
		CategoryId: 2,
		Image:      []string{"/uploads/images/item-02.jpg", "/uploads/images/item-03.jpg"},
		Name:       "Herschel supply co 25l",
		Price:      75,
		IsSale:     true,
		CreatedAt:  string(timeNow),
		ModifiedAt: string(timeNow),
		Review: []model.Review{
			{Id: 1, ProductId: 1, Comment: "bad", Rating: 3},
			{Id: 2, ProductId: 1, Comment: "Good", Rating: 5},
		},
		Rating: 4,
		PriceHistory: []string{
			"01/02/2020", 
			"01/02/2021",
		}, 
	})
}

func updateRating(product *model.Product) {
	RatingTotal := 0
	ReviewTotal := 0
	for _, value := range product.Review {
		ReviewTotal++
		RatingTotal += value.Rating
	}
	product.Rating = float64(RatingTotal) / float64(ReviewTotal)
}

func (r *ProductRepository) CreateNewProductRepo(product *model.Product) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	product.Id = nextID

	updateRating(product)
	r.products[nextID] = product //tạo mới một phần tử trong map, gán key bằng nextID

	return nextID
}

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	ProductRepo.products = make(map[int64]*model.Product)
	ProductRepo.autoID = 0

	ProductRepo.InitData("sql:45312")
}

func (r *ProductRepository) UpsertProductRepo(product *model.Product) int64 {
	if _, ok := r.products[product.Id]; ok {
		updateRating(product)
		r.products[product.Id] = product //tìm thấy thì update
		return product.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewProductRepo(product)
	}
}

func (r *ProductRepository) UpdateProductRepo(product *model.Product) error {
	if _, ok := r.products[product.Id]; ok {
		updateRating(product)
		r.products[product.Id] = product
		return nil //tìm được
	} else {
		return errors.New("product not found")
	}
}

func (r *ProductRepository) DeleteProductkByIdRepo(Id int64) error {
	if _, ok := r.products[Id]; ok {

		delete(r.products, Id)
		return nil
	} else {
		return errors.New("product not found")
	}
}

func (r *ProductRepository) GetAllProductRepo() map[int64]*model.Product {
	for _, product := range r.products {
		updateRating(product)
	}
	return r.products
}

// Product -> Image
func (r *ProductRepository) UpdateImageInProductRepo(product *model.Product, Id int64) error {
	// fmt.Println(product)
	
	if _, ok := r.products[Id]; ok {
		image := product.Image
		// fmt.Println(product.Image)
		r.products[Id].Image = append(r.products[Id].Image, image...)
		return nil //tìm được
	} else {
		return errors.New("product not found")
	}
}

func (r *ProductRepository) DeleteEachImageInProductRepo(product *model.Product, Id int64) error {
	// fmt.Println(product)
	imageTxt := []string{}
	image := product.Image
	for _,i := range r.products[Id].Image{
		temp := 0
		for _,j := range image{
			if(i==j){
				temp++
			}
		}
		if temp == 0 {
			imageTxt = append(imageTxt, i)
		}
	}
	if _, ok := r.products[Id]; ok {
		
		// fmt.Println(product.Image)
		if imageTxt != nil {
			r.products[Id].Image = imageTxt
			return nil
		}
		
		return nil //tìm được
	} else {
		return errors.New("product not found")
	}
}


// Product Review
func (r *ProductRepository) UpdateReviewInProductRepo(product *model.Product, Id int64) error {
	// fmt.Println(product)
	if _, ok := r.products[Id]; ok {
		review := product.Review
		r.products[Id].Review = append(r.products[Id].Review, review...)
		updateRating(product)
		for _, p := range product.Review {
			p.ProductId = Id
		}
		return nil //tìm được
	} else {
		return errors.New("product not found")
	}
}


func (r *ProductRepository) UpdatePriceStoreInProductRepo(product *model.Product, Id int64) error {
	// fmt.Println(product)
	
	if _, ok := r.products[Id]; ok {
		priceHistory := product.PriceHistory
		fmt.Println(product.PriceHistory)
		r.products[Id].PriceHistory = append(r.products[Id].PriceHistory, priceHistory...)
		r.products[Id].ModifiedAt = r.products[Id].PriceHistory[len(r.products[Id].PriceHistory)-1]
		fmt.Println(r.products[Id])
		return nil //tìm được
	} else {
		return errors.New("product not found")
	}
}