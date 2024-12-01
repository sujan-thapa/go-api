# Instructions for setting up and running the application.
### Requirements
```
1. Golang: https://go.dev/doc/install 
2. MySQL: https://www.apachefriends.org/download.html
```

### Setting up MySQL Database
```
1. Start MySQL and Apache server
2. Create a database and then table for storing time log of Toronto.

```
### Script or instructions for setting up the MySQL database and table.
```
1. Download and Install Xampp: https://www.apachefriends.org/download.html 
2. Start the MySQL server and go to the localhost/phpmyadmin
3. From there you can manually set up a database and create a table or you can execute the mysql query to create a database and table. i.e:
    i. CREATE DATABASE IF NOT EXISTS time_api;
    ii. CREATE TABLE time_log (id INT AUTO_INCREMENT PRIMARY KEY,timestamp DATETIME NOT NULL);

```

### Go Application
```
1. Go to the github and clone the project. https://github.com/sujan-thapa/go-api.git 
2. Install dependencies: i.e, 
    go mod init (project-name)
    go mod tidy: this will clean up the unused packages

3. Run the application
    go run main.go

4. Then using the http://localhost:8080/current-time, you can see the current timespan of Toronto in a JSON format.
```