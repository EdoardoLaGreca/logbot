# logbot

Please note that this repo has nothing to do with logging. It's just a Discord bot.

## Roadmap

 - [x] Working draft

## Compile & run

You can do pretty much everything using make.

**Remember that, before running, there must be a file named `DISCORD_TOKEN` in the same directory where you run the bot. That file must contain YOUR bot's discord token and nothing else.**

To compile run this:

```
make
```

To run without compiling run this:

```
make run
```

To get the dependencies (which is done automagically before compiling or running, so you shouldn't really use this):

```
make deps
```

## License

See [LICENSE](/LICENSE)
