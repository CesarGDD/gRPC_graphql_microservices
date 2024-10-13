# Usa la imagen base de Go
FROM golang:1.20-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo main.go en el contenedor
COPY main.go .

# Compila la aplicación Go
RUN go build -o app

# Expone el puerto 8080
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./app"]
