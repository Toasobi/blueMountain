package controller

import (
	"blueMountain/work3/model"
	"fmt"

	"blueMountain/work3/dao"
)

//1.查询查询服装尺码为'S'且销售价格在100以下的服装信息

func CGetSAndUnder100() {
	clothes := model.GetSAndUnder100(dao.DB)
	for i := 0; i < len(clothes); i++ {
		fmt.Printf("%v\n", clothes[i])
	}
}

//2.查询仓库容量最大的仓库信息
func CGetMaxCapacity() {
	storeM := model.GetMaxCapacity(dao.DB)
	fmt.Printf("%v\n", storeM)
}

//3.查询A类服装的库存总量
func CGetCapacity() {
	store := model.GetCapacity(dao.DB)
	fmt.Printf("%v\n", store.Capacity)
}

//4.查询服装编码以‘A’开始开头的服装
func CGetStartA() {
	Aclothes := model.GetStartA(dao.DB)
	for i := 0; i < len(Aclothes); i++ {
		fmt.Printf("%v\n", Aclothes[i])
	}
}

func CGetDisqualification() {
	suppliers := model.GetDisqualification(dao.DB)
	for i := 0; i < len(suppliers); i++ {
		fmt.Printf("%v\n", suppliers[i])
	}
}
