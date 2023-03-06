package conf

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"encoding/json"
	"github.com/pkg/errors"
)

var sddlRights = map[string]string{
	// Generic access rights
	"GA": "GENERIC_ALL",
	"GR": "GENERIC_READ",
	"GW": "GENERIC_WRITE",
	"GX": "GENERIC_EXECUTE",
	// Standard access rights
	"RC": "READ_CONTROL",
	"SD": "DELETE",
	"WD": "WRITE_DAC",
	"WO": "WRITE_OWNER",
	// Directory service object access rights
	"RP": "ADS_RIGHT_DS_READ_PROP",
	"WP": "ADS_RIGHT_DS_WRITE_PROP",
	"CC": "ADS_RIGHT_DS_CREATE_CHILD",
	"DC": "ADS_RIGHT_DS_DELETE_CHILD",
	"LC": "ADS_RIGHT_ACTRL_DS_LIST",
	"SW": "ADS_RIGHT_DS_SELF",
	"LO": "ADS_RIGHT_DS_LIST_OBJECT",
	"DT": "ADS_RIGHT_DS_DELETE_TREE",
	"CR": "ADS_RIGHT_DS_CONTROL_ACCESS",
	// File access rights
	"FA": "FILE_ALL_ACCESS",
	"FR": "FILE_GENERIC_READ",
	"FW": "FILE_GENERIC_WRITE",
	"FX": "FILE_GENERIC_EXECUTE",
	// Registry key access rights
	"KA": "KEY_ALL_ACCESS",
	"KR": "KEY_READ",
	"KW": "KEY_WRITE",
	"KX": "KEY_EXECUTE",
	// Mandatory label rights
	"NR": "SYSTEM_MANDATORY_LABEL_NO_READ_UP",
	"NW": "SYSTEM_MANDATORY_LABEL_NO_WRITE_UP",
	"NX": "SYSTEM_MANDATORY_LABEL_NO_EXECUTE",
}

var sddlInheritanceFlags = map[string]string{
	"P":  "DDL_PROTECTED",
	"AI": "SDDL_AUTO_INHERITED",
	"AR": "SDDL_AUTO_INHERIT_REQ",
}

var sddlAceType = map[string]string{
	"D":  "ACCESS DENIED",
	"OA": "OBJECT ACCESS ALLOWED",
	"OD": "OBJECT ACCESS DENIED",
	"AU": "SYSTEM AUDIT",
	"OU": "OBJECT SYSTEM AUDIT",
	"OL": "OBJECT SYSTEM ALARM",
	"A":  "ACCESS ALLOWED",
}

var sddlAceFlags = map[string]string{
	"CI": "CONTAINER INHERIT",
	"OI": "OBJECT INHERIT",
	"NP": "NO PROPAGATE",
	"IO": "INHERITANCE ONLY",
	"ID": "ACE IS INHERITED",
	"SA": "SUCCESSFUL ACCESS AUDIT",
	"FA": "FAILED ACCESS AUDIT",
}

var sddlSidsRep = map[string]string{
	"O":  "Owner",
	"AO": "Account operators",
	"PA": "Group Policy administrators",
	"RU": "Alias to allow previous Windows 2000",
	"IU": "Interactively logged-on user",
	"AN": "Anonymous logon",
	"LA": "Local administrator",
	"AU": "Authenticated users",
	"LG": "Local guest",
	"BA": "Built-in administrators",
	"LS": "Local service account",
	"BG": "Built-in guests",
	"SY": "Local system",
	"BO": "Backup operators",
	"NU": "Network logon user",
	"BU": "Built-in users",
	"NO": "Network configuration operators",
	"CA": "Certificate server administrators",
	"NS": "Network service account",
	"CG": "Creator group",
	"PO": "Printer operators",
	"CO": "Creator owner",
	"PS": "Personal self",
	"DA": "Domain administrators",
	"PU": "Power users",
	"DC": "Domain computers",
	"RS": "RAS servers group",
	"DD": "Domain controllers",
	"RD": "Terminal server users",
	"DG": "Domain guests",
	"RE": "Replicator",
	"DU": "Domain users",
	"RC": "Restricted code",
	"EA": "Enterprise administrators",
	"SA": "Schema administrators",
	"ED": "Enterprise domain controllers",
	"SO": "Server operators",
	"WD": "Everyone",
	"SU": "Service logon user",
}

