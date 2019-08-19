mysqldump -u root -p TestJTree > ./tests/sql/jtree_backup.sql
mysql -u "root" "-pwaterloo" < "./tests/sql/DropDatabase.sql"
mysql -u "root" "-pwaterloo" < "./tests/sql/CreateDatabase.sql"
mysql -u "root" "-pwaterloo" < "./tests/sql/CreateTables.sql"
mysql -u "root" "-pwaterloo" < "./tests/sql/CreateUsers.sql"


