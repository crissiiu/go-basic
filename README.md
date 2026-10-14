# Go Basic — Giáo án tổng kết 14 bài nền tảng

Tài liệu này hệ thống lại toàn bộ mã nguồn từ `lesson1` đến `lesson14`. Mạch học xuyên suốt là xây dựng chương trình bán hàng nhỏ: lưu tên và giá sản phẩm, tính hóa đơn, quản lý danh sách, cập nhật tồn kho, rồi xử lý lỗi và hiểu thời điểm thực thi.

Mục tiêu sau khi hoàn thành: tự đọc được các bài hiện có, giải thích được dữ liệu đi qua từng hàm như thế nào và viết lại một chương trình quản lý sản phẩm đơn giản mà không chép mẫu.

> Phạm vi hiện tại là Go cơ bản chạy trên terminal. Dự án chưa có HTTP, JSON, Gin, goroutine hay bài kiểm thử tự động. Các chủ đề này nằm trong phần học tiếp, không phải chức năng đã triển khai.

## Mục lục

- [1. Chuẩn bị và cách chạy](#1-chuẩn-bị-và-cách-chạy)
- [2. Lộ trình học](#2-lộ-trình-học)
- [3. Giáo án từng bài](#3-giáo-án-từng-bài)
- [4. Các điểm cần đối chiếu trong mã nguồn](#4-các-điểm-cần-đối-chiếu-trong-mã-nguồn)
- [5. Bài thực hành tổng kết](#5-bài-thực-hành-tổng-kết)
- [6. Checklist chốt kiến thức](#6-checklist-chốt-kiến-thức)
- [7. Hướng học tiếp](#7-hướng-học-tiếp)

## 1. Chuẩn bị và cách chạy

### Cấu trúc dự án

```text
go-basic/
├── README.md             # Giáo án tổng kết
├── lesson1/
│   ├── go.mod            # Khai báo module của bài
│   ├── main.go           # Chương trình thực thi
│   └── README.md         # Hiện đang trống
├── lesson2/
│   ├── go.mod
│   └── main.go
├── ...
└── lesson14/
    ├── go.mod
    └── main.go
```

Mỗi thư mục bài học là một module riêng, có `go.mod` riêng. Thư mục gốc chưa có `go.mod` hoặc `go.work`, vì vậy hãy chạy lệnh bên trong bài muốn học.

Các `go.mod` hiện khai báo `go 1.27.1`. Môi trường dùng để đối chiếu giáo án cũng có Go `1.27.1` trên Windows. Đây là phiên bản trong checkout này, không phải khẳng định về phiên bản Go mới nhất.

Trong VS Code, mở thư mục dự án, chọn **Terminal → New Terminal**, rồi chạy PowerShell:

```powershell
Set-Location 'E:\workspace\go\github.com\crissiiu\go-basic'
go version
Set-Location .\lesson1
go run .
```

Chuyển sang bài khác:

```powershell
Set-Location ..\lesson2
go run .
```

`go run .` biên dịch và chạy package trong thư mục hiện tại. Không cần cài Gin hay thư viện ngoài: các bài hiện chỉ dùng thư viện chuẩn `fmt` và `errors`.

### Module, package và import khác nhau thế nào?

| Thành phần | Vai trò | Ví dụ trong dự án |
|---|---|---|
| Module | Đơn vị quản lý mã nguồn và phụ thuộc; tên khai báo trong `go.mod` | `github.com/crissiiu/go-basic/lesson1` |
| Package | Nhóm các file Go cùng package trong một thư mục | `package main` |
| Import | Cho phép sử dụng một package khác | `import "fmt"` |
| `main()` | Điểm bắt đầu của chương trình thực thi | `func main() { ... }` |

Tên module không tự động đổi theo tên thư mục. `lesson2`, `lesson3` dùng đường dẫn module ngắn hơn các bài khác; `lesson6` đang khai báo trùng tên module của `lesson5`. Chạy từng bài độc lập vẫn được, nhưng cần thống nhất lại nếu sau này tổ chức nhiều module cùng một workspace hoặc import giữa các bài.

### Cách học mỗi bài

1. Đọc mục tiêu và lý thuyết trong tài liệu này.
2. Mở `main.go`, dự đoán kết quả trước khi chạy.
3. Chạy `go run .`, đối chiếu với dự đoán.
4. Thay một dữ liệu đầu vào và giải thích vì sao kết quả thay đổi.
5. Làm bài tập nhỏ cuối bài; chỉ chuyển tiếp khi có thể tự giải thích đoạn code chính.

## 2. Lộ trình học

| Chặng | Bài | Nội dung | Đầu ra cần đạt |
|---|---|---|---|
| 1. Đọc và điều khiển chương trình | 1–4 | Cấu trúc Go, biến, kiểu dữ liệu, điều kiện, vòng lặp | Tính tiền theo số lượng và chọn đúng mức giảm |
| 2. Chia nhỏ và quản lý dữ liệu | 5–8 | Hàm, nhiều giá trị trả về, array, slice, map | Tách phép tính thành hàm; duyệt và tra cứu dữ liệu |
| 3. Mô hình hóa và hành vi | 9–12 | Struct, con trỏ, method, interface | Cập nhật sản phẩm đúng nơi; dùng chung hành vi giữa nhiều kiểu |
| 4. Xử lý thất bại và thời điểm chạy | 13–14 | `error`, `nil`, `defer` | Giải thích nguyên nhân thất bại và thứ tự thực thi |

Không cần cố định số ngày học. Mỗi bài tương ứng một buổi học hoặc một lần ôn; tiêu chí hoàn thành là giải thích và tự viết lại được, không chỉ chạy thành công.

## 3. Giáo án từng bài

### Bài 1 — Chương trình Go đầu tiên

**Mã nguồn:** [lesson1/main.go](lesson1/main.go)

**Mục tiêu:** hiểu chương trình bắt đầu ở đâu và cách in ra terminal.

```go
package main

import "fmt"

func main() {
    fmt.Println("Nguyen Thanh Sieu")
}
```

- `package main`: khai báo package dành cho chương trình thực thi.
- `import "fmt"`: sử dụng package định dạng và xuất dữ liệu.
- `func`: bắt đầu khai báo hàm; `main` là tên hàm.
- `()`: danh sách tham số; ở đây không có tham số.
- `{ ... }`: khối lệnh thuộc hàm.
- `fmt.Println(...)`: gọi hàm `Println` của package `fmt`, in dữ liệu và xuống dòng.
- `//`: chú thích một dòng; không được thực thi.

Bài hiện in ba dòng: tên, ngành học và nội dung phát triển dự án. Các lệnh chạy lần lượt từ trên xuống trong `main`.

**Thực hành:** thêm một dòng giới thiệu mục tiêu học Go. Tự trả lời: nếu bỏ `import "fmt"` nhưng vẫn gọi `fmt.Println`, chương trình còn biên dịch được không? Vì sao?

### Bài 2 — Biến, hằng số và kiểu dữ liệu

**Mã nguồn:** [lesson2/main.go](lesson2/main.go)

**Mục tiêu:** lưu dữ liệu, phân biệt khai báo với gán và hiểu chuyển kiểu.

```go
productName := "Noodle"  // string
price := 4.2             // float64
quantity := 25           // int
total := price * float64(quantity)
```

Biến là tên gắn với một giá trị có kiểu. Trong ví dụ, Go suy ra kiểu từ giá trị khởi tạo. `price` là `float64`, còn `quantity` là `int`; phép nhân cần các toán hạng có kiểu phù hợp nên chuyển `quantity` sang `float64`.

`float64(quantity)` tạo giá trị số thực để tính toán; nó không đổi kiểu của biến `quantity` đang có.

| Cú pháp | Ý nghĩa |
|---|---|
| `quantity := 25` | Khai báo ngắn trong hàm; phải có ít nhất một biến mới ở cùng phạm vi |
| `quantity = 30` | Gán giá trị mới cho biến đã tồn tại |
| `var quantity int` | Khai báo biến `int`, giá trị ban đầu là `0` |
| `const taxRate = 0.1` | Khai báo hằng; không thể gán lại |
| `price * float64(quantity)` | Nhân đơn giá với số lượng đã chuyển kiểu |

Mã hiện có biến và chuyển kiểu; `const` mới được nhắc trong chú thích. Với dữ liệu hiện tại, tổng tiền là `105`.

**Thực hành:** đổi số lượng thành `10`, tính nhẩm rồi chạy lại. Thêm một hằng `shopName` và in cùng hóa đơn.

### Bài 3 — Điều kiện và dừng hàm sớm

**Mã nguồn:** [lesson3/main.go](lesson3/main.go)

**Mục tiêu:** kiểm tra dữ liệu đầu vào và chọn một trong nhiều mức giảm giá.

```go
if quantity <= 0 {
    fmt.Println("So luong khong hop le")
    return
}
```

Điều kiện trả về `bool`: `true` hoặc `false`. Khi số lượng không hợp lệ, `return` kết thúc lần gọi `PrintBill` hiện tại. Nó không kết thúc cả chương trình: `main` vẫn tiếp tục gọi `PrintBill` cho các số lượng kế tiếp.

Các toán tử cần nhớ: `==` là so sánh bằng, `!=` là khác, `>=` và `<=` là so sánh có bao gồm bằng. `=` là gán giá trị, không phải so sánh.

Trong bài này, `Total` tính tiền rồi gọi `Discount`. Hàm `Discount` trả về **số tiền sau giảm**:

| Số lượng | Quy tắc | Tiền phải trả với giá `4.2` |
|---|---|---|
| `0` | Không hợp lệ | In thông báo |
| `10` | Không giảm | `42` |
| `25` | Giảm 10% | `94.5` |
| `50` | Giảm 20% | `168` |

Phải kiểm tra ngưỡng `>= 50` trước `>= 20`. Nếu kiểm tra `>= 20` và trả về ngay, số lượng `50` cũng thỏa điều kiện đó và nhận nhầm mức giảm.

**Mở rộng lý thuyết:** `switch` chọn nhánh theo giá trị hoặc điều kiện. Mã hiện chưa dùng `switch`, dù chú thích đầu file có nhắc đến nó. Có thể viết lại quy tắc giảm như sau:

```go
switch {
case quantity >= 50:
    return total * 0.8
case quantity >= 20:
    return total * 0.9
default:
    return total
}
```

**Thực hành:** dự đoán các mốc `19`, `20`, `49`, `50`; giải thích vì sao thứ tự điều kiện quyết định kết quả.

### Bài 4 — Vòng lặp `for`, `continue` và `break`

**Mã nguồn:** [lesson4/main.go](lesson4/main.go)

**Mục tiêu:** lặp công việc và theo dõi sự thay đổi của biến điều khiển.

Go dùng `for` cho vòng lặp. Dạng `for quantity <= 60` kiểm tra điều kiện trước mỗi lượt. `quantity += 10` tương đương `quantity = quantity + 10`.

Trong bài, khi `quantity == 30`, chương trình tăng nó lên `40` rồi gọi `continue`. `continue` bỏ qua phần còn lại của lượt hiện tại và quay về kiểm tra điều kiện.

Luồng thực tế:

```text
0 → in thông báo không hợp lệ
10 → in hóa đơn
20 → in hóa đơn
30 → tăng thành 40, continue
40 → in hóa đơn
50 → in hóa đơn
60 → in hóa đơn
70 → điều kiện sai, kết thúc vòng lặp
```

Việc tăng biến trước `continue` là cần thiết ở cấu trúc hiện tại. Nếu bỏ dòng tăng này, `quantity` sẽ mắc ở `30` và vòng lặp không kết thúc.

`break` kết thúc vòng lặp gần nhất; bài hiện chưa dùng. `return` kết thúc hàm đang chạy. Ba lệnh có phạm vi tác động khác nhau.

**Thực hành:** thay vị trí bỏ qua từ `30` thành `20`. Viết thử dạng `for quantity := 0; quantity <= 60; quantity += 10` và xem bước tăng chạy thế nào khi gặp `continue`.

### Bài 5 — Hàm, tham số và nhiều giá trị trả về

**Mã nguồn:** [lesson5/main.go](lesson5/main.go)

**Mục tiêu:** tách tính toán khỏi việc trình bày kết quả.

```go
func CalculateBill(price float64, quantity int) (float64, float64, float64)
```

Đọc từ trái sang phải: khai báo hàm `CalculateBill`, nhận đơn giá và số lượng, trả về ba giá trị `float64`. Trong mã thực tế, ba giá trị lần lượt là tiền gốc, tiền giảm và tiền cuối cùng.

```go
total, discount, result := CalculateBill(price, quantity)
```

Thứ tự biến bên trái khớp thứ tự giá trị trong `return`. Tên biến ở nơi gọi không bắt buộc giống tên biến trong hàm; vị trí mới quyết định nhận giá trị nào.

Luồng gọi:

```text
main → PrintBill → CalculateBill → Total
                               → Discount
```

Từ bài 5, ý nghĩa các hàm thay đổi so với bài 3–4:

```text
Total        = price × quantity
Discount     = số tiền được giảm
CalculateBill trả về: total, discount, total - discount
```

Với `quantity = 25`: tiền gốc `105`, tiền giảm `10.5`, tiền cuối `94.5`.

`fmt.Printf` in theo mẫu: `%s` cho chuỗi, `%d` cho số nguyên, `%.2f` cho số thực với hai chữ số sau dấu thập phân, `\n` để xuống dòng. Định dạng hiển thị không thay đổi giá trị tính toán bên trong.

**Thực hành:** giải thích vì sao `Discount` dùng `0.2` ở bài này nhưng dùng `0.8` ở bài 3. Nếu chỉ chép công thức mà không đọc ý nghĩa giá trị trả về, hóa đơn sẽ sai thế nào?

### Bài 6 — Array, slice và `range`

**Mã nguồn:** [lesson6/main.go](lesson6/main.go)

**Mục tiêu:** quản lý nhiều số lượng và duyệt từng phần tử.

```go
quantities := []int{0, 10, 25, 30, 50}
quantities = append(quantities, 60)
for _, quantity := range quantities {
    // Xử lý từng số lượng
}
```

Array có độ dài nằm trong kiểu, ví dụ `[3]int{10, 20, 30}`. Slice dùng `[]int`, mô tả một vùng dữ liệu trên mảng nền và có độ dài có thể thay đổi qua các thao tác như `append`. Mã hiện dùng slice; array là kiến thức đối chiếu.

`range` trên slice trả về chỉ số và giá trị. `_` bỏ qua chỉ số. `quantity` là bản sao của giá trị phần tử trong từng lượt.

Điểm dễ nhầm: trong bài này, `quantity += 10` khi gặp `30` chỉ đổi biến cục bộ của lượt lặp. Nó không sửa `quantities`, không làm lần lặp sau nhận `40`. Sau `continue`, lượt tiếp theo nhận `50` từ slice.

Muốn sửa phần tử thật, cần chỉ số, chẳng hạn `quantities[i] = 40`. Phải giữ kết quả `append`, vì slice trả về có thể sử dụng mảng nền mới.

**Thực hành:** in `quantities` sau vòng lặp để xác nhận phần tử `30` còn nguyên. Sau đó viết một vòng lặp sửa phần tử thông qua chỉ số.

### Bài 7 — `len`, `cap`, `make` và mảng nền

**Mã nguồn:** [lesson7/main.go](lesson7/main.go)

**Mục tiêu:** phân biệt độ dài slice, sức chứa và chia sẻ dữ liệu.

```go
a := make([]int, 2, 3)
a[0] = 10
a[1] = 20
b := a
```

`make([]int, 2, 3)` tạo slice dài `2`, sức chứa `3`; hai phần tử ban đầu có giá trị `0`. `len` là số phần tử đang truy cập được, `cap` là số phần tử có thể chứa từ vị trí bắt đầu slice trong mảng nền hiện tại. Dù `cap(a) == 3`, truy cập `a[2]` khi `len(a) == 2` vẫn gây panic.

`b := a` sao chép mô tả slice, không sao chép toàn bộ mảng nền. Ban đầu cả hai cùng nhìn vào một mảng:

| Thao tác | Slice `a` | Slice `b` | Quan hệ mảng nền |
|---|---|---|---|
| `b := a` | `[10 20]`, len 2, cap 3 | `[10 20]`, len 2, cap 3 | Dùng chung |
| `a = append(a, 30)` | `[10 20 30]`, len 3, cap 3 | `[10 20]`, len 2, cap 3 | Vẫn dùng chung |
| `a = append(a, 40)` | `[10 20 30 40]`, len 4 | `[10 20]`, len 2, cap 3 | `a` chuyển sang mảng mới |
| `a[0] = 99` | `[99 20 30 40]` | `[10 20]` | Sửa `a` không đổi `b` |

Khi append vượt sức chứa, Go cấp phát mảng nền mới và sao chép dữ liệu. Không nên học thuộc một quy tắc nhân đôi sức chứa: chỉ cần biết sức chứa mới đủ cho các phần tử cần thêm.

Hai chú thích cuối file hiện dự đoán sai kết quả. Kết quả đúng của hai lệnh in slice cuối cùng là:

```text
[99 20 30 40]
[10 20]
```

**Thực hành:** chuyển `a[0] = 99` lên trước lần append vượt `cap`, dự đoán `b[0]`. Muốn bản sao độc lập ngay từ đầu, thử `b := make([]int, len(a))` rồi `copy(b, a)`.

### Bài 8 — Map và tra cứu theo khóa

**Mã nguồn:** [lesson8/main.go](lesson8/main.go)

**Mục tiêu:** phân biệt phần tử không tồn tại với phần tử có giá trị bằng không.

```go
prices := map[string]float64{"Milk": 2.5}
price, ok := prices["Milk"]
```

`map[string]float64` lưu cặp khóa chuỗi và giá trị số thực. Với khóa `"Milk"`, nhận `2.5, true`. Với khóa không tồn tại, nhận `0, false`.

Không thể kết luận khóa không tồn tại chỉ bằng `price == 0`: cửa hàng có thể có sản phẩm giá `0`. Biến `ok` mới thể hiện sự tồn tại của khóa.

`!ok` nghĩa là phủ định `ok`. Hàm `FindPrice` nhận map và tên sản phẩm, trả kết quả tra cứu cùng cờ tồn tại. Bài gọi với `"Milk"` và `"Sit"`, nên lần lượt in giá `2.5` và thông báo không tìm thấy.

**Thực hành:** thêm một sản phẩm giá `0`, tra cứu nó và một tên không tồn tại. Giải thích sự khác nhau của hai kết quả.

### Bài 9 — Struct và giá trị lấy ra từ map

**Mã nguồn:** [lesson9/main.go](lesson9/main.go)

**Mục tiêu:** gom thông tin liên quan thành một kiểu và hiểu khi nào cần ghi dữ liệu trở lại map.

```go
type Product struct {
    ID    int
    Name  string
    Price float64
    Stock int
}
```

Struct gom các trường có thể khác kiểu để biểu diễn một sản phẩm. `Product{ID: 1, Name: "Milk", ...}` tạo một giá trị; `product.Stock` truy cập trường tồn kho.

`map[int]Product` dùng số nguyên làm khóa và lưu giá trị struct. Khi lấy `product := products[id]`, biến `product` nhận một bản sao. Muốn lưu thay đổi vào map, cần gán lại:

```go
product.Stock -= 1
products[id] = product
```

**Đối chiếu đúng mã hiện tại:** `FindProduct` đã trừ `1` trên bản sao rồi trả về; `main` lại trừ `1` nữa và ghi vào map. Pizza ban đầu có kho `14`, cuối cùng in kho `12`. Hàm mang tên “tìm” đang có thêm hành vi giảm kho, khiến luồng nghiệp vụ khó theo dõi.

Khóa map `3` còn chứa sản phẩm có `ID: 2`. Khóa và trường `ID` là hai dữ liệu riêng; Go không tự đồng bộ chúng.

**Thực hành:** đề xuất cách để `FindProduct` chỉ tra cứu, còn giảm kho nằm ở một nơi duy nhất. Sau khi sửa trong bài tập, một lần mua một sản phẩm phải đưa kho từ `14` về `13`.

### Bài 10 — Con trỏ, `&`, `*` và `nil`

**Mã nguồn:** [lesson10/main.go](lesson10/main.go)

**Mục tiêu:** truyền địa chỉ để hàm cập nhật đúng sản phẩm gốc.

| Cú pháp | Cách đọc |
|---|---|
| `pizza Product` | Một giá trị kiểu `Product` |
| `&pizza` | Địa chỉ của biến `pizza` |
| `p *Product` | `p` có kiểu con trỏ tới `Product` |
| `*p` | Truy cập giá trị mà con trỏ `p` trỏ tới |
| `p == nil` | Con trỏ chưa trỏ tới sản phẩm nào |

Go vẫn truyền tham số theo giá trị. Khi truyền con trỏ, giá trị được sao chép là địa chỉ; hàm có thể dùng địa chỉ đó để sửa đối tượng gốc.

```go
func ReduceStock(pizza *Product, quantity int) bool {
    if pizza == nil || quantity <= 0 {
        return false
    }
    if pizza.Stock < quantity {
        return false
    }
    pizza.Stock -= quantity
    return true
}
```

Với con trỏ struct, `pizza.Stock` là cách viết thuận tiện thay cho `(*pizza).Stock`. `||` nghĩa là “hoặc” và đánh giá ngắn mạch: nếu vế trái đã đúng, vế phải không cần chạy.

Chương trình bắt đầu với kho `14`. Giảm `4` thành công, kho còn `10`. Giảm tiếp `20` thất bại và giữ nguyên kho. Truyền `nil` cũng thất bại nhưng được kiểm tra trước khi truy cập trường nên không panic. Khi in trực tiếp con trỏ struct, `fmt.Println` thường thể hiện dạng `&{...}` trong ví dụ này.

**Thực hành:** thử số lượng `0`, số âm và số lượng đúng bằng kho còn lại. Phân biệt “hàm trả `false`” với “chương trình bị panic”.

### Bài 11 — Method và receiver

**Mã nguồn:** [lesson11/main.go](lesson11/main.go)

**Mục tiêu:** gắn hành vi vào một kiểu và chọn receiver phù hợp.

```go
func (pizza Product) Print(ok bool) { /* ... */ }
func (pizza *Product) ReduceStock(quantity int) bool { /* ... */ }
```

Phần giữa `func` và tên method gọi là **receiver**: cho biết method thuộc kiểu nào. `Print` nhận bản sao `Product`; `ReduceStock` nhận con trỏ để thay đổi sản phẩm gốc.

Từ hàm thường `ReduceStock(&pizza, 4)` ở bài 10, chuyển thành `pizza.ReduceStock(4)` ở bài 11. Với biến `pizza` có thể lấy địa chỉ, Go hỗ trợ gọi pointer method theo cú pháp này.

```go
pizza.Print(pizza.ReduceStock(4))
```

Trong mã hiện tại, `ReduceStock(4)` cập nhật kho gốc trước khi `Print` in kết quả, nên dòng này in kho `10`. Lời gọi tiếp theo với `Restock(6)` in kho `16`. Kết quả chạy cả bài lần lượt là sản phẩm ban đầu có kho `14`, sau giảm còn `10`, sau nhập thành `16`, rồi thông báo thất bại với số lượng âm. Value receiver vẫn là bản sao: nếu sửa trường bên trong thân `Print`, thay đổi đó không cập nhật sản phẩm gốc.

Để dễ đọc và theo dõi từng bước cập nhật, có thể tách thành hai câu lệnh trong bài tập:

```go
ok := pizza.ReduceStock(4)
pizza.Print(ok)
```

`Restock` hiện còn kiểm tra `pizza.Stock < quantity`, nên có thể từ chối nhập thêm một lượng lớn hơn kho đang có. Đây không phải quy tắc nhập kho thông thường. Method này cũng chưa kiểm tra receiver `nil`; bài 13 xử lý đầy đủ hơn.

**Thực hành:** tách lời gọi lồng thành hai dòng, dự đoán lại kết quả. So sánh lý do `ReduceStock` cần kiểm tra đủ hàng còn `Restock` thông thường không cần điều kiện đó.

### Bài 12 — Interface và hành vi dùng chung

**Mã nguồn:** [lesson12/main.go](lesson12/main.go)

**Mục tiêu:** viết một hàm làm việc với nhiều kiểu có cùng hành vi.

```go
type Named interface {
    GetName() string
}

func Print(item Named) {
    fmt.Println(item.GetName())
}
```

Interface `Named` yêu cầu method `GetName() string`. `Product` và `Service` có trường dữ liệu khác tên nhưng đều cung cấp method đúng chữ ký, nên cả hai đều dùng được với `Print`.

Go thỏa mãn interface một cách ngầm định: không cần từ khóa `implements`. Hàm `Print` chỉ cần biết đối tượng lấy tên được, không cần biết nó là sản phẩm hay dịch vụ.

Các trường `id`, `name`, `serviceName` bắt đầu bằng chữ thường nên không được export ra ngoài package. `GetName` bắt đầu bằng chữ hoa nên được export.

Một điểm nối với bài 11: nếu `GetName` chỉ khai báo trên receiver `*Product`, thì `*Product` thỏa mãn interface này, còn giá trị `Product` không tự có method set tương ứng để gán vào `Named`. Mã hiện dùng value receiver nên cả `Product` và `*Product` đều có method cần thiết.

Kết quả bài hiện tại:

```text
Pizza
Delivery
```

**Thực hành:** thêm kiểu `Customer` có `GetName() string` rồi truyền vào `Print` mà không sửa hàm `Print`.

### Bài 13 — `error`: biết thất bại vì sao

**Mã nguồn:** [lesson13/main.go](lesson13/main.go)

**Mục tiêu:** thay kết quả đúng/sai bằng thông tin lỗi có ý nghĩa.

`bool` cho biết thành công hay thất bại, nhưng không phân biệt được sản phẩm thiếu, số lượng sai hay không đủ hàng. `error` biểu diễn nguyên nhân thất bại theo quy ước:

```go
err := pizza.ReduceStock(4)
if err != nil {
    fmt.Println("Loi:", err)
    return
}
```

- Trả `nil`: không có lỗi.
- `errors.New("...")`: tạo lỗi với thông báo cố định.
- `fmt.Errorf("... %d ...", value)`: tạo lỗi có dữ liệu cụ thể.
- Kiểm tra `err != nil` trước khi tiếp tục công việc phụ thuộc vào thành công.

`ReduceStock` kiểm tra lần lượt sản phẩm tồn tại, số lượng dương và đủ hàng; chỉ cập nhật kho sau khi qua tất cả điều kiện. `Restock` kiểm tra sản phẩm và số lượng dương rồi cộng kho.

| Thao tác trong `main` | Kết quả |
|---|---|
| Kho ban đầu | `14` |
| `ReduceStock(4)` | Thành công, còn `10` |
| `Restock(6)` | Thành công, lên `16` |
| `Restock(-2)` | Trả lỗi, vẫn `16` |
| `missing.Restock(6)` với con trỏ `nil` | Trả lỗi sản phẩm không tồn tại |

Có thể gọi pointer method trên con trỏ `nil`; có panic hay không phụ thuộc vào thân method. Ở đây, `Restock` kiểm tra `nil` trước khi truy cập trường nên trả lỗi an toàn.

`Print` là value method và không phải nơi nên truyền một con trỏ `nil` để in sản phẩm. Mã hiện xử lý lỗi của `missing` trực tiếp bằng `fmt.Println`.

**Thực hành:** gọi `ReduceStock(100)` để xem thông báo có cả số đang còn và số cần lấy. Xác nhận kho không đổi sau mọi thao tác thất bại.

### Bài 14 — `defer` và thứ tự thực thi

**Mã nguồn:** [lesson14/main.go](lesson14/main.go)

**Mục tiêu:** hiểu lệnh nào bị hoãn và giá trị nào được ghi nhận ngay.

`defer` đăng ký một lời gọi để thực thi khi hàm bao quanh kết thúc. Nó không chạy nền và không trì hoãn theo một khoảng thời gian.

File có ba ví dụ; `main` hiện chỉ gọi `Demo3()`. Để chạy từng ví dụ khác, đổi lời gọi trong `main`.

| Hàm | Thứ tự in | Điều cần hiểu |
|---|---|---|
| `Demo()` | `Bat dau` → `Dang xu ly` → `Don dep` | Lời gọi defer chạy khi hàm kết thúc |
| `Demo2()` | `C` → `B` → `A` | Nhiều defer chạy theo thứ tự đăng ký sau chạy trước |
| `Demo3()` | `C: 6` → `B: 6` → `A: 4` | Đối số được đánh giá lúc đăng ký defer |

Trong `defer fmt.Println("A:", quantity)`, giá trị `quantity` được ghi nhận khi đang là `4`. Việc gán lại `quantity = 6` không đổi đối số đã ghi nhận đó.

Ứng dụng thường gặp là đóng tài nguyên sau khi mở thành công, chẳng hạn `defer file.Close()`. Khi học tới xử lý file, cần xem xét cả lỗi mở và lỗi đóng theo yêu cầu chương trình.

**Thực hành:** dự đoán kết quả nếu thêm một defer in `D` sau các defer đang có. Tự giải thích sự khác nhau giữa “đánh giá đối số ngay” và “thực thi lời gọi sau”.

## 4. Các điểm cần đối chiếu trong mã nguồn

Đây là ghi chú đọc hiểu cho checkout hiện tại. Việc biên soạn README không tự sửa các file bài học.

| Vị trí | Hiện trạng | Điều cần ghi nhớ hoặc tự chỉnh khi thực hành |
|---|---|---|
| `lesson3/main.go` | Chú thích nhắc `switch`, mã chỉ dùng `if` | Phân biệt chủ đề được nhắc với cú pháp đã thực hành |
| `lesson3`–`lesson5` | `Discount` đổi ý nghĩa giá trị trả về | Bài 3–4 trả tiền sau giảm; từ bài 5 trả tiền được giảm |
| `lesson6/go.mod` | Tên module đang là `lesson5` | Tên module không tự theo tên thư mục |
| `lesson6/main.go` | Tăng biến `quantity` trong `range` | Không cập nhật phần tử slice hoặc bước duyệt |
| `lesson7/main.go` | Hai chú thích kết quả cuối sai | Append vượt cap tách mảng nền; phần tử `30` vẫn còn trong `a` |
| `lesson9/main.go` | Trừ kho trong cả `FindProduct` và `main` | Một lần chạy giảm từ `14` xuống `12`; cần tách tra cứu khỏi cập nhật |
| `lesson9/main.go` | Khóa `3` chứa sản phẩm `ID: 2` | Kiểm tra tính nhất quán khóa và mã sản phẩm |
| `lesson11/main.go` | Lồng lời gọi cập nhật vào lời gọi in | Mã hiện in kho sau cập nhật; tách hai câu lệnh giúp theo dõi luồng dễ hơn |
| `lesson11/main.go` | `Restock` kiểm tra kho hiện có và thiếu kiểm tra `nil` | Đối chiếu phiên bản trả `error` trong bài 13 |
| Toàn dự án | Chưa có file `*_test.go` | Chạy thành công chưa chứng minh mọi trường hợp đều đúng |

Các ví dụ dùng `float64` để học phép tính. Số thực nhị phân có thể có sai số; `%.2f` chỉ định dạng hiển thị. Khi mở rộng thành ứng dụng tiền tệ thực tế, cần chọn cách biểu diễn tiền và quy tắc làm tròn rõ ràng, ví dụ số nguyên theo đơn vị tiền nhỏ nhất.

## 5. Bài thực hành tổng kết

### Đề bài: quản lý kho và tính hóa đơn

Tự viết một chương trình terminal riêng, kết hợp kiến thức bài 1–14. Phần này là bài tập đề xuất, chưa có mã triển khai trong dự án.

**Dữ liệu:** sản phẩm gồm `ID`, `Name`, `Price`, `Stock`; lưu bằng `map[int]Product`.

**Yêu cầu:**

1. Tra cứu sản phẩm bằng ID; báo lỗi nếu không tồn tại.
2. Kiểm tra số lượng mua lớn hơn `0` và không vượt kho.
3. Tính tiền gốc, tiền giảm và tiền cuối theo quy tắc: từ `20` sản phẩm giảm 10%, từ `50` giảm 20%.
4. Chỉ giảm kho một lần khi giao dịch thành công; giao dịch thất bại không đổi kho.
5. Ghi sản phẩm đã cập nhật trở lại map.
6. In hóa đơn gồm tên, số lượng, ba khoản tiền và kho còn lại.
7. Dùng method cho thao tác kho và `error` để mô tả nguyên nhân thất bại.

### Gợi ý thứ tự xây dựng

Bắt đầu từ struct và map của bài 9. Viết hàm tìm kiếm chỉ đọc dữ liệu. Dùng pointer method trả `error` theo bài 13 để cập nhật biến sản phẩm đã tìm. Tính hóa đơn theo bài 5, rồi lưu lại vào map sau khi các kiểm tra cần thiết thành công.

Với `map[int]Product`, không lấy địa chỉ trực tiếp bằng `&products[id]`. Hãy lấy ra biến cục bộ, truyền địa chỉ của biến đó cho method khi cần, rồi gán kết quả trở lại map.

Không cần đưa interface hoặc defer vào nghiệp vụ nếu chưa có nhu cầu. Để ôn riêng, có thể dùng interface `Named` in tên nhiều kiểu và dùng một hàm demo defer để kiểm tra thứ tự chạy.

### Các tình huống phải tự kiểm tra

| Tình huống | Kết quả mong đợi |
|---|---|
| ID không tồn tại | Có lỗi; không thay đổi dữ liệu |
| Số lượng `0` hoặc âm | Có lỗi; không thay đổi kho |
| Số lượng lớn hơn kho | Có lỗi đủ thông tin; không thay đổi kho |
| Mua đúng bằng kho | Thành công; kho còn `0` |
| Mua `19`, `20`, `49`, `50` với kho đủ | Chọn đúng ranh giới giảm giá |
| Mua thành công hai lần liên tiếp | Lần sau sử dụng kho đã cập nhật từ lần trước |
| Tra cứu sản phẩm mà không mua | Kho giữ nguyên |

**Tiêu chí đạt:** chương trình chạy được, các tình huống trên đúng, và bạn giải thích được mỗi nơi sao chép giá trị, mỗi nơi sửa dữ liệu gốc, mỗi nhánh trả lỗi.

## 6. Checklist chốt kiến thức

- [ ] Giải thích được `package main`, `import`, `func main` và `go.mod`.
- [ ] Chạy từng bài từ đúng thư mục module bằng `go run .`.
- [ ] Phân biệt `:=`, `=`, `==`; hiểu khi nào cần chuyển `int` sang `float64`.
- [ ] Theo dõi được luồng `if`, `return`, `for`, `continue` và `break`.
- [ ] Đọc được chữ ký hàm và nhận đúng thứ tự nhiều giá trị trả về.
- [ ] Phân biệt tiền được giảm với tiền phải trả sau giảm.
- [ ] Phân biệt array và slice; giải thích được `len`, `cap`, `make`, `append`.
- [ ] Biết lúc nào hai slice dùng chung mảng nền và lúc nào append tạo mảng mới.
- [ ] Hiểu biến giá trị trong `range` là bản sao trong ví dụ của bài 6.
- [ ] Dùng `value, ok` khi tra cứu map.
- [ ] Hiểu lấy struct từ map rồi sửa phải ghi lại nếu muốn cập nhật map.
- [ ] Phân biệt `Product`, `*Product`, `&product`, `*pointer` và `nil`.
- [ ] Phân biệt function, method, value receiver và pointer receiver.
- [ ] Giải thích được interface qua ví dụ `Named`.
- [ ] Biết `nil` nghĩa là thành công khi trả `error` và xử lý lỗi trước bước tiếp theo.
- [ ] Dự đoán được cả thứ tự defer và giá trị đối số được ghi nhận.
- [ ] Tự hoàn thành bài quản lý kho mà không làm thay đổi dữ liệu khi thao tác thất bại.

Nếu còn vướng mục nào, quay về bài tương ứng và tự viết một ví dụ ngắn để kiểm chứng. Đọc hiểu kết quả thực tế quan trọng hơn việc nhớ tên cú pháp.

## 7. Hướng học tiếp

Sau khi hoàn thành checklist, có thể mở rộng theo thứ tự sau:

| Thứ tự | Chủ đề tiếp theo | Mục tiêu thực hành |
|---|---|---|
| 1 | Package và tổ chức nhiều file | Tách dữ liệu sản phẩm, nghiệp vụ và phần chạy chương trình |
| 2 | Testing với `testing` | Viết kiểm thử ranh giới giảm giá và quy tắc tồn kho |
| 3 | File và JSON | Lưu, đọc danh sách sản phẩm; áp dụng xử lý lỗi và defer |
| 4 | HTTP bằng `net/http` | Nhận request và trả response JSON cho sản phẩm |
| 5 | Gin | Tổ chức route, đọc path/query, bind dữ liệu và xử lý lỗi API |
| 6 | Context và goroutine | Hiểu hủy tác vụ, thời hạn và truy cập dữ liệu đồng thời |

Phần nền tảng được chốt khi bạn nối được toàn bộ luồng: **dữ liệu đầu vào → kiểm tra → tính toán → cập nhật đúng đối tượng → xử lý lỗi → xuất kết quả**.
