package seed

import (
	"log"

	"blog_api/api/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	// err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }
	err := db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	/*
		for i := range users {
			err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
			if err != nil {
				log.Fatalf("cannot seed users table: %v", err)
			}
			posts[i].AuthorID = users[i].ID

			err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
			if err != nil {
				log.Fatalf("cannot seed posts table: %v", err)
			}
		}
	*/
}
