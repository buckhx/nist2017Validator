## NIST 2017 Password Compliance Validator
### Project Overview
This project is a password validation tool that implements the National Institute of Standards and Technology's 
June 2017 guidelines on user supplied password security. It completely rewrites previously held policy by removing
rules related to the composition of passwords, mandatory expiration periods and hints. The new rule set implements
the following requirements for a password to be valid:
* The password has an character minimum
* There must be a minimum maximum character limit of at least 64 characters, though there should be no limit ideally
* Allow all ASCII compliant text, unicode support is optional
* The password cannot appear on a list of commonly used passwords

### Build Steps
* Clone this repository from GitHub.
```bash
git clone https://www.github.com/aminasian/nist2017Validator
```
* Change directory to your cloned repository.
```bash
cd $GOPATH/src/github.com/aminasian/nist2017Validatory
```
* Make sure all dependencies are installed with dep. If you don't have dep installed please see
[here](https://github.com/golang/dep)
```bash
dep ensure
```
* Build the project
```bash
go build -o nist-2017-validator
```

### How To Run
After following the above build steps you can easily run your validator by executing the following command
```bash
export PWS=path/to/passwords.txt
export COMMON_PWS=path/to/common/passwords.txt
cat $PWS | ./nist-2017-validator --common-pws=$COMMON_PWS
```
You can also use optional build flags listed below to customize your validation parameters within NIST's guidelines.
#### Runtime Flags
A number of build flags can be used to help customize password validation parameters.

|Flag|Mandatory|Usage|Default|
|:--:|:-------:|:---:|:-----:|
|common-pws|Yes|Path of file containing commonly used passwords to be checked against.|"""|
|max-char|No|Maximum number of characters to allow in a user supplied password. Cannot be less then 64|128|
|ascii-only|No|Allows to validate with all unicode compliant passwords.|true|

Example
```bash
cat user_passwords.txt | ./nist-2017-validator --common-pws=/Usr/Documents/common_pws.txt --max-char=64 --ascii-only=false
```