var sddlWellKnownSidsRep = map[string]string{
	"S-1-0":        "Null Authority",
	"S-1-0-0":      "Nobody",
	"S-1-1":        "World Authority",
	"S-1-1-0":      "Everyone",
	"S-1-2":        "Local Authority",
	"S-1-2-0":      "Local ",
	"S-1-2-1":      "Console Logon ",
	"S-1-3":        "Creator Authority",
	"S-1-3-0":      "Creator Owner",
	"S-1-3-1":      "Creator Group",
	"S-1-3-2":      "Creator Owner Server",
	"S-1-3-3":      "Creator Group Server",
	"S-1-3-4":      "Owner Rights ",
	"S-1-4":        "Non-unique Authority",
	"S-1-5":        "NT Authority",
	"S-1-5-1":      "Dialup",
	"S-1-5-2":      "Network",
	"S-1-5-3":      "Batch",
	"S-1-5-4":      "Interactive",
	"S-1-5-6":      "Service",
	"S-1-5-7":      "Anonymous",
	"S-1-5-8":      "Proxy",
	"S-1-5-9":      "Enterprise Domain Controllers",
	"S-1-5-10":     "Principal Self",
	"S-1-5-11":     "Authenticated Users",
	"S-1-5-12":     "Restricted Code",
	"S-1-5-13":     "Terminal Server Users",
	"S-1-5-14":     "Remote Interactive Logon ",
	"S-1-5-15":     "This Organization ",
	"S-1-5-17":     "This Organization ",
	"S-1-5-18":     "Local System",
	"S-1-5-19":     "NT Authority",
	"S-1-5-20":     "NT Authority",
	"S-1-5-32-544": "Administrators",
	"S-1-5-32-545": "Users",
	"S-1-5-32-546": "Guests",
	"S-1-5-32-547": "Power Users",
	"S-1-5-32-548": "Account Operators",
	"S-1-5-32-549": "Server Operators",
	"S-1-5-32-550": "Print Operators",
	"S-1-5-32-551": "Backup Operators",
	"S-1-5-32-552": "Replicators",
	"S-1-5-64-10":  "NTLM Authentication ",
	"S-1-5-64-14":  "SChannel Authentication ",
	"S-1-5-64-21":  "Digest Authentication ",
	"S-1-5-80":     "NT Service ",
	"S-1-5-80-0":   "All Services ",
	"S-1-5-83-0":   "NT VIRTUAL MACHINE\\Virtual Machines",
	"S-1-16-0":     "Untrusted Mandatory Level ",
	"S-1-16-4096":  "Low Mandatory Level ",
	"S-1-16-8192":  "Medium Mandatory Level ",
	"S-1-16-8448":  "Medium Plus Mandatory Level ",
	"S-1-16-12288": "High Mandatory Level ",
	"S-1-16-16384": "System Mandatory Level ",
	"S-1-16-20480": "Protected Process Mandatory Level ",
	"S-1-16-28672": "Secure Process Mandatory Level ",
	"S-1-5-32-554": "BUILTIN\\Pre-Windows 2000 Compatible Access",
	"S-1-5-32-555": "BUILTIN\\Remote Desktop Users",
	"S-1-5-32-556": "BUILTIN\\Network Configuration Operators",
	"S-1-5-32-557": "BUILTIN\\Incoming Forest Trust Builders",
	"S-1-5-32-558": "BUILTIN\\Performance Monitor Users",
	"S-1-5-32-559": "BUILTIN\\Performance Log Users",
	"S-1-5-32-560": "BUILTIN\\Windows Authorization Access Group",
	"S-1-5-32-561": "BUILTIN\\Terminal Server License Servers",
	"S-1-5-32-562": "BUILTIN\\Distributed COM Users",
	"S-1-5-32-569": "BUILTIN\\Cryptographic Operators",
	"S-1-5-32-573": "BUILTIN\\Event Log Readers ",
	"S-1-5-32-574": "BUILTIN\\Certificate Service DCOM Access ",
	"S-1-5-32-575": "BUILTIN\\RDS Remote Access Servers",
	"S-1-5-32-576": "BUILTIN\\RDS Endpoint Servers",
	"S-1-5-32-577": "BUILTIN\\RDS Management Servers",
	"S-1-5-32-578": "BUILTIN\\Hyper-V Administrators",
	"S-1-5-32-579": "BUILTIN\\Access Control Assistance Operators",
	"S-1-5-32-580": "BUILTIN\\Remote Management Users",
	"S-1-5-80-956008885-3418522649-1831038044-1853292631-2271478464": "Trusted Installer",
}

// ACLProcessor main struct with methods
type ACLProcessor struct {
	Rights permissions
	File   string
}

type entryACL struct {
	AccountSid        string   `json:"accountSID,omitempty"`
	AceType           string   `json:"aceType,omitempty"`
	AceFlags          []string `json:"aceFlags,omitempty"`
	Rights            []string `json:"rights,omitempty"`
	ObjectGUID        string   `json:"objectGUID,omitempty"`
	InheritObjectGUID string   `json:"inheritObjectGUID,omitempty"`
}

type permissions struct {
	Owner     string     `json:"owner,omitempty"`
	Primary   string     `json:"primary,omitempty"`
	Dacl      []entryACL `json:"dacl,omitempty"`
	DaclInher []string   `json:"daclInheritFlags,omitempty"`
	Sacl      []entryACL `json:"sacl,omitempty"`
	SaclInger []string   `json:"saclInheritFlags,omitempty"`
}

