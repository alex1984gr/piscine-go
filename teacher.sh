#1/bin/bash
KEYWORDS="saw|heard|weapon|car|license|object|testimony|suspect"
echo "Searching for relevant interviews..." MATCHED_FILES=$(grep -iE "$KEYWORDS" mystery/interviews/* | cut -d ':' -f1 | sort -u)
echo "Matched interviews:" for FILE in $MATCHED_FILES; do INTERVIEW_NUM=$(echo "$FILE" | grep -oE '[0-9]+') echo "Interview $INTERVIEW_NUM:" cat "$FILE" echo "-----------------------------" done
echo "Extracting suspects from interviews..." SUSPECTS=$(grep -i "suspect:" $MATCHED_FILES | cut -d ':' -f2 | xargs -n1 | sort | uniq)
echo "Checking suspects against crimescene..." for NAME in $SUSPECTS; do if grep -iq "$NAME" mystery/crimescene; then echo "Likely suspect: $NAME" export MAIN_SUSPECT="$NAME" break
