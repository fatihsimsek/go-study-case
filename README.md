# Go Study Case

In memory key-value store 

## Getting Started

* GET /get?key={key}: get value of this key
* POST /set { "key": "x", "value": "123"} : set key-value pair to in-memory
* POST /flush: flush in-memory to file
