# Mafia Game

Server supports multiple parallel game sessions, private chat for mafiosi during the night. Clients use redis as a message broker.

It's not for production. Security is something to think about here. Clients with redis password has too much control. [This](https://redis.io/docs/management/security/acl/) should be used to make it more secure.

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



