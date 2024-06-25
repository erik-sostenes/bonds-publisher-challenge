# Bonds Publisher API - Challenge

REST API that allows users to publish bonds for sale and buy bonds published by other users.

###### Docs

- [API documentation (Swagger)](docs/swagger.yaml)

#### Get token banxico

[banxico](https://www.banxico.org.mx/SieAPIRest/service/v1/token)

### how to execute the project?

To get started, follow the instructions on how to run the REST API

###### recomendations (optional)

Verify that you do not have any docker image running

```bash
docker stop $(docker ps -q)
```

##### Build app

```bash
docker-compose up --build
```

##### Generate the public and private RSA keys

```bash
make keys
```

##### Load private rsa key

```bash
 export PRIVATE_KEY=$(base64 -w 0 private_key.pem)
```

##### Load public rsa key

```bash
 export PUBLIC_KEY=$(base64 -w 0 public_key.pem)
```

##### Load environment variables

```bash
export $(grep -v ^# .env.example)
```

##### Run it!

```bash
make run
```

### how run test?

Follow the instructions below to run the tests

##### Generate the public and private RSA keys

```bash
make keys
```

##### Load private rsa key

```bash
 export PRIVATE_KEY=$(base64 -w 0 private_key.pem)
```

##### Load public rsa key

```bash
 export PRIVATE_KEY=$(base64 -w 0 private_key.pem)
```

##### Load environment variables

```bash
export $(grep -v ^# .env.example)
```

##### Migrate

```bash
make migrate
```

##### Run it!

```bash
make test
```

###### Unit tests

#### Structure

```go
package bond

import (
	"strconv"
	"testing"
)

func TestX(t *testing.T) {
	// Table Driven Testing
	tdt := map[string]struct{}{
	    {},	// Test Case
    	}{}

	// Setup
	x := struct{}{}

	// Cleanup function (Optional)
	t.Cleanup(func() {})

	// Subtests
	for name, tsc := range tdt {
		t.Run(name, func(t *testing.T){
			t.Log(tsc)
			// Cleanup function for each subtest (Optional)
			t.Cleanup(func() {})
		})
	}
}
```

### Hexagonal Architecture

The hexagonal architecture is based on three principles and techniques:

1. Explicitly separate User-Side, Business Logic, and Server-Side.
2. Dependencies are going from User-Side and Server-Side to the Business Logic.
3. We isolate the boundaries by using Ports and Adapters.

```
├───cmd 						👉🏼 (execute commands)
│   ├───bootstrap               👉🏼 (bootstrap package that builds the program with its full set of components)
│   └───http					👉🏼 (the directory name for each application should match the name of the executable)
├───internal            		👉🏼 (private packages)
│   ├───app                     👉🏼 (core business)
│   │   ├───(module)            👉🏼 (represents a boundary)
│   │   │   ├───business        👉🏼 (business logic layer)
│   │   │   │   ├───domain      👉🏼 (data transfer objects, business objects, errors, entities, value objects)
│   │   │   │   ├───ports       👉🏼 (business contracts)
│   │   │   │   └───services    👉🏼 (business logic)
│   │   │   └───infrastructure  👉🏼 (layer infrastructure)
│   │   │       ├───driven      👉🏼 (output adapters)
│   │   │       └───drives      👉🏼 (input adapters)
├───pkg                         👉🏼 (explicitly communicate that the code in that directory is safe for use by others -> global code or libraries)
```

### User-Side

- This is the side through which the user or external programs will interact with the application.

### Business Logic

- This is the part that we want to isolate from both left and right sides. It contains all the code that concerns and implements business logic.

### Server-Side

- This is where we’ll find what your application needs, what it drives to work. It contains essential infrastructure details such as the code that interacts with your database, makes calls to the file system, or code that handles HTTP calls to other applications on which you depend for example.

```

|------------|                   |----------------|             |-------------|
| User side  | =====[port]=====> | Business logic | <==[port]== | server side |
|------------|                   |----------------|             |-------------|

```

#### References

- [The Pattern: Ports and Adapters (‘’Object Structural’’) ](https://alistair.cockburn.us/hexagonal-architecture/)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Hexagonal Architecture: three principles and an implementation example](https://blog.octo.com/hexagonal-architecture-three-principles-and-an-implementation-example)
