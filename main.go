package main

import (
    "fmt"
    "regexp"
)

const (
    maxNumber = 10
)

var (
    romanNumerals = map[int]string{
        1: "I",
        2: "II",
        3: "III",
        4: "IV",
        5: "V",
        6: "VI",
        7: "VII",
        8: "VIII",
        9: "IX",
        10: "X",
    }
)

func isNumber(number string) bool {
    if number == "" {
        return false
    }

    if number[0] == '0' {
        return false
    }

    return number[0] >= '1' && number[0] <= '9'
}

func isOperation(operation string) bool {
    return operation == "+" || operation == "-" || operation == "*" || operation == "/"
}

func convertRoman(number int) string {
    return romanNumerals[number]
}

func calculate(a, b int, operation string) int {
    switch operation {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        return a / b
    }

    return 0
}

func main() {
    fmt.Println("Введите выражение:")

    expression := ""
    for {
        expression = fmt.Scanln()
        if expression == "" {
            break
        }

        numbers := regexp.MustCompile(`^\d+|\w+`).FindAllString(expression, -1)
        if len(numbers) != 3 {
            fmt.Println("Неверная строка")
            continue
        }

        a, b, operation := numbers[0], numbers[1], numbers[2]
        if !isNumber(a) || !isNumber(b) || !isOperation(operation) {
            fmt.Println("Неверные числа или операция")
            continue
        }

        result := calculate(int(a), int(b), operation)
        if operation == "/" && result < 1 {
            fmt.Println("Результат меньше единицы")
            continue
        }

        fmt.Println(result)
    }
}
