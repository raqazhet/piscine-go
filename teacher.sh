NUMBER_1=$(head -n 179 .streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $NUMBER_1
cat ./interviews/interview-"$NUMBER_1"
echo $MAIN_SUPECT
