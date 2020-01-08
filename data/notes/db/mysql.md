General
-------

Login to mysql server:

    mysql --defaults-file=/etc/mysql/debian.cnf
    
Show processes:

    SHOW PROCESSLIST;

Identify the user you are logged in as:

    SELECT USER();          # logged in as
    SELECT CURRENT_USER();  # connected as

Obtain some user info:

    SELECT host,user,password,Grant_priv,Super_priv FROM mysql.user;

Change user password:

    SET PASSWORD FOR 'debian-sys-maint'@'localhost' = PASSWORD('ThePassword');
    FLUSH PRIVILEGES;

Grant privileges for user:

    GRANT SELECT ON bacula.* TO 'bakstat'@'1.2.3.4' IDENTIFIED BY 'ThePassword';
    
Show grants for user:

    SHOW GRANTS FOR 'bakstat'@'1.2.3.4';
    
Replication
-----------

Master/slave status:

    SHOW MASTER STATUS\G
    SHOW SLAVE STATUS\G
    
Change IP address of slave on master:

    UPDATE mysql.user SET host='NEW_IP' WHERE host='OLD_IP';
    
Stop slave, skip one error, start slave (both threads - Slave_IO_Running and Slave_SQL_Running):

    STOP SLAVE;
    SET GLOBAL SQL_SLAVE_SKIP_COUNTER=1;
    START SLAVE;

More

* MySQL High Availability
 * Chapter 3. MySQL Replication Fundamentals 
* https://www.digitalocean.com/community/tutorials/how-to-set-up-master-slave-replication-in-mysql
