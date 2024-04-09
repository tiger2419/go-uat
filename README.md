







#类型强转案例
例1、int转float64  
  var age int = 25  强转-> floa64(age)

例2、interface转struct

type Cat struct{
}

  var cats chan interface{} 
  cat := <-cats
  cat.(实际的struct类型) 例如 转 Cat结构体  cat.(Cat)
