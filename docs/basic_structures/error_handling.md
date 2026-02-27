# Обработка ошибок

```
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}
    
func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    fmt.Println("Результат:", result)
}
```

## Как это работает?

 Для работы с ошибками в go есть специальный тип данных "error". Рассмотрим принцип работы на примере выше. У нас есть 
 функция, которая на вход принимает два числа, а на выход - число и ошибку. По какой причине функция которая делит одно 
 число на другое может вернуть ошибку? Например - если делитель - 0. Такая ситуация возможна и ее стоит сразу обработать, 
 иначе работа завершится с паникой. 
 Для того чтобы обработать ошибку добавляем проверку:

 ```
     if b == 0 {
        return 0, errors.New("деление на ноль")
    }
```

Первое возвращаемое значение - 0, поскольку деление не удалось, но мы должны вернуть какое-то значение, а 0 - это дефолтное
значение для float64. Вторым параметром возвращаем ошибку. Внимание! - По тексту ошибки должно быть понятно - в чем ее причина.

 ```
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
```

а в том случае, если переменная не ноль, мы можем вернуть результат деления и "nil"

 ```
    return a / b, nil
```

В дальнейшем, при вызове этой функции мы можем легко проверить, завершилась ли она ошибкой, и в случае если да - обработать ее нужным образом:

 ```
   result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    fmt.Println("Результат:", result)
```

Интерфейс error

Есть спец интерфейс, который содержит один метод Error с возвращаемым типом string. Значение по умолчанию nil

 ```
type error interface{
 
    Error() string
}
```

Любой тип, реализующий интерфейс error, может представлять тип ошибки.

Допустим, у нас есть принятый формат возвращаемых ошибок:

 ```
{
    "error_type": "",
    "error_text": "",
    "error_info": ""
}
```

В таком случае для начала реализуем структуру:

 ```
type CustomError struct {
	ErrorType string `json:"error_type"`
	ErrorText string `json:"error_text"`
	ErrorInfo string `json:"error_info"`
}
```

Затем нужна реализация интерфейса Error() и конструктор ошибки:

 ```
func (e CustomError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func NewCustomError(t, text, info string) error {
	return CustomError{
		ErrorType: t,
		ErrorText: text,
		ErrorInfo: info,
	}
}
```

Сама функция которая может вернуть ошибку и ее вызов:


 ```
func GetCustomError(isError bool) error {
	if isError {
		return NewCustomError(
			"VALIDATION_ERROR",
			"invalid input",
			"field email is required",
		)
	}
	return nil
}

func main() {
	err := GetCustomError(true)

	if err != nil {
		fmt.Println(err) // выведет JSON
	}
}
```

И вот такой ответ получим:

 ```
{"error_type":"VALIDATION_ERROR","error_text":"invalid input","error_info":"field email is required"}

```