package main

import (
	"fmt"
	_ "bufio"
	"os"
	_ "io/ioutil"
	"log"
	_ "errors"
	// "jszwec/csvutil"
)

const (
	ADD_PRODUCT = iota
	FILE_NAME = "productfile.csv"
	RESULT_TRUE = 1
	RESULT_FALSE = -1
)

// todo: キャメルケースに修正すること
type productInfo struct {
	no int				//データNo.
	name string			//商品名
	cost_price int 		//原価
	selling_price int	//売価
	list_price int		//定価
	stock int			//在庫
	product_code string //商品コード
}

func main() {

	result := menu()
	os.Exit(result)
	
}

//　メニュー画面 
func menu() int {
	
	println("===========商品管理システム===========")
	
	// todo: ファイルを作る処理
	// ファイルの存在確認
	err := Exists();
	if err != nil {
		log.Printf("debug: must make a file.")
		createCSVFile();
	}

	fmt.Println(" [追加: 1, 削除: 2, 更新: 3, 終了: 0]\n\n\n")
	
	return 0
}

// ファイル存在確認関数
func Exists() (error) {
	_, err := os.Stat(FILE_NAME)
	if os.IsNotExist(err) {
		log.Printf("debug: not exists file.")
		return err
	}
	return nil
}

// ファイル作成関数
func createCSVFile() error {
	_, err := os.Create(FILE_NAME)
	if err != nil {
		return err
	}
	return nil
}
