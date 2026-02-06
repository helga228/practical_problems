# Работа с горутинами и каналами

```
    // Горутины
    go func() {
        // параллельное выполнение
    }()
    
    // Каналы
    ch := make(chan int)
    go func() { ch <- 42 }()
    value := <-ch
    
    // Select для работы с несколькими каналами
    select {
    case msg := <-ch1:
        // ...
    case msg := <-ch2:
        // ...
    case <-time.After(time.Second):
        // таймаут
    }
```

## Как это работает?