# Go

## Linux环境搭建Golang开发环境

### 安装Golang SDK

- 将go1.11.1.linux-amd64.tar.gz传输到Linux服务器。
- 将go1.11.1.linux-amd64.tar.gz拷贝到 /opt 目录下。**[ cp  go1.11.1.linux-amd64.tar.gz   /opt ]**
- 移动到opt目录下。**[ cd  /opt ]**
- 解压压缩文件，解压之后可以看见一个go目录**[ tar  -zxvf  go1.11.1.linux-amd64.tar.gz ]**
- 移动到go/bin目录下。**[ cd  go/bin ]**
- 验证golang SDK版本。**[ ./go  version ]**

### 配置Golang环境变量

- 切换root权限编辑profile文件。**[vim  /etc/profile]**

```shell
export GOROOT = /opt/go				
# GOROOT是go的安装路径。
export PATH = $PATH:$GOROOT/bin		
# 配置go的可执行文件的路径，用以使用go命令和go工具。
export GOPATH = $HOME/goproject		
# GOPATH是编译后二进制的存放目的地和import包时的搜索路径 (其实就是工作目录)
# GOPATH之下主要包含三个目录: bin、pkg、src
# bin目录主要存放可执行文件; pkg目录存放编译好的库文件, 主要是*.a文件; src目录下主要存放go的源文件
```

- 若需要生效，需要注销重新登录。

### 开发Golang程序

```shell
# 方法一
go bulid hello.go
./hello
# 方法二
go run hello.go
```

## Golang概述

### 程序执行流程分析

- 对源码先执行**[ go  build ]**编译后，再执行，Golang的执行流程如图：



- 对源码直接执行**[ go  run ]**，Golang的执行流程如图：



两种执行方式的区别：

1. 如果先编译生成可执行文件，那么可以将该可执行文件拷贝到没有go开发环境的机器上，仍然可以运行。
2. 如果直接执行源代码，那么如果要在另一台机器上运行，需要Golang开发环境，否则无法执行。
3. 在编译时，编译器会将程序运行以来的库文件包含在可执行文件中，所以可执行文件变大了。

### 开发注意事项

- Golang源文件以“go”为扩展名。
- Golang应用程序的执行入口是main()函数。
- Golang严格区分大小写。
- Golang方法由一条条语句构成，**每个语句后边不需要分号**。（Golang会在每行后自动加分号）
- Golang编译器是一行行编译的，因此我们一行就写一条语句，不能把多条语句写在同一行，否则报错。（若强迫症非要一行写多句，除最后一句之外每句都加上分号，不建议，破坏了Golang简洁性。）
- Golang定义的**变量**或者**import的包**如果没有使用到，代码编译不能通过。

###转义字符

- **\t**：制表符（通常用作排版）
- **\n**：换行符
- **\\\\**：一个\
- **\ ""**：一个""
- **\r**：一个回车

### 注释

```go
//行注释
/*
块注释（快注释里面不允许有块注释嵌套）
*/
```

## Golang基本数据类型

### 整数类型

|  类型   | 有无符号 |      占用存储空间       |                  表数范围                  |        备注        |
| :-----: | :------: | :---------------------: | :----------------------------------------: | :----------------: |
|   int   |    有    | 32位[int32]/64位[int64] |                                            |                    |
|  int8   |    有    |      1 byte[8 bit]      |                 -128 ~ 127                 |                    |
|  int16  |    有    |     2 byte[16 bit]      |               -32768 ~ 32767               |                    |
|  int32  |    有    |     4 byte[32 bit]      |          -2147483648 ~ 2147483647          |                    |
|  int64  |    有    |     8 byte[64 bit]      | -9223372036854775808 ~ 9223372036854775807 |                    |
|  uint   |    无    | 32位[int32]/64位[int64] |                                            |                    |
|  uint8  |    无    |      1 byte[8 bit]      |                  0 ~ 255                   |                    |
| uint16  |    无    |     2 byte[16 bit]      |                 0 ~ 65535                  |                    |
| uint32  |    无    |     4 byte[32 bit]      |               0 ~ 4294967295               |                    |
| uint64  |    无    |     8 byte[64 bit]      |          0 ~ 18446744073709551615          |                    |
|  rune   |    有    |       与int32等价       |                与int32等价                 |   一个Unicode码    |
|  byte   |    无    |       与uint8等价       |                与uint8等价                 | 存储字符时选用byte |
| uintptr |    无    | 32位[int32]/64位[int64] |                                            |  存储指针用的类型  |

