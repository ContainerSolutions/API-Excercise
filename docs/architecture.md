# Modules Architecture and Design Decisions

I separated the project in the following modules: `cmd`, `endpoint`,
`parser`, `storage`, `titanic`, `util`. Below is a description of each of
these:

## cmd

`cmd` is the entry-point for the project and contains the `main` function and
all the CLI commands that are used to run the project.

I decided to use `Cobra` to help creating the CLI and, as suggested on their
documentation, I created different files for each CLI command. Right now it
has only two commands: `listen`, that runs the REST API server, and
`populate`, that will populate a configured Postgres instance with the
titanic dataset.

Also, I decided to create all `.go` files from this package under folder
`server`. This way, if you run `go install ./...`, instead of creating
a binary named `cmd`, it creates a binary named `server`.

## endpoint

`endpoint` contains the implementation of the server REST API and its
configuration. I separated the implementation of the endpoints in a different
package, named `apiv1`; this way, if someone want to implement a different
api in the future, its only necessary to change a few lines on the `endpoint`
package. The usage is kinda straightforward:

```go
svrCfg := endpoint.BuildSvrConfig("dev") // can be "test" or "prod"
server := endpoint.BuildSvr(svrCfg)
server.Serve() // listens on port :8080
```

## titanic

`titanic` contains the implementation of the `Person` model, that represents
people from the Titanic dataset. I made to sure to have no calls to
`postgres` in the `titanic` package, such that if you want to change to a
different database, you don't have to change anything in the business/model
layer.

## parser

`parser` contains the implementation of the Titanic CSV dataset. You
might parse a different Titanic dataset the following way:
```go
parser.Parse("newdataset.csv")
```

## storage

`storage` contains the integration between the models and the database. For
this project I decided to use Postgres, but made sure to separate its usage
in a different package named `postgres`. Also, I separated all the
interaction with table `people` in a specific DAO (data-acess object); if I
need to add interaction with another table, I would just need to write a new
DAO dedicated to this new table. The biggest upside of this design is
testability: I was able to easily test the `PersonDAO` without interference
from any other packages.

## util

`util` contains general-utility functions that I had to use across all other
packages.

