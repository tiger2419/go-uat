#类型强转案例
例1、int转float64  
var age int = 25  强转-> floa64(age)
var index int = 30
var bill float32 = float32(index)
例2、interface转struct

type Cat struct{
}

  var cats chan interface{} 
  cat := <-cats
  cat.(实际的struct类型) 例如 转 Cat结构体  cat.(Cat)


基本数据类型转string
方式一:
  var data
  fmt.Sprintf("%d",data)
方式二:
  strconv包的format函数
  strconv.FormatBool()
  strconv.FormatFloat()
  strconv.FormatInt()
  strconv.FormatUint()
-----------------------------------------
string转基本数据类型
base代表进制 2->二进制 ...
bitSize代表整型类型 有0 8 16 32 64 分别对应 int int8 int16 int32 int64
这里bitSize没有 4 24
strconv.ParseInt(s string, base int, bitSize int)
strconv.ParseUint()
strconv.ParseFloat()
strconv.ParseInt()
strconv.Atoi()  //默认十进制
--------------------------------------
获取终端的方式
调用 fmt.Scanln(),例如下
  var scanln string
  fmt.Scanln(&scanln)

调用fmt.Scanf(),一般用于一次性输入多个
var name string
var age int
var salary float64
fmt.Scanf("%s %d %f",&name,&age,&salary)
==========================================================
位运算规则:
正数,先转换成补码,用补码运算  
  - 右移, 低位移除后,缺位的高位用符号位的值填补,因为正数的符号是0,所以用0来补充缺位
  - 左移, 符号位不变,低位缺位用0补充
负数,先转补码再用补码运算
  - 右移, 低位移除后,缺位的高位用符号位的值填补,因为负数的符号是1,所以用1来补充高位的缺位
  - 左移, 符号位不变,低位缺位用0补充
上面的正负数左移规则是针对未高位溢出的情况。
当高位溢出时符号位也参与运算,例如
  - var i int8 = 5 --> 00000101
  上面的i最多支持8位,当第一个1移到最左端时,就是一个负数了,但仍然是一个补码;第二个1都移走时,i就永远是0了
==========================================================






==========================================================


