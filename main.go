package main

import (
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/repository"
	"backend-food/pkg/infrastucture/db"
	"fmt"
)

func main() {
	db, err := db.NewDB()
	fmt.Println(err)
	repo := repository.NewUserRepository(db)
	//	repo.CreateUser(entity.Users{Username: "abc123", Password: "12dfgf34556"})
	l, _ := repo.FindUserList(entity.Users{Username: "lao123455"})
	for _, u := range l {
		fmt.Println(u)
	}

	// //repo.CreateUser(entity.Users{Username: "abc123", Password: "12dfgf34556"})
	//err = repo.UpdateUser(entity.Users{ID: l[0].ID}, entity.Users{Password: "156", Username: "lao123455"})
	fmt.Println(err)

}
