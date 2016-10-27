# Contributing

Welcome to the cdatools project! Below you will find instructions
on how to begin contributing to this repository.

## Contributing Code

1. Clone the repo.
2. Install dependencies using `go get`:
    - http://github.com/mooveweb/gokogiri/xml ??
    - http://github.com/moovweb/gokogiri/xpath ??
    - http://github.com/pebbe/util
    - http://github.com/pborman/uuid
	- http://github.com/stretchr/testify/assert
    - http://gopkg.in/check.v1 ??

Make sure the tests pass before you start:

    go test ./...
    
Then create an issue for your change if it does not already exist. Try to
describe some user stories as checklist items in the description so there is a
definition of done.

When you're ready to start work on the issue, go to the project board and move it
to "In Progress".

Before making changes, create your own branch. The name should reflect the 
change. For example, the branch that made this contribution guide is named 'contributing'.

Also, make good [commit messages](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html) while you work.

## Writing Tests
Open hqmf\_qrda\_oids.json and search for the file name of the template you are making. This file will give a name to the oid.
In the patient data folder, given by one of the team members, `ag` for the name in the measures folder.
Choose a file that contains a result in the search and search in the file for that particular record.

Now go into the patient folder and `ag` for the even more specific name for a particular record.

Take both of these json objects and put them in the respective measures and patients folders under fixtures. Give them the
same name that describes them well.

Then copy an existing test method and change the variables to the file names of the fixture files.

Use the xpath strategy of checking values in the template to assure that the template is working as expected.

## Pull Requests
Make your change, with new passing tests.

Push your branch up to remote. Go on github and make a pull request to have it 
merged into master.

Others will give constructive feedback.
This is a time for discussion and improvements,
and making the necessary changes will be required before we can
merge the contribution.