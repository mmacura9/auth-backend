package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ChooseCruise/choosecruise-backend/domain"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Queries: New(db), db: db}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("tx err %v, rollback error %v", err, rberr)
		}
		return err
	}

	return tx.Commit()
}

type UserRepositoryStruct struct {
	db Store
}

func NewUserRepository(db Store) domain.UserRepository {
	return UserRepositoryStruct{db: db}
}

func (urs UserRepositoryStruct) Create(c context.Context, user *domain.User) error {
	usr := CreateUserParams{
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		Password:  user.Password,
		BirthDate: user.BirthDate,
	}
	fmt.Println(usr)
	_, err := urs.db.CreateUser(c, usr)
	return err
}

func (urs UserRepositoryStruct) Fetch(c context.Context) ([]domain.User, error) {

	return nil, nil
}

func (urs UserRepositoryStruct) GetByEmail(c context.Context, email string) (domain.User, error) {
	usr, err := urs.db.GetUserByEmail(c, email)
	if err != nil {
		return domain.User{}, err
	}
	output := domain.User{
		ID:        usr.ID,
		Username:  usr.Username,
		Password:  usr.Password,
		Email:     usr.Email,
		FullName:  usr.FullName,
		BirthDate: usr.BirthDate,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		LastLogin: usr.LastLogin,
	}
	return output, err
}

func (urs UserRepositoryStruct) GetByID(c context.Context, id string) (domain.User, error) {
	return domain.User{}, nil
}

// func (r *GetEventsFromUserRow) Scan(scanFn func(dest ...interface{}) error) error {
// 	var eventTimeStr string
// 	err := scanFn(&r.ID, &r.IDUser, &r.Title, &eventTimeStr, &r.ID_2, &r.Username, &r.Fullname, &r.Password)
// 	if err != nil {
// 		return err
// 	}

// 	parsedTime, err := time.Parse(time.RFC3339Nano, eventTimeStr)
// 	if err != nil {
// 		return err
// 	}

// 	r.EventTime = parsedTime

// 	return nil
// }
