# GO TELEGRAM BOT

## Запуск на windows


### 1. запустить ngrok
    ```
    ngrok http 8080
    ```

### 2. запустить приложение
    ``` 
    go run main.go
    ```

### 3. запустить прокси
    ```
    netsh interface portproxy add v4tov4 listenport=8081 listenaddress=0.0.0.0 connectport=8080 connectaddress=127.0.0.1
    ```