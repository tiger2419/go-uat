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
  var data
  fmt.Sprintf("%d",data)



