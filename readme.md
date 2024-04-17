## Riddle Solver

### My Task (en)

Task 26: Develop an algorithm for solving the problem and implement it as a program in the ANSI C++ language. Color some
cells so that there are no repeating numbers in each row or column.
Painted cells should not touch each other. All unpainted cells must connect to each other horizontally or vertically, so
so that a single continuous space of unpainted cells emerges.

Your task. Painted cells can collide with each other. At least one unpainted cell must remain in each row and column.

![task.png](readme_images%2Ftask.png)

An example and its solution:

![example.png](readme_images%2Fexample.png)

### My Task (ua)

Завдання 26: Розробіть алгоритм вирішення задачі та реалізуйте його у вигляді програми мовою ANSI C++. Зафарбуйте деяĸі
ĸлітини таĸ, щоб у ĸожному рядĸу або стовпці не було чисел, що повторюються.
Зафарбовані ĸлітини не повинні стиĸатися одна з одною. Усі незафарбовані ĸлітини повинні з'єднуватися одна з одною
сторонами по горизонталі або по вертиĸалі таĸ,
щоб вийшов єдиний безперервний простір із незафарбованих ĸлітин.

Ваше завдання. Зафарбовані ĸлітини можуть стиĸатися одна з одною. У ĸожному рядĸу та ĸожному стовпці повинна залишатися
принаймні одна незафарбована ĸлітина.

![task.png](readme_images%2Ftask.png)

Приĸлад та його вирішення:

![example.png](readme_images%2Fexample.png)

### How to run

``` sh
docker-compose up --build -d
```

docs: http://localhost:8081/docs/index.html

### Curl client

In order to use the server you need to get a token
and set it in the header of the request.
Example of header with token
```sh
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjMzYTkyMzA3LTNiM2QtNDQwNy1hZGI1LWJiMTNkNmYzNjkxYSIsImVtYWlsIjoidXNlckBleGFtcGxlLmNvbSIsImlzc3VlZF9hdCI6IjIwMjQtMDQtMTZUMTc6NTY6MzAuNDI1NDAwNyswMzowMCIsImV4cGlyZWRfYXQiOiIyMDI0LTA0LTE2VDE5OjM2OjMwLjQyNTQwMDcrMDM6MDAifQ.qG_dSc64hrfgvD4JFqtZ4Urr20SDEA3MmbcMjrqHVo0'
```

To get a token you need to register a user
```sh
curl -X 'POST' \
  'http://localhost:8081/auth/create' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "user@example.com",
  "password": "string"
}'
```

And Then login to get a token
```sh
curl -X 'POST' \
  'http://localhost:8081/auth/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "user@example.com",
  "password": "string"
}'
```


Test the server with (**Don`t forget about token**)

```sh
curl -X 'GET' \
  'http://localhost:8081/ping' \
  -H 'accept: application/json'
```

A Big riddle

```sh
curl -X 'POST' \
'http://localhost:8081/solve' \
-H 'accept: application/json' \
-H 'Content-Type: application/json' \
-d '[
[1, 1, 4, 3, 4, 1, 3, 2, 2],
[1, 1, 2, 3, 2, 1, 3, 2, 2],
[3, 2, 1, 4, 3, 3, 2, 1, 3],
[4, 3, 4, 2, 3, 1, 1, 2, 4],
[4, 2, 1, 1, 2, 3, 3, 4, 1],
[2, 2, 3, 3, 4, 4, 4, 1, 2],
[2, 3, 3, 1, 3, 2, 2, 4, 1],
[4, 4, 2, 1, 3, 1, 2, 3, 3],
[4, 4, 2, 1, 1, 1, 2, 3, 3]
]'
```

A small riddle

```sh
curl -X 'POST' \
'http://localhost:8081/solve' \
-H 'accept: application/json' \
-H 'Content-Type: application/json' \
-d '[
[4, 2, 4, 8],
[8, 6, 6, 8],
[4, 2, 6, 6],
[2, 2, 6, 6]
]'
```

Get the next solution

```sh
curl -X 'GET' \
  'http://localhost:8081/solution' \
  -H 'accept: application/json'
```