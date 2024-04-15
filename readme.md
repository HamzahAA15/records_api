# Records API Service

Records API Service is a service that enables you to get the data recorded of student's marks. It uses a persistence postgresql data source, during the initiation it will do migration for DB schema & inject seeder data into records table if it have not exist yet.

**Note:** View this readme file using preview mode (Ctrl-K + V) or (Command-K + V) for better readability.

## Features

- Modular project structure with dependency injection on the repository, service & controller layers.
- Using Docker Compose to ease the experience of using this service
- `Ready to use DB, since we run migration and data seeder only for you :)`

## Usage

Run the following command:

```bash
docker-compose build
docker-compose up

accesss through `http://localhost:8080/`
```

import the JSON collection of request from the attachment of email into API platform such as Postman
 
## API LIST

### Records Endpoints
<details>

**Get list of records**
- **URL:** `/api/records`
- **Method:** `GET`
- **Description:** Retrieves a list of records available as per criteria requested.
- **Request Body:** 
  - minCount is minimal count of sum of Marks, must be lower and not equal to maxCount
  - maxCount is maximal count of sum of marks, must be higher and not equal to minCount
  - startDate is starting date of the data recorded, must be earlier and not equal to endDate
  - endDate is end date of the data recorded, must be later and not equal to startDate
- **Response:**
  - Returns success message, status code, and list of records with details like id, totalMarks, and createdAt.
  - Returns an error message and status code

</details>

## DB Schema
<details>
<summary>Click to toggle the database schema</summary>

```sql
CREATE TABLE IF NOT EXISTS records (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    marks INTEGER[] NOT NULL,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

```
</details>
