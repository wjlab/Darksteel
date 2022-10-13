package process

import (
	"strings"
)

var domainLdap string

//将格式wanliu1.com 转换为 DC=wanliu1,DC=com
func DcFormatConversion(domain string) string {
	a := strings.Split(domain, ".")
	for _, j := range a {
		domainLdap += "DC=" + j + ","
	}
	return domainLdap[0 : len(domainLdap)-1]
}
