# Task3_v2
Требуется разработать консольную утилиту, которая принимает несколько диапазонов чисел, имя файла для вывода, таймаут ограничивающий исполнение команды по времени. 
Утилита находит в заданных дипазонах все простые числа и выводит их в файл. Например для диапазона 11:20 простые числа это 11, 13, 17, 19

Пример команды:

find_primes --file testfile.txt --timeout 10 --range  1:10 --range 200000:3000000 --range 400:500
file имя файла для вывода найденных простых чисел
timeout значение в секундах, по истечении которого программа должна прекратить свое исполнение
range диапазон чисел, в пределах которого программа должна найти простые числа
Замечания к реализации
Результатом должен быть текст программы на Go в одном файле main.go, для простоты проверки.

При разработке предусмотреть, чтобы каждый диапазон обрабатывался в отдельной гороутине, вывод в файл тоже производился в отдельной гороутине. 
Общение между гороутинами предусмотреть каналом/каналами.