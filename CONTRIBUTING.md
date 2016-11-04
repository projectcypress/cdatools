## Table of Contents
1. [Introduction](#introduction)
1. [Set Up](#set-up)
    - [List of current dependencies](#list-of-current-dependencies)
1. [Writing Tests](#writing-tests)
1. [Pull Requests](#pull-requests)

## Introduction

Welcome to the cdatools project! Below you will find instructions
on how to begin contributing to this repository.

## Set Up
1. Clone the repo.
1. Install [dependencies](#list-of-current-dependencies) using `go get`
    - A short way of pulling them all in is `go get -u ./...`
1. Make sure the tests pass before you start using `go test ./...`

_Then create an issue for your change if it does not already exist. Try to describe some user stories as checklist items in the description so there is a definition of done._

_When you're ready to start work on the issue, go to the project board and move it to "In Progress"._

_Before making changes, create your own branch. The name should reflect the change. For example, the branch that made this contribution guide is named 'contributing'._

Also, make good [commit messages](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html) while you work.

### List of current dependencies
```
http://github.com/jbowtie/gokogiri/xml
http://github.com/jbowtie/gokogiri/xpath
http://github.com/pebbe/util
http://github.com/pborman/uuid
http://github.com/stretchr/testify/assert
http://gopkg.in/check.v1
http://github.com/jteeuwen/go-bindata/...
```

## Writing Tests
1. Open hqmf\_qrda\_oids.json and search for the file name of the template you are making. This file will give a name to the oid.
1. In the patient data folder, given by one of the team members, `ag` for the name in the measures folder.
1. Choose a file that contains a result in the search and search in the file for that particular record.
1. Now go into the patient folder and `ag` for the even more specific name for a particular record.
1. Take both of these json objects and put them in the respective measures and patients folders under fixtures. Give them the same name that describes them well.
1. Then copy an existing test method and change the variables to the file names of the fixture files.
1. Use the xpath strategy of checking values in the template to assure that the template is working as expected.

## Pull Requests
1. Make your change, with new passing tests.
1. Push your branch up to remote. Go on github and make a pull request to have it merged into master.
1. Others will give constructive feedback. This is a time for discussion and improvements, and making the necessary changes will be required before we can merge the contribution.