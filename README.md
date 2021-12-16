![](https://socialify.git.ci/MariamFahmy98/Distributed-Marketplace-System/image?descriptionEditable=Marketplace%20System%20where%20sellers%20and%20customers%20to%20sell%20and%20buy%20their%20desired%20products%20online%20and%20facilitates%20the%20shopping%20from%20many%20different%20sources.&font=Inter&forks=1&issues=1&logo=https%3A%2F%2Fraw.githubusercontent.com%2FMariamFahmy98%2FDistributed-Marketplace-System%2Fmain%2Ffrontend%2Fpublic%2Flogo512.png&pattern=Solid&pulls=1&stargazers=1&theme=Dark)

# Table of contents
* [Distributed Marketplace System](#distributed-marketplace-system)
* [Features](#features)
* [Built With](#built-with)
* [Getting Started](#getting-started)

# Distributed Marketplace System
It is online marketplace where it allows the sellers and customers to sell and buy their desired products online and facilitates the shopping from many different sources. Each user can create his own store where he can add his products so that other customers can buy their desired products. The store owner can buy products from the other users’ store and add the purchased products to his own store.

# Features
   - User can create his own store.
   - User can add new products to his store specifying the needed price.
   - User can modify the current products in his store.
   - User can delete a specific product in his store.
   - User can search for products in other stores.
   - User can purchase products from another user’s store, the money required is transfered from the buyer’s account to the seller’s account, and the purchased product is added to the buyer’s store.
   - User can view his account info such as current cash balance, a list of purchased products, a list of sold products, and the products that haven’t been sold yet.
   - User can sell the products of other stores.
   - Other stores can sell user’s products.

# Built With
   - [Gin](https://github.com/gin-gonic/gin#gin-web-framework) - HTTP web framework written in Go (Golang).
   - [GORM](https://gorm.io/index.html) - Golang ORM for mapping data directly to PostgreSQL.
   - [PostgreSQL](https://www.postgresql.org/) - Open source object-relational database system.
   - [JSON Web Token](https://jwt.io/) - A standard to securely authenticate HTTP requests.
   - [React](https://reactjs.org/) - JavaScript library for building user interfaces.
   - [Redux](https://redux.js.org/) - JavaScript library to help better manage application state.
   - [Docker](https://www.docker.com/) - A set of platform as a service (PaaS) products that use OS-level virtualization.
   - [Kubernetes](https://kubernetes.io/) - An open-source system for automating deployment, scaling, and management of containerized applications.

# Getting Started
## Dependencies
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

- Then create a database by running: 
  ```
  CREATE DATABASE ds_db;
  ```
<!-------------------------------------------------------------------------------------------------->
## Running the code
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
- List all services in the namespace to get marketplace service:
  ```
  kubectl get services
  ```
- Get the kubernetes URL for a marketplace service in your local minikube cluster:
  ```
  minikube service [marketplace-service-name]
  ```
- To have a look on the database stored in postgresql-db-0:
  ```
  kubectl exec -ti postgresql-db-0 -- psql -U postgres -p 5432 postgres
  ```
  ***Same in case of postgreql-db-1***
  
