# concurrency01
Golang Basics course Concurrency - 02

### Задание

- Даны два канала. В первый пишутся числа.
- Нужно, чтобы числа читались из первого по мере поступления, что-то с ними происходило (допустим, возводились в квадрат) и результат записывался во второй канал. 

### Логика работы задачи

- Запускаем две горутины.
- В одной пишем в первый канал.
- Во второй читаем из первого канала и пишем во второй.

- Главное — не забыть закрыть каналы, чтобы ничего нигде не заблокировалось.
- После этого нужно вывести в консоль содержимое второго канала.

