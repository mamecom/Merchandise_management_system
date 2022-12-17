package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"io"
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
		err = CreateCSVFile();
		log.Println(err)
		if err != nil {
			return 0
		} else {
			fmt.Printf("ファイルを作成しました。")
			// ヘッダー実装
			header := [][]string{
				[]string{"ID","商品名", "原価", "売価", "定価", "在庫数", "商品コード"},
			}
			CsvFileWrite(header)
		}	
	} else {
		log.Printf("debug: already make a file.")
	}
	
	AddProduct();
	if err != nil {
		log.Printf("alert: can`t write file")
	}

	// err = CsvFileRead();
	// if err != nil {
	// 	log.Printf("alert: can`t read file")
	// }

	fmt.Printf("\n[追加: 1, 削除: 2, 更新: 3, 終了: 0]\n")

	
	
	return 0
}



// NOTE: ファイル存在確認関数
func Exists() error {
	_, err := os.Stat(FILE_NAME)
	if !os.IsNotExist(err) {
		// log.Fatal(err)
		return nil
	}
	return err
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
func CsvFileRead(list *[][]string) error {
	// MEMO: Close不要
	f, err := os.Open(FILE_NAME)		//ファイルの有無に関わらず書き込み権限で開くもの
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	data := csv.NewReader(f)

	for {
		row, err := data.Read() 		// csvを1行ずつ読み込む
		  if err == io.EOF {
			break 						// 読み込むべきデータが残っていない場合，Readはnil, io.EOFを返すため、ここでループを抜ける
		  }
		fmt.Println(row)
	}

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
	// var product_name string
	// var cost_price int
	// var selling_price int
	// var list_price int
	// var stock int
	// var product_code string

	records := [][]string{
		
	}

	CsvFileRead(&records)

	data, err := ioutil.ReadFile(FILE_NAME)
    if err != nil {
        fmt.Println(err)
    }

	records = data

    fmt.Println(string(data))
	

	// CsvFileRead(records);

	// log.Println(records)

	// //商品情報入力
	// fmt.Print("商品名：")
	// fmt.Scanf("%s", &product_name)

	// //原価値情報入力
	// fmt.Print("原価：")
	// fmt.Scanf("%d", &cost_price)
	// // if product.cost_price != int {
	// // 	return errors.New("数値ではありませんでした。")
	// // }
	// //売価値情報入力
	// fmt.Print("売価：")
	// fmt.Scanf("%d", &selling_price)

	// //定価値情報入力
	// fmt.Print("定価：")
	// fmt.Scanf("%d", &list_price)

	// //在庫数情報入力
	// fmt.Print("在庫数：")
	// fmt.Scanf("%d", &stock)

	// //商品コード情報入力
	// fmt.Print("商品コード：")
	// fmt.Scanf("%s", &product_code)
	
	id := len(rows)
	// add := []string{strconv.Itoa(id), product_name, strconv.Itoa(cost_price), strconv.Itoa(selling_price), strconv.Itoa(list_price), strconv.Itoa(stock), product_code}
	add := []string{strconv.Itoa(id), "abc", "100", "100", "100", "100", "100"}
	fmt.Println(add)

	// records = append(records, add)

	// CsvFileWrite(records)

	return nil

}