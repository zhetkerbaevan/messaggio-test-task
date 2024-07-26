# Message processing microservice  
This microservice accepts messages via HTTP API, saves them to PostgreSQL, sends them to Kafka, and marks them as processed.  
It also provides an API to fetch statistics on processed messages.  
## Setup
1. Clone repository
```sh
git clone https://github.com/zhetkerbaevan/messaggio-test-task.git
cd messaggio-test-task
```
2. Install Dependencies
 ```sh
go mod tidy
```
3. Database Configuration  
* Set up PostgreSQL database.  
* Configure connection details in internal/config/config.go  
4. Run Migrations
 ```sh
make migrate-up
```
5. Start application
 ```sh
make run
```
## API Endpoints
POST /api/v1/message - Send a new message (Only content).      
GET /api/v1/statistics - Retrieve statistics about processed messages.    
   
