# Программа по генерации резюме
Программа генерирует резюме страницу, по вашим данным из Вконтакте. Получаемый файл отображается на локальном сервере.

## Получение токена Вконтакте
1. Переходим по [ссылке](https://oauth.vk.com/authorize?client_id=6178269&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=friends,notify,photos,wall,email,mail,groups,stats,offline&response_type=token&v=5.74)
2. Нажимаем **Продолжить как..**
3. В адресной строке копируем токен после **access_token=** до **&expires_in=...**
## Заполнение .env файла
Вставляем ваш id Вконтакте и полученный токен в ковычки (id должен представлять собой набор цифр)
```no-highlight
method = "users.get"
vk_id = ""
token = ""
```
## Установка дополнительных библиотек
В файле **go.mod** присутствуют версии всех дополнительных библиотек.
При желании добавить внешнюю библиотеку используйте команду **go get**.
1. `go get github.com/joho/godotenv`

## Как запустить
1. `go run main.go`

