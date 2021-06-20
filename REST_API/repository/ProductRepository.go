package repository

import (
	"fmt"
	"time"
	"errors"

	"github.com/CCNIT1998/OCG/REST_API/model"
)

type ProductRepository struct {
	products map[int64]*model.Product
	autoID int64
}

var ProductRepo ProductRepository

func (r *ProductRepository) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ProductRepository) InitData(connection string) {
	fmt.Println("Connect to ", connection)
	timeNow := time.Now().Local().Format("01/02/2006")
	r.CreateNewProductRepo(&model.Product{
		Id: 1,
		CategoryId: 2,
		Image: []string{"/uploads/images/item-02.jpg", "/uploads/images/item-03.jpg"},
		Name: "Herschel supply co 25l",
		Price: 75,
		IsSale: true,
		CreatedAt: string(timeNow),
		ModifiedAt: string(timeNow),
	})
}

func (r *ProductRepository) CreateNewProductRepo(product *model.Product) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	product.Id = nextID
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
		r.products[product.Id] = product //tìm thấy thì update
		return product.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewProductRepo(product)
	}
}


func (r *ProductRepository) UpdateProductRepo(product *model.Product) error {
	if _, ok := r.products[product.Id]; ok {
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
	return r.products
}


func (r *ProductRepository) UpdateProductRepo1(product *model.Product, Id int64) error {
	// fmt.Println(product)
	if _, ok := r.products[Id]; ok {
		image := product.Image
		r.products[Id].Image = append(r.products[Id].Image, image...)
		return nil //tìm được
	} else {
		return errors.New("product not found")
	}
}
