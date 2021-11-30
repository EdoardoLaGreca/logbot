# logbot

Please note that this repo has nothing to do with logging. It's just a Discord bot.

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/EdoardoLaGreca/logbot/Makefile%20CI)

If the build fails, it's the fault of GitHub CI whose version of Go is out of date.

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

To get the dependencies (which is done automagically before compiling or running, so you really don't want to use this):

```
make deps
```

## License

See [LICENSE](/LICENSE)
