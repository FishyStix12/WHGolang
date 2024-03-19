# WHGolang <br />
# By: Nicholas Fisher <br />
**Disclaimer: These scripts should only be used in cases where the user has permission to use these scripts on the subject systems!** <br />

**Import Note: Please read this statement carefully: By downloading any of the scripts in this repository, you, as the user, take full responsibility for storing, and using these scripts. You also take full responsibility for any misuse of this malicious codes. Finally, Please note that any data the Trojan extracts will be posted to a GitHub repository, and if that repository is public, all the extracted data will be available for the whole world to see.** <br />

# Go Linux Setup <br />
![image](https://github.com/FishyStix12/White-Hat-Go-/assets/102126354/05e0dca4-a04f-4ed1-842e-5ab143ace30a) <br />
**Please run these scripts in your home directory!** <br />

**The Following List gives a short description of all the scripts in this group:** <br />
1. goconfig.sh - This Bash script installs the Go programming language on a Linux system using the `sudo apt install -y golang` command. It then sets up environment variables (`GOROOT`, `GOPATH`, and updates `PATH`) in the `~/.bashrc` file to configure the Go development environment. Finally, it reloads the `~/.bashrc` file using the `source ~/.bashrc` command to apply the changes immediately. <br />
2. gomods.sh - This Bash script installs the latest versions of several Go programs from their respective GitHub repositories. The programs being installed are: <br />
               1. dalfox by hahwul for web application vulnerability scanning. <br />
               2. qsreplace by tomnomnom for URL query string manipulation. <br />
               3. hakrawler by hakluke for website crawling and analysis. <br />
               4. httprobe by tomnomnom for HTTP/HTTPS probe. <br />
               5. wildcheck by theblackturtle for wildcard domain detection. <br />
               6. gau by lc for Google Analytics data collection. <br />
               7. assetfinder by tomnomnom for finding domains and subdomains. <br />

# Fundamentals of Go Programming <br />
![image](https://github.com/FishyStix12/White-Hat-Go-/assets/102126354/4650c928-f66a-494b-9c05-e02af87698e9) <br />
**The Following List gives a short description of all the scripts in this group:** <br />
1. hello_world.go - The script is a simple Go program that demonstrates the basic structure of a Go application. It starts by declaring that it belongs to the `main` package, imports the "fmt" package for formatting input and output, defines the `main` function as the entry point of the program, and then prints "Hello, White Hat Gophers!" to the console using the `fmt.Println` function. This script serves as a starting point for learning how to write and execute Go programs. <br />
