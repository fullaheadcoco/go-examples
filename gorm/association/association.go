package association

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string
	BillingAddress  Address
	ShippingAddress Address
	Emails          []Email
	Languages       []Language
}
type Address struct {
	gorm.Model
	UserID   uint
	Address1 string
	Address2 string
}
type Email struct {
	gorm.Model
	UserID uint
	Email  string
}
type Language struct {
	gorm.Model
	UserID uint
	Name   string
}

func Test() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Email{})
	db.AutoMigrate(&Language{})

	user := User{
		Name:            "jinzhu",
		BillingAddress:  Address{Address1: "Billing Address - Address 1"},
		ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
		Emails: []Email{
			{Email: "jinzhu@example.com"},
			{Email: "jinzhu-2@example.com"},
		},
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	result := db.FirstOrCreate(&user)
	// BEGIN TRANSACTION;
	// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1"), ("Shipping Address - Address 1") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
	// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com"), (111, "jinzhu-2@example.com") ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "languages" ("name") VALUES ('ZH'), ('EN') ON DUPLICATE KEY DO NOTHING;
	// INSERT INTO "user_languages" ("user_id","language_id") VALUES (111, 1), (111, 2) ON DUPLICATE KEY DO NOTHING;
	// COMMIT;

	if result.RowsAffected == 0 {

		//db.Preload("Languages").First(&user)
		marshal, err := json.Marshal(user.Languages)
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", string(marshal))

		err = db.Model(&user).Association("Languages").Replace([]Language{{Name: "asdf"}, {Name: "aaaasdf"}, {Name: "asffff"}})
		if err != nil {
			return
		}

		//db.Preload("Languages").First(&user)
		marshal, err = json.Marshal(user.Languages)
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", string(marshal))

		err = db.Model(&user).Association("Languages").Replace([]Language{{Name: "qwer"}, {Name: "aaaaa"}})
		if err != nil {
			return
		}

		//db.Preload("Languages").First(&user)
		marshal, err = json.Marshal(user.Languages)
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", string(marshal))

	}

}
