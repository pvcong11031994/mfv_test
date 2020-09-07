package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type UserTransaction struct {
	tableName       struct{}  `sql:"user_transactions,alias:user_transactions" pg:",discard_unknown_columns"`
	ID              uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserAccountID   uint32    `sql:"type:int REFERENCES user_accounts(id)" json:"account_id"`
	Amount          float32   `json:"amount"`
	Bank            string    `json:"bank"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserTransactionReponse struct {
	ID              int64     `json:"id"`
	UserAccountID   int64     `json:"account_id"`
	Amount          float32   `json:"amount"`
	Bank            string    `json:"bank"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
}

func (u *UserTransaction) Prepare() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *UserTransaction) PrepareUpdate() {
	u.UpdatedAt = time.Now()
}

func (p *UserTransaction) Validate() error {

	if p.UserAccountID < 1 {
		return errors.New("Required User Account ID")
	}
	if p.Amount < 1 {
		return errors.New("Required Amount")
	}
	if p.TransactionType == "" {
		return errors.New("Required Transaction Type")
	}
	return nil
}

func (p *UserTransaction) ValidateUpdate() error {

	if p.UserAccountID < 1 {
		return errors.New("Required User Account ID")
	}
	return nil
}

func (p *UserTransaction) SavePost(db *gorm.DB) (*UserTransaction, error) {
	var err error
	err = db.Debug().Model(&UserTransaction{}).Create(&p).Error
	if err != nil {
		return &UserTransaction{}, err
	}
	return p, nil
}

func (p *UserTransaction) FindUserTransactionByUserId(db *gorm.DB, userID uint64) (*[]UserTransactionReponse, error) {
	var err error
	posts := []UserTransactionReponse{}
	err = db.Debug().Model(&UserTransaction{}).
		Table("user_transactions").
		Joins("JOIN user_accounts ON user_accounts.id = user_transactions.user_account_id").
		Joins("JOIN users ON users.id = user_accounts.user_id").
		Select("user_transactions.id, user_accounts.id as user_account_id, user_transactions.amount, user_accounts.bank, user_transactions.transaction_type, user_transactions.created_at").
		Where("users.deleted_at IS NULL").
		Where("users.id = ?", userID).
		Limit(100).
		Find(&posts).Error
	if err != nil {
		return &[]UserTransactionReponse{}, err
	}
	return &posts, nil
}

func (p *UserTransaction) FindUserTransactionByUserIdUserTransactionId(db *gorm.DB, userID uint64, userAccountID uint64) (*[]UserTransactionReponse, error) {
	var err error
	posts := []UserTransactionReponse{}
	err = db.Debug().Model(&UserTransaction{}).
		Table("user_transactions").
		Joins("JOIN user_accounts ON user_accounts.id = user_transactions.user_account_id").
		Joins("JOIN users ON users.id = user_accounts.user_id").
		Select("user_transactions.id, user_accounts.id as user_account_id, user_transactions.amount, user_accounts.bank, user_transactions.transaction_type, user_transactions.created_at").
		Where("users.deleted_at IS NULL").
		Where("users.id = ?", userID).
		Where("user_accounts.id = ?", userAccountID).
		Find(&posts).Error
	if err != nil {
		return &[]UserTransactionReponse{}, err
	}

	return &posts, nil
}

func (p *UserTransaction) UpdateAPost(db *gorm.DB) (*UserTransaction, error) {

	var err error
	err = db.Debug().Model(&UserTransaction{}).Where("id = ?", p.ID).Updates(UserTransaction{UpdatedAt: time.Now()}).Error
	if err != nil {
		return &UserTransaction{}, err
	}
	return p, nil
}

func (p *UserTransaction) DeleteAPost(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&UserTransaction{}).Where("id = ?", pid).Take(&UserTransaction{}).Delete(&UserTransaction{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
