grant SELECT on JTree.* to 'select'@'%' identified by 'passwords';
flush privileges;
grant SELECT,INSERT, UPDATE on JTree.* to 'update'@'%' identified by 'passwordu';
flush privileges;