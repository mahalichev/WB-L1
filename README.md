## 1. Какой самый эффективный способ конкатенации строк?
Так как строки в Go представляют собой неизменяемую последовательность байтов, использование оператора "+" при конкатенации каждый раз создает новую строку, что является неэффективным методом.

Для обеспечения наибольшей эффективности используется strings.Builder, который имеет собственный буфер для записи строк, вследствие чего при конкатенации строка создается только 1 раз:
```go
builder := strings.Builder{}
builder.Grow(10)
builder.WriteString("abc")
builder.WriteString("def")
builder.WriteString("ghij")
str := builder.String()
```
Если строки для конкатенации находятся в slice, пакет strings предоставляет функцию Join(), которая также использует strings.Builder для конкатенации строк:
```go
strs := []string{"abc", "def", "ghij"}
str := strings.Join(strs, "")
```
## 2. Что такое интерфейсы, как они применяются в Go?
### Что такое интерфейс
Интерфейс - абстрактный тип данных, представляющий собой контракт, описывающий методы, которыми должен обладать объект, чтобы удовлетворял данному интерфейсу:
```go
// Интерфейсу Reader удовлетворяет любой объект, реализующий метод Square() int
type Squarer interface {
    Square() int
}
```
Go использует "Duck typing", что означает определение факта реализации определённого интерфейса объектом без явного указания или наследования этого интерфейса,
а по реализации полного набора его методов.
### Использование интерфейса в качестве типа параметра функции
Интерфейс можно указать в качестве типа параметра функции. Тогда функция будет принимать любой аргумент, реализующий интерфейс:
```go
type Number int

func (number *Number) Square() int {
    return int(number * number)
}

func PrintSquare(number Squarer) {
    fmt.Println(number.Square())
}

func main() {
    number := Number(5)
    PrintSquare(number) // stdout: 25
}
```
### Пустой интерфейс
Пустой интерфейс interface{} не имеет методов. Переменная с типом interface{} может содержать значение любого типа:
```go
var a interface{} = 5
```
### Приведение типа и переключатель типов
iface является корневым типом, представляющий интерфейс. Его определение выглядит следующим образом:
```go
type iface struct {
    tab  *itab
    data unsafe.Pointer
}

```
Тип данных itab описывает как интерфейс, так и тип данных, на который он указывает. Благодаря этому интерфейс можно привести к тому типу данных, на которые он указывает. Для этогу существует такой механизм, как приведение типа.

Приведение типа позволяет получить внутреннее значение интерфейса:
```go
var a interface{} = 5
if val, ok := a.(int); ok {
    fmt.Println(val * 2) // stdout: 10
}
```
Для приведения типа существует переключатель типов:
```go
switch value := unknown.(type) {
    case string:
        fmt.Println("This is string:", value)
    case int:
        fmt.Println("This is integer:", value)
    default:
        fmt.Println("Can't detect type")
}
```
### Интерфейсы и параметры типов
При использовании дженериков, для ограничения типов данных параметра типа можно указать  интерфейс. В данном интерфейсе определены типы данных, которые может принимать параметр типа. Одним из таких интерфейсов является cmp.Ordered:
```go
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string
}

func IsGreater[T Ordered](a, b T) bool {
    return a > b
}

func main() {
	fmt.Println(IsGreater(5, 3)) // stdout: true
	fmt.Println(IsGreater(10.3, 12.5)) // stdout: false
}
```

## 3. Чем отличаются RWMutex от Mutex?
**Mutex** является обычным блокировщиком, который разрешает работать с критическим местом только одной горутине.
```go
mutex.Lock()
counter++
mutex.Unlock()
```
**RWMutex** основан на Mutex и предоставляет дополнительные возможности работы с критическим местом, такие как RLock() и RUnlock(). Также как и с Mutex, только одна горутина может выполнять изменения в критическом месте (функции Lock() и Unlock()). Однако, с использованием RWMutex может быть несколько горутин-считывателей. Пока все считыватели не разблокируют RWMutex, блокировка для записи будет невозможна.
```go
type Counter struct {
	counter int
	rwMut   sync.RWMutex
}

func (counter *Counter) PrintSquare() {
	counter.rwMut.RLock()
	fmt.Println(counter.counter * counter.counter)
	counter.rwMut.RUnlock()
}

func (counter *Counter) Write() {
	counter.rwMut.Lock()
	counter.counter++
	counter.rwMut.Unlock()
}
```
## 4. Чем отличаются буферизированные и небуферизированные каналы?
Небуферизированный канал не имеет буфера. Поведение горутин при использовании небуферизированного канала следующее:
* При отправке данных в канал горутина блокируется, пока не будут получены данные другой горутиной;
* При получении данных горутина блокируется, пока другая горутина не отправит данные в канал.

Создание небуферизированного канала:
```go
ch := make(chan int)
ch := make(chan int, 0)
```
Буферизированный канал имеет буфер, ёмкость которого больше 0. Поведение горутин при использовании буферизированного канала следующее:
* Отправка данных в канал выполняется до тех пор, пока в буфере есть место. При попытке отправить данные в заполненный канал, горутина блокируется, пока данные не будут получены другой горутиной;
* Получение данных из канала выполняется до тех пор, пока буфер не пуст. При попытке получить данные из пустого буфера, горутина блокируется, пока в канал не будут отправлены данные другой горутиной.

