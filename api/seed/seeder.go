package seed

import (
	"log"
	"mfv_test/api/models"

	"github.com/jinzhu/gorm"
)

//Data sample
var users = []models.Users{
	models.Users{
		Name:  "Phan Van Cong",
		Email: "pvc1@gmail.com",
	},
	models.Users{
		Name:  "Phan Van Cong 2",
		Email: "pvc2gmail.com",
	},
	models.Users{
		Name:  "Phan Van Cong 3",
		Email: "pvc3gmail.com",
	},
	models.Users{
		Name:  "Phan Van Cong 4",
		Email: "pvc4gmail.com",
	},
}

var userAccount = []models.UserAccount{
	models.UserAccount{
		Name: "User Account Phan Van Cong",
		Bank: "VCB",
	},
	models.UserAccount{
		Name: "User Account Phan Van Cong 2",
		Bank: "ACB",
	},
	models.UserAccount{
		Name: "User Account Phan Van Cong 3",
		Bank: "VIB",
	},
	models.UserAccount{
		Name: "User Account Phan Van Cong 4",
		Bank: "VIB",
	},
}

var posts = []models.UserTransaction{
	models.UserTransaction{
		Amount:          100000.00,
		TransactionType: "deposit",
	},
	models.UserTransaction{
		Amount:          100001.00,
		TransactionType: "deposit",
	},
	models.UserTransaction{
		Amount:          100002.00,
		TransactionType: "withdraw",
	},
	models.UserTransaction{
		Amount:          100003.00,
		TransactionType: "withdraw",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.UserTransaction{}, &models.UserAccount{}, &models.Users{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Users{}, &models.UserAccount{}, &models.UserTransaction{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.Users{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		userAccount[i].UserID = users[i].ID
		err = db.Debug().Model(&models.UserAccount{}).Create(&userAccount[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		posts[i].UserAccountID = users[i].ID

		err = db.Debug().Model(&models.UserTransaction{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}

}
