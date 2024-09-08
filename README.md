# BUCKET TELEGRAM BOT
Telegram bot for saving notes
special libraries to work with telegramAPI are not used 
gets updates with webhooks

## Stack
- Go
- Docker compose
- PostgreSQL with [pgx/v5](https://github.com/jackc/pgx)
- PgBouncer


## local run
### create .env file
```comandline
make env
```

### launch ngrok
#### linux
```commandline
ngrok http 8080
```

#### windows
```commandline
ngrok http 8081
```
```commandline
netsh interface portproxy add v4tov4 listenport=8081 listenaddress=0.0.0.0 connectport=8080 connectaddress=127.0.0.1
```

### run app
```commandline
make run
```

### shutdown
```commandline
make off
```