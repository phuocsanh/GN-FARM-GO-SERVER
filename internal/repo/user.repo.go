package repo

import (
	"go_ecommerce/global"
	"go_ecommerce/internal/database"
)

// type UserRepo struct {}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser()string {
// 	return "sanh 123"
// }

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("user_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	// fmt.Print("row", row)
	// user, err := u.sqlc.GetUserByEmailSQLC(ctx,email)
	// if err != nil {
	// 	return false
	// }
	// return user.UserID != NumberNull
	return  1!=0
}

func NewUserRepository() IUserRepository { return &userRepository{
	sqlc: database.New(global.Mdbc),
} }
