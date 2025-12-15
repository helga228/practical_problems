package int_task

import "fmt"

func SimpleCalculations() {
	a := 15
	b := 4

	sum := a + b
	diff := a - b
	prod := a * b
	quot := a / b
	rem := a % b
	bitAnd := a & b
	fmt.Printf("Сумма: %d + %d = %d\n", a, b, sum)
	fmt.Printf("Разность: %d - %d = %d\n", a, b, diff)
	fmt.Printf("Произведение: %d * %d = %d\n", a, b, prod)
	fmt.Printf("Частное: %d / %d = %d\n", a, b, quot)
	fmt.Printf("Остаток: %d / %d = %d\n", a, b, rem)
	fmt.Printf("Побитовое И: %d & %d = %d\n \n", a, b, bitAnd)
}

func Overflow() {
	// В int8 входит диапазон от -128 до 127
	var x int8 = 127
	overInt := x + 1
	fmt.Printf("Переполнение: int8(%d) + 1 даст %d\n", x, overInt)

	// В uint8 входит диапазон от 0 до 255
	var y uint8 = 255
	overUint := y + 1
	fmt.Printf("Переполнение: uint8(%d) + 1 даст %d\n \n", y, overUint)
}

func Bitwise() {
	num := 10

	leftShift := num << 2
	fmt.Printf("Сдвиг %d на 2 бита влево даст %d\n", num, leftShift)

	rightShift := num >> 1
	fmt.Printf("Сдвиг %d на 1 бит вправо даст %d\n", num, rightShift)

	bitNot := ^num
	fmt.Printf("Инверсия %d даст %d\n \n", num, bitNot)
}

func TypeConversion() {
	bigNum := int64(1000)
	smallNum := int8(100)

	Int64ToInt32 := int32(bigNum)
	fmt.Printf("Преобразование int64(%d) в int32:  %d\n", bigNum, Int64ToInt32)

	Int8ToInt16 := int16(smallNum)
	fmt.Printf("Преобразование int8(%d) в int16:  %d\n", smallNum, Int8ToInt16)

	Int64ToInt8 := int8(bigNum)
	fmt.Printf("Преобразование int64(%d) в int8 (перед преобразованием стоит добавлять проверки для предотвращения таких случаев):  %d\n \n", bigNum, Int64ToInt8)
}

func ProcessNumberDigits() {
	num := 444
	originalNum := num

	var digits []int
	for num > 0 {
		digit := num % 10
		digits = append([]int{digit}, digits...)
		num = num / 10
	}
	sum := digits[0]
	for i := 1; i < len(digits); i++ {
		sum += digits[i]
	}

	product := digits[0]
	for i := 1; i < len(digits); i++ {
		product *= digits[i]
	}

	fmt.Printf("Сумма цифр числа %d равна: %d\n", originalNum, sum)
	fmt.Printf("Произведение цифр числа %d равна: %d\n \n", originalNum, product)
}

func RoundAndCompareFloat() {
	resultOne := (0.1 + 0.2) * 100.0 / 3.0 // выполнится во время компиляции

	numbers := []float64{0.1, 0.2, 100.0, 3.0}
	resultTwo := (numbers[0] + numbers[1]) * numbers[2] / numbers[3] // вычисления происходят во время выполнения

	fmt.Println("Результат вычисления  (0.1 + 0.2) * 100.0 / 3.0 без слайса:", resultOne)
	fmt.Println("Результат вычисления  (0.1 + 0.2) * 100.0 / 3.0 через слайс:", resultTwo)
}
