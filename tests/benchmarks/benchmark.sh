#!/bin/bash
# make build
# for i in `seq 1 6`;
#  do
# mysql -u "root" "-pwaterloo" "Jtree" < "delete.sql"
# ten=10
# echo $i
# j=$(echo "$ten^$i" | bc)
# process=$(echo "./bin/jtree -g="$j" &> /dev/null")
# $process
sum=0
for i in `seq 1 100`;
        do
            add=$(curl -X POST -w %{time_total} -o /dev/null -s -H 'Content-Type: application/json' http://127.0.0.1:8000/Jtree/metadata/0.1.0/query -d '{"selected_fields": ["*"],"selected_tables": ["samples"],"selected_conditions": []}')
            sum=$(echo "$sum + $add" | bc)
        done 
hun=100
sum=$(echo "$sum / $hun"| bc -l)
echo $sum
# kill $!
# done 
