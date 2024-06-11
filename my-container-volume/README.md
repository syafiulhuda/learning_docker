Implementasi volume pada docker

dengan volume: -v my-postgres-admin:/var/lib/postgresql/data

1. buat container dengan volume diatas
2. buat table di pgadmin dengan ketentuan yang berlaku
3. hapus container, kemudian buat container baru tapi dengan penamaan yang berbeda
4. cek apakah table di pgadmin dapat terkoneksi dengan container kita
