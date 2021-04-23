package model_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/okh8609/gin_blog/internal/model"
	"github.com/okh8609/gin_blog/pkg/setting"
)

func TestGorm(t *testing.T) {
	sss, err := setting.NewSetting()
	if err != nil {
		log.Fatalf("setting.NewSetting err: %v", err)
	}
	var config setting.DatabaseSetting
	err = sss.ReadSection("Database", &config)
	if err != nil {
		log.Fatalf("sss.ReadSection err: %v", err)
	}

	db, err := model.NewDBEngine(&config)
	if err != nil {
		log.Fatalf("model_test.TestGorm err: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&model.BlogArticle{})

	// Create
	db.Create(&model.BlogArticle{
		Title:   "MyTitle",
		Desc:    "MyDepiction",
		Content: "MyContent...",
	})

	// Read
	var data model.BlogArticle
	// db.First(&data, 1)                      // find data with integer primary key
	db.First(&data, "Title = ?", "MyTitle") // find data with ... // 要小寫?..沒差!
	// 若上兩行都執行的話 會把他們的條件合併起來!!
	fmt.Println(data)

	// Update - update one field
	db.Model(&data).Update("Title", "MyTitle2")
	fmt.Println(data)

	// Update - update multiple fields
	db.Model(&data).Updates(model.BlogArticle{Title: "MyTitle123", Desc: "MyDepiction123"}) // non-zero fields
	db.Model(&data).Updates(map[string]interface{}{"Title": "MyTitle888", "Desc": "MyDepiction888"})
	fmt.Println(data)

	// // Delete - delete product
	// db.Delete(&data, 1)

}
