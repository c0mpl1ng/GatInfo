package Common

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Params struct {
	UrlFilePath    string
	Url            string
	CertFilePath   string
	Cert           string
	ResultFilePath string
	City           string
	Region         string
}

var rootCmd = &cobra.Command{
	Use:   "GatInfo",
	Short: "GatInfo 是一个护网资产收集工具",
	Long:  " GatInfo 是一个护网资产收集工具",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			fmt.Println("Arguments:", args)
		}
	},
}

func init() {

	rootCmd.CompletionOptions.DisableDefaultCmd = true

}

var param Params

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
