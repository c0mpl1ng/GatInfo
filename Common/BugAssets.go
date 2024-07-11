package Common

import (
	"GatInfo/Modules/BugAssets"
	"github.com/spf13/cobra"
)

var GetBugAssets = &cobra.Command{
	Use:   "GetBugAssets",
	Short: "通过已知有漏洞的通用组件指纹收集资产",
	Run: func(cmd *cobra.Command, args []string) {
		if param.City != "" {
			BugAssets.Entry(param.City, "City")
		} else if param.Region != "" {
			BugAssets.Entry(param.Region, "Region")
		}

	},
}

func init() {

	rootCmd.AddCommand(GetBugAssets)
	GetBugAssets.Flags().StringVarP(&param.City, "City", "", "", "获取指定城市资产")
	GetBugAssets.Flags().StringVarP(&param.Region, "Region", "", "", "获取指定地区资产")

}
