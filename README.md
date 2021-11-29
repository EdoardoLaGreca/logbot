# logbot

Please note that this repo has nothing to do with logging. It's just a Discord bot.

## Roadmap

 - [x] Working draft
 - [x] `grp` command to get the one of the rising posts from a subreddit
 - [x] `help` command

## Compile & run

You can do pretty much everything using make.

**Remember that, before running, there must be a file named `DISCORD_TOKEN` in the same directory where you run the bot. That file must contain YOUR bot's discord token and nothing else.**

Dependencies:

 - [Golang](https://go.dev/)
 - `make`
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
