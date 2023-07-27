package cmd

import (
	"darksteel/ldap"
	"github.com/spf13/cobra"
)

var computerIpCmd = &cobra.Command{
	Use:   "computerip",
	Short: "查询域内计算机的ip地址",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		ldap.LdapInit(domain, dc, password, user, allDelegate, searchValue, "computerip", outputContent, ldapSizeLimit, outputFile, allLdap, file)
	},
}

func init() {
	computerIpCmd.Flags().StringVarP(&user, "user", "u", "", "域内的用户名")
	computerIpCmd.Flags().StringVarP(&password, "pass", "p", "", "对应的密码或哈希用户")
	computerIpCmd.Flags().StringVarP(&file, "file", "f", "", "从文件中查询机器对应ip")
	computerIpCmd.Flags().IntVarP(&ldapSizeLimit, "ldapsizelimit", "l", 0, "查询LDAP最大数目（默认为0）")
	computerIpCmd.Flags().StringVarP(&outputFile, "outputfile", "o", "", "输出文件的位置（默认当前目录）")
	rootCmd.AddCommand(computerIpCmd)
}
