# 1. Golang

[golang标准库中文版](https://studygolang.com/pkgdoc)

## 1.1. Go语言变量

### 1.1.1. 声明变量的语法

---

```golang
var identifier type
```

``var``: 声明变量关键字

``identifier``: 变量名称

``type``: 变量类型

### 1.1.2. 批量声明变量

---

```golang
var (
    name string
    age  int
    b    bool
)
```

### 1.1.3. 变量的初始赋值

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

### 1.1.4. 初始化多个变量

---

初始化多个变量，中间用逗号<kbd>,</kbd>分隔

```golang
var name, age, b = "名字", 12, true
```

### 1.1.5. 变量的零值

---

没有明确初始值的变量声明会被赋予它们的**零值**<br>

> ``int``和``float``类型变量的**零值**为``0``<br>
> ``string``类型变量的**零值**为``""``<br>
> ``bool``类型变量的**零值**为``false``<br>
> 切片、函数、指针变量的**零值**为``nil``<br>

### 1.1.6. 短变量声明

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

### 1.1.7. 匿名函数

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

## 1.2. Go语言常量

**常量**：在程序**编译阶段**就确定下来的值，而程序在**运行时**无法改变该值。<br>
在Go程序中，**常量**只能是数值类型（包括整型、浮点型和复数行）、布尔类型、字符串类型等。<br>
不曾使用的常量，在编译的时候，是不会报错

### 1.2.1. 定义常量的语法

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

### 1.2.2. 常量``iota``关键字，枚举类型

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

## 1.3. Go语言数据类型

---

### 1.3.1. 基本数据类型

---

#### 1.3.1.1. 布尔类型

布尔型的值只可以是常量 ``true`` 或者 ``false`` <br>
例子：``var b bool = true``

---

#### 1.3.1.2. 数值类型

- 整数：int
  - 有符号：最高位表示符号位，0正数，1负数，其余位数表示数值
    - int8, int16, int32, int64
  - 无符号：所有位表示数值
    - uint8, uint16, uint32, uint64
- 浮点：float
  - float32
  - float64
- 复数：complex

---

#### 1.3.1.3. 字符串类型

> 字符串就是一串固定长度的字符链接起来的字符序列<br>
> Go的字符串是由单个字节连接起来的<br>
> GO语言的字符串的字节使用UTF-8编码表示Unicode文本<br>

1. 概念：多个byte的集合，理解为一个字符序列<br>

2. 语法：使用双引号<kbd>"</kbd>或反引号<kbd>`</kbd>来创建

3. 转义字符： <kbd>\\</kbd>
   1. ``\n`` 换行符
   2. ``\t`` 制表符

---

### 1.3.2. 复合数据类型

1. 指针类型（Pointer）
2. 数组类型
3. 结构化类型（struct）
4. Channel类型
5. 函数类型
6. 切片类型
7. 接口类型（interface）
8. Map类型

---

#### 数组类型（Array）

> 数组一旦定义长度后，长度就不能变，不能访问超过长度的下标

1. 概念：存储一组相同**数据类型**的数据结构<br>理解为容器，存储一组数据
2. 语句：
    1. `var 数组名 [长度] 数据类型`
    2. `var 数组名 = [长度] 数据类型{元素1, 元素2,...}`：在定义数组时赋值
    3. `var 数组名 = [长度] 数据类型{下标:元素1, 下标：元素2,...}`：在定义数组时，通过下标给指定位置赋值
    4. `数组名 := [...] 数据类型 {元素...}`：根据元素给数组定义长度
3. 通过下标访问
    1. 下标：索引：index
    2. 从0开始，到[长度]-1
4. 长度和容量
    1. `len(array/map/slice/string)`：长度
    2. `cap(array)`： 容量

5. 数组的遍历
    1. 方法一：通过下标索引，依次取值
    2. 方法二：通过for循环，配合下标

        ```golang
            for i := 0; i < len(数组); i++ {
                数组[i]
            }
        ```

    3. 方法三：使用`range`，词义：范围
        [菜鸟教程：range](https://www.runoob.com/go/go-range.html)

        ```golang
            var arr1 = [5]int{1, 2, 3, 4, 5}
            for index, value := range arr1 {
                fmt.Printf("下标是%d，数值是%d\n", index, value)
            }
        ```

6. 数组的排序
    1. 排序算法：冒泡排序，插入排序，选择排序，希尔排序，堆排序，快速排序。。。。
    2. 冒泡排序：（Bubble Sort)，从小到大
        > 依次比较两个相邻的元素，如果他们的顺序（如从大到小）就把他们交换过来

        ```golang
            arr := [5]int{15, 23, 8, 10, 7}
            for i := 1; i < len(arr); i++ {
                for j := 0; j < len(arr)-i; j++ {
                    if arr[j] > arr[j+1] {
                        arr[j], arr[j+1] = arr[j+1], arr[j]
                    }
                }
                fmt.Println(arr)
            }
        ```

7. 多维数组
    1. 语法：`var 数组名 [长度] [长度]`
    2. **多层嵌套数组**

#### 切片（Slice）

1. 定义语法
    1. 正常语法
        1. `var 切片名 []类型`
        2. `切片名 := []int{元素1, 元素2, .....}`
    2. 使用`make()`函数来创建切片
        1. `var 切片名 []类型 = make([]类型, 长度)`
        2. `make([]类型, 长度, 容量)`

2. `append()`在切片末尾增加元素
    1. 语法：`切片名 = append(切片名, 元素1, 元素2, ....)`
    2. 在切片末尾添加另一个切片的所有元素：`切片名 = append(切片1, 切片2...)`后面三个<kbd>.</kbd>不能省略

3. 从已有的数组上进行切片

    ```golang
        arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
        s1 := arr[:5]  //1-5
        s2 := arr[3:8] //4-8
        s3 := arr[5:]  //6-10
        s4 := arr[:]   //1-10
    ```

---

### 1.3.3. 基本数据类型之间的转换：Type Convert

go语言是静态语言，**定义，赋值，运算必须类型一致**

1. 语法格式：``Type(Value)``
2. 常数（整数，浮点数）：有需要的时候，自动转型
3. 变量：需要手动转型 ``类型(值)``

---

## 1.4. Go运算符

1. 算术运算符

    > <kbd>+</kbd> 加法
    > <kbd>-</kbd> 减法
    > <kbd>*</kbd> 乘法
    > <kbd>/</kbd> 除法，取商
    > <kbd>%</kbd> 除法，取余
    > <kbd>++</kbd> 给自己加
    > <kbd>--</kbd> 给自己减1

2. 关系运算符（结果为bool类型）

    > <kbd>></kbd> 大于
    > <kbd><</kbd> 小于
    > <kbd>>=</kbd> 大于等于
    > <kbd><=</kbd> 小于等于
    > <kbd>==</kbd> 等于
    > <kbd>!=</kbd> 不等于

3. 逻辑运算符（结果为bool类型）

    |      运算符      | 描述        |
    | :-------------: | :---------- |
    |  <kbd>&&</kbd>  | 逻辑<font color='#FF9999'>**与**</font>运算符。<br>所有的操作数都是``true``，结果才为``true``<br>有一个``false``，结果就为``false``（一假则假，全真才真） |
    | <kbd>\|\|</kbd> | 逻辑<font color='#FF9999'>**或**</font>运算符。<br>所有的操作数都是``false``，结果才为``false``<br>有一个``true``，结果就为``true``（一真为真，全假为假） |
    |  <kbd>!</kbd>   | 逻辑<font color='#FF9999'>**非**</font>运算符。<br>`!true -> false`<br>`!false -> true` |

4. 位运算符

    [位运算符_bilibili](https://www.bilibili.com/video/BV1Db411s7in?p=23&spm_id_from=pageDriver&vd_source=2fccd62dc87437c40cb47db5cee75f89)

5. 赋值运算符

    **运算符**   | **描述**                       | **示例**
    :---------:|:----------------------------|:------------------------
    **`=`**   | 把=右边的值，赋值给=左边的变量             | `C = A + B` 将A + B的值赋值给C
    **`+=`**  | 把=右边的值与自身的值**相加**后，赋值给=左边的变量 | `C += A` 相当于 C = C + A
    **`-=`**  | 把=右边的值与自身的值**相减**后，赋值给=左边的变量 | `C -= A` 相当于 C = C - A
    **`*=`**  | 把=右边的值与自身的值**相乘**后，赋值给=左边的变量 | `C *= A` 相当于 C = C* A
    **`/=`**  | 把=右边的值与自身的值**相除**后，赋值给=左边的变量 | `C /= A` 相当于 C = C / A
    **`%=`**  | 把=右边的值与自身的值**取模**后，赋值给=左边的变量 | `C %= A` 相当于 C = C % A
    **`<<=`** | **左移位**并赋值运算符                | `C <<= 2` 相当于 C = C << 2
    **`>>=`** | **右移位**并赋值运算符                | `C >>= 2` 相当于 C = C >> 2
    **`&=`**  | **按位与**并赋值运算符                | `C &= 2` 相当于 C = C & 2
    **`\|=`**  | **按位或**并赋值运算符                | `C \|= 2` 相当于 C = C \| 2

## 1.5. Go键盘输入&打印输出

[fmt中文文档](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.3.html)

### 1.5.1. 输出

1. `fmt.Print()`：打印输出
2. `fmt.Printf()`：格式化打印输出
3. `fmt.Println()`：打印输出后换行
4. 格式化打印输出占位符：

    **占位符**      | **说明**          | **举例**                                           | **输出**
    :------------:|:---------------|:------------------------------------------------|:---------------
    **%v**       | 打印默认值           | Printf\("%v", "渣男"\)                             | 渣男
    **%T**       | 打印类型            | Printf\("%T", "渣男"\)                             | string
    **%%**       | 打印百分号           | Printf\("%%"\)                                   | %
    **%t**       | 打印true 或 false  | Printf\("%t\\n", 2 < 4\)                         | true
    **%d**       | 打印10进制          | Printf\("%d\\n", 13\)                            | 13
    **%f,%\.2f** | 打印浮点数，小数点后2位浮点数 | Printf\("%f,%0\.2f\\n", 3\.1415926, 3\.1415926\) | 3\.141593,3\.14
    **%q**       | 打印带双引号的字符串      | Printf\("%q\\n", "渣男"\)                          | 渣男

    ```golang
    func main() {
        fmt.Printf("name: %v\n", "渣男")                 //打印值
        fmt.Printf("%T\n", "渣男")                       //打印类型
        fmt.Printf("%%\n")                             //打印%号
        fmt.Printf("%t\n", 2 < 4)                      //打印bool值
        fmt.Printf("%d\n", 13)                         //打印十进制
        fmt.Printf("%f,%0.2f\n", 3.1415926, 3.1415926) //打印浮点数，打印小数点后2位浮点数
        fmt.Printf("%q\n", "渣男")                       //输出带双引号的字符串，由Go语法安全地转义
    }
    ```

### 1.5.2. 输入

1. 使用`fmt.Scanln()`来输入

```golang
func main() {
    var x int
    var y float64
    fmt.Println("请输入一个整数，一个浮点数：")
    fmt.Scanln(&x, &y) //读取键盘的输入，通过操作内存地址，赋值给x和y
    fmt.Printf("x: %d,y: %f\n", x, y)
}
```

## 1.6. Go逻辑判断

### 1.6.1. 程序的流程结构

- 程序的流程控制结构一共有三种：

1. 顺序结构：从上向下，逐行执行
2. 选择结构：条件满足，某些代码才会执行，0-1次
    1. 分支语句：`if`, `switch`, `select`
3. 循环结构：条件满足，某些代码会被反复的执行多次。0-N次
    1. 循环语句：`for`

### 1.6.2. `if`语句

- `if`语句语法格式：

    ```golang
    if 布尔表达式 {
        /* 在布尔表达式为 true 时执行 */
    }
    ```

    ```golang
    if 布尔表达式 {
        /* 在布尔表达式为 true 时执行 */
    } else {
        /* 在布尔表达式为 false 时执行 */
    }
    ```

    ```golang
    if 布尔表达式1 {
        /* 在布尔表达式为1 true 时执行 */
    } else if 布尔表达式2 {
        /* 在布尔表达式1为 false 时，布尔表达式2为 true 时执行 */
    } else {
        /* 在上面两个布尔表达式都为 false 时，执行 */
    }
    ```

    ```golang
    if 初始化语句; 条件判断{
        /* 初始化语句只作用在if语句 */
    }
    if num := 4; num > 5{
        /* 输出语句 */
    }
    ```

---

### 1.6.3. `switch`语句

- `switch`语句：

    ```golang
    switch 变量名{
        case 数值1: 分支1
        case 数值2: 分支2
        case 数值3: 分支3
        ......
        default: 最后一个分支
        /* 没有匹配项时执行 */
            
    }
    ```

    ```golang
    num := 3
    switch num {
    case 1:
    fmt.Println("第一季度")
    case 2:
    fmt.Println("第二季度")
    case 3:
    fmt.Println("第三季度")
    case 4:
    fmt.Println("第四季度")
    }
    /* 输出"第三季度" */
    ```

- 注意事项：
    1. `switch`可以作用在其他类型上，`case`后的数值必须和`switch`作用的变量类型一直
    2. `case`是无序的
    3. `case`后的数值是唯一的
    4. `default`是可选的操作

- `switch`其他写法

    1. 省略`switch`后的变量，相当于直接作用在`true`上

        ```golang
        /* 成绩：
        [0-59]，不及格c
        [60-69]，及格
        [70-79]，中
        [80-89]，良好
        [90-100]，优秀 */
        score := 88
        switch {
        case score >= 0 && score < 60:
        fmt.Println(score, "不及格")
        case score >= 60 && score < 69:
        fmt.Println(score, "及格")
        case score >= 70 && score < 79:
        fmt.Println(score, "中等")
        case score >= 80 && score < 89:
        fmt.Println(score, "良好")
        case score >= 90 && score <= 100:
        fmt.Println(score, "优秀")
        default:
        fmt.Println("成绩有误。。。")
        }
        ```

    2. `case`后可以同时跟随多个数值

        ```golang
            letter := "A"
            switch letter {
            case "A", "E", "I", "O", "U":
            fmt.Println(letter, "是元音")
            case "N", "M":
            fmt.Println("M或N")
            default:
            fmt.Println("其他")
            }
        ```

    3. `switch`后可以多一条初始化语句

        ```golang
            switch 初始化语句; 变量{
            }
        ```

- `switch`中的`break`和`fallthrough`语句

    1. `break`
        1. 可以使用在`switch`中，也可以使用在`for`循环中。
        2. 作用：强制结束`case`语句，从而结束`switch`分支。
    2. `fallthrough`
        1. 用于穿透`switch`.
        2. 作用：当`switch`中某个`case`匹配成功后，执行该`case`语句，如果`case`语句中有`fallthrough`，那么后面紧邻的`case`，无需匹配，直接穿透执行.
        3. `fallthrough`只能在某个case语句的最后一行

---

### 1.6.4. `for`循环语句

1. `for`语法

    ```golang
        for init; condition; post {}
    ```

    > `init`：初始化语句，只执行一次。
    > `condition`：条件判断。为`true`执行循环，为`false`结束循环。
    > `post`：每次循环结束后执行。用于变量的变换。

    ```golang
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d:hello, world.\n", i)
    }
    /* 打印5次hello, world. */
    ```

2. `for`语句的其他写法

    1. 省略`init`和`post`<br>相当于`while(条件)`

        ```golang
            for condition {}
        ```

    2. 省略`init`，`condition`和`post`<br>相当于`while(true)`，一直循环

        ```golang
            for {}
        ```

3. `for`循环练习题

    ```golang
    package main

    import "fmt"

    func main() {
        /*
        for循环的练习题：
        练习1：打印58-23数字
        练习2：求1-100的和
        练习3：打印1-100内，能够被3整除，但是不能被5整除的数字，统计被打印的数字的个数，每行打印5个
        */
        // 练习1：打印58
        for i := 58; i >= 23; i-- {
            fmt.Printf("%d ", i)
        }
        fmt.Println("\n>--------------------")
        // 练习2：求1-100的和
        num := 0
        for i := 1; i <= 100; i++ {
            num += i
        }
        fmt.Println("1-100的和：", num)
        fmt.Println(">--------------------")
        // 练习3：打印1-100内，能够被3整除，但是不能被5整除的数字，统计被打印的数字的个数，每行打印5个
        count := 0 //计数器
        for i := 1; i <= 100; i++ {
            if i%3 == 0 && i%5 != 0 {
                count++
                fmt.Print(i, "\t")
                if count%5 == 0 {
                    fmt.Println()
                }
            }
        }
        fmt.Println()
        fmt.Println("count -->", count)
    }
    ```

4. 多层`for`循环嵌套：99乘法表

    1. 99乘法表

        ```golang
        package main
        import "fmt"
        func main() {
            // 打印99乘法表
            // 控制行
            for i := 1; i < 10; i++ {
                // 控制列
                for j := 1; j <= i; j++ {
                    fmt.Printf("%d * %d = %d\t", j, i, j*i)
                }
                fmt.Println()
            }
        }
        ```

    2. 水仙花数

        ```golang
        package main
        import "fmt"
        func main() {
            /*
                水仙花数：三位数：[100-999]
                    每个位上的数字的立方和，刚好等于该数字本身，那么就叫水仙花数，一共有4个
                    比如：153
                    1*1*1 + 5*5*5 + 3*3*3 = 1 + 125 + 27 = 153
            */
            for i := 100; i < 1000; i++ {
                x := i / 100     //百位
                y := i / 10 % 10 //十位
                z := i % 10      //个位
                if x*x*x+y*y*y+z*z*z == i {
                    fmt.Println(i)
                }
            }
            fmt.Println("-------------")
            for x := 1; x < 10; x++ {
                for y := 0; y < 10; y++ {
                    for z := 0; z < 10; z++ {
                        if x*x*x+y*y*y+z*z*z == x*100+y*10+z*1 {
                            fmt.Println(x*100 + y*10 + z*1)
                        }
                    }
                }
            }
        }
        ```

    3. 2-100内的素数

        ```golang
            func main() {
                /*
                    打印2-100内的素数（只能被1和本身整除）
                */
                for i := 2; i <= 100; i++ {
                    flag := true
                    for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
                        if i%j == 0 {
                            flag = false
                            break
                        }
                    }
                    if flag {
                        fmt.Println(i)
                    }
                }
            }
        ```

5. `break`和`continue`

    1. `break`语句
        > 跳出循环，提前结束循环
    2. `continue`语句
        > 提前跳出某次循环，直接执行下个循环

---

## `goto`语句

1. 作用：可以无条件地转移到过程中指定的行

2. 语法结构：

    ```golang
        goto label:
        ..
        ..
        label: statement;
    ```

---

## 生成随机数

使用`math/rand`包

```golang
func main() {
    // setp1：设置种子数，设置为时间戳
    rand.NewSource(time.Now().UnixNano())
    for i := 0; i < 3; i++ {
        // setp2：调用生成随机数的函数
        fmt.Println("-->", rand.Intn(10))
        // 获取指定范围的随机数[4,10]-->[0,6]+4
        fmt.Println("<==", rand.Intn(4)+4) //取到的随机就是[4,10]之间的
    }
}
```
