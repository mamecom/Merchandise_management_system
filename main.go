package main

import (
	"fmt"

)

func main() {

	menu()
}

func menu() {
	println("===========商品管理システム===========")
	println(" [追加: 1, 削除: 2, 更新: 3, 終了: 0]\n")
	
	var str int

	for {
		fmt.Scan(&str)

		switch str {
		case 1:
			println("a")
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
