# cultist-txt

This is a Twitter bot, which periodically posts text entries from [Cultist Simulator](https://weatherfactory.biz/cultist-simulator/).

## Data

The text entries are xtracted and cleaned using a shell script. They are then stored in an SQLite database. This is a simpler solution than trying to repair and parse the game's JSON files.

The script requires pcregrep and sqlite3 on the path and the game's [core data folder](https://cultistsimulator.gamepedia.com/Modding#Modding_game_files) should be copied into the project's data directory.

*TODO: Copyright and permission notice*
