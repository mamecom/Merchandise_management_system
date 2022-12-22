package main

import (
	"fmt"
	"os"
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

	if err := FileInit(); err != nil {
		os.Exit(1)
	}
	result := Menu()
	os.Exit(result)
	
}

//　メニュー画面 
func Menu() int {
	var selecter int
	
	println("===========商品管理システム===========")
	records := ReadCsv(FILE_NAME)
	DisplayRecords(records)

	for {
		fmt.Printf("\n[追加: 1, 削除: 2, 更新: 3, 終了: 0]: ")
		fmt.Scan(&selecter)

		switch selecter {
			case 1:
				AddProduct()
			case 2:
				DeleteProducts()
			case 3:
				UpdateProductsInfo()
			case 0:
				return 0
			default:			
			}
	}
}

func FileInit() error {
	csvExist := IsCsvExist();
	if !csvExist {
		log.Println("debug: must make a file.")

		if err := CreateCSV(); err != nil {
			return err
		} else {
			fmt.Println("ファイルを作成しました。")
			MakeHeader()
		}
	}
	return nil
}

func MakeHeader()  {
	header := []string{"No","商品名", "原価", "売価", "定価", "在庫数", "商品コード"}
	WriteCsv(header)
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
func WriteCsv(record []string) error {
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
		return err
    }
    writer.Flush()

	return nil
}

func DisplayRecords( records [][]string ) {
	for _, record := range records {
		fmt.Println(record)
	}
}

// !FIXME: 入力時に型と違う値が入ると入力キャンセルされ、飛ばされる
func AddProduct() error {
	csvData := ReadCsv(FILE_NAME)
	productName, costPrice, sellingPrice, listPrice, stock, productCode := InputProductInfo()
	
	lastArr := csvData[len(csvData) - 1]
	id := lastArr[0]
	intId, _ := strconv.Atoi(id)
	intId = intId + 1
	add := []string{strconv.Itoa(intId), productName, strconv.Itoa(costPrice), strconv.Itoa(sellingPrice), strconv.Itoa(listPrice),strconv.Itoa(stock), productCode}

	if err := WriteCsv(add); err != nil {
		return err
	}

	return nil

}

// OPTIMIZE: HACK: 戻り値をまとめて返すようにする
func InputProductInfo() (string, int, int, int, int, string) {
	var productName, productCode string
	var costPrice, sellingPrice, listPrice, stock int

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

	return productName, costPrice, sellingPrice, listPrice, stock, productCode
	
}

func DeleteProducts() error {
	var delNo int
	cansel := 0

	fmt.Printf("削除したいNoは：")
	fmt.Scanf("%d", &delNo)

	if delNo > cansel {
		if err := Remove(delNo); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func Remove(delNo int) error {
	records := ReadCsv(FILE_NAME)
	os.Remove(FILE_NAME)
	err := CreateCSV()
	if err != nil {
		log.Fatal(err)
		return err
	}
	for cnt, record := range records {
		if cnt != delNo {
			if err := WriteCsv(record); err != nil {
				log.Fatal(err)
				return err
			}
		}
	}

	return nil
}

// TODO: 指定した要素のみを更新できるように修正
func UpdateProductsInfo() error {
	var updateNo int

	fmt.Printf("更新したいNoは：")
	fmt.Scanf("%d", &updateNo)

	if err := UpdateProducts(updateNo); err != nil {
		return err
	}

	return nil
}

// func UpdateNo( records [][]string ) {
// 	for no, record := range records {
// 		record[0] = strconv.Itoa(no)
// 	}
// }

func UpdateProducts(updateNo int ) error {

	records := ReadCsv(FILE_NAME)
	os.Remove(FILE_NAME)
	if err := CreateCSV(); err != nil {
		log.Fatal(err)
		return err
	}
	for cnt, record := range records {
		if cnt == updateNo {
			productName, costPrice, sellingPrice, listPrice, stock, productCode := InputProductInfo()
			updateInfo := []string{strconv.Itoa(cnt), productName, strconv.Itoa(costPrice), strconv.Itoa(sellingPrice), strconv.Itoa(listPrice),strconv.Itoa(stock), productCode}
			if err := WriteCsv(updateInfo); err != nil {
				log.Fatal(err)
				return err
			}
		} else {
			if err := WriteCsv(record); err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}