// checkSIDsFile check file of SIDs where data saved in SID,User
func checkSIDsFile(filePath string, sid string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Split(scanner.Text(), ",")[0] == sid {
			return strings.Split(scanner.Text(), ",")[1]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sid
}

// sidReplace replace identification account: sid/wellkhownsid/usersid
func (app *ACLProcessor) sidReplace(str string) string {
	if len(str) > 2 {
		if x, ok := sddlWellKnownSidsRep[str]; ok {
			return x
		} else if app.File != "" {
			return checkSIDsFile(app.File, str)
		}
		return str
	}
	return app.replacer(sddlSidsRep, str)[0]
}

// replacer chunk string with 2 letters, add to array and then resolve
func (app *ACLProcessor) replacer(maps map[string]string, str string) []string {
	var temp, result []string
	if len(str) > 2 {
		for j := 0; j < len(str)-1; j = j + 2 {
			temp = append(temp, fmt.Sprintf("%s%s", string(str[j]), string(str[j+1])))
		}
	} else {
		temp = append(temp, str)
	}
	for _, v := range temp {
		if x, ok := maps[v]; ok {
			result = append(result, x)
		} else {
			result = append(result, v)
		}
	}
	return result
}

/*
	splitBodyACL Convert values from string to struct with replace strings

Base format Rights: (ace_type;ace_flags;rights;object_guid;inherit_object_guid;account_sid)
*/
func (app *ACLProcessor) splitBodyACL(str string) entryACL {
	splitACL := strings.Split(str, ";")
	return entryACL{
		AceType:           app.replacer(sddlAceType, splitACL[0])[0],
		AceFlags:          app.replacer(sddlAceFlags, splitACL[1]),
		Rights:            app.replacer(sddlRights, splitACL[2]),
		ObjectGUID:        splitACL[3],
		InheritObjectGUID: splitACL[4],
		AccountSid:        app.sidReplace(splitACL[5]),
	}
}

func (app *ACLProcessor) splitBody(body string) []entryACL {
	var entryACLInternalArr []entryACL
	for _, y := range strings.Split(body, "(") {
		if y != "" {
			ace := strings.TrimSuffix(y, ")")
			entryACLInternalArr = append(entryACLInternalArr, app.splitBodyACL(ace))
		}
	}
	return entryACLInternalArr
}

func (app *ACLProcessor) parseBody(body string) ([]string, []entryACL) {
	var inheritFlagArr []string
	var entryACLInternalArr []entryACL
	if strings.Index(body, "(") != 0 {
		inheritFlag := body[0:strings.Index(body, "(")]
		ace := body[strings.Index(body, "("):]
		if len(inheritFlag) > 2 {
			for j := 0; j < len(inheritFlag)-1; j = j + 2 {
				inheritFlagArr = append(inheritFlagArr, app.replacer(sddlInheritanceFlags, fmt.Sprintf("%s%s", string(inheritFlag[j]), string(inheritFlag[j+1])))[0])
			}
		}
		entryACLInternalArr = app.splitBody(ace)
	} else {
		entryACLInternalArr = app.splitBody(body)
	}
	return inheritFlagArr, entryACLInternalArr
}

func (app *ACLProcessor) parseSDDL(sddrArr []string) {
	for _, y := range sddrArr {
		sddlSplit := strings.Split(y, ":")
		letter := sddlSplit[0]
		body := sddlSplit[1]
		switch letter {
		case "O":
			app.Rights.Owner = app.sidReplace(body)
		case "G":
			app.Rights.Primary = app.sidReplace(body)
		case "D":
			app.Rights.DaclInher, app.Rights.Dacl = app.parseBody(body)
		case "S":
			app.Rights.SaclInger, app.Rights.Sacl = app.parseBody(body)
		default:
			log.Fatal("Unresolved group")
		}
	}
}

// slice SDDL create slice objects from str to array of strings
func (app *ACLProcessor) sliceSDDL(indecs []int, str string) {
	var sddlArr []string
	for i := 0; i < len(indecs)-1; i++ {
		sl := str[indecs[i]:indecs[i+1]]
		sddlArr = append(sddlArr, sl)
	}
	app.parseSDDL(sddlArr)
}

// FindGroupIndex used for find index of group Owner, Primary, DACL, SACL
func (app *ACLProcessor) findGroupIndex(str string) error {
	groups := []string{"O:", "G:", "D:", "S:"}
	var result []int
	for _, i := range groups {
		if strings.Index(str, i) != -1 {
			result = append(result, strings.Index(str, i))
		}
	}
	if result == nil {
		return errors.New("Can't find any group")
	}
	result = append(result, len(str))
	app.sliceSDDL(result, str)
	return nil
}

// Processor main function in gosddl package
func (app *ACLProcessor) Processor(str string) (sddl string, err error) {
	err = app.findGroupIndex(str)
	if err != nil {
		return "", err
	}
	body, err := json.Marshal(app.Rights)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(body), nil
}
