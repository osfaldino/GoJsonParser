ENGLISH

# This is a project for learning Go
## Here is a basic json parser

To parse json you need to create a config:

```go
var cfg Config
```
### We call the function and check it for errors:

```go
err := cfg.ParseConfig("cfg.json")
if err != nil {
  fmt.Printf("Error: %s", err)
  os.Exit(1)
}
```
### Printing the values:

```go
fmt.Printf("AppName: %s\nVersion: %s\n", cfg.AppName, cfg.Version)

for _, user := range cfg.Users {
  fmt.Println(user)
}
```
### JSON example:
```json
{
    "appName": "JsonLearn",
    "version": "1.0.0",
    "debug": true,
    "users": ["Alice", "Bob", "Charlie"]
}
```
---------------------
РУССКИЙ

# Это проект для обучения Go
## Здесь представлен базовый парсер json

Для парса json нужно создать конфиг:

```go
var cfg Config
```
### Вызываем функцию и проверяем ее на наличие ошибок:

```go
err := cfg.ParseConfig("cfg.json")
if err != nil {
  fmt.Printf("Ошибка: %s", err)
  os.Exit(1)
}
```
### Выводим значения:

```go
fmt.Printf("AppName: %s\nVersion: %s\n", cfg.AppName, cfg.Version)

for _, user := range cfg.Users {
  fmt.Println(user)
}
```
### Пример json:
```json
{
    "appName": "JsonLearn",
    "version": "1.0.0",
    "debug": true,
    "users": ["Alice", "Bob", "Charlie"]
}
```
