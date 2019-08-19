mysqldump -u root -p JTree > ./sql/jtree_backup.sql
mysql -u "root" "-pwaterloo" < "./sql/DropDatabase.sql"
mysql -u "root" "-pwaterloo" < "./sql/CreateDatabase.sql"
mysql -u "root" "-pwaterloo" < "./sql/CreateTables.sql"
mysql -u "root" "-pwaterloo" < "./sql/CreateUsers.sql"


