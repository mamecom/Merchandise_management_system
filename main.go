package main

import (
	"fmt"
	"os"
	// "io/ioutil"
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

	err := FileInit()
	if err != nil {
		os.Exit(1)
	}
	result := Menu()
	os.Exit(result)
	
}

// TODO: アッパーキャメルで実装
// TODO: 型判定＋エラーハンドリング
// TODO: CSV形式で読み書きできるよう実装
//　メニュー画面 
func Menu() int {
	var selecter int
	
	println("===========商品管理システム===========")
	records := ReadCsv(FILE_NAME)
	DisplayRecords(records)

	fmt.Printf("\n[追加: 1, 削除: 2, 更新: 3, 終了: 0]: ")
	for {
		fmt.Scan(&selecter)

		switch selecter {
			case 1:
				AddProduct();
			case 2:
			case 3:
			case 0:
				return 0
			default:
				fmt.Printf("[追加: 1, 削除: 2, 更新: 3, 終了: 0]: ")
			}
	}
}

func FileInit() error {
	csvExist := IsCsvExist();
	if csvExist {
		log.Printf("debug: already make a file.")			
	} else {
		log.Printf("debug: must make a file.")

		err := CreateCSV();
		if err != nil {
			return err
		} else {
			fmt.Printf("ファイルを作成しました。")
			// ヘッダー実装
			header := []string{"No","商品名", "原価", "売価", "定価", "在庫数", "商品コード"}
			WriteCsv(header)
		}
	}
	return nil
}

// NOTE: ファイル存在確認関数
func IsCsvExist() bool {
	_, err := os.Stat(FILE_NAME)
	return !os.IsNotExist(err)
}

// NOTE: ファイル作成関数
func CreateCSV() error {
	_, err := os.Create(FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// NOTE: ファイル読み込み関数
func ReadCsv(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

// NOTE: ファイル書き込み関数
func WriteCsv(record []string) {
	// レコード追加
    file, err := os.OpenFile(FILE_NAME, os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Fatal("Error:", err)
    }
    defer file.Close()

    // 追加するレコードはstringのスライスで定義
    writer := csv.NewWriter(file)
    err = writer.Write(record)
    if err != nil {
        log.Fatal("Error:", err)
    }
    writer.Flush()
}

func DisplayRecords( records [][]string ) {
	for _, record := range records {
		fmt.Println(record)
	}
}

func UpdateNo( records [][]string ) {
	for no, record := range records {
		record[0] = strconv.Itoa(no)
	}
}

// !WARNING: 入力時に型と違う値が入ると入力キャンセルされ、飛ばされる
func AddProduct() error {
	var productName, productCode string
	var costPrice, sellingPrice, listPrice, stock int

	csvData := ReadCsv(FILE_NAME)

	//商品情報入力
	fmt.Print("商品名：")
	fmt.Scanf("%s", &productName)

	//原価値情報入力
	fmt.Print("原価：")
	fmt.Scanf("%d", &costPrice)
	// if product.cost_price != int {
	// 	return errors.New("数値ではありませんでした。")
	// }
	//売価値情報入力
	fmt.Print("売価：")
	fmt.Scanf("%d", &sellingPrice)

	//定価値情報入力
	fmt.Print("定価：")
	fmt.Scanf("%d", &listPrice)

	//在庫数情報入力
	fmt.Print("在庫数：")
	fmt.Scanf("%d", &stock)

	//商品コード情報入力
	fmt.Print("商品コード：")
	fmt.Scanf("%s", &productCode)
	
	id := len(csvData)
	add := []string{strconv.Itoa(id), productName, strconv.Itoa(costPrice), strconv.Itoa(sellingPrice), strconv.Itoa(listPrice),strconv.Itoa(stock), productCode}

	WriteCsv(add)

	return nil

}

func DeleteProducts() {

}