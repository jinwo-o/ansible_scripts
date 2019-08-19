grant SELECT on TestJTree.* to 'select'@'%' identified by 'passwords';
flush privileges;
grant SELECT, INSERT, UPDATE on TestJTree.* to 'update'@'%' identified by 'passwordu';
flush privileges;