# README #

## goenv ##
* goenv is a simple tool that mimics python's virtualenv, for use in go projects.
* Version 0.1

### Installation ###

Using go 1.9.2 or higher:

```
#!bash

git install github.com/greatercommonsfernando/goenv
```
### Usage ###

Once goenv is installed, you may use it similarly to virtualenv

```
#!bash

goenv ~/goprojects/test_go_project
cd ~/goprojects/test_go_project
source bin/activate.sh
```

On Windows:

```
goenv c:\Users\yourhome\goprojects\test_go_project
cd c:\Users\yourhome\goprojects\test_go_project
call bin\activate.bat

```


Note: you may source the activate script from any location

```
#!bash

source ~/projects/test_go_project/bin/activate
```

### Dependencies ###

* bash (Linux, MacOS) or cmd (Windows)
* find
* install
* A working Go 1.3+ installation

### Contribution guidelines ###

* Have a bug or an issue? Have a pull request? File an issue on the issue tracker!