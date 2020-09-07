package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Users struct {
	tableName struct{}   `sql:"users,alias:users" pg:",discard_unknown_columns"`
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"size:255;not null;unique" json:"name"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"default:NULL" json:"deleted_at,omitempty"` // When deleted_at null -> not show display
}

func (u *Users) PrepareUpdate() {
	u.UpdatedAt = time.Now()
}

func (u *Users) UpdateAUser(db *gorm.DB, uid uint32) (*Users, error) {

	db = db.Debug().Model(&Users{}).Where("id = ?", uid).Take(&Users{}).UpdateColumns(
		map[string]interface{}{
			"name":       u.Name,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Users{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&Users{}).Where("id = ?", uid).Find(&u).Error
	if err != nil {
		return &Users{}, err
	}

	fmt.Println(u.DeletedAt)
	return u, nil
}

func (u *Users) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Users{}).Where("id = ?", uid).Take(&Users{}).Delete(&Users{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
