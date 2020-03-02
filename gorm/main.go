package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm/models"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main01() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&Product{})
	// insert
	db.Create(&Product{
		Code:  "L1212",
		Price: 1000,
	})

	// 读取
	var product Product
	// 判断数据库是否存在
	productIsExiste := db.HasTable(&product)
	if !productIsExiste {
		db.CreateTable(&product)
	}
	db.First(&product, 1)                 // 查询id为1的product
	db.First(&product, "code=?", "L1212") //查询code为l1212的product
	// 更新
	db.Model(&product).Update("Price", 2000)
	//删除
	db.Delete(&product)
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var animal = models.Animal{
		Name: "anziguoer",
		Age:  56,
	}
	//db.DropTableIfExists(&animal)
	//db.AutoMigrate(&animal)
	//db.Create(&animal)
	// 查找
	//var animals []models.Animal
	//db.Find(&animals)
	//var jsonResult []byte
	//
	//jsonResult, err = json.MarshalIndent(animals, "", "  ")
	//fmt.Println(string(jsonResult))

	// 查找其他的用法
	//db.Find(&animal, "name = 'demo'")
	//db.Where("name = 'demo'").Find(&animal).Update("name", "hello")
	//db.Where(5).Find(&animal).Update(
	//	map[string]interface{}{
	//		"name": "animal",
	//		"age":  10,
	//	})
	//db.Where("id = ?", 6).Find(&animal).Update(models.Animal{
	//	Name: "lnart",
	//	Age:  2,
	//})

	// in where
	db.Where("name in (?)", []string{"anziguoer", "demo", "lnart"}).Find(&animal)

	fmt.Println(animal)
}
