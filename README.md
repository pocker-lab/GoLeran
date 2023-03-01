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

### 数组（Array）

1. 定义数组`[数组长度]值类型{值1, 值2, ...}`

    ```golang
    // 确定长度的数组定义
    arr1 := [5]int{0, 1, 2, 3, 4}
    fmt.Printf("arr1-->%v, %T\n", arr1, arr1)
    // arr1-->[0 1 2 3 4], [5]int

    // 不确定长度的数组定义，有go自动确定长度
    arr2 := [...]int{0, 21, 23, 3, 4, 1, 12, 1233, 123, 1, 2, 31, 23, 12}
    fmt.Printf("arr2-->%v, %T\n", arr2, arr2)
    // arr2-->[0 21 23 3 4 1 12 1233 123 1 2 31 23 12], [14]int

    // 先定义数组及长度，后面赋值
    var arr3 = new([10]int)
    arr3[5] = 5
    fmt.Printf("arr3-->%v, %T\n", arr3, arr3)
    // arr3-->&[0 0 0 0 0 5 0 0 0 0], *[10]int
    ```

2. 循环数组

    - 利用 for 循环数组

        ```golang
        // 利用 for 循环数组
        arr1 := [...]string{"狗子", "猫", "老虎"}

        for i := 0; i < len(arr1); i++ {
            fmt.Printf("%d-->%v\n", i, arr1[i])
            /*
                0-->狗子
                1-->猫
                2-->老虎
            */
        }
        ```

    - 利用 for + range 循环数组  
    range 是获取数组的所有索引，以及对应的值

        ```golang
        // 利用 for + range 循环数组
        // range 是获取数组的索引，对应的值
        for i, v := range arr1 {
            fmt.Printf("%d-->%v\n", i, v)
            /*
                0-->狗子
                1-->猫
                2-->老虎
            */
        }
        ```

3. `len()`和`cap()`

    > len(数组)：数组的长度  
    > cap(数组)：数组的容量长度

4. 多维数组

    > 二维数组

    ```golang
    // 二维数组，一个数组里面嵌套另一个数组
    er := [3][3]int{
        {0, 1, 2},
        {1, 2, 3},
        {2, 3, 4},
    }
    fmt.Printf("er: %v\n%T\n", er, er)
    /*
        er: [[0 1 2] [1 2 3] [2 3 4]]
        [3][3]int
    */
    ```

    > 三维数组

    ```golang
    // 三维数组
    san := [3][3][3]int{
        {{0, 1, 2}, {1, 2, 3}, {2, 3, 4}},
        {{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
        {{2, 3, 4}, {3, 4, 5}, {4, 5, 6}},
    }
    fmt.Printf("san: %v\n%T", san, san)
    /*
        san: [[[0 1 2] [1 2 3] [2 3 4]] [[1 2 3] [2 3 4] [3 4 5]] [[2 3 4] [3 4 5] [4 5 6]]]
        [3][3][3]int
    */
    ```

### 切片（slice）

1. 定义`[]值类型{值1, 值2, ...}`  
使用`make([]int, 长度, 容量)`创建的切片是指定类型的默认值  
如`int`类型的默认值为`0`，`string`类型的默认值为`""`

    ```golang
    // 定义空切片
    var sl1 []int
    // sl1: []int,[]
    // 定义切片并指定初始长度和容量
    sl2 := make([]int, 4, 4)
    // 
    // sl2: []int,[0 0 0 0]
    ```

2. 以数组为蓝本创建切片

    ```golang
    arr := [4]int{0, 1, 2, 3}
    // 以数组为蓝本创建切片
    sl1 := arr[:]   //[]int,[0 1 2 3]
    sl2 := arr[:3]  //[]int,[0 1 2]
    sl3 := arr[2:]  //[]int,[2 3]
    sl4 := arr[1:4] //[]int,[1 2 3]

    fmt.Printf("sl1-->%T,%v\n", sl1, sl1)
    fmt.Printf("sl2-->%T,%v\n", sl2, sl2)
    fmt.Printf("sl3-->%T,%v\n", sl3, sl3)
    fmt.Printf("sl4-->%T,%v\n", sl4, sl4)
    ```

3. 切片的用法

    1. 添加值`append()`

        ```golang
        sl1 := []int{1, 2, 3}
        // 在后面添加 4,5,6
        sl1 = append(sl1, 4, 5, 6)
        // sl1: [1 2 3 4 5 6]
        ```

    2. 复制切片`copy`

        将`sl2`的值复制到`sl1[3:]`的位置

        ```golang
        sl1 := make([]int, 6)
        sl2 := []int{4, 5, 6}
        copy(sl1[3:], sl2)
        // sl1: [0 0 0 4 5 6]
        // sl2: [4 5 6]
        ```

### map

`[key类型]值类型{key:值, key:值， ...}`

1. 定义

    ```golang
    // 定义map
    var m1 map[string]string
    // m1: map[string]string,map[]

    // 简短定义map
    m2 := map[string]string{}
    // m2: map[string]string,map[]

    // 使用make定义map
    m3 := make(map[string]string)
    // m3: map[string]string,map[]
    ```

2. 使用 map ，添加元素

    ```golang
    m1 := map[string]string{}
    m1["name"] = "渣男"
    m1["sex"] = "男"
    // m1: map[name:渣男 sex:男]
    ```

3. 使用 delete 删除元素

    ```golang
    m1 := map[string]string{
        "name": "渣男",
        "sex":  "男",
    }
    // 删除"sex"键值对
    delete(m1, "sex")
    // m1: map[name:渣男]
    ```

4. 循环map

    ```golang
    m1 := map[string]string{
        "name": "渣男",
        "sex":  "男",
    }
    for k, v := range m1 {
        fmt.Println(k, v)
        /*
        name 渣男
        sex 男
        */
    }
    ```

---

### 指针 `*`

[bilibili](https://www.bilibili.com/video/BV1mg4y187pS/?spm_id_from=333.999.0.0&vd_source=2fccd62dc87437c40cb47db5cee75f89)

---

### 函数方法 `func`

- 函数定义语法：

    ```golang
        func 函数名(传参1 传参2 类型,[不定长参...类型])(返回值 类型, 返回值 类型){
            函数体
            return 返回值
        }
    ```

1. 定义一个函数，并传入两个类型相同参数，返回一个参数

    ```golang
    func add(num1, num2 int) (num int) {
        num = num1 + num2
        return num
    }
    ```

2. 定义一个函数，传入不定长参数，返回一个参数

    ```golang
    func get(num1 ...int) (number1 int) {
        for _, v := range num1 {
            number1 += v
        }
        return number1
    }
    ```

3. 匿名函数

    ```golang
    func main() {
        ff := func() {
            num1 := 16
            fmt.Print(num1)
        }
        ff() // 16
    }
    ```

4. 递归函数

    ```golang
    // 求1-n的和，递归函数
    func getSum(n int) int {
        if n == 1 {
            return 1
        }
        return getSum(n-1) + n
    }
    ```

5. `defer`语句，延迟执行

    ```golang
    func main() {
        defer info() // 虽然调用函数在第一行，但是会在最后执行
        fmt.Print("1\n")
        fmt.Print("2\n")
        fmt.Print("3\n")
    }

    func info() {
        fmt.Print("我在这里")
    }
    ```

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
