#!/bin/bash

echo 'Extracting entries from files...'

# Store current dir so we can run this from elsewhere
dir=$(dirname "$0")/data
entries="$dir/entries"
sql="$dir/init.sql"
db="$dir/core.db"

# Refresh the output file
echo '' > $entries

# Grab all descriptions and startdescription entries
# This is the vast majority of the data
# We get the strings within quotes and save them to a file
pcregrep -rhM -o1 '(?:startdescription|description):\s*\"((?:\n|.)*?)\"' $dir/core/* >> $entries

# Grab all drawmessage entries from the mansus file
# First grab the drawmessages object, then quoted strings within it
# Picks up some messy data. Dealt with by removing short lines
pcregrep -rhM -o1 'drawmessages:\s*{((?:\n|.)*?)}' $dir/core/decks/mansus.json | pcregrep -rhM -o1 '\"((?:\n|.)*?)\"' >> $entries

# Replace all carriage returns with <br>'s
# This is consistent with other rich text styling used in the entries
# It will bring multiline strings together and ensure newlines are reliable separators
# This must be done before sorting!
sed -zi 's/\r\n/<br>/g' $entries

# Sort and keep only the unique lines
sort -u $entries --output=$entries

# Remove more cruft lines, based on their beginnings
sed -ri '/^BEGIN|^\. |^,/d' $entries

# Replace empty variables with appropriate placeholders
sed -i 's/#PREVIOUSCHARACTERNAME#/The Aspirant/g' $entries
sed -i 's/#LAST_DESIRE#/Temptation/g' $entries
sed -i 's/#LAST_FOLLOWER#/Renira/g' $entries
sed -i 's/#LAST_BOOK#/The Rose of Hypatia/g' $entries

# Remove the in-game helper text within square brackets [] - this was a difficult decision!
sed -ri 's/\[.*?\].*//g' $entries

# Remove game data variable info, between # and | symbols
sed -i 's/#.*|//g' $entries

# Remove @ symbols
sed -i 's/@//g' $entries

# Trim all the lines
sed -i 's/^[ \t]*//;s/[ \t]*$//' $entries

# Escape single quotes for SQL
sed -i 's/\x27/\x27\x27/g' $entries

# Remove short lines
# Using 24 chars here to remove some cruft
sed -ri '/^.{,24}$/d' $entries

# Generate SQL for SQLite. Slow single line inserts get the job done
echo 'Saving to DB...'
echo 'DROP TABLE IF EXISTS entries' | sqlite3 $db
echo 'CREATE TABLE entries(id INTEGER PRIMARY KEY, entry TEXT NOT NULL, postcount INTEGER NOT NULL);' | sqlite3 $db
echo '' > $sql
while read entry
do
    echo "INSERT INTO entries (entry, postcount) VALUES (\"$entry\", 0);" >> $sql
done < $entries
sqlite3 $db < $sql

# Clean up
rm $sql
rm $entries

echo 'Done!'