- 整数类型分为：**有符号**和**无符号**，int和uint的大小和具体系统有关
- **用于存放整数值，默认声明类型为int，默认值：0**

**Tips：在保证程序正常运作的情况下，应尽量使用占用空间较小的数据类型。**

### 浮点类型

|  类型   |  占用存储空间  |        表数范围        |
| :-----: | :------------: | :--------------------: |
| float32 | 4 byte[32 bit] |  -3.403E38 ~ 3.403E38  |
| float64 | 8 byte[64 bit] | -1.798E308 ~ 1.798E308 |

- **用于存放小数值，都是有符号的，默认声明类型为float64，默认值：0**
- 浮点类型有**固定的范围和字段长度**，不受具体OS(操作系统)的影响。
- 浮点类型的存储分为三部分：符号位、指数位、尾数位
- 浮点类型常量有两种表示形式
  - 十进制数形式：5.12、.512
  - 科学计数法形式：5.1234e2 = 5. 12*10的2次方、5.12E-2 = 5. 12/10的2次方

**Tips：尾数可能会丢失，float64比float32更精确，开发中，推荐使用float64。**

```go
var num1 float32 = -123.0000901
var num2 float64 = -123.0000901
fmt.Println("num1:" , num1 , "num2" , num2)
// 结果
// num1: -123.00009 num2 -123.0000901
```

### 字符类型

​	Golang中没有专门的字符类型，如果要存储单个字符(字母)，一般使用byte来保存。

- **字符类型：英文字母-1字节(byte)，汉字-3字节(int)，默认值：“”**
- 字符的本质是一个整数，直接输出，是按照该字符对应的UTF-8编码的码值，Golang使用UTF-8编码。
- 如果需要按照字符的方式输出，需要使用格式化输出`fmt.Printf("%c \n" ,ch)`

```go
var char_1 byte = 'a'
var char_2 byte = '0'
fmt.Println(char_1)				// 97
fmt.Println(char_2)				// 48
fmt.Printf("%c \n" ,char_1)		// a
fmt.Printf("%c \n" ,char_2)		// 0
```

- 保存的字符对应的码值大于255时，需要考虑使用int类型存储，否则报错。

```go
var char_3 byte = '北'
fmt.Printf("%c \n" ,char_3 )	// 报错 overflows byte
```

**Tips：字符类型是可以进行运算的，相当于一个整数**。

### 布尔类型

- **布尔类型：一个字节，默认值：0**
- 布尔类型又称bool类型，bool类型数据只允许取值 true 和 false 。

**Tips：布尔类型适用于逻辑运算，一般用于程序流程控制。**

### 字符串类型

> 传统：字符串是一串固定长度的字符连接起来的字符序列。

> Golang：字符串是由单个字节连接起来的。由字节组成的。

- Golang中字符串是不可变得，一旦赋值了，字符串就不能修改。
- 字符串的两种表示形式

```go
// 使用双引号，会识别转义字符
str1 := "hello world"
// 使用反引号，字符串的原生形式输出，可以防止攻击，输出源码等效果。
str2 := `
hello world
hello golang
`
```

- 字符串拼接

```go
var str = "hello" + "world"
str += "haha~"
```

**Tips：由于Golang中字符串的字节使用的是UTF-8编码标识的Unicode文本，中文乱码的问题不会出现。**

#### strings相关方法

> 查找子串是否在指定的字符串中

```go
strings.Contains("hello_world","hello")			 // true

```

> 统计一个字符串有几个指定的子串

```go
strings.Count("ceheese", "e")					 // 4

```

> 不区分大小写的字符串比较(==是区分字母大小写的)

```go
strings.EqualFold("abc", "Abc")					 // true

```

> 返回子串在字符串第一次出现的 index 值，如果没有返回-1

```go
strings.Index("NLT_abc", "abc") 				 // 4

```

> 返回子串在字符串最后一次出现的 index值，如果没有返回-1

```go
strings.LastIndex("go golang", "go")			 // 3

```

> 将指定的子串替换成 另外一个子串

```go
strings.Replace("go go", "go", "go语言", n)	    // go语言 go语言
// n 可以指定你希望替换几个，如果 n=-1 表示全部替换

```

> 按照指定的某个字符为分割标识，将一个字符串拆分成字符串数组。

