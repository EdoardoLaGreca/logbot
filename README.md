# logbot

Please note that this repo has nothing to do with logging. It's just a Discord bot.

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/EdoardoLaGreca/logbot/Makefile%20CI)

If the build fails, it's probably the fault of GitHub CI whose version of Go is out of date.

## Roadmap

 - [x] Working draft
 - [x] `grp` command to get the one of the rising posts from a subreddit
 - [x] `help` command

## Compile & run

You can do pretty much everything using make.

**Remember that, before running, there must be a file named `DISCORD_TOKEN` in the working directory. That file must contain YOUR bot's discord token and nothing else.**

Dependencies:

 - [Golang](https://go.dev/)
 - `make`
 - `git` (to update the repo)
 - external Golang libraries which will be installed automatically

To compile:

```
make build
```

To run without compiling:

```
make run
```

To update to the latest commit (this will just update the tracked repo files, it won't build an updated version):

```
make update
```

To clean the repo from files generated through the Makefile:

```
make clean
```

## License

See [LICENSE](/LICENSE)
