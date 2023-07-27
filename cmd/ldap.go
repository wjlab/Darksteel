package cmd

import (
	"darksteel/ldap"
	"github.com/spf13/cobra"
)

var ldapCmd = &cobra.Command{
	Use:   "ldap",
	Short: "ldap查询",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		ldap.LdapInit(domain, dc, password, user, allDelegate, searchValue, integrate, outputContent, ldapSizeLimit, outputFile, allLdap, fuzz)
	},
}

func init() {
	ldapCmd.Flags().StringVarP(&user, "user", "u", "", "域内的用户名")
	ldapCmd.Flags().StringVarP(&password, "pass", "p", "", "对应的密码或哈希用户")
	ldapCmd.Flags().StringVarP(&allDelegate, "delegate", "w", "", "all \n  输出所有委派 \nuw \n  非约束委派 \ncw \n  约束委派 \nbw \n  基于资源的约束委派")
	ldapCmd.Flags().StringVarP(&searchValue, "searchldap", "f", "", "自定义查询ldap")
	ldapCmd.Flags().StringVarP(&outputContent, "outputcontent", "t", "", "要查询的字段（可以写多个）")
	ldapCmd.Flags().StringVarP(&integrate, "integrate", "m", "", "user \n  查询域内的所有用户 \ncomputer \n  查询域内的所有计算机 \nscomputer \n  查询域内存活的计算机 \ndc \n  查询域内的所有域控 \nspn \n  查询域内的所有SPN \nou \n  查询域内的所有OU \nmssql \n  查询域内的所有mssql服务 \nasreproast \n  查询域内可以使用 as-rep roast 的用户 \nmaq \n  查询域内maq的值\nadmins \n  查询域管理员\nenterprise \n  查询企业管理员\nexchangecomputer \n  查询exchange服务器\nexchangesystem \n  查询 Exchange Trusted Subsystem\nexchangeorgmanager \n  查询 Exchange Organization Management\ntrustdomain \n  查询信任域\nadminsdholder \n  查询设置了AdminSDHolder权限的用户\nsidhistory \n  查询已设置SIDHistory的用户\ncacomputer \n  查询 adcs\nesc1 \n  受到esc1威胁的模板\nesc2 \n  受到esc2威胁的模板\ncomputerip \n  查询域中计算机的ip地址\nsddl \n  查询配置错误的acl")
	ldapCmd.Flags().IntVarP(&ldapSizeLimit, "ldapsizelimit", "l", 0, "查询LDAP最大数目（默认为0）")
	ldapCmd.Flags().StringVarP(&outputFile, "outputfile", "o", "", "输出文件的位置（默认当前目录）")
	ldapCmd.Flags().BoolVarP(&allLdap, "allLdap", "a", false, "查询所有内容")
	ldapCmd.Flags().StringVarP(&fuzz, "fuzz", "z", "", "模糊查询")
	rootCmd.AddCommand(ldapCmd)
}
