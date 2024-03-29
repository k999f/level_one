1.	Какой самый эффективный способ конкатенации строк?

Самым эффективным методом конкатенации строг является использование strings.Builder. Согласно [документации](https://pkg.go.dev/strings#Builder): 
>A Builder is used to efficiently build a string using Write methods. It minimizes memory copying.

Также это подтверждают [бенчмарки](https://shantanubansal.medium.com/golang-how-to-efficiently-concatenate-strings-f2e51564f8d)

2.	Что такое интерфейсы, как они применяются в Go?

Интерфейсы описывают поведение других типов, указывая набор методов, которые должны быть реализованы для поддержки этого поведения. Чтобы тип данных удовлетворял интерфейсу, он должен реализовать все методы типа, требуемые этим интерфейсом. 
Интерфейсы применяются для уменьшения дублирования кода, когда приходится иметь дело с несколькими типами данных, выполняющими одну и ту же задачу. Также при помощи интерфейсов можно генерировать мок-объекты и уменьшать связанность кода.

3.	Чем отличаются RWMutex от Mutex?

RWMutex позволяет нескольким горутинам одновременно выполнять чтение, благодаря дополнительным методам - RLock и RUnlock (в отличие от обычного Mutex, который блокирует как одновременное чтение, так и запись).

4.	Чем отличаются буферизированные и не буферизированные каналы?

Тем, что буферизованные каналы имеют буфер, который хранит определенное количество значений. Это приводит к тому, что запись в небуферизованный канал блокирует горутину до тех пор, пока из канал не будет выполнено чтение, а буферизованный канал не блокирует горутину до тех пор, пока буфер не будет заполнен.

5.	Какой размер у структуры struct{}{}?

0 бит

6.	Есть ли в Go перегрузка методов или операторов?

Нет

7.	В какой последовательности будут выведены элементы map[int]int?
Пример:

```
m[0]=1
m[1]=124
m[2]=281
```

Если выполнить вывод через fmt.Println, то map будет отсортирована по ключам (сортировку выполняет сама функция печати). Но при итерации по элементам map порядок ее обхода случайный, и, соответственно, если выводить значения map по отдельности в цикле, то порядок их вывод тоже будет случайным.

8.	В чем разница make и new?

new выделяет память и возвращает указатель на нулевое значение типа данных; <br>
make возвращает инициализированное значение (применяется только для создания map, каналов и слайсов).

9.	Сколько существует способов задать переменную типа slice или map?

<b>slice</b>

Первый способ<br>
```s1 := []int{1, 2, 3, 4, 5}```

Второй способ<br>
```s2 := make([]int, 5)```

Третий способ
```
a := [6]int{0, 1, 2, 3, 4, 5} <br>
s3 := a[1:]
```

Четвертый способ (создастся указатель *int[])<br>
```s4 := new([]int)```

<b>map</b>

Первый способ<br>
```m1 := map[string]int{"one": 1, "two": 2, "three": 3}```

Второй способ<br>
```m2 := make(map[string]int)```

10.	Что выведет данная программа и почему?

```
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

Программа выведет 
```
1
1
```
Потому что внутри функции update будет меняться указатель, являющийся копией указателя p (т.к. значения в функцию передаются путем копирования).

11.	Что выведет данная программа и почему?

```
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

В случайном порядке напечатает числа от 0 до 4, а затем будет deadlock, потому что wg.Done прменяется не к wg, созданной в начале функции main, а к ее копии внутри анонимной функции.

12.	Что выведет данная программа и почему?

```
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```
Выведет 0, потому что внутри if будет создана новая переменная n, которая будет существовать только внутри фигурных скобок if.

13.	Что выведет данная программа и почему?

```
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

Выведет ``[100 2 3 4 5]``, потому что внутри someAction мы изменяем копию слайса (который представляет собой структуру, содержащую указатель на базовый массив, len - количесвто элементов в слайсе, и cap - длину базового массива). Заменив элемент с индексом 0, мы замениил его в базовом массиве (на который указывает и копия слайса внутри someAction и сам слайс в функции main). Затем, применив appen, мы меняем копию слайса в someAction, но не сам слайс в main.

14.	Что выведет данная программа и почему?

```
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

Выведет ``[b b a][a a]`` т.к. внутри анонимной функции используется копия slice, к которой применяется append, и теперь структура копии слайса внутри анонимной функции указывает на новый базовый массив (в 2 раза больше предыдущего), и значения для индексов 0 и 1 меняются, соответтсвенно, тоже в новом базовом массиве. Структура пременной slice в функции main все так же указывает на старый базовый массив. 