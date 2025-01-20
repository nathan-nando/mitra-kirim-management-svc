package user

import "time"

type User struct {
	ID          int        `db:"id"`
	Name        string     `db:"name"`
	Title       string     `db:"title"`
	Email       string     `db:"email"`
	Phone       string     `db:"phone"`
	Gender      string     `db:"gender"`
	Img         string     `db:"img"`
	Status      int        `db:"status"`
	CreatedDate time.Time  `db:"created_date"`
	CreatedBy   string     `db:"created_by"`
	UpdatedDate *time.Time `db:"created_date"`
	UpdatedBy   string     `db:"created_by"`
}
