package main

import (
	"fmt"
	"bufio"
	"os"
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

func add_product() {
	var product Product_Info

	var X_price int

	scanner := bufio.NewScanner(os.Stdin)

	//商品情報入力
	fmt.Print("商品名：")
	scanner.Scan()
	product.name = scanner.Text()

	//原価値情報入力
	fmt.Print("原価：")
	fmt.Scan(&X_price)
	product.cost_price = X_price

	//売価値情報入力
	fmt.Print("売価：")
	fmt.Scan(&X_price)
	product.selling_price = X_price

	//定価値情報入力
	fmt.Print("定価：")
	fmt.Scan(&X_price)
	product.list_price = X_price

	//在庫数情報入力
	fmt.Print("在庫数：")
	fmt.Scan(&X_price)
	product.stock = X_price

	//商品コード情報入力
	fmt.Print("商品コード：")
	scanner.Scan()
	product.product_code = scanner.Text()

}