```go
strings.Split("hello,wrold,ok", ",")			 // str[0]=hello,str[1]=wrold,str[2]=ok

```

> 将字符串的字母进行大小写的转换

```go
strings.ToLower("Go") 							 // go 
strings.ToUpper("Go") 							 // GO

```

> 将字符串左右两边的空格去掉

```go
strings.TrimSpace(" tn a lone gopher ntrn ")	 // tn a lone gopher ntrn

```

> 将字符串左右两边指定的字符去掉

```go
strings.Trim("! hello! ", " !") 				 // hello

```

> 将字符串左边指定的字符去掉 

```go
strings.TrimLeft("! hello! ", " !") 			 // hello!  

```

> 将字符串右边指定的字符去掉 

```go
strings.TrimRight("! hello! ", " !") 			 // !hello

```

> 判断字符串是否以指定的字符串开头

```go
strings.HasPrefix("ftp://192.168.10.1", "ftp")   // true

```

> 判断字符串是否以指定的字符串结束

```go
strings.HasSuffix("NLT_abc.jpg", "abc") 		 //false

```

### 指针类型

​	所有像int、float、bool、string、array和struct这些类型都属于**值类型**，这些类型的变量**直接指向存在内存中的值**，值类型的变量的值存储在栈中。

​	指针类型用于存储一个地址，这个地址指向的空间才是真正的值，可以通过 &获取对应的内存地址。

```go
var a int = 10
fmt.Println("a address:", &a)			// a address: 0xc420084008
// b 是一个指针变量
// b 的类型是 *int
// b 的本身的值是 &a
var b *int = &a
fmt.Printf("b:%v\n",b)					// b:0xc420084008
fmt.Printf("b address:%v\n",&b)			// b address:0xc42008c020
fmt.Printf("b value:%v\n",*b)			// b value:10

```

- 不同类型的指针不能互相转化，例如*int, int32, 以及int64。
- 任何普通指针类型*T和uintptr之间不能互相转化。
- 指针变量不能进行运算, 比如C/C++里面的++, --运算。

### 值类型和引用类型

- 值类型：基本数据类型 int 系列, float 系列, bool, string 、数组和结构体 struct，**变量直接存储值，内存通常在栈内分配。**
- 引用类型：指针、slice 切片、map、管道 chan、interface 等都是引用类型，**变量存储的是一个地址，这个地址指向的空间才是真正存储数据值，内存通常在堆内分配，没有任何变量引用这个地址，地址对应的数据空间成为一个垃圾，由GC来回收。**

### 数据类型转换

Golang在不同类型的变量之间赋值时需要显式转换，也就是说Golang中的数据类型不能自动转换。

#### 基本数据类型之间相互转换

> 高精度向低精度转换时，编译时不会报错，但是转换的结果出问题按溢出处理。
> 按溢出处理和我们希望的结果不一样，所以在转换时，需要考虑范围。

- **int 调精度**

```go
var a int8 = 100
var b int16 = int8(a)
var c int32 = int32(a)	
var d int64 = int64(a)	
fmt.Println(a,b,c,d)			// 100 100 100 100
var e int64 = 10000000
var f int8 = int8(e)
var g int16 = int16(e)	
var h int32 = int32(e)	
fmt.Println(e,f,g,h)			// 10000000 -128 -27008 10000000 

```

- **float 调精度**

```go
var a float32 = 11.11
var b float64 = float64(a)
fmt.Println(a,b)				// 11.11 11.109999656677246
var c float64 = 11.11
var d float32 = float32(c)
fmt.Println(c,d)				// 11.11 11.11

```

- **int 转 float**

```go
var a int = 100
var b float32 = float32(a)
var c float64 = float64(a)
fmt.Println(a,b,c)				// 100 100 100

```

**Tips：数据在处理时，要注意类型。**

```go
var a int32 = 12
var b int64
var c int8
b = a + 20						// 编译不通过
c = a + 20						// 编译不通过
b = int64(a) + 20				// 编译通过
c = int8(a) + 20				// 编译通过
fmt.Println(b,c)
var d int8 
var e int8 
d = int8(a) + 127				// 编译通过，但是结果不是127+12，按溢出处理
e = int8(a) + 128				// 编译不通过
fmt.Println(d,e)

```

#### 基本数据类型和String相互转换

> 基本数据类型转String类型

- 使用`fmt.Sprintf("%参数",表达式)`

