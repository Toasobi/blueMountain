package main

import (
	"blueMountain/work3/mysql/controller"
	"blueMountain/work3/mysql/dao"
	"blueMountain/work3/mysql/model"
)

func main() {
	err := dao.InitDao()
	if err != nil {
		panic(err)
	}

	//程序退出关闭数据库
	sqlDB, _ := dao.DB.DB()
	defer sqlDB.Close()

	//1.查询查询服装尺码为'S'且销售价格在100以下的服装信息
	controller.CGetSAndUnder100()

	//2.查询仓库容量最大的仓库信息
	controller.CGetMaxCapacity()

	//3.查询A类服装的库存总量
	controller.CGetCapacity()

	//4.查询服装编码以‘A’开始开头的服装
	controller.CGetStartA()

	//5.查询服装质量等级有不合格的供应商信息
	controller.CGetDisqualification()

	//6.把服装尺寸为'S'的服装的销售价格均在原来基础上提高10%
	model.Expensiver(dao.DB)

	//7.删除所有服装质量等级不合格的供应情况
	model.Delete(dao.DB)

	//8.插入（这里挑服装表进行插入）
	model.Insert(dao.DB)
}
