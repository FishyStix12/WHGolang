# White Hat Go
# By: Nicholas Fisher

**Disclaimer: These scripts should only be used in cases where the user has permission to use these scripts on the subject systems!** <br />

**Import Note: Please read this statement carefully: By downloading any of the scripts in this repository, you, as the user, take full responsibility for storing, and using these scripts. You also take full responsibility for any misuse of this malicious codes. Finally, Please note that any data the Trojan extracts will be posted to a GitHub repository, and if that repository is public, all the extracted data will be available for the whole world to see.** <br />

# Go Linux Setup <br />
![image](https://github.com/FishyStix12/White-Hat-Go-/assets/102126354/05e0dca4-a04f-4ed1-842e-5ab143ace30a) <br />
**Please run these scripts in your home directory!** <br />

**The Following List gives a short description of all the scripts in this group:** <br />
1. goconfig.sh - This Bash script installs the Go programming language on a Linux system using the `sudo apt install -y golang` command. It then sets up environment variables (`GOROOT`, `GOPATH`, and updates `PATH`) in the `.bashrc` file to configure the Go development environment. Finally, it reloads the `.bashrc` file using the `source .bashrc` command to apply the changes immediately. <br />
2. gomods.py - This script is a Python program designed to install specified Golang modules. It defines a function `install_golang_module` that takes a Golang module name as input and installs it if it's not already installed. The script iterates through a list of Golang modules to install and calls the `install_golang_module` function for each module. It uses environment variables and system commands (`os.system`) to execute the installation commands. <br />
