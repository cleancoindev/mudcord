# Contributing to mudcord

I intend for this to be pretty easy going.


## General code style and semantics

* Decently commented (check out the other comments, it doesn't need to be too much)
* Standard Go code formatting (easily achieved by running `go fmt`)
* Item structs are start with "Item", room structs start with "Room", etc. ("ItemFoo", "RoomBar")


## Structure of the project

### Branches

* `master`: The current running version of the bot (production)
* `release`: The latest stable changes
* Anything else will be tested and then merged to `release`, then after a bit `release` will be merged to `master`

### Files

Each file name is pretty self descriptive.

I'll update this with more indepth guides to structuring things if there is interest.


## Workflow

Just standard stuff.

* Make a fork
* Make a branch with a relevant name
* Do your stuff on that branch
* Test out your changes (make your own discord bot client to run it, and invite it to a server to test)
* File a PR


## Compiling and running the bot

To compile just make sure your in the right directory and have installed the dependencies logrus and discordgo, and it's as simple as `go build`

Before you run it, you'll need to make your own client and add a bot to it on the Discord dev panel.

Then just set the environment variable `MUDCORD_TOKEN` to your secret token and run the executable.
