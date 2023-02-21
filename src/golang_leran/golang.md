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

在**函数中**，简洁赋值语句 ``:=`` 可在类型明确的地方代替``var``声明。

**函数外**的每个语句都必须以关键字开始（``var``, ``func``等等）<br>

> 注意：短变量声明 ``:=`` 只能使用在**函数内部**，函数外不能使用

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
在Go程序中，**常量**只能是数值类型（包括整型、浮点型和复数行）、布尔类型、字符串类型等。<br>
不曾使用的常量，在编译的时候，是不会报错

### 定义常量的语法

> 常量不能用  ``:=``  语法声明<br>
> **常量名需要大写**<br>
> <font color=#FF9999>一组常量中，如果某个常量没有初始值，默认和上一行一致</font>
---

1. 定义一个**常量**使用``const``关键字，语法格式如下：

    ```golang
    const constantName [type] = value
    ```

    > ``const``: 定义常量关键字<br>
    > ``constantName``: 常量名称，**规范大写**<br>
    > ``type``: 常量类型<br>
    > ``value``: 常量的值<br>

2. 定义一组**常量**

    ```golang
    func main() {
        // 一组常量中，如果某个常量没有初始值，默认和上一行一致
        const (
            NAME = "名字"
            AGE  = 18
            B    = false
            C               // 结果： false
        )
        const NAME1, AGE2 = "名字", 18
    }
    ```

### 常量``iota``关键字，枚举类型

iota: 特殊的**常量**，可以被编译器自动修改的常量<br>
    1. 每当定义一个``const``，``iota``的初始值为``0``<br>
    2. 每当定义一个**常量**，就会自动累加``1``<br>
    3. 直到下一个``const``出现，**清零**<br>

---

```golang
func main() {
    const(
        A1 = iota   // 结果 0
        B1          // 结果 1
        C1          // 结果 2
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

## Go语言数据类型

---

### 基本数据类型

- 布尔类型：bool
- 数值类型：Numeric Types
  - int8, int16, int32, int64, int
  - uint8, uint16, uint32, uint64, uint
  - float32, float64
  - complex64, complex128
  - byte:uint8
  - rune:int32
- 字符串：string

---

#### 布尔类型

布尔型的值只可以是常量 ``true`` 或者 ``false`` <br>
例子：``var b bool = true``

---

#### 数值类型

- 整数：int
  - 有符号：最高位表示符号位，0正数，1负数，其余位数表示数值
    - int8:
    - int16:
    - int32:
    - int64:
  - 无符号：所有位表示数值
    - uint8
    - uint16
    - uint32
    - uint64
- 浮点：float
  - float32
  - float64
- 复数：complex

---

#### 字符串类型

> 字符串就是一串固定长度的字符链接起来的字符序列<br>
> Go的字符串是由单个字节连接起来的<br>
> GO语言的字符串的字节使用UTF-8编码表示Unicode文本<br>

1. 概念：多个byte的集合，理解为一个字符序列<br>

2. 语法：使用双引号<kbd>"</kbd>或反引号<kbd>`</kbd>来创建

3. 转义字符： <kbd>\\</kbd>
   1. ``\n`` 换行符
   2. ``\t`` 制表符

---

### 复合数据类型

1. 指针类型（Pointer）
2. 数组类型
3. 结构化类型（struct）
4. Channel类型
5. 函数类型
6. 切片类型
7. 接口类型（interface）
8. Map类型

---

### 基本数据类型之间的转换：Type Convert

go语言是静态语言，**定义，赋值，运算必须类型一致**

1. 语法格式：``Type(Value)``
2. 常数（整数，浮点数）：有需要的时候，自动转型
3. 变量：需要手动转型 ``类型(值)``

---

## 运算符

<https://www.bilibili.com/video/av47467197/?>

### 算术运算符

### 关系运算符

### 逻辑运算符

### 位运算符

### 赋值运算符

### 优先级运算符
