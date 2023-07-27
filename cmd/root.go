package cmd

import (
	"darksteel/conf"
	"github.com/spf13/cobra"
	"os"
)

var (
	dc            string
	domain        string
	user          string
	password      string
	allDelegate   string
	searchValue   string
	outputContent string
	integrate     string
	ldapSizeLimit int
	outputFile    string
	allLdap       bool
	fuzz          string
	userFile      string
	passwordFile  string
	roastModule   string
	targetUser    string
	format        string
	enctype       string
	ticket        string
	blastModule   string
	threads       int
	verbose       bool
	userPass      string
	file          string
)

var rootCmd = &cobra.Command{
	Use:   "darksteel",
	Short: "",
	Long:  "自动化域内信息搜集、kerberos利用工具",
}

func Execute() {
	conf.Banner()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dc, "dc", "d", "", "域控地址")
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "n", "", "域名")
}
