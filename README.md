# Mafia Game

Server supports multiple parallel game sessions, private chat for mafiosi during the night. Clients use redis as a message broker.

## Run Server

```
$ docker compose up
```

## Run Client

```
$ cd client
$ go run . --username mafioso123
```

Type `/help` to show available commands.

## Run Bot

```
$ cd client
$ go run . --username mafioso412
```



