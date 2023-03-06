package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"encoding/json"
	"fmt"
	ber "github.com/go-asn1-ber/asn1-ber"
	"github.com/go-ldap/ldap/v3"
	"golang.org/x/sys/windows"
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

type Dacl struct {
	AccountSID string   `json:"accountSID"`
	AceType    string   `json:"aceType"`
	AceFlags   []string `json:"aceFlags"`
	Rights     []string `json:"rights"`
	ObjectGUID string   `json:"objectGUID"`
}

type Data struct {
	Owner   string `json:"owner"`
	Primary string `json:"primary"`
	Dacl    []Dacl `json:"dacl"`
}

type aclPermissions struct {
	CN   string
	NAME string
}

// 判断数组中是否包含某字符串
func in(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

// 结构体去重
type messageContentsNormal []aclPermissions

func removeDuplicate(personList messageContentsNormal) messageContentsNormal {
	resultMap := map[string]bool{}
	for _, v := range personList {
		data, _ := json.Marshal(v)
		resultMap[string(data)] = true
	}
	result := messageContentsNormal{}
	for k := range resultMap {
		var t aclPermissions
		json.Unmarshal([]byte(k), &t)
		result = append(result, t)
	}
	return result
}

type ControlMicrosoftSDFlags struct {
	Criticality  bool
	ControlValue int32
}

func (c *ControlMicrosoftSDFlags) GetControlType() string {
	return "1.2.840.113556.1.4.801"
}

func (c *ControlMicrosoftSDFlags) Encode() *ber.Packet {
	packet := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "Control")
	packet.AppendChild(ber.NewString(ber.ClassUniversal, ber.TypePrimitive, ber.TagOctetString, "1.2.840.113556.1.4.801", "Control Type"))
	packet.AppendChild(ber.NewBoolean(ber.ClassUniversal, ber.TypePrimitive, ber.TagBoolean, true, "Criticality"))
	p2 := ber.Encode(ber.ClassUniversal, ber.TypePrimitive, ber.TagOctetString, nil, "Control Value(SDFlags)")
	seq := ber.Encode(ber.ClassUniversal, ber.TypeConstructed, ber.TagSequence, nil, "SDFlags")
	seq.AppendChild(ber.NewInteger(ber.ClassUniversal, ber.TypePrimitive, ber.TagInteger, c.ControlValue, "Flags"))
	p2.AppendChild(seq)

	packet.AppendChild(p2)
	return packet
}

func (c *ControlMicrosoftSDFlags) String() string {
	return fmt.Sprintf("Control Type: %s (%q)  Criticality: %t  Control Value: %d", "1.2.840.113556.1.4.801",
		"1.2.840.113556.1.4.801", c.Criticality, c.ControlValue)
}

