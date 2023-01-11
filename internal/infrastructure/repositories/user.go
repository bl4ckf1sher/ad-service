package repositories

import (
	"context"
	"github.com/bl4ckf1sher/ad-service/internal/domain"
	"github.com/bl4ckf1sher/ad-service/internal/infrastructure/postgres"
	"github.com/google/uuid"
)

type User struct {
	connect *postgres.Connect
}

func NewUserRepository(connect *postgres.Connect) *User {
	return &User{connect: connect}
}

func (u User) Get(c context.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User

	err := u.connect.Db.GetContext(c, &user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) GetAll(c context.Context) (*[]domain.User, error) {
	var users []domain.User

	err := u.connect.Db.SelectContext(c, &users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u User) Create(c context.Context, user domain.User) (err error) {
	createUserQuery := "INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, now(), now())"

	res := u.connect.Db.MustExecContext(
		c,
		createUserQuery,
		uuid.New(),
		user.Role, user.Email, user.Password, user.Name, user.Surname,
	)

	_, err = res.RowsAffected()

	return
}

func (u User) Update(c context.Context, user domain.User) (err error) {
	updateUserQuery := "UPDATE users SET (role, email, password, name, surname, updated_at) = " +
		"($1, $2, $3, $4, $5, now()) WHERE id=$6"
	res := u.connect.Db.MustExecContext(
		c,
		updateUserQuery,
		user.Role, user.Email, user.Password, user.Name, user.Surname, user.ID,
	)

	_, err = res.RowsAffected()

	return
}

func (u User) Delete(c context.Context, id uuid.UUID) (err error) {
	res := u.connect.Db.MustExecContext(c, "DELETE FROM users WHERE id=$1", id)

	_, err = res.RowsAffected()

	return
}