```go
var a int64 = 99
var b float64 = 3.1415926
var c bool = true
var d byte = 'h'
var e string
e = fmt.Sprintf("%d",a)
fmt.Printf("type:%T value:%q\n",e,e)			// type:string value:"99"
e = fmt.Sprintf("%f",b)
fmt.Printf("type:%T value:%q\n",e,e)			// type:string value:"3.141593"
e = fmt.Sprintf("%t",c)
fmt.Printf("type:%T value:%q\n",e,e)			// type:string value:"true"
e = fmt.Sprintf("%c",d)
fmt.Printf("type:%T value:%q\n",e,e)			// type:string value:"h"

```

- 使用`strconv`包的函数

```go
// 常用
e = strconv.FormatInt(a,10)
fmt.Printf("type:%T value:%q\n",e,e)			// type:string value:"99"
e = strconv.FormatFloat(b,'f',10,64)
fmt.Printf("type:%T value:%q\n",e,e)			// type:string value:"3.1415926000"
e = strconv.FormatBool(c)
fmt.Printf("type:%T value:%q\n",e,e)			// type:string value:"true"
// Itoa 整数转字符串
e = strconv.Itoa(int(a))
fmt.Printf("type:%T value:%q\n",e,e)			// type:string value:"99"

```

- []byte 转 string

```go
var str string
str = string([]byte{99,98,97})
fmt.Printf("type:%T value:%q\n",str,str)		// type:string value:"cba"
str = string([]byte{'c','b','a'})
fmt.Printf("type:%T value:%q\n",str,str)		// type:string value:"cba"

```

> String类型转基本数据类型

- 使用`strconv`包的函数

```go
var a string = "123456"
var b int64
var c int
b, _ = strconv.ParseInt(a,10,64)
fmt.Printf("type:%T value:%v\n",b,b)			// type:int64 value:123456
-------------------------------------
var a string = "true"
var b bool
b, _ = strconv.ParseBool(a)
fmt.Printf("type:%T value:%v\n",b,b)			// type:bool value:true
-------------------------------------
var a string = "123.456"
var b float64
b, _ = strconv.ParseFloat(a,64)
fmt.Printf("type:%T value:%v\n",b,b)			// type:float64 value:123.456
// Atoi	字符串转整数
var a string = "123456"
b, err := strconv.Atoi(a)
fmt.Printf("type:%T value:%v\n",b,b)			// type:int value:123456

```

- string 转 []byte 或者 []rune

```go
// 转成[]byte后，可以处理英文和数字，但是不能处理中文。
var str1 string = "cba"
var arr1 []byte = []byte(str)
for _,v := range arr1 {							
    fmt.Printf("type:%T value:%q\n",v,v)		
}
// type:uint8 value:'c'
// type:uint8 value:'b'
// type:uint8 value:'a'
// 转成[]rune后，可以处理中文。
var str2 string = "牛B"
var arr2 []rune = []rune(str2)
for _,v := range arr2 {							
    fmt.Printf("type:%T value:%q\n",v,v)		
}
// type:int32 value:'牛'
// type:int32 value:'B'

```

**Tips：在将String转换成基本数据类型时，确保String类型能够转成有效的数据，可以把“123”转换成一个整数类型，但是不可以把“hello”转换成一个整数类型，如果这样做Golang直接将其转换成0，其他类型一个道理。**

## Golang变量和常量

### 标识符命名规则

- 由26个英文字母大小写，0-9，_组成
- 数字不可以开头
- 标识符不能包含空格
- 不能以系统**保留关键字**作为标识符（一共25个）
- Golang严格区分大小写`var num int `和`var Num int`中num和Num是两个不同变量。

### 常量使用方式

​	**常量是指变量一旦定义，在编译阶段它的值就是确定了的，在运行阶段是不可以被改变的**，golang 中 number|boolean|string 类型可以定义为常量。

```go
const PI = 3.14
const Pi float64 = 3.1415

```

### 变量使用方式

#### 局部变量

**局部变量：定义在函数内部的变量，只能在其被声明的函数内部被访问。**

- **单变量**

```go
// 指定变量类型，声明后若不赋值，使用默认值。
var x int
fmt.Println("x=", x)
// 类型推导，根据值自行判定变量类型。
var y = 10.11
fmt.Println("y=", y)
// 省略var,左侧的变量不应该是已经声明过的，否则会导致编译错误。
// 下面方式等价于 var i string	i = "hello"
z := "hello"
fmt.Println("z=", z)

```

- **多变量**

