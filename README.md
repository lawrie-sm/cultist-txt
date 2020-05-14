# cultist-txt

A Twitter bot which periodically posts text from [Cultist Simulator's](https://weatherfactory.biz/cultist-simulator/) in-game descriptions.

### More Information

Text entries are extracted from the game files using a shell script. They are  stored in an SQLite database. The script requires pcregrep and sqlite3. The game's [core data folder](https://cultistsimulator.gamepedia.com/Modding#Modding_game_files) should be copied into the project's data directory.

The bot will grab a random entry that hasn't yet been posted, format it correctly, and post it to Twitter using [anaconda](https://github.com/ChimeraCoder/anaconda). It will do this until every entry has been posted.

*TODO: Copyright and permission notice*
