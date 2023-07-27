package cmd

import (
	"darksteel/blast"
	"github.com/spf13/cobra"
)

var blastCmd = &cobra.Command{
	Use:   "blast",
	Short: "爆破域内用户",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		blast.SetupSession(domain, dc, user, userFile, password, passwordFile, userPass, threads, blastModule, outputFile, verbose)
	},
}

func init() {
	blastCmd.Flags().StringVarP(&user, "user", "u", "", "域内的用户名")
	blastCmd.Flags().StringVarP(&userFile, "userfile", "U", "", "用户字典")
	blastCmd.Flags().StringVarP(&password, "pass", "p", "", "单个密码")
	blastCmd.Flags().StringVarP(&passwordFile, "passfile", "P", "", "密码字典")
	blastCmd.Flags().StringVarP(&userPass, "upfile", "F", "", "用户名和密码对应的字典,内容以\":\"分割")
	blastCmd.Flags().IntVarP(&threads, "threads", "t", 20, "线程数量")
	blastCmd.Flags().StringVarP(&blastModule, "blastmodule", "m", "", "userenum -userfile user.txt\n  用户枚举\npassspray -userfile user.txt -pass password\n  密码喷洒\nblastpass -user username -passfile password.txt\n  单用户爆破密码\nuserpass -upfile userpass.txt\n  用户密码组合爆炸")
	blastCmd.Flags().StringVarP(&outputFile, "outputfile", "o", "", "输出文件的位置（默认当前目录）")
	blastCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "显示失败信息")
	rootCmd.AddCommand(blastCmd)
}
