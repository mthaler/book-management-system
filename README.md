# book-managment-system

Simple book managment system using [fiber](https://github.com/gofiber/fiber), [gorm](https://gorm.io/index.html) and sqlite from [this tutorial](https://tutorialedge.net/golang/basic-rest-api-go-fiber/).

The application provides a REST API to manage books.

## Usage

Create book:

```bash
$ curl -X POST -H "Content-Type: application/json" --data "{\"title\": \"Angels and Demons\", \"author\": \"Dan Brown\", \"rating\": 4}" http://localhost:3000/api/v1/book
```

List all books:

```bash
$ curl http://localhost:3000/api/v1/book
```

List book with a given ID:

```bash
$ curl http://localhost:3000/api/v1/book/1
```

Update book:

```bash
$ curl -X POST -H "Content-Type: application/json" --data "{\"title\": \"Angels and Demons\", \"author\": \"Dan Brown\", \"rating\": 5}" http://localhost:3000/api/v1/book/1
```

Delete book with a given ID:

```bash
$ curl -X DELETE http://localhost:3000/api/v1/book/1
```

## TODO

Implement update

