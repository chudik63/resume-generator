# Программа по генерации резюме
Программа генерирует резюме страницу, по вашим данным из Вконтакте. Получаемый файл отображается на локальном сервере.
Для запуска проследуйте инструкциям ниже.

## Получение токена Вконтакте
1. Переходим по [ссылке](https://oauth.vk.com/authorize?client_id=6178269&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=friends,notify,photos,wall,email,mail,groups,stats,offline&response_type=token&v=5.74)
2. Нажимаем **Продолжить как..**
3. В адресной строке копируем токен после **access_token=** до **&expires_in=...**

## Клонируем репозиторий
Необходимо самим создать .env файл и заполнить его, используя полученный токен и ваш 
id страниц (id должен представлять собой набор цифр).
```
git clone https://git.miem.hse.ru/biv23x-ps/vivinaumov.git
cd vivinaumov
touch .env
echo "method = "users.get"
     vk_id = "YOUR ID"
     token = "YOUR TOKEN" " >> .env
```

## Установка дополнительных библиотек
В файле **go.mod** присутствуют версии всех дополнительных библиотек.
При желании добавить внешнюю библиотеку используйте команду `go get`.
```
go get github.com/joho/godotenv
```

## Как запустить
```
go run main.go
```
Переходим на **localhost:8080**