func SearchSddl(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var app conf.ACLProcessor
	var listSddl []aclPermissions
	var completelyAcl []aclPermissions
	var changePasswordAcl []aclPermissions
	var dcsyncAcl []aclPermissions
	var selfAcl []aclPermissions
	var addDcsyncAcl []aclPermissions
	searchSddl := ldap.NewSearchRequest(
		domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["aclUser"],
		[]string{"nTSecurityDescriptor", "cn"},
		[]ldap.Control{&ControlMicrosoftSDFlags{ControlValue: 7}},
	)

	searchSddls, err := (*l).SearchWithPaging(searchSddl, 10000)
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range searchSddls.Entries {
		a := aclPermissions{
			entry.GetAttributeValue("cn"),
			entry.GetAttributeValue("nTSecurityDescriptor"),
		}
		listSddl = append(listSddl, a)
	}

	// 进一步转换成sddl string
	for _, j := range listSddl {
		header := (*reflect.SliceHeader)(unsafe.Pointer(&j.NAME))
		advapi32Dll := windows.NewLazyDLL("advapi32.dll")
		ConvertSecurityDescriptorToStringSecurityDescriptorW := advapi32Dll.NewProc("ConvertSecurityDescriptorToStringSecurityDescriptorW")
		var sddl *uint16
		r, _, _ := syscall.Syscall6(
			ConvertSecurityDescriptorToStringSecurityDescriptorW.Addr(),
			5,
			uintptr(unsafe.Pointer(header.Data)),
			1,
			0xff,
			uintptr(unsafe.Pointer(&sddl)),
			0,
			0)
		_ = r
		defer windows.LocalFree(windows.Handle(unsafe.Pointer(sddl)))
		//fmt.Println(j.CN)
		// sddl解析
		acl, err := app.Processor(windows.UTF16PtrToString(sddl))
		if err != nil {
			panic(err)
		}
		//fmt.Println(acl)
		//json格式输出
		var c Data
		err = json.Unmarshal([]byte(acl), &c)
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			return
		}
		for _, k := range c.Dacl {
			switch {
			//判断完全控制
			case in("ADS_RIGHT_DS_CREATE_CHILD", k.Rights) && in("ADS_RIGHT_DS_DELETE_CHILD", k.Rights) && in("ADS_RIGHT_ACTRL_DS_LIST", k.Rights) && in("ADS_RIGHT_DS_SELF", k.Rights) && in("ADS_RIGHT_DS_READ_PROP", k.Rights) && in("ADS_RIGHT_DS_WRITE_PROP", k.Rights) && in("ADS_RIGHT_DS_DELETE_TREE", k.Rights) && in("ADS_RIGHT_DS_LIST_OBJECT", k.Rights) && in("ADS_RIGHT_DS_CONTROL_ACCESS", k.Rights) && in("DELETE", k.Rights) && in("READ_CONTROL", k.Rights) && in("WRITE_DAC", k.Rights) && in("WRITE_OWNER", k.Rights) && !in("ACE IS INHERITED", k.AceFlags):
				//fmt.Println(k.Rights)
				if strings.Contains(k.AccountSID, "S-1-5-21-") {
					filter1 := fmt.Sprintf("(&(objectSid=" + k.AccountSID + "))")
					msDSConstraint := ldap.NewSearchRequest(
						domain,
						ldap.ScopeWholeSubtree,
						ldap.NeverDerefAliases,
						ldapSizeLimit,
						0,
						false,
						filter1,
						[]string{"cn"}, // Attributes []string
						nil,
					)
					sr1, _ := (*l).SearchWithPaging(msDSConstraint, 10000)
					for _, j1 := range sr1.Entries {
						for _, i1 := range j1.GetAttributeValues("cn") {
							if i1 == "Enterprise Admins" || i1 == "Organization Management" || i1 == "Exchange Trusted Subsystem" || i1 == "Exchange Servers" {
								continue
							} else {
								a := aclPermissions{
									j.CN,
									i1,
								}
								completelyAcl = append(completelyAcl, a)
							}
						}
					}
				}
				break
			//判断修改密码权限
			case in("ADS_RIGHT_DS_WRITE_PROP", k.Rights) && in("WRITE_DAC", k.Rights) && !in("ACE IS INHERITED", k.AceFlags):
				if strings.Contains(k.AccountSID, "S-1-5-21-") {
					filter1 := fmt.Sprintf("(&(objectSid=" + k.AccountSID + "))")
					msDSConstraint := ldap.NewSearchRequest(
						domain,
						ldap.ScopeWholeSubtree,
						ldap.NeverDerefAliases,
						ldapSizeLimit,
						0,
						false,
						filter1,
						[]string{"cn", "objectClass"}, // Attributes []string
						nil,
					)
					sr1, _ := (*l).SearchWithPaging(msDSConstraint, 10000)
					for _, j1 := range sr1.Entries {
						for _, i1 := range j1.GetAttributeValues("objectClass") {
							if i1 != "user" {
								continue
							} else if j.CN == "" {
								for _, i2 := range j1.GetAttributeValues("cn") {
									a := aclPermissions{
										j.CN,
										i2,
									}
									addDcsyncAcl = append(addDcsyncAcl, a)
								}
							} else {
								for _, i2 := range j1.GetAttributeValues("cn") {
									a := aclPermissions{
										j.CN,
										i2,
									}
									changePasswordAcl = append(changePasswordAcl, a)
								}
							}
						}
					}
				}
				break
			//判断DCSync权限
			case in("ADS_RIGHT_DS_CONTROL_ACCESS", k.Rights) && !in("ACE IS INHERITED", k.AceFlags):
				//fmt.Println(k.Rights)
				if strings.Contains(k.AccountSID, "S-1-5-21-") {
					filter1 := fmt.Sprintf("(&(objectSid=" + k.AccountSID + "))")
					msDSConstraint := ldap.NewSearchRequest(
						domain,
						ldap.ScopeWholeSubtree,
						ldap.NeverDerefAliases,
						ldapSizeLimit,
						0,
						false,
						filter1,
						[]string{"cn", "objectClass"}, // Attributes []string
						nil,
					)
					sr1, _ := (*l).SearchWithPaging(msDSConstraint, 10000)
					for _, j1 := range sr1.Entries {
						for _, i1 := range j1.GetAttributeValues("objectClass") {
							if i1 != "user" {
								continue
							} else if "Domain Admins" == j.CN || "Enterprise Admins" == j.CN || "Organization Management" == j.CN || "Exchange Trusted Subsystem" == j.CN || "Exchange Servers" == j.CN {
								continue
							} else {
								for _, i2 := range j1.GetAttributeValues("cn") {
									a := aclPermissions{
										j.CN,
										i2,
									}
									dcsyncAcl = append(dcsyncAcl, a)
								}
								dcsyncAcl = removeDuplicate(dcsyncAcl)
							}
						}
					}
				}
				break
			//判断Self-Membership权限
			case in("ADS_RIGHT_DS_SELF", k.Rights) && !in("ACE IS INHERITED", k.AceFlags):
				//fmt.Println(k.Rights)
				if strings.Contains(k.AccountSID, "S-1-5-21-") {
					filter1 := fmt.Sprintf("(&(objectSid=" + k.AccountSID + "))")
					msDSConstraint := ldap.NewSearchRequest(
						domain,
						ldap.ScopeWholeSubtree,
						ldap.NeverDerefAliases,
						ldapSizeLimit,
						0,
						false,
						filter1,
						[]string{"cn", "objectClass"}, // Attributes []string
						nil,
					)
					sr1, _ := (*l).SearchWithPaging(msDSConstraint, 10000)
					for _, j1 := range sr1.Entries {
						for _, i1 := range j1.GetAttributeValues("objectClass") {
							if i1 != "user" {
								continue
							} else {
								for _, i2 := range j1.GetAttributeValues("cn") {
									a := aclPermissions{
										j.CN,
										i2,
									}
									selfAcl = append(selfAcl, a)
								}
							}
						}
					}
				}
				break
			}
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Acl :\n", outputFile)
		for _, i := range completelyAcl {
			if i.CN == "" {
				process.OutFile("\t"+i.NAME+" 完全控制 ------> "+domain, outputFile)
			} else {
				process.OutFile("\t"+i.NAME+" 完全控制 ------> "+i.CN, outputFile)
			}
		}
		for _, i := range changePasswordAcl {
			process.OutFile("\t"+i.NAME+" 修改密码 ------> "+i.CN, outputFile)
		}
		for _, i := range dcsyncAcl {
			process.OutFile("\t"+i.NAME+" 拥有DCSync权限", outputFile)
		}
		for _, i := range selfAcl {
			process.OutFile("\t"+i.NAME+" 将自己添加到 ------> "+i.CN, outputFile)
		}
		for _, i := range addDcsyncAcl {
			process.OutFile("\t"+i.NAME+" 拥有添加DCSync权限", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Acl save file to: %s\n", outputFile)
	} else {
		fmt.Printf("[*] Acl :\n")
		for _, i := range completelyAcl {
			if i.CN == "" {
				fmt.Println("\t" + i.NAME + " 完全控制 ------> " + domain)
			} else {
				fmt.Println("\t" + i.NAME + " 完全控制 ------> " + i.CN)
			}
		}
		for _, i := range changePasswordAcl {
			fmt.Println("\t" + i.NAME + " 修改密码 ------> " + i.CN)
		}
		for _, i := range dcsyncAcl {
			if i.CN == "" {
				fmt.Println("\t" + i.NAME + " 拥有DCSync权限")
			} else {
				continue
			}
		}
		for _, i := range selfAcl {
			fmt.Println("\t" + i.NAME + " 将自己添加到 ------> " + i.CN)
		}
		for _, i := range addDcsyncAcl {
			fmt.Println("\t" + i.NAME + " 拥有添加DCSync权限")
		}
		fmt.Printf("\n")
	}
}
