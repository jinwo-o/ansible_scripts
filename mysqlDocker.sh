docker exec -i mysqldb mysql -u root -pwaterloo -e "CREATE DATABASE JTree"
docker exec -i mysqldb mysql -u root -pwaterloo JTree < ./sql/jtree_backup.sql
docker exec -i mysqldb mysql -u root -pwaterloo -e "grant SELECT on JTree.* to 'select'@'%' identified by 'passwords';flush privileges;grant SELECT,INSERT, UPDATE on JTree.* to 'update'@'%' identified by 'passwordu';flush privileges;"