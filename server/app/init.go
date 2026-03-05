package app

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/rsj-rishabh/urbanClapClone/server/config"
	"github.com/rsj-rishabh/urbanClapClone/server/app/model"
)
func (a *App) InitializeDB() {

    cfg := config.GetConfig()

    dsn := fmt.Sprintf(
        "%s:%s@tcp(mysql:3306)/%s?charset=%s&parseTime=True&loc=Local",
        cfg.DB.Username,
        cfg.DB.Password,
        cfg.DB.Name,
        cfg.DB.Charset,
    )

    var err error

    for i := 0; i < 30; i++ {

        a.DB, err = gorm.Open(cfg.DB.Dialect, dsn)

        if err == nil {

            // Ping DB to ensure it is ready
            err = a.DB.DB().Ping()

            if err == nil {
                fmt.Println("Connected to MySQL")
                return
            }
        }

        fmt.Println("Waiting for MySQL...")
        time.Sleep(5 * time.Second)
    }

    panic("Database connection failed")
}

func (a *App) DBMigrate() {
	// Drop the table if it exists
	a.DB.AutoMigrate().DropTable(&model.User{})
	a.DB.AutoMigrate().DropTable(&model.Service{})
	a.DB.AutoMigrate().DropTable(&model.Booking{})
	a.DB.AutoMigrate().DropTable(&model.CityServiceMapping{})

	// Migrate the schema
	a.DB.AutoMigrate(&model.User{}, &model.Service{}, &model.Booking{}, &model.CityServiceMapping{})

	// Create users table


	// Create services table
	a.DB.Create(&model.Service{
		Id:          1,
		Name:        "AC Maintanence",
		Description: "Any type of AC maintanence such as filter cleaning, part replacement, etc.",
		Category:    "Electronics",
		ImageName:   "air_conditioning.jpg",
		Price:       80,
	})
	a.DB.Create(&model.Service{
		Id:          2,
		Name:        "Plumbing",
		Description: "Sanitary and household plumbing. No sewage service.",
		Category:    "Household",
		ImageName:   "plumbing.jpg",
		Price:       100,
	})
	a.DB.Create(&model.Service{
		Id:          3,
		Name:        "Saloon",
		Description: "Haricut, massage, nailwork, makeup, etc.",
		Category:    "Personal Care",
		ImageName:   "saloon.jpg",
		Price:       25,
	})
	a.DB.Create(&model.Service{
		Id:          4,
		Name:        "Furniture Repair",
		Description: "Furniture frame repair, drilling, fitting new furniture, etc.",
		Category:    "Household",
		ImageName:   "furniture_repair.jpg",
		Price:       70,
	})
	a.DB.Create(&model.Service{
		Id:          5,
		Name:        "Exterminator",
		Description: "Pest control, wildlife evac, alligator emergency, etc.",
		Category:    "Animal/Pet",
		ImageName:   "pest_control.jpg",
		Price:       150,
	})

	// Create booking table
	a.DB.Create((&model.Booking{
		UserId:      1,
		ServiceId:   1,
		Date:        "2022-02-15",
		StartTime:   "12:30",
		EndTime:     "13:30",
		IsCancelled: false,
	}))
	a.DB.Create((&model.Booking{
		UserId:      1,
		ServiceId:   2,
		Date:        "2022-02-15",
		StartTime:   "16:30",
		EndTime:     "17:30",
		IsCancelled: false,
	}))
	a.DB.Create((&model.Booking{
		UserId:      2,
		ServiceId:   3,
		Date:        "2022-02-15",
		StartTime:   "16:30",
		EndTime:     "17:30",
		IsCancelled: false,
	}))

	// Create CityServiceMapping table
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Newyork",
		ServiceId: 3,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Newyork",
		ServiceId: 2,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "LA",
		ServiceId: 2,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "LA",
		ServiceId: 3,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Boston",
		ServiceId: 1,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Boston",
		ServiceId: 2,
	}))
}
