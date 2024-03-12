## Riddle Solver

Test the server with

```sh
curl -X 'GET' \
  'http://localhost:8084/ping' \
  -H 'accept: application/json'
```

A Big riddle

```sh
curl -X 'POST' \
'http://localhost:8084/solve' \
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
'http://localhost:8084/solve' \
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
  'http://localhost:8084/solution' \
  -H 'accept: application/json'
```