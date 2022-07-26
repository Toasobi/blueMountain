package model

import (
	"log"

	"gorm.io/gorm"
)

//定义一个服装类
type Cloth struct {
	Id      string
	Size    string
	Price   int
	Variety string
}

//定义一个仓库类
type Store struct {
	Id       string
	Capacity int
}

//定义一个供应情况类
type Availability struct {
	Id_of_clothes   string
	Id_of_providers string
	Level           string
}

//定义一个供应商表
type Supplier struct {
	Id   string
	Name string
}

//1.查询查询服装尺码为'S'且销售价格在100以下的服装信息
func GetSAndUnder100(db *gorm.DB) (clothes []*Cloth) {
	if err := db.Where("size = ? AND price <= ?", "S", 100).Find(&clothes).Error; err != nil {
		log.Fatal(err)
		return
	}
	return
}

//2.查询仓库容量最大的仓库信息
func GetMaxCapacity(db *gorm.DB) (store *Store) {
	var stores []*Store
	var temp int
	var index int
	db.Find(&stores)
	temp = stores[0].Capacity
	for i := 1; i < len(stores); i++ {
		if stores[i].Capacity > temp {
			temp = stores[i].Capacity
			index = i
		}
	}
	return stores[index]
}

//3.查询A类服装的库存总量
func GetCapacity(db *gorm.DB) (store *Store) {
	if err := db.Where("Id = ?", "A").Find(&store).Error; err != nil {
		log.Fatal(err)
		return
	}
	return
}

//4.查询服装编码以‘A’开始开头的服装
func GetStartA(db *gorm.DB) (clothes []*Cloth) {
	if err := db.Where("Id LIKE ?", "%A%").Find(&clothes).Error; err != nil {
		log.Fatal(err)
		return
	}
	return
}

//5.查询服装质量等级有不合格的供应商信息
func GetDisqualification(db *gorm.DB) (suppliers []*Supplier) {
	var availabilities []*Availability = make([]*Availability, 200)
	var str []string = make([]string, 100)
	if err := db.Where("Level = ?", "不合格").Find(&availabilities).Error; err != nil {
		log.Fatal(err)
		return
	}
	for i := 0; i < len(availabilities); i++ {
		str[i] = availabilities[i].Id_of_providers
	}
	if err := db.Where("Id IN ?", str).Find(&suppliers).Error; err != nil {
		log.Fatal()
		return
	}
	return
}

//6.把服装尺寸为'S'的服装的销售价格均在原来基础上提高10%
func Expensiver(db *gorm.DB) {
	var clothes []*Cloth = make([]*Cloth, 100)
	db.Where("Size = ?", "S").Find(&clothes)
	for i := 0; i < len(clothes); i++ {
		clothes[i].Price += int(0.1 * float64(clothes[i].Price))
	}
	db.Save(&clothes)

}

//7.删除所有服装质量等级不合格的供应情况
func Delete(db *gorm.DB) {
	if err := db.Delete(&Availability{}, "Level = ?", "不合格").Error; err != nil {
		log.Fatal(err)
		return
	}

}

//8.插入（这里挑服装表进行插入）
func Insert(db *gorm.DB) {
	cloth := Cloth{Id: "B1", Size: "M", Price: 100, Variety: "B"}
	if err := db.Create(&cloth).Error; err != nil {
		log.Fatal(err)
		return
	}
}
