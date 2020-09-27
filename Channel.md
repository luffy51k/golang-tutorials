
### Định nghĩa

```bash
    <-Channel<-
```

ex:

results chan<- int: dữ liệu từ biến result sẽ đi vào channel
worker <-chan int: dữ liệu từ biến worker sẽ dc nhận từ channel

### channel parameter function


```go
/* channel 2 chiều */
func example(values chan <kieu_du_lieu>, ...) {
    //
}

/* channel 1 chiều: in */
func example(values <-chan <kieu_du_lieu>, ...) {
    //
}

/* channel 1 chiều: out */
func example(values chan<- <kieu_du_lieu>, ...) {
    //
}
```