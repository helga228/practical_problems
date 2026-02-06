# Switch

```
    switch x {
    case 1:
        // ...
    case 2, 3:
        // ...
    default:
        // ...
    }
    
    // switch без выражения (как if-else)
    switch {
    case x > 0:
        // ...
    case x < 0:
        // ...
    }
    
    // type switch
    switch v := i.(type) {
    case int:
        // ...
    case string:
        // ...
    }
```

## Как это работает?