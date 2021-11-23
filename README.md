# surge

## Run Docker image
### Navigate to the root of the project and build Docker image as a superuser
```shell
docker image build -t surge .
```
### Run the built Docker image
```shell
docker container run -p 8080:8080 surge
```


## Run PostgreSQL using Docker
### navigate to the database folder and run the following command
```shell
docker-compose up
```
