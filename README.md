# GoLang course step by step

## Introduction

This repository is focussed in show step by step the implementations of the course of Golang in Edutin
to ensure all the concepts are properly learned chaptere by chapter.

## Basic commands

First of all I suggest init the repository and make a initial commit with no files. Them start by adding this line to the commands with the user name and repository name you are using. This will initialize a module to use. This can be found in [LordCeilan]: <https://github.com/LordCeilan/introToGo>

For more Module creation use this syntax: 

```bash
go mod init github/{username}/{repositoryName}/{moduleName}
```

To run the code you just edited use the next command:

```bash
go run main.go
```

To run all the edited files run:

```bash
go run .
```

The next command will create the compilate files, they are another version of the .dll files or the .exe.  Now we are using .mod. Also if you wish an .exe file using this command will solve your issue.

```bash
go build -o name.go
```

To create a binary ommit the .go extension in the last command

```bash
go build -o name
```

If using the terminal then just run through a shell

```bash
./main
```

To run a compilated file use:

```bash
./main.exe
```
