# AppCheck
## Godfrey Macasero
AppCheck is a web application that has an html frontend and a Go backend. It lets you see what applications you currently have installed in your Ubuntu or Debian machine.

# User Stories
- [x] Users should be able to:
    - [x] See a list of application downloads in their linux OS
    - [x] Click on a downloaded application on the list and upgrade it
    - [x] Click on a downloaded application on the list and uninstall it
    - [x] Search up an application
    - [x] Install an application
    - [x] See the current version of the application clicked
- [] When running on a EC2 instance make a script that automatically runs the file on startup(currently runs when you ssh into the EC2 instance)

# Instructions
```bash
git clone https://github.com/gmac220/project1.git
sudo go run main.go
On browser put localhost on the url
```
### [Presentation](https://gitpitch.com/gmac220/project1/master)