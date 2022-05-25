package main

import (
	"fmt"
	"sesi7-gorm/database"

	// "sesi7-gorm/models"
	"sesi7-gorm/repository"
	"strings"

	"gorm.io/gorm"
)

func main() {
	db := database.StartDB()
	user(db)
	// product(db)

}

func user(db *gorm.DB) {
	userRepo := repository.NewUserRepo(db)

	// user := models.User{
	// 	Email: "test@gmail.com",
	// }

	// err := userRepo.CreateUser(&user)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// fmt.Println("Created success")

	// user := models.User{
	// 	ID:    1,
	// 	Email: "nani@gmail.com",
	// }

	// err := userRepo.UpdateUser(&user)
	// if err != nil {
	// 	fmt.Println("errorupdate : ", err.Error())
	// 	return
	// }

	user, err := userRepo.DeleteUser(21)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}
	fmt.Println("User Success for delete :", user)

	// fmt.Println("User updated")

	employees, err := userRepo.GetUsersWithProducts()
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	// employees.Print()

	for k, emp := range *employees {
		fmt.Println("User :", k+1)
		emp.Print()
		fmt.Println(strings.Repeat("=", 10))
	}

	// emp, err := userRepo.GetUserByID(3)

	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// emp.Print()
}

func product(db *gorm.DB) {
	productRepo := repository.NewProductRepo(db)

	// product := models.Product{
	// 	Name:   "Hoodie",
	// 	Brand:  "Zara",
	// 	UserID: 1,
	// }

	// //create product
	// err := productRepo.CreateProduct(&product)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("Created Product Success !")

	products, err := productRepo.GetAllProduct()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for k, product := range *products {
		fmt.Println("Product :", k+1)
		product.Print()
		fmt.Println(strings.Repeat("=", 10))

		fmt.Println("selesai")
	}
}
