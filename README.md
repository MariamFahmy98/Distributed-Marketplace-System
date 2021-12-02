## Installing Go (Ubuntu 20.04):

- Downloading and Extracting the Go binary archive in the /usr/local directory:
  `wget -c https://dl.google.com/go/go1.17.3.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local`

- Adding the location of the Go directory to the $PATH environment variable:
  `sudo nano ~/.profile`
  - Appending the following line in `~/.profile` :
    `export PATH=$PATH:/usr/local/go/bin`
  - Save the file, and load the new PATH environment variable into the current shell session:
    `source ~/.profile`
- Verifying Go Installation by `go version`, output should be something like this:
  > go version go1.17.3 linux/amd64

## Installing PostgreSQL

```
sudo apt update
sudo apt install postgresql postgresql-contrib
```

## Creating Database in PostgreSQL

To create a new database in PostgreSQL, you need to access the PostgreSQL database shell **_psql_**

- The installation procedure created a user account called **_postgres_**, we need to to switch over to the **_postgres_** account on your server by running: `sudo -i -u postgres`
- You can access the Postgres prompt by running: `psql`
- Creating a password for the created user account by running: `\password postgres`, then enter the password: 12345
- Then create a database by running: `CREATE DATABASE gorm;`

## Some usefull PostgreSQL commands

Entering the new created database by: `\c gorm`
Exit from **_psql_** by: `\q`

## Running the code

- Get the dependencies for code in the current directory by running `go get .`
- Running the code: `go run .`
- In a new terminal, use curl to make a request to your running web service:
  - Displaying all data stored in postgres table: `curl http://localhost:8080/users`
  - Adding new user to the database: `curl http://localhost:8080/users --include --header "Content-Type: application/json" --request "POST" --data '{"id": 5,"name": "Essam","email": "essam272@gmail.com","password": "hey"}'`
