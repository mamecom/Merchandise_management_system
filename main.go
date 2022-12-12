package main

import (
	"fmt"

)

type Product_Info struct {
	name string			//商品名
	cost_price int 		//原価
	selling_price int	//売価
	list_price int		//定価
	stock int			//在庫
	product_code string //商品コード
}

func main() {

	menu()
}

//　メニュー画面 
func menu() {
	println("===========商品管理システム===========")
	println(" [追加: 1, 削除: 2, 更新: 3, 終了: 0]\n")
	
	var str int

	for {
		fmt.Scan(&str)

		switch str {
		case 1:
			add_product()
		case 2:
			println("b")
		case 3:
			println("c")
		case 0:
			return
		default:
			return
		}
	}
}

