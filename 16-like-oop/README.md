### Struct

GO không có Class, nhưng có Struct. Struct bao gồm: fields và medthods <-> class.

##### field

```go
type Creature struct {
    Name string // field
    Real bool
}
```

##### method

Mô tả hành vi cụ thể của `Struct`. ex:

```go
func (c Creature) Dump() {
    fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

c:= Creature{Name: "Cat", Real: true} // c:= Creature{"Cat", true}
c.Dump()
```

##### Embedding (nhúng)

Tương tự như kế thừa trong OOP.

```go



type FlyingCreature struct {
  Creature
  WingSpan int
}

dragon := FlyingCreature{
    Creature{"Dragon", false, },
    15,
}
 
fmt.Println(dragon.Name)
fmt.Println(dragon.Real)
fmt.Println(dragon.WingSpan)

dragon.Dump()
```

##### Interface

Các Struct có chung tên method thì có thể dùng Interface. Ví dụ,
sinh viên và nhân viên có hàm những sau này:
- Sinh viên: SayHi(), Sing() và BorrowMoney().
- Nhân viên: SayHi(), Sing() và SpendSalary().
Trong ta thấy rằng sinh viên và nhân viên có 2 hàm giống nhau đó là đó là SayHi() và Sing() chung ta có thể tạo một Interface khác đó là con người có những method SayHi() và Sing() là được gọi là Interface và sinh viên và nhân viên có thể lấy interface đó để dùng được.


