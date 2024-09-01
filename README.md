# GO TELEGRAM BOT


## Запуск через docker
```commandline
sudo docker build -t bucket-app .
sudo docker run -p 8080:8080  bucket-app
```



## Запуск на локальной машине

### 1. запустить ngrok
```commandline
ngrok http 8080
```

### 2. запустить приложение
```commandline
go run main.go
```

### 3. запустить прокси (только для windows)
```commandline
netsh interface portproxy add v4tov4 listenport=8081 listenaddress=0.0.0.0 connectport=8080 connectaddress=127.0.0.1
```
