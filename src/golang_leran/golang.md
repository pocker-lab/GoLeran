# Golang

## Go语言变量

### 声明变量的语法

---

```golang
var identifier type
```

``var``: 声明变量关键字

``identifier``: 变量名称

``type``: 变量类型

### 批量声明变量

---

```golang
var (
    name string
    age  int
    b    bool
)
```

### 变量的初始赋值

---

```golang
var 变量名 类型 = 初始值
// 相当于python的变量赋值
```

如果初始化值已存在，则可以省略类型；<br>
变量会从初始值中获得类型。

```golang
var 变量名 = 初始值
// 有初始值的情况下声明变量
```

### 初始化多个变量

---

初始化多个变量，中间用逗号<kbd>,</kbd>分隔

```golang
var name, age, b = "名字", 12, true
```

### 变量的零值

---

没有明确初始值的变量声明会被赋予它们的**零值**<br>

> ``int``和``float``类型变量的**零值**为``0``<br>
> ``string``类型变量的**零值**为``""``<br>
> ``bool``类型变量的**零值**为``false``<br>
> 切片、函数、指针变量的**零值**为``nil``<br>

### 短变量声明

---

在**函数中**，简洁赋值语句``:=``可在类型明确的地方代替``var``声明。

**函数外**的每个语句都必须以关键字开始（``var``, ``func``等等）<br>

> 注意：短变量声明``:=``只能使用在**函数内部**，函数外不能使用

```golang
func main() {
    name := "名字"
    age := 18
    b := true
}
```

### 匿名函数

---

如果我们接收到多个变量，有一些变量使用不到，可以使用下划线<kbd>_</kbd>表示变量名称，这种变量叫做**匿名变量**。例如：

```golang
package main

import "fmt"

func getNameAndAge() (string, int) {
    return "名字", 18
}

func main() {
    name, _ := getNameAndAge()
    fmt.Printf("name: %v\n", name)
}

//结果： name: 名字
```

---

## Go语言常量

**常量**：在程序**编译阶段**就确定下来的值，而程序在**运行时**无法改变该值。<br>
在Go程序中，**常量**可以是数值类型（包括整形、浮点型和复数行）、布尔类型、字符串类型等。

### 定义常量的语法

> 常量不能用``:=``语法声明<br>
> **常量名需要大写**
---

1. 定义一个**常量**使用``const``关键字，语法格式如下：

    ```golang
    const constantName [type] = value
    ```

    > ``const``: 定义常量关键字<br>
    > ``constantName``: 常量名称，**规范大写**<br>
    > ``type``: 常量类型<br>
    > ``value``: 常量的值<br>

2. 定义多个**常量**

   ```golang
   func main() {
       const (
           NAME = "名字"
           AGE  = 18
           B    = false
       )
   }
   ```

3. 多个**常量**赋初始值

    ```golang
    const NAME, AGE = "名字", 18
    ```

### 常量``iota``计数器

在**const**语句块中首个**常量**声明为``iota``后，开始值为``0``，后面的**常量**都会实现++，无需再次赋值``iota``<br>
遇到其他**const**语句块时被重置为``0``

---

```golang
func main() {
    const(
        A1 = iota   // 结果 0
        B1 = iota   // 结果 1
        C1 = iota   // 结果 2
    )
    // 使用 _ 占位符跳过某些值
    const(
        A1 = iota   // 结果 0
        _           // 使用 _ 占位符跳过 1
        C1 = iota   // 结果 2
    )
    // iota 声明中间插队
    const(
        A1 = iota   // 结果 0
        B1 = 100    // 结果 100
        C1 = iota   // 结果 2
    )
}
```

## Go语言字符串

一个Go语言字符串时一个任意**字节的常量序列**

### Go语言字符串字面量

---

字符串字面量使用双引号<kbd>"</kbd>或反引号<kbd>`</kbd>来创建。
