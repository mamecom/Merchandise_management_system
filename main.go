package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"encoding/csv"
	"strconv"
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

	result := Menu()
	os.Exit(result)
	
}

// TODO: アッパーで
// TODO: 型判定＋エラーハンドリング
// TODO: CSV形式で読み書きできるよう実装
//　メニュー画面 
func Menu() int {
	
	println("===========商品管理システム===========")
	
	err := Exists();
	if err != nil {
		log.Printf("debug: must make a file.")
		err := CreateCSVFile();
		if err != nil {
			fmt.Printf("ファイルを作成しました。")
		} else {
			return 0
		}
	}
	AddProduct();
	// records := [][]string{
	// 	[]string{"名前", "年齢", "出身地", "性別"},
	// 	[]string{"山本", "24", "兵庫", "男"},
	// 	[]string{"鈴木", "29", "神奈川", "女"},
	// 	[]string{"佐藤", "27", "鹿児島", "男"},
	// }
	// err = CsvFileWrite(records);
	if err != nil {
		log.Printf("alert: can`t write file")
	}

	err = CsvFileRead();
	if err != nil {
		log.Printf("alert: can`t read file")
	}

	fmt.Printf("\n[追加: 1, 削除: 2, 更新: 3, 終了: 0]\n")
	
	return 0
}



// NOTE: ファイル存在確認関数
func Exists() error {
	_, err := os.Stat(FILE_NAME)
	if os.IsNotExist(err) {
		log.Fatal(err)
		return err
	}
	return nil
}

// NOTE: ファイル作成関数
func CreateCSVFile() error {
	_, err := os.Create(FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// NOTE: ファイル読み込み関数
func CsvFileRead() error {
	data, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(string(data))
	return nil
}

// NOTE: ファイル書き込み関数
func CsvFileWrite(record [][]string) error {
	
	// MEMO: Close不要
	f, err := os.Create(FILE_NAME)		//ファイルの有無に関わらず書き込み権限で開くもの
	if err != nil {
		log.Fatal(err)
		return err
	}
	
	w := csv.NewWriter(f)
	w.WriteAll(record)

	return nil
}

func AddProduct() error {
	var product_name string
	var cost_price int
	var selling_price int
	var list_price int
	var stock int
	var product_code string


	records := [][]string{
		[]string{"ID","商品名", "原価", "売価", "定価", "在庫数", "商品コード"},
	}
	//商品情報入力
	fmt.Print("商品名：")
	fmt.Scanf("%s", &product_name)

	//原価値情報入力
	fmt.Print("原価：")
	fmt.Scanf("%d", &cost_price)
	// if product.cost_price != int {
	// 	return errors.New("数値ではありませんでした。")
	// }
	//売価値情報入力
	fmt.Print("売価：")
	fmt.Scanf("%d", &selling_price)

	//定価値情報入力
	fmt.Print("定価：")
	fmt.Scanf("%d", &list_price)

	//在庫数情報入力
	fmt.Print("在庫数：")
	fmt.Scanf("%d", &stock)

	//商品コード情報入力
	fmt.Print("商品コード：")
	fmt.Scanf("%s", &product_code)
	
	id := len(records)
	add := []string{strconv.Itoa(id), product_name, strconv.Itoa(cost_price), strconv.Itoa(selling_price), strconv.Itoa(list_price), strconv.Itoa(stock), product_code}
	records = append(records, add)

	CsvFileWrite(records)

	return nil

}