# Это проект для обучения Go
## Здесь представлен базовый парсер json

Для парса json нужно создать конфиг:

```go
var cfg Config
```
Вызываем функцию и проверяем ее на наличие ошибок:

```go
err := cfg.ParseConfig("cfg.json")
if err != nil {
  fmt.Printf("Ошибка: %s", err)
  os.Exit(1)
}
```
Выводим значения:

```go
fmt.Printf("AppName: %s\nVersion: %s\n", cfg.AppName, cfg.Version)

for _, user := range cfg.Users {
  fmt.Println(user)
}
```
