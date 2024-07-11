package Common

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetGys = &cobra.Command{
	Use:   "GetGys",
	Short: "获取供应商信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run GatGys...")
	},
}

func init() {
	rootCmd.AddCommand(GetGys)
	//GetGys.Flags().StringVarP(&param.filepath, "filepath", "f", "", "从文本中读取目标公司名称")
}
