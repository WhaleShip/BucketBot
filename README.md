# BUCKET TELEGRAM BOT
Telegram bot for saving notes
special libraries to work with telegramAPI are not used 
gets updates with webhooks

## Stack
- Go
- Docker compose
- PostgreSQL with [pgx/v5](https://github.com/jackc/pgx)
- PgBouncer


## Local run
### Create .env file
```bash
make env
```

### Launch ngrok
#### Linux
```bash
ngrok http 8080
```

### Write your ngrok address and bot's token into [config](./config/app_config) file

#### Windows
```bash
ngrok http 8081
```
```bash
netsh interface portproxy add v4tov4 listenport=8081 listenaddress=0.0.0.0 connectport=8080 connectaddress=127.0.0.1
```

### Run app
```bash
make run
```

### Shutdown
```bash
make off
```