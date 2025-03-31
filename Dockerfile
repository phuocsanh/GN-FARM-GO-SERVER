FROM golang:alpine

# Cài đặt các công cụ cần thiết
RUN apk add --no-cache git

# Thiết lập thư mục làm việc
WORKDIR /app

# Copy toàn bộ mã nguồn vào container
COPY . .

# Tải xuống các dependencies của Go
RUN go mod download

# Cài đặt Air
RUN go install github.com/air-verse/air@latest

# Expose port (nếu ứng dụng của bạn chạy trên một port cụ thể)
EXPOSE 8002

# Sử dụng lệnh air để phát triển mà không cần build lại toàn bộ
CMD ["/go/bin/air", "-c", ".air.toml"]
