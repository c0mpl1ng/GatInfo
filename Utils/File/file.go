package File

import (
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 获取程序文件所在路径
func GetCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取程序执行路径
func GetCurrentAbPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	return dir
}
func ReadFileToArray(filePath string) []string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	lines := strings.Split(string(content), "\n")
	return lines
}

func WriteArrayToFile(filePath string, arr []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, elem := range arr {
		_, err := file.WriteString(elem + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// excel操作，创建excel文件
func CreateExcel() *excelize.File {
	return excelize.NewFile()
}

// 创建一个sheet
func CreateSheet(f *excelize.File, sheet string) int {
	index, _ := f.NewSheet(sheet)
	return index
}

// 删除一个sheet
func DeleteSheet(f *excelize.File, sheet string) bool {
	if err := f.DeleteSheet(sheet); err != nil {
		log.SetPrefix("[-] ")
		log.Println("DeleteSheet :" + err.Error())
		return false
	}
	return true
}

// 写入数据
func WriteExcel(f *excelize.File, sheet string, pos string, content string) bool {
	if err := f.SetCellValue(sheet, pos, content); err != nil {
		log.SetPrefix("[-] ")
		log.Println("WriteExcel :" + err.Error())
		return false
	}
	return true
}

// 保存excel
func SaveExcel(f *excelize.File, filepath string) bool {
	if err := f.SaveAs(filepath); err != nil {
		log.SetPrefix("[-] ")
		log.Println("SaveExcel :" + err.Error())
		return false
	}
	return true
}