Создание буферизированного канала:
```go
ch := make(chan int, 3)
```
## 5. Какой размер у структуры struct{}{}?
Тип данных struct{} называется пустой структурой. Данная структура не имеет полей, поэтому struct{}{} имеет размер равный 0. Все указатели на неё указывают на одно и то же место в памяти. struct{} широко
используется, например, при создания множеств (map[int]struct{} - множество значений типа int), а также в каналах для уведомления горутин.
## 6. Есть ли в Go перегрузка методов или операторов?
Go не предоставляет возможности перегрузки методов и операторов. При попытке реализации двух методов с одинаковым названием и получателем, возникнет ошибка компиляции.
## 7. В какой последовательности будут выведены элементы map[int]int?
Так как map представляет собой хеш-таблицу, порядок расположения элементов зависит от хэшей ключей добавляемых элементов. Поэтому порядок расположения элементов в map не зависит
от порядка добавления данных и каждый раз может быть разным.
## 8. В чем разница make и new?
**make** используется для создания slice, map и каналов с указанными параметрами(например, для слайсов длина и ёмкость). Возвращает экземпляр необходимого типа данных.
```go
slice := make([]string, 2, 5)
//typeof slice == []string
```
**new** выделяет память для экземпляра указанного типа данных (работает с любым типом данных, кроме map и каналов - указатели в их структурах будут проинициализированы nil). Возвращает указатель на
выделенную память. При этом поля структур или же необходимый тип данных проинициализированы нулевыми значениями данных типов.
```go
type Human struct {
    Name string
    Surname string
    Age int
}
human := new(Human)
// typeof human == *Human
```
## 9. Сколько существует способов задать переменную типа slice или map?
Существует 5 способов создать переменную типа slice:
```go
var slice []int
slice := []int{1, 2, 3}
slice := make([]int, 5)
slice := make([]int, 5, 10)
slice := new([]int)
```
Существует 3 корректных способа создать переменную типа map:
```go
mmap := map[int]int{
    1: 10,
    2: 20,
    3: 30,
}
mmap := make(map[int]int)
mmap := make(map[int]int, 5)
```
Существует 2 неккоректных способа создать map:
```go
var mmap map[int]int
mmap := new(map[int]int)
```
В данном случае mmap будет указывать на nil, поэтому при любой попытке добавить элемент в map будет паника.
## 10. Что выведет данная программа и почему?
```go
func update(p *int) {
    b := 2
    p = &b
}

func main() {
    var (
        a = 1
        p = &a
    )
    fmt.Println(*p)
    update(p)
    fmt.Println(*p)
}
```
1. Переменная p является указателем на a;
2. При разыменовывании указателя функция fmt.Println(*p) выводит "1";
3. Вызывается функция update со значением, хранящимся в указателе p. Сам update имеет копию указателя p - локальную переменную p с переданным значением адреса a;
4. Локальной p присваивается адрес переменной b. Функция завершается и локальные переменные удаляются из памяти. Переменная p в функции main() продолжает ссылаться на a;
5. При разыменовывании указателя функция fmt.Println(*p) выводит "1".

Ответ:
```console
1
1
```
## 11. Что выведет данная программа и почему?
```go
func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(wg sync.WaitGroup, i int) {
            fmt.Println(i)
            wg.Done()
        }(wg, i)
    }
    wg.Wait()
    fmt.Println("exit")
}
```
Так как в анонимные горутины передается копия WaitGroup, то вызов wg.Done() не будет оказывать влияния на wg из main(), поэтому после завершения анонимных горутин произойдет deadlock.

Пример вывода:
```console
3
2
0
4
1
fatal error: all goroutines are asleep - deadlock!
```
Чтобы исправить ошибку, можно использовать замыкание:
```go
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("exit")
}
```
или передавать в функцию адрес wg и работать с указателем:
```go
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
```
Пример вывода исправленной программы:
```console
0
4
3
1
2
exit
```
## 12. Что выведет данная программа и почему?
```go
func main() {
    n := 0
    if true {
        n := 1
        n++
    }
    fmt.Println(n)
}
```
Конструкция if имеет свою область видимости, в которой с помощью ":=" заново объявляется и инициализируется переменная n.
Данная переменная доступна только в блоке if, и при выходе из него fmt.Println(n) выведет "0".

Ответ:
```console
0
```
## 13. Что выведет данная программа и почему?
```go
func someAction(v []int8, b int8) {
    v[0] = 100
    v = append(v, b)
}

func main() {
    var a = []int8{1, 2, 3, 4, 5}
    someAction(a, 6)
    fmt.Println(a)
}
```
Slice - структура, состоящая из 3 полей: длина slice, ёмкость slice и указатель на массив. При передаче slice в качестве аргумента, функция работает с копией данной структуры. Из этого
следует, что присваивание значения внесет изменение в оригинальный slice, а операции, при которых изменяется ёмкость/длина/указатель на массив, не окажут влияния на оригинальный slice
(в данном примере такой функцией является append()).

Ответ:
```console
[100 2 3 4 5]
```
## 14. Что выведет данная программа и почему?
```go
func main() {
    slice := []string{"a", "a"}

    func(slice []string) {
        slice = append(slice, "a")
        slice[0] = "b"
        slice[1] = "b"
        fmt.Print(slice)
    }(slice)
    fmt.Print(slice)
}
```
Основываясь на сказанном ранее, так как анонимная функция работает не с замыканием, а с параметром slice, при выполнении append() все остальные изменения (в том числе изменения значений
элементов) будут проводиться с локальным slice.

Ответ:
```console
[b b a][a a]
```