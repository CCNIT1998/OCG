package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/CCNIT1998/OCG/REST_API/model"
)

type UserRepository struct {
	users  map[int64]*model.User
	autoID int64 //đây là biến đếm tự tăng gán giá trị cho id của User
}

var UserRepo UserRepository //Khai báo biến toàn cục, global variable

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	UserRepo = UserRepository{autoID: 0}
	UserRepo.users = make(map[int64]*model.User)
	UserRepo.InitData("sql:45312")
}

//Pointer receiver ~ method trong Java. Đối tượng chủ thể là *BookRepo
func (r *UserRepository) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *UserRepository) CreateNewUser(user *model.User) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	user.Id = nextID
	r.users[nextID] = user //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}

func (r *UserRepository) InitData(connection string) {
	fmt.Println("Connect to ", connection)
	timeNow := time.Now().Local().Format("01/02/2006")
	r.CreateNewUser(&model.User{
		FirstName: "Administrator",
		Id: 1,
		LastName: "",
		Username: "admin",
		Email: "admin@gmail.com",
		Password: "admin",
		Avatar: "https://robohash.org/eaquequasincidunt.png?size=50x50&set=set1",
		Gender: "Genderfluid",
		Phone: "933-658-1213",
		Birthday: "1994-03-23",
		Status: "true",
		CreatedAt: string(timeNow),
		ModifiedAt: string(timeNow),
	})
}

func (r *UserRepository) GetAllUserRepo() map[int64]*model.User {
	return r.users
}

func (r *UserRepository) FindUserByIdRepo(Id int64) (*model.User, error) {
	if user, ok:= r.users[Id]; ok{
		return user, nil //tìm được
	} else {
		return nil, errors.New("user not found")
	}
}

func (r *UserRepository) UpsertUserRepo(user *model.User) int64 {
	if _, ok := r.users[user.Id]; ok {
		r.users[user.Id] = user //tìm thấy thì update
		return user.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewUser(user)
	}
}

func (r *UserRepository) UpdateUserRepo(user *model.User) error {
	if _, ok := r.users[user.Id]; ok {
		r.users[user.Id] = user
		return nil //tìm được
	} else {
		return errors.New("user not found")
	}
}

func (r *UserRepository) DeleteUserkByIdRepo(Id int64) error {
	if _, ok := r.users[Id]; ok {
		delete(r.users, Id)
		return nil
	} else {
		return errors.New("user not found")
	}
}




// //Cập nhật average rating của Book
// func (r *BookRepository) Update(bookId int64, averageRating float32) error {
// 	//TODO: cập nhật dữ liệu ở đây
// }
