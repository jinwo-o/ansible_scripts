	if [ $CI ]
	then
        echo "Travis Complete"
    else	
        mysql -u "root" "-pwaterloo" < "./tests/sql/DeleteData.sql" 
	fi