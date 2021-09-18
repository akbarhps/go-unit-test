# Go Unit Test

Sumber Tutorial:
[Youtube](https://www.youtube.com/watch?v=t9QJPE5vwhs) |
[Slide](https://docs.google.com/presentation/d/1XxMEaA-JsPHr9BUw2oIOPlEL_psI3EaUFUpuvdlDB_Q/edit#slide=id.p)

## Testing Package

---

- Di bahasa pemrograman lain, biasanya untuk implementasi unit test, kita butuh library atau framework khusus untuk unit test
- Berbeda dengan Go-Lang, di Go-Lang sudah untuk unit test sudah disediakan sebuah package khusus bernama testing
- Selain itu untuk menjalankan unit test, di Go-Lang juga sudah disediakan perintah nya
- Hal ini membuat implementasi unit testing di golang sangat mudah dibanding dengan bahasa pemrograman yang lain
- https://golang.org/pkg/testing/

## testing.T

---

- Go-Lang menyediakan sebuah struct yang bernama `testing.T`
- Struct ini digunakan untuk unit test di Go-Lang

## testing.M

---

- `testing.M` adalah struct yang disediakan Go-Lang untuk mengatur life cycle testing
- Materi ini nanti akan kita bahas di chapter Main

## testing.B

---

- `testing.B` adalah struct yang disediakan Go-Lang untuk melakukan benchmarking
- Di Go-Lang untuk melakukan benchmark (mengukur kecepatan kode program) pun sudah disediakan, sehingga kita tidak perlu menggunakan library atau framework tambahan

## Aturan File Test

---

- Go-Lang memiliki aturan cara membuat file khusus untuk unit test
- Untuk membuat file unit test, kita harus menggunakan akhiran _test
- Jadi kita misal kita membuat file `hello_world.go`, artinya untuk membuat unit testnya, kita harus membuat file `hello_world_test.go`

## Aturan Function Test

---

- Selain aturan nama file, di Go-Lang juga sudah diatur untuk nama function unit test
- Nama function untuk unit test harus diawali dengan nama `Test`
- Misal jika kita ingin mengetest function `HelloWorld`, maka kita akan membuat function unit test dengan nama `TestHelloWorld`
- Tak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan di test, yang penting adalah harus diawalin dengan kata `Test`
- Selanjutnya harus memiliki parameter `(t *testing.T)` dan tidak mengembalikan return value

## Menjalankan Unit Test

---

- Untuk menjalankan unit test kita bisa menggunakan perintah :  
`go test`
- Jika kita ingin lihat lebih detail function test apa saja yang sudah di running, kita bisa gunakan perintah :   
`go test -v`
- Dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah :  
`go test -v -run TestNamaFunction`

## Menjalankan Semua Unit Test

---

- Jika kita ingin menjalankan semua unit test dari top folder module nya, kita bisa gunakan perintah :  
`go test ./...`

## Menggagalkan Unit Test

---

- Menggagalkan unit test menggunakan panic bukanlah hal yang bagus
- Go-Lang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan `testing.T`
- Terdapat function `t.Fail()`, `t.FailNow()`, `t.Error()` dan `t.Fatal()` jika kita ingin menggagalkan unit test

## t.Fail() dan t.FailNow()

---

- Terdapat dua function untuk menggagalkan unit test, yaitu `t.Fail()` dan `t.FailNow()`. Lantas apa bedanya?
- `t.Fail()` akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal
- `t.FailNow()` akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test

```go
func TestSum(t *testing.T) {
  expect := 5 + 5
  result := Sum(5, 5)

  if result != expect {
    t.Fail()
    t.FailNow()
  }
}
```

## t.Error(args...) dan t.Fatal(args...)

---

- Selain `t.Fail()` dan `t.FailNow()`, ada juga `t.Error()` dan `t.Fatal()`
- `t.Error()` function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil function `t.Fail()`, sehingga mengakibatkan unit test dianggap gagal
- Namun karena hanya memanggil `t.Fail()`, artinya eksekusi unit test akan tetap berjalan sampai selesai
- `t.Fatal()` mirip dengan `t.Error()`, hanya saja, setelah melakukan log error, dia akan memanggil `t.FailNow()`, sehingga mengakibatkan eksekusi unit test berhenti

```go
func TestSum(t *testing.T) {
  expect := 5 + 5
  result := Sum(5, 5)

  if result != expect {
    t.Error("Expecting", expect, ", got:", result)
    t.Fatal("Expecting", expect, ", got:", result)
  }
}
```

## Assertion

---

- Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan
- Apalagi jika result data yang harus di cek itu banyak
- Oleh karena itu, sangat disarankan untuk menggunakan `Assertion` untuk melakukan pengecekan
- Sayangnya, Go-Lang tidak menyediakan package untuk assertion, sehingga kita butuh menambahkan library untuk melakukan assertion ini

## Testify

---

- Salah satu library assertion yang paling populer di Go-Lang adalah Testify
- Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test
- https://github.com/stretchr/testify 
- Kita bisa menambahkannya di Go module kita :
`go get github.com/stretchr/testify`

## assert vs require

---

- Testify menyediakan dua package untuk assertion, yaitu `assert` dan `require`
- Saat kita menggunakan `assert`, jika pengecekan gagal, maka assert akan memanggil `t.Fail()`, artinya eksekusi unit test akan tetap dilanjutkan
- Sedangkan jika kita menggunakan `require`, jika pengecekan gagal, maka require akan memanggil `t.FailNow()`, artinya eksekusi unit test tidak akan dilanjutkan

```go
assert.Equal(t, expect, result, "result is not the same as expected")

require.Equal(t, expect, result, "result is not the same as expected")
```

## Skip Test

---

- Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
- Di Go-Lang juga kita bisa membatalkan eksekusi unit test jika kita mau
- Untuk membatalkan unit test kita bisa menggunakan function `t.Skip()`

```go
if runtime.GOOS == "windows" {
  t.Skip("this test can't run on windows")
}
  ```

## Before dan After Test

---

- Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi
- Jikalau kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka membuat manual di unit test function nya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya
- Untungnya di Go-Lang terdapat fitur yang bernama testing.M
- Fitur ini bernama Main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan untuk melakukan Before dan After di unit test

## testing.M

---

- Untuk mengatur ekeskusi unit test, kita cukup membuat sebuah function bernama TestMain dengan parameter testing.M
- Jika terdapat function TestMain tersebut, maka secara otomatis Go-Lang akan mengeksekusi function ini tiap kali akan menjalankan unit test di sebuah package
- Dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau
- Ingat, function TestMain itu dieksekusi hanya sekali per Go-Lang package, bukan per tiap function unit test

```go
func TestMain(m *testing.M) {
  fmt.Println("Before test")

  m.Run() // run tests on this package

  fmt.Println("After test")
}
```

## Subtest

---

- Go-Lang mendukung fitur pembuatan function unit test di dalam function unit test
- Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman yang lainnya
- Untuk membuat sub test, kita bisa menggunakan function Run()

```go
func TestAll(t *testing.T) {
  t.Run("Test Sum Function", func(t *testing.T) {
    result := Sum(10, 10)
    assert.Equal(t, 20, result, "result is not the same as expected")
  })
  t.Run("Test Multiply Function", func(t *testing.T) {
    result := Multiply(10, 10)
    assert.Equal(t, 100, result, "result is not the same as expected")
  })
}
```

## Menjalankan Hanya Sub Test

---

- Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah :
`go test -run TestNamaFunction`
- Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah :
`go test -run TestNamaFunction/NamaSubTest`
- Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah :
`go test -run /NamaSubTest`

## Table Test

---

- Sebelumnya kita sudah belajar tentang sub test
- Jika diperhatikan, sebenarnya dengan sub test, kita bisa membuat test secara dinamis
- Dan fitur sub test ini, biasa digunaka oleh programmer Go-Lang untuk membuat test dengan konsep table test
- Table test yaitu dimana kita menyediakan data beruba slice yang berisi parameter dan ekspektasi hasil dari unit test
- Lalu slice tersebut kita iterasi menggunakan sub test

```go
func TestTableSum(t *testing.T) {
  table := []struct {
    TestName   string
    TestParam1 int
    TestParam2 int
    TestExpect int
  }{
    {
      TestName:   "TestSum(10+10)",
      TestParam1: 10,
      TestParam2: 10,
      TestExpect: 20,
    }, {
      TestName:   "TestSum(20+10)",
      TestParam1: 20,
      TestParam2: 10,
      TestExpect: 30,
    },
  }

  for _, test := range table {
    t.Run(test.TestName, func(t *testing.T) {
      result := Sum(test.TestParam1, test.TestParam2)
      assert.Equal(t, test.TestExpect, result, "result is not the same as expected")
    })
  }
}
```

## Mock

---

- Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data yang sudah kita program diawal
- Mock adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di testing
- Misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil API Call ke third party service. Hal ini sangat sulit untuk di test, karena unit testing kita harus selalu memanggil third party service, dan belum tentu response nya sesuai dengan apa yang kita mau
- Pada kasus seperti ini, cocok sekali untuk menggunakan mock object

## Testify Mock

---

- Untuk membuat mock object, tidak ada fitur bawaan Go-Lang, namun kita bisa menggunakan library testify yang sebelumnya kita gunakan untuk assertion
- Testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object
- Namun, perlu diperhatikan, jika desain kode program kita jelek, akan sulit untuk melakukan mocking, jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik
- Mari kita buat contoh kasus

## Aplikasi Query Ke Database

---

- Kita akan coba contoh kasus dengan membuat contoh aplikasi golang yang melakukan query ke database
- Dimana kita akan buat layer Service sebagai business logic, dan layer Repository sebagai jembatan ke database
- Agar kode kita mudah untuk di test, disarankan agar membuat kontrak berupa Interface

```go
var categoryRepository = &repository.CategoryRepositoryMock{
  Mock: mock.Mock{},
}

var categoryService = CategoryService{
  Repository: categoryRepository,
}

func TestCategoryService_GetSuccess(t *testing.T) {
  result := entity.Category{
    Id:   "1",
    Name: "Elektronik",
  }

  categoryRepository.Mock.On("FindById", "1").Return(result)
  category, err := categoryService.Get("1")

  assert.Nil(t, err)
  assert.NotNil(t, category)
  assert.Equal(t, &result, category)
}
```

## Benchmark

---

- Selain unit test, Go-Lang testing package juga mendukung melakukan benchmark
- Benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita
- Benchmark di Go-Lang dilakukan dengan cara secara otomatis melakukan iterasi kode yang kita panggil berkali-kali sampai waktu tertentu
- Kita tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh testing.B bawaan dari testing package

## testing.B

---

- `testing.B` adalah struct yang digunakan untuk melakukan benchmark. 
- `testing.B` mirip dengan `testing.T`, terdapat function `Fail()`, `FailNow()`, `Error()`, `Fatal()` dan lain-lain
- Yang membedakan, ada beberapa attribute dan function tambahan yang digunakan untuk melakukan benchmark
- Salah satunya adalah attribute N, ini digunakan untuk melakukan total iterasi sebuah benchmark

## Cara Kerja Benchmark

---

- Cara kerja benchmark di Go-Lang sangat sederhana
- Dimana kita hanya perlu membuat perulangan sejumlah N attribute
- Nanti secara otomatis Go-Lang akan melakukan eksekusi sejumlah perulangan yang ditentukan secara otomatis, lalu mendeteksi berapa lama proses tersebut berjalan, dan disimpulkan performa benchmark nya dalam waktu

## Benchmark Function

---

- Mirip seperti unit test, untuk benchmark pun, di Go-Lang sudah ditentukan nama function nya, harus diawali dengan kata Benchmark, misal BenchmarkHelloWorld, BenchmarkXxx
- Selain itu, harus memiliki parameter `(b *testing.B)`
- Dan tidak boleh mengembalikan return value
- Untuk nama file benchmark, sama seperti unit test, diakhiri dengan _test, misal `hello_world_test.go`

```go
func BenchmarkSum(b *testing.B) {
  for i := 1; i < b.N; i++ {
    Sum(10000000000, 100000000)
  }
}
``` 

## Menjalankan Benchmark

---

- Untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test, namun ditambahkan parameter bench :
`go test -v -bench=.`
- Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah :
`go test -v -run=NotMathUnitTest -bench=.`
- Kode diatas selain menjalankan benchmark, akan menjalankan unit test juga, jika kita hanya ingin menjalankan benchmark tertentu, kita bisa gunakan perintah :
`go test -v -run=NotMathUnitTest -bench=BenchmarkTest`
- Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah :
`go test -v -bench=./...`

## Sub Benchmark

---

- Sama seperti `testing.T`, di `testing.B` juga kita bisa membuat sub benchmark menggunakan function `Run()`

```go
func BenchmarkAll(b *testing.B) {
  b.Run("Benchmark Sum", func(b *testing.B) {
    for i := 1; i < b.N; i++ {
      Sum(10000000000, 100000000)
    }
  })

  b.Run("Benchmark Multiply", func(b *testing.B) {
    for i := 1; i < b.N; i++ {
      Multiply(10000000000, 100000000)
    }
  })
}
```

## Menjalankan Hanya Sub Benchmark

---

- Saat kita menjalankan benchmark function, maka semua sub benchmark akan berjalan
Namun jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah :
`go test -v -bench=BenchmarkNama/NamaSub`

## Table Benchmark

---

- Sama seperti di unit test, programmer Go-Lang terbiasa membuat table benchmark juga
- Ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data berbeda-beda tanpa harus membuat banyak benchmark function

```go
func BenchmarkMultiplyTable(b *testing.B) {
  table := []struct {
    TestName   string
    TestParam1 int
    TestParam2 int
    TestExpect int
  }{
    {
      TestName:   "TestMultiply(10*10)",
      TestParam1: 10,
      TestParam2: 10,
      TestExpect: 100,
    },
    {
      TestName:   "TestMultiply(20*10)",
      TestParam1: 20,
      TestParam2: 10,
      TestExpect: 200,
    },
  }

  for _, test := range table {
    b.Run(test.TestName, func(b *testing.B) {
      for i := 0; i < b.N; i++ {
        result := Multiply(test.TestParam1, test.TestParam2)
        assert.Equal(b, test.TestExpect, result, "result is not the same as expected")
      }
    })
  }
}
```