#!/bin/sh

if [ -d "/usr/share/games/fortune" ] 
then
    FORTUNE_FOLDER="/usr/share/games/fortune" 
elif [ -d "/usr/share/fortune" ]
then
    FORTUNE_FOLDER="/usr/share/fortune" 
else
    echo "No fortune folder found. Is fortune installed?"
    exit 1
fi

sudo cp ./my-fortune ${FORTUNE_FOLDER}/my-fortune
cd ${FORTUNE_FOLDER}
sudo strfile my-fortune 1>/dev/null
echo "${FORTUNE_FOLDER}/my-fortune was created"
fortune -f my-fortune

cd - 1>/dev/null
