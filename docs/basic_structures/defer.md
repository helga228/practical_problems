# Отложенные вызовы (defer)

```
    defer file.Close()
    defer mu.Unlock()
    defer func() {
        // восстановление после паники
        if r := recover(); r != nil {
            log.Println("Recovered:", r)
        }
    }()
```

## Как это работает?