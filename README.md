# sleep-sort 

sleep_sort package is implementation of Sleep Sort algorithm in Golang.

# How to run
```
go test ./... 
go run main.go -2 1 3 0 
```
or using docker 
```
docker-compose -f sleep-sort.yaml up --build
docker run -it sleep-sort_sleep-sort /bin/sh
go test ./... 
go run main.go -2 1 3 0
exit
docker-compose -f sleep-sort.yaml down
```