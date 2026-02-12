# Объявление структур

```
type person struct {
    Name string
    Age  int
}
```

## Как это работает?

Структура в го это составной тип данных, содержание которого определяет разработчик.
Для объявления структуры используются ключевые слова type и struct.
Каждое поле структуры имеет название и тип данных. 

Для обращения к полю структуры после ее имени ставится точка и указывается поле:

```
personName := person.name
```

Можно объявлять анонимные структуры

```
var person struct {
    Name string
    Age  int
}
```

Но только в том случае если нам нужен только один экземпляр, можно использовать для константных значений, например:

```
var DefaultConfig = struct {
    Host    string
    Port    int
    Timeout int
}{
    Host:    "localhost",
    Port:    8080,
    Timeout: 30,
}
```
Так-же есть возможность объявлять анонимные поля структуры (неочевидная работа, я бы избегала такого использования)

```
type person struct{
    string
    int
}
```

Можно создавать указатели на структуры (хорошо использовать когда в структуре много данных)

```
func increment_age(user *person){
    user.age += 1
    fmt.Println(*user)   // {Tom 42}
}
 
func main() {
      
    tom := person {"Tom",  41,}
    // в функцию передается адрес структуры tom
    increment_age(&tom)
     
    // изменение структуры внутри функции по указателю приведет к изменению оригинальной структуры
    fmt.Println(tom)    // {Tom 42}
}
```

Структуры можно использовать как поля других структур

```
type person struct{ 
    name string
    age int
}
 
type account struct{
    login string
    password string
    person
}
```

Из интересного:

Структуры можно сравнивать

```
p1 := Person{"Alice", 30}
p2 := Person{"Alice", 30}
fmt.Println(p1 == p2) // true!
```

Пустая структура не занимает память TODO - дописать где это может пригодиться

```
empty := struct{}{}
fmt.Println(unsafe.Sizeof(empty)) // 0
```

В структурах можно использовать теги, очень полезная штука, можно валидировать данные

```
type Validation struct {
    Email string `validate:"email,required"`
    Age   int    `validate:"min=18,max=100"`
}
```

или использовать альтернативные имена для других структур данных

```
type User struct {
Name string `json:"name" xml:"fullname" db:"user_name"`
// Теги доступны через reflect
}
```
