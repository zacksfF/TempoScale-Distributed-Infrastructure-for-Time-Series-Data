``Still Working on Not fixit``
# TempoScale: Distributed Infrastructure for Time-Series Data

## Description 
TempoScale is an innovative distributed infrastructure designed to efficiently store, manage, and query large volumes of time-series data. With the proliferation of IoT devices, financial transactions, and continuous logging systems, the demand for robust and scalable time-series data storage solutions has never been higher. TempoScale addresses this need by providing a high-performance, fault-tolerant, and scalable platform specifically optimized for time-series data.
## Define The business problems
TempoScale tackles key challenges in time-series data management. Traditional databases struggle with high data volumes and performance bottlenecks, which TempoScale solves with a scalable, distributed architecture. It enables fast data retrieval through advanced indexing and compression, ensuring real-time insights. Data integrity and reliability are maintained via replication and consistent storage mechanisms.

TempoScale simplifies complex infrastructure management with automated deployment using Docker and Kubernetes. Robust APIs and client libraries ensure seamless integration with existing systems. Efficient storage techniques reduce costs, while integration with Prometheus and Grafana enhances monitoring and performance visibility. Security is bolstered through encryption and access control, ensuring data protection and regulatory compliance. Automation minimizes operational inefficiencies, and customizable configurations allow tailored data management solutions.

In summary, TempoScale offers a high-performance, reliable, and scalable platform for managing large-scale time-series data, reducing operational complexity and costs while improving security and observability.
## Query data
```
-- query table data
SELECT citus_set_coordinator_host('citus_coordinator', 5432);
```
```
public ❯ -- query table data
SELECT * FROM air;
+---------------------+------------+------------+-------------+----------+
| time                | station    | visibility | temperature | pressure |
+---------------------+------------+------------+-------------+----------+
| 2023-01-11T06:40:00 | XiaoMaiDao | 55         | 68          | 76       |
| 2023-01-13T06:33:17 | XiaoMaiDao | 56         | 69          | 77       |
| 2023-01-11T07:40:00 | DaMaiDao   | 65         | 68          | 76       |
+---------------------+------------+------------+-------------+----------+
Query took 0.036 seconds.
```
``
Probably U come to run any example it will be show you like this 
Output 

![alt text](Screen%20Shot%202024-06-08%20at%2021.15.02.png)


## Setting Up

1. Clone the repository:
   ```sh
   git clone https://github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data.git
   cd TempoScale-Distributed-Infrastructure-for-Time-Series-Data
   ```

2. Start the development environment:
   ```sh
   docker-compose -p temposcale -f dev.docker-compose.yml up
   ```

3. Access the application shell:
   ```sh
   docker exec -it temposcale /bin/sh
   ```

4. Access the PostgreSQL console:
   ```sh
   docker exec -it temposcale_citus_coordinator psql -U golang -d temposcale_db
   ```

## Deployment

Build and deploy the Docker image:
```sh
docker build -f Dockerfile -t zacksfF/temposcale:latest --platform linux/amd64 .
docker tag zacksfF/temposcale:latest bartmika/temposcale:latest
docker push zacksfF/temposcale:latest
```

## Vulnerability Scanning

Perform a vulnerability scan on the Go project and Docker image:
```sh
govulncheck ./...
trivy image zacksfF/temposcale:latest
trivy repo github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data.git
```

## Updating Dependencies

Upgrade all dependencies in the project:
```sh
go get -u ./...
go mod tidy
```

## Usage

```
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/util/dtos"
	rpc_client "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/util/rpc"
)

// DESCRIPTION:
// The purpose of this application is to connect to a running `stockyard` server
// instance on your machine and create an entity.

func main() {
	
	// Sample data to use in our example code.
	ipAddress := "127.0.0.1"
	port := "8000"
	deviceName1 := "temperate-sensor-1" // Give your entity any unique name you like.
	deviceName2 := "backyard-birdfeeder-phototimer"

	// Connect to a running server from this appolication.
	applicationAddress := fmt.Sprintf("%s:%s", ipAddress, port)
	client, err := rpc_client.NewClient(applicationAddress, 3, 15*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the remote call.
	entity, err := client.InsertEntity(deviceName1, dtos.EntityObservationDataType, "")
	if err != nil {
		log.Fatal(err)
	}

	if entity == nil {
		log.Fatal("error as nothing was returned")
	}

	// See the results.
	log.Println("id", entity.ID)
	log.Println("uuid", entity.UUID)
	log.Println("name", entity.Name)
	log.Println("data_type", entity.DataType)
	log.Println("meta", entity.Meta)

	////
	//// Create a `phototimer` collection.
	////

	// Execute the remote call.
	entity, err = client.InsertEntity(deviceName2, dtos.EntityTimeKeyDataType, "")
	if err != nil {
		log.Fatal(err)
	}

	if entity == nil {
		log.Fatal("error as nothing was returned")
	}

	// See the results.
	log.Println("id", entity.ID)
	log.Println("uuid", entity.UUID)
	log.Println("name", entity.Name)
	log.Println("data_type", entity.DataType)
	log.Println("meta", entity.Meta)
}
```

## License

This project is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for details.
```

This README provides instructions for installing necessary tools, setting up the development environment, deploying the application, and performing vulnerability scans and dependency updates. Make sure to adjust any specific details according to your project requirements.
