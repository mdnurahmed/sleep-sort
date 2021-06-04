# sleep-sort 

This is concurrency problem. In this algorithm we create different threads for each of the elements in the input array and then each thread sleeps for an amount of time which is proportional to the value of corresponding array element.

Hence, the thread having the least amount of sleeping time wakes up first and the number gets printed and then the second least element and so on. The largest element wakes up after a long time and then the element gets printed at the last. Thus the output is a sorted one.

Go doesn't have thread . Instead it has goroutines which is like lightweight thread. Go scheduler maps the goroutines on the OS thread underneath.

As you can see this algorithm can only work for non-negative numbers . To make it work for negative numbers we gotta shift the values of the array so they all become positive , sort it and then shift the values back to get the actual array .   

sleep_sort package is implementation of Sleep Sort algorithm in Golang.

# Complexity 

O(n) time complexity and O(1) space complexity . The code will take m unit of time where m is the largest value in the array after shifting . 

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