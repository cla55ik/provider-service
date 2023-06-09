# Provider Service

Итоговая аттестация курса "Go-разработчик" Skillbox

## Сборка проекта

`go build web/main.go`


## Конфигурация

Файл конфигурации `/config/config.yaml`

```yaml
# Настройки сервера
server:
  addr: "localhost"
  port: "8282"

# Настройки симулятора
simulator:
  addr: "localhost"
  port: "8787"

# Папка для файлов с данными
data:
  path: "data/"
```

#### Важно!

Для корректного отображения данных на странице `index.html` необходимо, чтобы номер порта в настройках server и номер порта в файле `main.js` совпадали

```js
let apiPath = 'http://localhost:8282';
```

## Запуск сервиса

Симулятор интегрирован в проект. Выбор режима запуска сервиса опеделяется флагом -mode.

Для корректной работы необходимо запустить симулятор и основной сервер.

Запуск симулятора

`./main -mode="simulator"`

Запуск сервера

`./main -mode="server"`

### Просмотр данных

Открыть файл `/web/index.html`

### Пример запроса на сервер

`localhost:8282`

Сервер отдает данные в формате JSON для отображения на странице `index.html`

Структура ответа

```json
{
"status":true,
"data":{Данные для отображения},
"error":""
}
```
