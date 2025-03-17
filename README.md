# Go Client for Stack Auth API

[![Go Reference](https://pkg.go.dev/badge/github.com/BlaisePopov/stack-auth.svg)](https://pkg.go.dev/github.com/BlaisePopov/stack-auth)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/BlaisePopov/stack-auth/blob/main/LICENSE)

Go-клиент для взаимодействия с [Stack Auth](https://stack-auth.com) — открытым набором инструментов для аутентификации и управления пользователями. Клиент поддерживает все методы API, включая работу с пользователями, сессиями и правами доступа.

## Установка

Для установки пакета выполните:

```bash
go get github.com/BlaisePopov/stack-auth
```

## Примеры использования

### Инициализация клиента

```go
package main

import (
    "github.com/BlaisePopov/stack-auth/api"
    "github.com/BlaisePopov/stack-auth/base_http_client"
)

func main() {
    stackAuth := api.NewClient(base_http_client.Config{
        ProjectID:       "your_project_id",
        SecretServerKey: "your_secret_server_key",
    })
}
```

### Получение информации о пользователе

```go
user, err := stackAuth.Users.GetUser("f2c50a5c-84ff-4076-8c24-0a536db98bcd")
if err != nil {
    panic(err)
}
println(user.DisplayName)
```

## Документация API
Официальная документация Stack Auth:  
[https://docs.stack-auth.com/next/rest-api/server/api-v-1](https://docs.stack-auth.com/next/rest-api/server/api-v-1)

## Лицензия
Этот проект распространяется под лицензией MIT. Подробности см. в файле [LICENSE](LICENSE).