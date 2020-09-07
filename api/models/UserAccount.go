package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserAccount struct {
	tableName struct{}  `sql:"user_accounts,alias:user_accounts" pg:",discard_unknown_columns"`
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	UserID    uint32    `sql:"type:int REFERENCES users(id)" json:"user_id"`
	Bank      string    `json:"bank"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *UserAccount) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *UserAccount) SaveUser(db *gorm.DB) (*UserAccount, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &UserAccount{}, err
	}
	return u, nil
}

func (u *UserAccount) FindAllUsers(db *gorm.DB) (*[]UserAccount, error) {
	var err error
	users := []UserAccount{}
	err = db.Debug().Model(&UserAccount{}).
		Limit(100).
		Find(&users).Error
	if err != nil {
		return &[]UserAccount{}, err
	}
	return &users, err
}

func (u *UserAccount) FindUserByID(db *gorm.DB, uid uint32) (*UserAccount, error) {
	var err error
	err = db.Debug().Model(UserAccount{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserAccount{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &UserAccount{}, errors.New("User Not Found")
	}
	return u, err
}
