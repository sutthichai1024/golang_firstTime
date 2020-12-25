### Condition (IF Else)
 - Golang if else จะสามารถ ประกาศตัวแปร ได้ แต่จะใช้ได้แค่ใน if or else if 
 - if else ไม่ต้องมี วงเว็บ ครอบ เงือนไข

### Condition (switch case) 3 type
 -  switch ทั่วไปคือ เข้า case ที่ ตามตัวแปรที่สนใจ แต่ไม่ต้องใส่ break ส่วนตัวแปร fallthrough เป็นตัวแปรที่มีเพื่อให้ตกลงไปใน case ถัดไปได้ คล้าย continue

 ### array 
  - ex. var a [3]int ,b := [5]int{5, 2, 1, 4, 6}

  ### Slices 
  - เป็น array ที่ ขนาดความกว้างของชุดข้อมูลประเภท Slices นั้นสามารถยืดหดได้ 

 ### maps 
  - ชุดคู่ของข้อมูลประเภท key/value คล้าย slice แต่ที่ตัวตำแหน่งหรือคีย์นั้นเปลี่ยนเป็นตัวแปรประเภทต่างๆได้

  ### Range
   - Range การวนซ้ำตามลำดับบน array, slice, maps เราสามารถวนไปตามขนาดของที่มีข้างในได้ range นั้นจะ return index หรือ key และ ค่าของตำแหน่งนั้นๆ