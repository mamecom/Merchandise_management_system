package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	ADD_PRODUCT = iota
	FILE_NAME   = "productfile.csv"
	FILE_SEARCH = "search.csv"
)

func main() {

	if err := FileInit(FILE_NAME); err != nil {
		os.Exit(1)
	}
	result := Menu()
	os.Exit(result)

}

// 　メニュー画面
func Menu() int {
	var selecter int
	isDisplay := true

	DisplayRecords(FILE_NAME)

	for {
		fmt.Printf("\n[追加: 1, 削除: 2, 更新: 3, ソート: 4, 検索: 5, 終了: 0]: ")
		fmt.Scan(&selecter)

		switch selecter {
		case 1:
			AddProduct()
		case 2:
			DeleteProducts()
		case 3:
			UpdateProductsInfo()
		case 4:
			SortRecords()
		case 5:
			SearchRecord()
			isDisplay = false
		case 0:
			return 0
		default:
		}
		if isDisplay {
			DisplayRecords(FILE_NAME)
		} else {
			isDisplay = true
		}
	}
}

func FileInit(fileName string) error {
	csvExist := IsCsvExist()
	if !csvExist {
		log.Println("debug: must make a file.")

		if err := CreateCSV(fileName); err != nil {
			return err
		} else {
			fmt.Println("ファイルを作成しました。")
			MakeHeader(fileName)
		}
	}
	return nil
}

func MakeHeader(fileName string) {
	header := []string{"No", "商品名", "原価", "売価", "定価", "在庫数", "商品コード"}
	WriteCsv(header, fileName)
}

// NOTE: ファイル存在確認関数
func IsCsvExist() bool {
	_, err := os.Stat(FILE_NAME)
	return !os.IsNotExist(err)
}

// NOTE: ファイル作成関数
func CreateCSV(fileName string) error {
	_, err := os.Create(fileName)
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
func WriteCsv(record []string, filename string) error {
	// レコード追加
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer file.Close()

	// 追加するレコードはstringのスライスで定義
	writer := csv.NewWriter(file)
	if err = writer.Write(record); err != nil {
		log.Fatal("Error:", err)
		return err
	}
	writer.Flush()

	return nil
}

func WriteCsvs(records [][]string, fileName string) {
	// レコード追加
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer file.Close()

	// 追加するレコードはstringのスライスで定義
	w := csv.NewWriter(file)
	if err := w.WriteAll(records); err != nil {
		log.Fatal("Error:", err)
	}
}

func DisplayRecords(fileName string) {
	var records [][]string

	if fileName != FILE_NAME {
		println("----------------------検索結果------------------------")
		records = ReadCsv(fileName)
		os.Remove(fileName)
	} else {
		println("----------------------商品管理システム------------------------")
		records = ReadCsv(fileName)
	}

	w := csv.NewWriter(os.Stdout)
	w.Comma = '\t'
	if err := w.WriteAll(records); err != nil {
		log.Fatal("Error:", err)
	}
	w.Flush()

	fmt.Println("--------------------------------------------------------------")
}

// !FIXME: 入力時に型と違う値が入ると入力キャンセルされ、飛ばされる
func AddProduct() error {
	csvData := ReadCsv(FILE_NAME)
	productName, costPrice, sellingPrice, listPrice, stock, productCode := InputProductInfo()

	lastArr := csvData[len(csvData)-1]

	id := lastArr[0]
	intId, _ := strconv.Atoi(id)
	intId = intId + 1
	add := []string{strconv.Itoa(intId), productName, strconv.Itoa(costPrice), strconv.Itoa(sellingPrice), strconv.Itoa(listPrice), strconv.Itoa(stock), productCode}

	if err := WriteCsv(add, FILE_NAME); err != nil {
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
		strDelNo := strconv.Itoa(delNo)
		// strDelNo[0][0] = delNo

		Remove(strDelNo)
	}

	return nil
}

func Remove(delNo string) {
	records := ReadCsv(FILE_NAME)
	isDeltered := false

	for index, record := range records {
		if record[0] == delNo { //MEMO: Noを同じ値かを判定
			os.Remove(FILE_NAME)
			if err := CreateCSV(FILE_NAME); err != nil {
				log.Fatal(err)
			}
			removedRecords := append(records[:index], records[index+1:]...)
			WriteCsvs(removedRecords, FILE_NAME)
			isDeltered = true
			break
		}
	}
	if !isDeltered {
		fmt.Println("データが存在しませんでした。")
	}
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

func UpdateProducts(updateNo int) error {

	records := ReadCsv(FILE_NAME)
	os.Remove(FILE_NAME)
	if err := CreateCSV(FILE_NAME); err != nil {
		log.Fatal(err)
		return err
	}
	for cnt, record := range records {
		if cnt == updateNo {
			productName, costPrice, sellingPrice, listPrice, stock, productCode := InputProductInfo()
			updateInfo := []string{strconv.Itoa(cnt), productName,
				strconv.Itoa(costPrice), strconv.Itoa(sellingPrice),
				strconv.Itoa(listPrice), strconv.Itoa(stock), productCode,
			}
			if err := WriteCsv(updateInfo, FILE_NAME); err != nil {
				log.Fatal(err)
				return err
			}
		} else {
			if err := WriteCsv(record, FILE_NAME); err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}

func SortRecords() {
	var sortType string
	var sorted [][]string

	for {
		fmt.Printf("[ソートパターンを選択してください。昇順: x, 降順: y, キャンセル: c]: ")
		fmt.Scan(&sortType)
		switch sortType {
		case "y":
			sorted = BubbleSort(sortType)
		case "x":
			sorted = BubbleSort(sortType)
		case "c":
			return
		default:
			return
		}
		WriteCsvs(sorted, FILE_NAME)
		DisplayRecords(FILE_NAME)
	}
}

func BubbleSort(sortType string) [][]string {
	topToBottom := "y"
	bottomToTop := "x"

	records := ReadCsv(FILE_NAME)

	for i := 1; i < len(records); i++ {
		for j := 1; j < len(records)-i; j++ {
			if sortType == topToBottom {
				if records[j][0] < records[j+1][0] {
					records[j], records[j+1] = records[j+1], records[j]
				}
			} else if sortType == bottomToTop {
				if records[j][0] > records[j+1][0] {
					records[j], records[j+1] = records[j+1], records[j]
				}
			}
		}
	}
	return records
}

func SearchRecord() {
	var searchStr string

	records := ReadCsv(FILE_NAME)

	fmt.Print("検索する文字を入力してください: ")
	fmt.Scan(&searchStr)
	searchRecords := SearchExec(records, searchStr)
	CreateAndWriteSearchCSV(searchRecords, FILE_SEARCH)
	DisplayRecords(FILE_SEARCH)
}

func CreateAndWriteSearchCSV(searchRecords [][]string, fileName string) {
	CreateCSV(fileName)
	WriteCsvs(searchRecords, fileName)
}

func SearchExec(records [][]string, str string) [][]string {
	searchRecords := [][]string{}

	// ヘッダー格納
	header := []string{"No", "商品名", "原価", "売価", "定価", "在庫数", "商品コード"}
	searchRecords = append(searchRecords, header)

	for _, record := range records {
		for _, data := range record {
			if data == str {
				searchRecords = append(searchRecords, record)
				break
			}
		}
	}
	return searchRecords
}
