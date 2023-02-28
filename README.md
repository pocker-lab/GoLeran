# go

## 变量声明

1. 申明变量，显式声明  
`var 变量名 类型`
2. 申明变量并赋值  
`var 变量名 类型 = 变量值`
3. 简短申明变量，隐式声明  
`变量名 := 变量值`

## 注释

1. 单行注释 `\\`
2. 多行注释 `\**\`

## 包

> 一个文件夹下只能有一个包  
> 这个文件夹就是这个包  
> 可以有多个go文件

1. 新建包  
新建一个文件夹，如testA  
在文件夹里新建一个go文件，如A.go  
在go文件的第一行写上【`package 文件夹名`】，表明这是一个包文件  
此时目录结构为：

    ```shell
    go_do
    │  go.md
    │  go.mod
    │  main.go
    │  README.md
    |
    └─testA
        A.go
    ```

1. 引用包

    在main.go文件中引用本地包文件

    ```golang
    import (
        "go_do/testA"   // 应用包文件
        "go_do/testA" // 引用包文件，并起别名
    )
    ```

## 公有与私有

> 公有与私有的区分是以【**首字母大小写划分**】

一个包里公有的东西可以被其他包所使用  
私有的东西只能在自己所在的文件内使用  

## 基本数据类型

### 整数类型

1. `uint`：不含负数

    ```golang
    var uint1 uint = 99
    fmt.Printf("uint1的类型%T，值是：%v\n", uint1, uint1)
    // uint1的类型uint，值是：99
    ```

2. `int`：包含负数

    ```golang
    var int1 int = -99
    fmt.Printf("int1的类型%T，值是：%v\n", int1, int1)
    // int1的类型int，值是：-99
    ```

### 浮点类型

1. `float32`
2. `float64`

    ```golang
    var f64 float64 = 3.222
    fmt.Printf("f64的类型%T，值是：%v\n", f64, f64)
    // f64的类型float64，值是：3.222
    ```

### 字符串类型（string）

1. `string`：`"`双引号包裹起来的是字符串类型`"`

    ```golang
    var str string = "字符串"
    fmt.Printf("str的类型%T，值是：%v\n", str, str)
    // str的类型string，值是：字符串
    ```

### 布尔类型（bool）

1. `true`

    ```golang
    var t1 bool = true
    fmt.Printf("t1的类型%T，值是：%v\n", t1, t1)
    // t1的类型bool，值是：true
    ```

2. `false`

    ```golang
    var f1 bool = false
    fmt.Printf("f1的类型%T，值是：%v\n", f1, f1)
    // f1的类型bool，值是：false
    ```

### 数据类型转换

> [strconv_api_文档](https://devdocs.io/go/strconv/index#pkg-overview)

#### string 转 int

> `strconv.Atoi()`

```golang
// string转int
var str1 string = "123321"
val, _ := strconv.Atoi(str1)
fmt.Printf("val --> %T, %v", val, val)
// val --> int, 123321
```

#### string 转 int64

```golang
// string 转 int64
var str1 string = "123321"
val, _ := strconv.ParseInt(str1, 10, 64)
fmt.Printf("val --> %T, %v", val, val)
// val --> int64, 123321
```

#### string 转 float64

> 转float32的话 `strconv.ParseFloat(str1, 32)`

```golang
// string 转 float64
var str1 string = "3.1415926"
val, _ := strconv.ParseFloat(str1, 64)
fmt.Printf("val --> %T, %v", val, val)
// val --> float64, 3.1415926
```

#### int 转 string

```golang
// int 转 string
var num1 int = 123321
val := strconv.Itoa(num1)
fmt.Printf("val --> %T, %v", val, val)
// val --> string, 123321
```

#### int64 转 string

```golang
// int64(以十进制) 转 string
var num1 int64 = 123321
val := strconv.FormatInt(num1, 10)
fmt.Printf("val --> %T, %v", val, val)
// val --> string, 123321
```

## 复杂数据类型

### 结构 `struct`

### 接口 `interface`

### 数据 `[数组长度]值类型{值1, 值2, ...}`

### 切片 slice `[]值类型{值1, 值2, ...}`

### map `[key类型]值类型{key:值, key:值， ...}`

### 指针 `*`

### 函数 `func`

### 管道 `chan`

## 流程控制

### 递增递减

1. `++` 自增1
2. `--` 自减1

```golang
var num int = 5
num++    // 5 + 1 = 6
num--    // 6 - 1 = 5
```

### `if`条件判断

1. 初始化语句只在if语句中起作用
2. 布尔表达式的结果必须是`bool`类型

```golang
if bool1 {
    /* 在bool1结果为 true 时执行 */
} else if bool2 {
    /* 在bool1为 false 时，bool2为 true 时执行 */
} else {
    /* 在上面两个布尔表达式都为 false 时，执行 */
}
```

### `switch`选择语句

1. `switch`语句：

    ```golang
    switch [初始化语句;] 变量{
        case 数值1, [数值2，数值3][布尔类型判断]: 
            /* 符合case后执行的操作 */
        case 数值1, [数值2，数值3][布尔类型判断]: 
            /* 符合case后执行的操作 */
            [fallthrough] 
            // 执行到fallthrough时会接着执行下一个case语句
        case 数值1, [数值2，数值3][布尔类型判断]: 
            /* 符合case后执行的操作 */
        ......
        [default:] 
        /* 没有匹配项时执行 */
    }
    ```

2. 注意事项：
    1. `switch`可以作用在其他类型上，`case`后的数值必须和`switch`作用的变量类型一直
    2. `case`是无序的
    3. `case`后的数值是唯一的
    4. `default`是可选的操作

3. `fallthrough`：
    1. 用于穿透`switch`.
    2. 作用：当`switch`中某个`case`匹配成功后，执行该`case`语句，如果`case`语句中有`fallthrough`，那么后面紧邻的`case`，无需匹配，直接穿透执行.
    3. `fallthrough`只能在某个case语句的最后一行

### `for`循环语句

```golang
for 初始化语句; bool判断语句; 自增/自减 {
    /* 执行语句 */
}
```

```golang
for val := 0; val < 10; val++ {
    fmt.Print(val) // 0,1, ...,8, 9
}
```

### 跳转语句

- `break`：用于中断结束`for`循环或跳出`switch`语句

    ```golang
    var a int = 1
    for a < 10 {
        fmt.Printf("a 的值为 : %d\n", a)
        a++
        if a > 5 {
            /* 在 a 大于 5 的时候使用 break 语句跳出循环  */
            break
        }
    }
    ```

- `continue`：跳过当前循环的剩余语句，然后继续进行下一轮循环。

    ```golang
    var a int = 1
    for a < 10 {
        if a == 5 {
            /* 跳过此次循环 数值5被跳过了 */
            a = a + 1
            continue
        }
        fmt.Printf("a 的值为 : %d\n", a)
        a++
    }
    ```

- `goto`：将控制转移到被标记的语句

    ```golang
        var i int = 1
        for i < 10 {
            fmt.Printf("%v-->", i)
            i++
            if i == 5 {
                goto B
            }
        }
    B:
        fmt.Println("\n这里是B")
        /* 在A是5的时候，会跳转到B这里*/
    ```
