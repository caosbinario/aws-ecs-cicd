# Imagen base
FROM golang:1.21.3

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente al contenedor
COPY . .

# Instala las dependencias y compila el programa
RUN go build -o server .

# Exponer el puerto 8080
EXPOSE 8080

# Comando para ejecutar la API
CMD ["./server"]
