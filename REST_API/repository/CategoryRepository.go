package repository

import (
	"errors"
	"fmt"

	"github.com/CCNIT1998/OCG/REST_API/model"
)

type CategoryRepository struct {
	categorys map[int64]*model.Category
	autoID    int64
}

var CategoryRepo CategoryRepository

func (r *CategoryRepository) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *CategoryRepository) InitData(connection string) {
	fmt.Println("Connect to ", connection)
	r.CreateNewCategoryRepo(&model.Category{
		Id:   0,
		Name: "All",
	})

	r.CreateNewCategoryRepo(&model.Category{
		Id:   1,
		Name: "Women",
	})

	r.CreateNewCategoryRepo(&model.Category{
		Id:   2,
		Name: "Men",
	})

	r.CreateNewCategoryRepo(&model.Category{
		Id:   3,
		Name: "Kids",
	})
}

func (r *CategoryRepository) CreateNewCategoryRepo(category *model.Category) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	category.Id = nextID
	r.categorys[nextID] = category //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	CategoryRepo.categorys = make(map[int64]*model.Category)
	CategoryRepo.autoID = 0

	CategoryRepo.InitData("sql:45312")
}

func (r *CategoryRepository) UpsertCategoryRepo(category *model.Category) int64 {
	if _, ok := r.categorys[category.Id]; ok {
		r.categorys[category.Id] = category //tìm thấy thì update
		return category.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewCategoryRepo(category)
	}
}

func (r *CategoryRepository) UpdateCategoryRepo(category *model.Category) error {
	if _, ok := r.categorys[category.Id]; ok {
		r.categorys[category.Id] = category
		return nil //tìm được
	} else {
		return errors.New("category not found")
	}
}

func (r *CategoryRepository) DeleteCategoryByIdRepo(Id int64) error {
	if _, ok := r.categorys[Id]; ok {
		delete(r.categorys, Id)
		return nil
	} else {
		return errors.New("category not found")
	}
}

func (r *CategoryRepository) GetAllCategoryRepo() map[int64]*model.Category {
	return r.categorys
}
