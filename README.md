# mustakuusi-go-gin
for learning purposes

## Start Server
```bash
go run main.go 3004
```

## Get Games
```text
curl -X 'GET' \
  'http://localhost:3004/games' \
  -H 'accept: */*'
```

## Get Characters
```text
curl -X 'GET' \
  'http://localhost:3004/characters' \
  -H 'accept: */*'
```
