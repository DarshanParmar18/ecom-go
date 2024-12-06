package user

import (
	"database/sql"
	"fmt"

	"github.com/darshanparmar18/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUsers() (*types.User, error) {
	user := new(types.User)
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		user, err = scanRowIntoUsers(rows)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

// --------------------------------- //  Get the user by his email  // --------------------------------- //
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	user := new(types.User)

	stmt, err := s.db.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		user, err = scanRowIntoUsers(rows) // helper func used here
		if err != nil {
			return nil, err
		}
	}

	// check if user exists
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	// if user exists the return the user
	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	// empty signatures
	return nil, nil
}
func (s *Store) CreateUser(user types.User) error {
	// empty signatures
	return nil
}

// --------------------------------- helper function to scan into User rows --------------------------------- //
func scanRowIntoUsers(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	if err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	); err != nil {
		return nil, err
	}

	return user, nil
}
