package Common

import (
	"GatInfo/Modules/Subdomain"
	"GatInfo/Utils/File"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func removeDuplicates(strs []string) []string {
	unique := make(map[string]struct{})
	result := []string{}
	for _, str := range strs {
		if _, ok := unique[str]; !ok {
			result = append(result, str)
			unique[str] = struct{}{}
		}
	}
	return result
}

var GetSubdomain = &cobra.Command{
	Use:   "GetSubdomain",
	Short: "通过根域获取证书，通过证书获取资产",
	Run: func(cmd *cobra.Command, args []string) {
		var targets []string
		var FullInfo Subdomain.FullUrlInfo
		if len(args) == 0 {
			cmd.Help()
		}
		if param.Url != "" {
			targets = append(targets, param.Url)
		}
		if param.UrlFilePath != "" {
			targets = append(targets, File.ReadFileToArray(param.UrlFilePath)...)
		}
		if param.Cert != "" {
			targets = append(targets, param.Cert)
		}
		if param.CertFilePath != "" {
			targets = append(targets, File.ReadFileToArray(param.CertFilePath)...)
		}
		if targets == nil {
			log.SetPrefix("[-] ")
			log.Fatalf("没有目标")
		}
		var AllUrlList []string
		if param.Url != "" || param.UrlFilePath != "" {
			FullInfo = Subdomain.GetSubdomain(targets)
		}
		if param.Cert != "" || param.CertFilePath != "" {
			FullInfo = Subdomain.GetSubdomainBySpecifyCert(targets)
		}

		CertInfo := FullInfo.CertInfo
		FofaInfo := FullInfo.FofaInfo
		excelFile := File.CreateExcel()
		File.DeleteSheet(excelFile, "Sheet1")
		//写入domain->证书信息
		File.CreateSheet(excelFile, "Cert")
		File.WriteExcel(excelFile, "Cert", "A1", "domain")
		File.WriteExcel(excelFile, "Cert", "B1", "CertName")

		count := 0
		for _, certMap := range CertInfo {
			for key, value := range certMap.(map[string]interface{}) {
				value := value.([]string)
				for _, target := range value {
					File.WriteExcel(excelFile, "Cert", fmt.Sprintf("A%d", count+2), key)
					File.WriteExcel(excelFile, "Cert", fmt.Sprintf("B%d", count+2), target)
					//log.SetPrefix("debug: ")
					//log.Println(fmt.Sprintf("%d,%s,%s", count, key, target))
					count += 1
					AllUrlList = append(AllUrlList, target)
				}

			}

		}

		//写入证书->urls信息
		File.CreateSheet(excelFile, "Fofa")
		File.WriteExcel(excelFile, "Fofa", "A1", "CertName")
		File.WriteExcel(excelFile, "Fofa", "B1", "Url")

		count = 0
		for _, urlMap := range FofaInfo {
			for key, value := range urlMap.(map[string]interface{}) {
				value := value.([]string)
				for _, target := range value {
					File.WriteExcel(excelFile, "Fofa", fmt.Sprintf("A%d", count+2), key)
					File.WriteExcel(excelFile, "Fofa", fmt.Sprintf("B%d", count+2), target)
					//log.SetPrefix("debug: ")
					//log.Println(fmt.Sprintf("%d,%s,%s", count, key, target))
					count += 1
					AllUrlList = append(AllUrlList, target)
				}

			}

		}

		//写入所有url到一个表
		AllUrlList = removeDuplicates(AllUrlList)
		File.CreateSheet(excelFile, "AllUrls")
		File.WriteExcel(excelFile, "AllUrls", "A1", "urls")

		for c, url := range AllUrlList {
			//log.SetPrefix("debug: ")
			//log.Println(fmt.Sprintf("%d,%s", c, url))
			File.WriteExcel(excelFile, "AllUrls", fmt.Sprintf("A%d", c+2), url)
		}

		//保存excel

		if param.ResultFilePath == "" {
			dir := File.GetCurrentAbPathByExecutable()
			File.SaveExcel(excelFile, dir+"/FullInfo.xlsx")
			log.SetPrefix("[+] ")
			log.Println("SubdomainList 已写入:" + dir + "/FullInfo.xlsx")
		} else {
			File.SaveExcel(excelFile, param.ResultFilePath)
			log.SetPrefix("[+] ")
			log.Println("SubdomainList 已写入:" + param.ResultFilePath)
		}
	},
}

func init() {

	rootCmd.AddCommand(GetSubdomain)
	GetSubdomain.Flags().StringVarP(&param.UrlFilePath, "DomainFile", "", "", "从文本中读取目domain,一行一个")
	GetSubdomain.Flags().StringVarP(&param.Url, "Domain", "", "", "单个domain目标")
	GetSubdomain.Flags().StringVarP(&param.CertFilePath, "CertFile", "", "", "从文本中读取目cert,一行一个")
	GetSubdomain.Flags().StringVarP(&param.Cert, "Cert", "", "", "单个Cert目标")
	GetSubdomain.Flags().StringVarP(&param.ResultFilePath, "Output", "o", "", "输出文件路径")
}
