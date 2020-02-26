# RSS Web App

A web app that allows users search for RSS feed items using keywords

## Getting Started

To get started with this project, git clone the repository by running

```
git clone https://github.com/benjorah/rss_web_app.git
```

### Prerequisites

To get the most out of the project, **Go** version 1.11 is required (due to go module support). The file go.mod also contains al;l the dependency for this project.
For copnverting performance profiles to pdf, the 3rd party tool **graphviz**. Get this buy running the follwoing command if on a mac
```
brew install graphviz
```
For full-text search support you would require MySQL 5.6 or higher installed. You can achive this by installing the **MAMP** server locally https://www.mamp.info/en/downloads/

### Installing

Right after installing downloading the project, cd into your project folder, build the project (which also gets the required dependency) using the following command

```
go build .
```
Make a copy of the **.env.example** and name it **.env** and fill with the required information.

Start your local MAMP server and access to phpMydmin from your browser (usually found at http://localhost:8888/phpMyAdmin/).

Find a file named **schema.sql** and import into phpMydmin to create your database and table

Run the app to start the server with the following command (while in your project directory)

```
./rss_web_app
```

To fetch the RSS feed before starting the server, start the app with the  following command ( you should do this the first time around in order to get some data in your database)


```
./rss_web_app -fetchrss
```

#End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Explain how to run the automated tests for this system

### Running unit tests

To run the unit test in the system with the following command

```
go test
```

### Running benchmark tests

This tests the efficiency of our function by calling it multiple times and recording the execution time. This gives a sense of how well we implement our functions

```
go test -bench .
```

We can generate memory and cpu profiles using the following command 

```
go test -cpuprofile rssreaderPkgCpu.prof -memprofile rssreaderPkgMem.prof -bench=. ./pathToPackage/...
```
This creates/edits two files **cpu.prof** and **mem.prof**

To convert this files to pdf for easy viewing, run the follwoing 2 commands

```
go tool pprof --pdf ~/path/to/project/cpu.prof > cpu.pdf
go tool pprof --pdf ~/path/to/project/mem.prof > mem.pdf
```

The memeory and cpu profile for the rssreader package have been generated and stored as rssreaderPkgMem.* and rssreaderPkgCpu.* respectively

## Built With

* [GoFeed](https://github.com/mmcdole/gofeed) - For parsing RSS
* [Go-Sql-Driver](https://github.com/go-sql-driver/mysql) - An sql driver for Golang


## Authors

* **Onwuorah Okechukwu** 

