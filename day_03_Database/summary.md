# **Aqua Developer Batch 2 Day 03**

Tanggal: 24 Agustus 2022

Materi: Relational Database

## **Definisi**
Relational database adalah database yang berbentuk tabel dimana satu tabel bisa jadi berelasi dengan table lain. Tabel terdiri atas column dan row, dimana tiap column berisi informasi/data yang penting

## **RDBMS**
Relational Database Management System merupakan tools yang digunakan untuk memudahkan proses relational database, contoh dari RDBMS ini ada beberapa, misalkan: PostgreSQL, MariaDB, Oracle, dll.

## **DDL (Data Definition Language)**
Berguna untuk mendefinisikan struktur database, table, schema, constraint, dll. Terdapat beberapa syntax yang dapat dipakai pada DDL, antara lain yaitu:
```
1.  CREATE -> digunakan untuk generate bisa database, table, dll
    contoh: CREATE TABLE A (id int, name char(50));

2.  ALTER -> digunakan untuk mengubah struktur yang sudah ada
    contoh: ALTER TABLE A ADD COLUMN B TYPE INT;

3.  DROP -> digunakan untuk menghapus database, tabel, dll
    contoh: DROP FROM A WHERE ID = 1;
```

## **DML (Data Manipulation Language)**
Berguna untuk melakukan manipulasi data dalam tabel, misalkan menambahkan data, mengupdate data, hingga menghapus data. Berikut syntax yang sering digunakan dalam DML, yaitu:
```
1.  SELECT -> digunakan untuk mendapatkan data
    contoh: SELECT * FROM A;

2.  INSERT -> digunakan untuk menambahkan data
    contoh: INSERT INTO A (name) VALUES ('halo');

3.  UPDATE -> digunakan untuk mengupdate data yang sudah ada
    contoh: UPDATE A SET name = 'world' WHERE ID = 1;

4.  DELETE -> digunakan untuk menghapus data dalam tabel
    contoh: DELETE FROM A WHERE ID = 1;
```

## **DCL (Data Control Language)**
Secara sederhana digunakan untuk memberikan control terhadap suatu user terhadap aksesnya menggunakan database. Untuk syntax-nya ada ```GRANT, REVOKE```, dll

## **JOIN**
Digunakan untuk melakukan penggabungan beberapa tabel minimal 2, terdapat beberapa jenis join, antara lain ```INNER JOIN, LEFT JOIN, RIGHT JOIN, FULL JOIN```, dll namun itu tadi yang sering dipakai dalam keseharian. Untuk ```INNER JOIN``` digunakan untuk mengambil data yang beririsan dati 2 tabel, ```LEFT JOIN``` digunakan untuk mengambil data beririsan dan data yang ada di sebelah kiri, ```RIGHT JOIN``` digunakan untuk mengambil data beririsan dan data yang ada di sebelah kanan.
```
Contoh penggunaan

SELECT *
FROM A
    JOIN B 
        ON A.b_id = B.id;
```

## **AGGREGATION**
Digunakan untuk mengembalikan sebuah result dari beberapa nilai dalam sebuah tabel, biasa dalam column. Ada beberapa yang sering digunakan antara lain ```SUM, COUNT, AVG, MAX, dan MIN```

## **SUBQUERY**
Digunakan untuk menuliskan sebuah query dalam query yang lain, tujuannya adalah untuk mempersingkat proses pengambilan data.
```
contoh

SELECT subquery.name
FROM (SELECT name FROM A) AS subquery;
```

## **FUNCTION**
Digunakan untuk memudahkan/mengoptimalkan penggunaan query yang panjang. Caranya dengan menjadikan query panjang tersebut kedalam sebuah function, sehingga ketika ingin menggunakan query tersebut tidak perlu menulis ulang query, melainkan tinggal memanggil function yang telah didefinisikan.
```
contoh

CREATE FUNCTION get_name(CHAR) RETURNS A AS
'SELECT id, name, phone FROM A WHERE name = $1'
LANGUAGE 'sql';
```
