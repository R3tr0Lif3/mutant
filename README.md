# mutant
_Challenge :  Get a string of DNA and validate if it is mutant or human_
## Starting 🚀

_Get a copy of the project on your local machine_

### pre-requirements 📋

### Install 🔧

* Install go version >=1.14 [Guide Install GO ](https://golang.org/doc/install)
* Verify go installation with command go version
```
go version 

```
_the following text should appear_
```
go1.14.4 windows/amd64

```
* Install redis [Guide Install Redis ](https://redis.io/topics/quickstart)
* Verify redis installation with command redis-server --version
```
redis-server --version
```
_The following text should appear_
```
Redis server v=5.0.7 sha=00000000:0 malloc=jemalloc-5.2.1 bits=64...

```
_Verify your redis with the command redis-cli_
```
redis-cli

```

_The following text should appear_
```
127.0.0.1:6379> 

```
_By default server connection is 127.0.0.1:6379 you can change it in the file db/redis.go_

```go
const server = "localhost:6379"
```

_You can run the project with the command go run main.go_

```
go run main.go

```
_Default Configuration: localhost:8080_

## Code Coverage 🛠️

_You can execute test with the command go test -v ./app/... -cover_
```
go test -v ./app/... -cover

```
_The following text should appear_

```console
=== RUN   TestGoLeftRight
--- PASS: TestGoLeftRight (0.00s)
=== RUN   TestGoTopButton
--- PASS: TestGoTopButton (0.00s)
=== RUN   TestGoAB
--- PASS: TestGoAB (0.00s)
=== RUN   TestGoBA
--- PASS: TestGoBA (0.00s)
=== RUN   TestGoXY
--- PASS: TestGoXY (0.00s)
=== RUN   TestGoYX
--- PASS: TestGoYX (0.00s)
=== RUN   TestIsMutantLogic
--- PASS: TestIsMutantLogic (0.01s)
PASS
coverage: 94.4% of statements
```

## Running Test in cloud server ⚙️

### End point /mutant

http://ec2-54-167-27-243.compute-1.amazonaws.com:8080/mutant

### End point /stats

http://ec2-54-167-27-243.compute-1.amazonaws.com:8080/stats

## Autor ✒️

### ❤️ By [Diego Guerrero @R3tr0lif3](https://github.com/R3tr0Lif3)