```go
// 指定变量类型，声明后若不赋值，使用默认值。
var x1 , x2 , x3 int
fmt.Println("x1=", x1 , "x2=" , x2 , "x3=", x3)
// 类型推导，根据值自行判定变量类型。
var y1 , y2 , y3 = 10.11 , "world" , 888
fmt.Println("y1=", y1 , "y2=", y2 , "y3=", y3)
// 省略var,左侧的变量不应该是已经声明过的，否则会导致编译错误。
// 下面方式等价于 
// var i string	
// i = "hello"
z1 , z2 , z3 := 10.11 , "world" , 888
fmt.Println("z1=", z1 , "z2=", z2 , "z3=", z3)

```

#### 全局变量

```go
// 定义全局变量，全局变量定义之后可以不去使用，编译不报错。
var n1 = 100
var n2 = "hello"
var (
	n3 = 200
    n4 = "world"
)
// 全局变量不支持在函数外进行赋值，:= 只支持局部变量操作
var str string  //定义了一个全局变量str
str = "test"   	//全局变量不支持这种操作
// 等价
str := "test"	//全局变量不支持这种操作，赋值语句不能再函数体外		 

```

### 分组定义

​	如果想定义多个常量，变量或者导入多个包，可以使用分组格式。

```go
import(
    "fmt"
    "os"
)
const(
    pi = 3.145
    prefix = "GO"
)
var(
    pi float32
    prefix string
)

```

### 特殊变量

> 变量 `_` 用来抛弃忽略某个值，golang 中不允许存在没有使用的变量

- 函数多返回值抛弃

```go
_, r := divide(10, 3)

```

- 导入包不使用(但是会调用导入包的 init方法)

```go
import(
    "fmt"
    _ "os"
)

```

> 关键字 iota 可以生成枚举类型，枚举值从0开始，每次加1

```go
const(
    x = iota					// x == 0
    y = iota 					// y == 1
    z = iota 					// z == 2
    w 							// 如果常量名称后面没有任何表达式，w 等价于 w = iota, w == 3
)
const v = iota 					// v == 0
const(
    e, f, g = iota, iota, iota 	// 如果 iota 在同一行，则 e == f == g == 0
)

```

### 作用域和可见性

- 不同的作用域，允许定义同名的变量。	
- 使用变量就近原则，如果局部变量没有，就去全局变量中找。
- 如果变量名、函数名、常量名首字母大写，则可以被其他包访问；如果首字母小写，则只能在本包中使用。[简单的理解成，首字母大写是公开的，首字母小写是私有的，在 golang 没有 public , private 等关键字。]

### 字节大小和数据类型

```go
var x int64 = 10
fmt.Printf("x的类型是：%T ，x占用的字节数是 %d 。", x , unsafe.Sizeof(x))
// 结果
// x的类型是：int64 ，x占用的字节数是 8 。

```



- 常量不能取地址
- goto不能进方法内部
- panic仅有最后一个可以被revover捕获
- 代码块中使用：= 相当于创造了个临时变量
- 如果类型实现String()，％v和％v格式将使用String()的值。因此，对该类型的String()函数内的类型使用％v会导致无限递归。

```go
//sync.Once实现
var once sync.Once

func GetIns2() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

```

- 给定：[1,3],[2,6],[8,10],[15,18] 返回：[1,6],[8,10],[15,18]

```go
type Interval struct {
	Start int
	End   int
}
func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	res := make([]Interval, 0)
	swap := Interval{}
	for k, v := range intervals {
		if k == 0 {
			swap = v
			continue
		}
		if v.Start <= swap.End {
			swap.End = v.End
		} else {
			res = append(res, swap)
			swap = v
		}
	}
	res = append(res, swap)
	return res
}

```

- map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。 定义的是var list map[string]Test，注意哦Test不是指针，而且map我们都知道是可以自动扩容的，那么原来的存储name的Test可能在地址A，但是如果map扩容了地址A就不是原来的Test了，所以go就不允许我们写数据。你改为var list map[string]*Test
- 有方向的channel不可被关闭
- golang中的map，的 key 可以是很多种类型，比如 bool, 数字，string, 指针, channel , 还有 只包含前面几个类型的 interface types, structs, arrays。 
  显然，slice， map 还有 function 是不可以了，因为这几个没法用 `==` 来判断，即不可比较类型。
- `recover()`的调用仅当它在`defer`函数中被直接调用时才有效。











