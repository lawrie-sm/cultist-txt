# cultist-txt

A Twitter bot which periodically posts text from [Cultist Simulator's](https://weatherfactory.biz/cultist-simulator/) in-game descriptions.

### More Information

The text entries are extracted and cleaned using a shell script. They are then stored in an SQLite database. The script requires pcregrep and sqlite3, and the game's [core data folder](https://cultistsimulator.gamepedia.com/Modding#Modding_game_files) should be copied into the project's data directory.

The bot itself is entirely contained in main.go. It will grab a random entry that hasn't been posted, format it correctly. and post it to Twitter using [anaconda](https://github.com/ChimeraCoder/anaconda).

*TODO: Copyright and permission notice*
