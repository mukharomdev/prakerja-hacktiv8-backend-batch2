GORM, yang merupakan ORM (Object-Relational Mapping) untuk Go, mendukung beberapa struktur data yang umum digunakan dalam pengembangan aplikasi basis data relasional. Berikut adalah beberapa struktur data yang biasa digunakan di GORM:

1. **Model**: Model adalah representasi dari tabel dalam basis data relasional. Setiap model dalam GORM biasanya dipetakan ke sebuah tabel dalam basis data. Setiap instance dari model ini mewakili satu baris dalam tabel.

2. **Field/Column**: Setiap atribut dalam model GORM mewakili sebuah kolom dalam tabel basis data. Kolom ini dipetakan ke field-field dalam struct Go yang digunakan untuk mendefinisikan model.

3. **Primary Key (PK)**: Primary key adalah kolom (atau kombinasi kolom) yang unik mengidentifikasi setiap baris dalam sebuah tabel. Dalam GORM, primary key biasanya didefinisikan dengan menggunakan tag `gorm:"primaryKey"` pada field-field struct.

4. **Foreign Key (FK)**: Foreign key adalah kolom dalam sebuah tabel yang mengacu pada primary key dari tabel lain. Dalam GORM, hubungan antar tabel didefinisikan menggunakan hubungan antar model, di mana field yang mengacu pada primary key model lain diberi tag `gorm:"foreignKey"`.

5. **Index**: Index digunakan untuk meningkatkan kinerja operasi pencarian dan pengurutan data dalam tabel. Dalam GORM, Anda dapat membuat indeks menggunakan tag `gorm:"index"` pada field-field struct yang ingin Anda indeks.

6. **Unique Constraint**: Unique constraint memastikan bahwa setiap nilai dalam kolom tertentu adalah unik di dalam tabel. Dalam GORM, Anda dapat menetapkan unique constraint dengan menggunakan tag `gorm:"unique"` pada field-field struct yang ingin Anda jadikan unik.

7. **Default Value**: Default value adalah nilai yang diberikan secara otomatis ke sebuah kolom jika tidak ada nilai yang diberikan saat menambahkan data baru. Dalam GORM, Anda dapat menetapkan default value dengan menggunakan tag `gorm:"default:..."` pada field-field struct yang ingin Anda beri nilai default.

8. **Auto Increment**: Auto increment adalah cara untuk secara otomatis menambahkan nilai ke primary key kolom yang memiliki tipe data numerik setiap kali baris baru ditambahkan ke dalam tabel. Dalam GORM, Anda dapat menetapkan auto increment dengan menggunakan tag `gorm:"autoIncrement"` pada field-field struct yang menjadi primary key.

Struktur data di atas merupakan beberapa konsep utama yang digunakan dalam GORM untuk mendefinisikan dan memanipulasi struktur basis data relasional menggunakan Go. Dengan menggunakan GORM, Anda dapat dengan mudah memodelkan dan berinteraksi dengan basis data menggunakan representasi objek dalam Go.

# chapter 2
GORM adalah ORM (Object Relational Mapper) untuk bahasa pemrograman Go. GORM menyediakan abstraksi yang memudahkan pengembang untuk berinteraksi dengan database menggunakan objek-objek Go alih-alih SQL mentah. Dalam implementasinya, GORM menggunakan beberapa konsep struktur data yang penting. Berikut adalah beberapa konsep struktur data yang digunakan di GORM:

1. **Struct**:
   Struct adalah struktur data utama yang digunakan untuk merepresentasikan tabel dalam database. Setiap tabel biasanya diwakili oleh satu struct, dan kolom-kolom dalam tabel diwakili oleh field-field dalam struct tersebut. Struct ini juga digunakan untuk mendefinisikan relasi antar tabel, seperti one-to-many atau many-to-many.

   Contoh:
   ```go
   type User struct {
       ID        uint
       Name      string
       Email     string
       Orders    []Order // One-to-many relationship
   }

   type Order struct {
       ID        uint
       UserID    uint
       Amount    float64
   }
   ```

2. **Map**:
   GORM menggunakan map untuk menyimpan pengaturan konfigurasi dan nilai-nilai dinamis lainnya. Misalnya, saat membuat query, pengembang bisa menggunakan map untuk menyimpan pasangan kunci-nilai yang merepresentasikan kondisi query.

   Contoh:
   ```go
   var users []User
   db.Where(map[string]interface{}{"name": "John", "email": "john@example.com"}).Find(&users)
   ```

3. **Slice**:
   Slice digunakan untuk menyimpan kumpulan record hasil query atau untuk mendefinisikan relasi banyak-ke-banyak. Slice juga sering digunakan untuk mengelola daftar operasi batch, seperti penyimpanan atau penghapusan batch.

   Contoh:
   ```go
   var users []User
   db.Find(&users)
   ```

4. **Context**:
   GORM mendukung penggunaan `context.Context` dari standar library Go untuk mengelola timeout, cancelation, dan passing value antar layer aplikasi. Ini sangat berguna dalam pengelolaan request yang mungkin memerlukan pembatalan atau pengaturan waktu maksimal.

   Contoh:
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
   defer cancel()

   var user User
   db.WithContext(ctx).First(&user, 1)
   ```

5. **Interface**:
   GORM menggunakan interface untuk mendefinisikan kontrak atau perilaku yang diimplementasikan oleh struct. Misalnya, GORM menyediakan interface untuk hook lifecycle seperti `BeforeCreate`, `AfterCreate`, dll., yang dapat diimplementasikan oleh struct untuk menjalankan logika tertentu sebelum atau sesudah operasi database.

   Contoh:
   ```go
   func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
       u.CreatedAt = time.Now()
       return
   }
   ```

Dengan memanfaatkan berbagai struktur data ini, GORM mampu menyediakan fungsionalitas ORM yang fleksibel dan kuat, memungkinkan pengembang untuk mengelola data dan relasi database dengan mudah.