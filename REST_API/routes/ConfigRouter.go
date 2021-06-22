package routes

import (
	"github.com/CCNIT1998/OCG/REST_API/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigUserRouter(router *fiber.Router) {

	(*router).Get("/", controller.GetAllUser) //Liệt kê
	(*router).Get("/:id", controller.FindUserById) //Xem chi tiết một bản ghi
	(*router).Post("", controller.CreateUser) //INSERT: Tạo một bản ghi
	(*router).Put("", controller.UpsertUser) //UPSERT: Cập nhật một bản ghi nếu tìm thấy còn không tạo mới
	(*router).Patch("", controller.UpdateUser) //UPDATE: Cập nhật một bản ghi
	(*router).Delete("/:id", controller.DeleteUserById) //Xoá một bản ghi
}


func ConfigProductRouter(router *fiber.Router) {

	(*router).Post("", controller.CreateNewProduct) //INSERT: Tạo một bản ghi
	(*router).Put("", controller.UpsertProduct) //UPSERT: Cập nhật một bản ghi nếu tìm thấy còn không tạo mới
	(*router).Patch("", controller.UpdateProduct) //UPDATE: Cập nhật một bản ghi
	(*router).Delete("/:id", controller.DeleteProductById) //Xoá một bản ghi
	(*router).Get("", controller.GetAllProduct) //Liệt kê
	(*router).Patch("/image/:id", controller.UpdateImageInProduct) // update image
	// data test UpdateImageInProduct
	// {
	// 	"id": 1,
	// 	"image": [
	// 		"/uploads/images/img1.jpg",
	// 		"/uploads/images/img2.jpg",
	// 		"/uploads/images/img3.jpg"
	// 	]
	// }
	(*router).Patch("/delete/image/:id", controller.DeleteEachImageInProduct) // delete image
	(*router).Patch("/review/:id", controller.UpdateReviewInProduct) // update review
	(*router).Patch("/historyprice/:id", controller.UpdatePriceStoreInProduct) // update History price

	// {
	// 	"priceHistory": [
	// 		"01/04/2021",
	// 		"02/05/2021"
	// 	]
	// }
}



func ConfigCategoryRouter(router *fiber.Router) {

	(*router).Post("", controller.CreateNewCategory) //INSERT: Tạo một bản ghi
	(*router).Put("", controller.UpsertCategory) //UPSERT: Cập nhật một bản ghi nếu tìm thấy còn không tạo mới
	(*router).Patch("", controller.UpdateCategory) //UPDATE: Cập nhật một bản ghi
	(*router).Delete("/:id", controller.DeleteCategoryById) //Xoá một bản ghi
	(*router).Get("", controller.GetAllCategory) //Liệt kê
	
}