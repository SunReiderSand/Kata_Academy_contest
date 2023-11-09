# Kata_Academy_contest
**Калькулятор**

**Описание**

Калькулятор выполняет операции сложения, вычитания, умножения и деления над двумя числами.

**Требования**

* Числа должны быть от 1 до 10 включительно.
* Калькулятор должен работать с арабскими и римскими цифрами.
* Калькулятор должен выводить результат в том же формате, в котором были введены числа.
* При вводе некорректного выражения или операции калькулятор должен выводить сообщение об ошибке.

**Использование**

Для выполнения операции введите выражение в одну строку. Например:

```
1 + 2
```

Вывод:

```
3
```

**Алгоритм**

Калькулятор работает следующим образом:

1. Ввод выражения от пользователя.
2. Проверка корректности выражения.
3. Преобразование чисел в нужный формат.
4. Выполнение расчета.
5. Вывод результата.

**Функции**

* `isNumber()` - проверяет, является ли строка допустимым числом.
* `isOperation()` - проверяет, является ли строка допустимой операцией.
* `convertRoman()` - преобразует арабское число в римское.
* `calculate()` - выполняет расчет.

**Пример работы**

```
Введите выражение:
1 + 2

3
```

```
Введите выражение:
III + IV

VII
```

```
Введите выражение:
10 - 2

8
```

```
Введите выражение:
5 * 6

30
```

```
Введите выражение:
10 / 2

5
```

```
Введите выражение:
10 / 0

Деление на ноль
```

```
Введите выражение:
1 + II

Неверная строка
```