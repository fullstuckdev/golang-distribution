# Menggunakan image golang versi 1.19 dengan base image alpine Linux yang ringan
FROM golang:1.19-alpine

# Menetapkan direktori kerja di dalam container menjadi /app
WORKDIR /app

# Menyalin seluruh file dari direktori saat ini (.) ke direktori kerja container (.)
COPY . .

RUN go mod init app2

RUN go build -o main .

CMD ["./main"]