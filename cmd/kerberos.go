package cmd

import (
	"darksteel/kerberos"
	"github.com/spf13/cobra"
)

var kerberosCmd = &cobra.Command{
	Use:   "kerberos",
	Short: "kerberos利用",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		kerberos.KerberosInit(domain, dc, password, user, roastModule, targetUser, format, enctype, outputFile, ldapSizeLimit, ticket)
	},
}

func init() {
	kerberosCmd.Flags().StringVarP(&user, "user", "u", "", "域内的用户名")
	kerberosCmd.Flags().StringVarP(&password, "pass", "p", "", "对应的密码或哈希用户")
	kerberosCmd.Flags().StringVarP(&roastModule, "roastmodule", "m", "", "asreproast \n  as-rep roast利用\nkerberoast \n  kerberoasting利用")
	kerberosCmd.Flags().StringVarP(&targetUser, "targetuser", "t", "", "脆弱易攻击的用户")
	kerberosCmd.Flags().StringVarP(&format, "format", "f", "hashcat", "输出格式为John the Ripper 或 Hashcat 可利用格式")
	kerberosCmd.Flags().StringVarP(&enctype, "enctype", "e", "rc4", "加密类型：RC4、AES128、AES256")
	kerberosCmd.Flags().StringVarP(&ticket, "ticket", "k", "", "使用票证认证时，输入票证的路径")
	kerberosCmd.Flags().IntVarP(&ldapSizeLimit, "ldapSizelimit", "l", 0, "查询LDAP最大数目（默认为0）")
	rootCmd.AddCommand(kerberosCmd)
}
