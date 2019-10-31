# AppCheck
## Godfrey Macasero
AppCheck is a web application that has an html frontend and a Go backend. This application puts a GUI to the apt command. It lets you see what applications you currently have installed in your Ubuntu or Debian machine. You can install, upgrade, or delete files that are in your Ubuntu or Debian machine.

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
## Cloning
```bash
git clone https://github.com/gmac220/project1.git
cd project1
```

## Running
```bash
sudo go run main.go
```
## Served up on [localhost](http://localhost)
Put localhost on your browser or click link above.

### [Presentation](https://gitpitch.com/gmac220/project1/master)