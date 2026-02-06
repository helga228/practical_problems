# Специальные конструкции

```
// Паника и восстановление
panic("something went wrong")
recover()

// Пустой идентификатор _
_, err := doSomething()

// Инициализация пакета
func init() {
    // выполняется при импорте пакета
}

// Теги структур
type User struct {
    Name string `json:"name" db:"user_name"`
}
```

## Как это работает?