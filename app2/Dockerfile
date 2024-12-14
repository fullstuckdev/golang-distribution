# Menggunakan image golang versi 1.19 dengan base image alpine Linux yang ringan
FROM golang:1.19-alpine

# Menetapkan direktori kerja di dalam container menjadi /app
WORKDIR /app

# Menyalin seluruh file dari direktori saat ini (.) ke direktori kerja container (.)
COPY . .

# Menginisialisasi modul Go dengan nama 'app2'
# Ini akan membuat file go.mod yang diperlukan untuk manajemen dependensi
RUN go mod init app2

# Melakukan kompilasi kDockerode Go menjadi binary executable bernama 'main'
RUN go build -o main .

# Menjalankan binary executable 'main' saat container dijalankan
CMD ["./main"]