# My Note


### 11-Con trỏ

*01- Các cách khởi tạo con trỏ ?*

```go
// cách 1: khai báo tường minh
var p *int
// cách 2: sử dụng built-in new() function
ptr := new(int)
```

*02- pointer tới ...?*

pointer trỏ về biến:

```go
var a = 5
var p = &a // gán con trỏ p vào địa chỉ biến a
```

pointer trỏ tới pointer:

*03- lấy / gán giá trị thông qua con trỏ?*

```go
var a = 5
var p = &a // gán con trỏ p vào địa chỉ biến a
*p = 7
```

*04- hai pointer trỏ về 1 biến?*

```go
var a = 75
var p1 = &a
var p2 = &a

if p1 == p2 {
    fmt.Println("Both pointers p1 and p2 point to the same variable.")
}
```
