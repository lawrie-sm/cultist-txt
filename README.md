# cultist-txt

This is a [Twitter bot](https://twitter.com/cultist_txt), which periodically posts text entries from [Cultist Simulator](https://weatherfactory.biz/cultist-simulator/).

## Data

The text entries have been extracted and cleaned using a shell script, [extract.sh](https://github.com/lawrie-sm/cultist-txt-private/blob/master/extract.sh). They are then stored in an SQLite database. This is a simpler solution than trying to repair and parse the game's JSON files.

The script requires pcregrep and sqlite3 on the path and the game's [core data folder](https://cultistsimulator.gamepedia.com/Modding#Modding_game_files) should be copied into the project's data directory.

*TODO: Copyright and permission notice*
