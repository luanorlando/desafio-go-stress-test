# --- Stage 1: Build do Binário ---
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código fonte
COPY . .

# Compila o binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -o load-tester main.go

# --- Stage 2: Imagem Final de Execução ---
FROM alpine:latest

WORKDIR /root/

# Copia o binário compilado do stage anterior
COPY --from=builder /app/load-tester .

# Define o binário como ponto de entrada padrão
ENTRYPOINT ["./load-tester"]