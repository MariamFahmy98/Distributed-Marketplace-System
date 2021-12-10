# Table of contents
* [Dependencies](#dependencies)
* [Running the code](#running-the-code)

# Dependencies
<b>1- Installing Go (Ubuntu 20.04): </b>

- Downloading and Extracting the Go binary archive in the /usr/local directory:
  ```
  wget -c https://dl.google.com/go/go1.17.3.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
  ```

- Adding the location of the Go directory to the $PATH environment variable:
  ```
  sudo nano ~/.profile
  ```
  -  Appending the following line in ``` ~/.profile ``` :
     ```
     export PATH=$PATH:/usr/local/go/bin
     ```
  - Save the file, and load the new PATH environment variable into the current shell session:
     ```
     source ~/.profile
     ```
- Verifying Go Installation by 
  ```
  go version
  ```
  The output should be something like this:  
  > go version go1.17.3 linux/amd64
<!-------------------------------------------------------------------------------------------------->
</br>
<b>2- Installing PostgreSQL</b>

- Using apt to install it:
  ```
  sudo apt update
  sudo apt install postgresql postgresql-contrib
  ```
<!-------------------------------------------------------------------------------------------------->
</br>
<b>3- Creating Database in PostgreSQL</b>

To create a new database in PostgreSQL, you need to access the PostgreSQL database shell ***psql***
- The installation procedure created a user account called ***postgres***, we need to to switch over to the ***postgres*** account on your server by running:
  ```
  sudo -i -u postgres
  ```
- You can access the Postgres prompt by running: 
  ```
  psql
  ```
- Creating a password for the created user account by running: 
  ```
  \password postgres
  ```
  Then enter the password that exist in ***.env*** file

- Then create a database by running: 
  ```
  CREATE DATABASE ds_db;
  ```

<!-------------------------------------------------------------------------------------------------->

# Running the code
- Start local minikube cluster by running:
  ```
  minikube start
  ```
- Navigate into the k8s directory and create kubernetes resources in the cluster:
  ```
  cd k8s
  kubectl apply -f .
  ```
- List all pods in the namespace to make sure that all pods are running:
  ```
  kubectl get pods
  ```
- Dump the marketplace server pod logs by running:
  ```
  kubectl logs -f [marketplace-pod-name]
  ```
