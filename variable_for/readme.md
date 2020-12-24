### variable in go 
1. ประกาศได้ 3 แบบ var
    - 1 ประกาศแบบระบุ ประเภทตัวแปร เช่น var i string , var j float64, var k int
    - 2 ประกาศแบบไม่ระบุ ประเภท เช่น var i = "hi", var j = 6.45, var k = 5
    - 3 ประกาศแบบลดรูป เช่น i := "hi" , j := 6.45 k := 5
2. ประกาศ const

### loop in go
1. for loop ex. for i := 0; i <= 5; i++ { ... }
2. while loop ex. i := 1 for i <= 5 { fmt.Println(i)  i = i + 1 }
3. do while loop ex.  j := 0 for { fmt.Println(j)  if !(j > 0) { break } }