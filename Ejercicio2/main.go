package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func mostrarProducto(producto Producto) {
	fmt.Println("Nombre:", producto.nombre)
	fmt.Println("Precio:", producto.precio)
	fmt.Println("Cantidad:", producto.cantidad)
	fmt.Println("Valor total:", producto.precio*float64(producto.cantidad))

	if producto.cantidad > 0 {
		fmt.Println("Hay stock disponible")
	} else {
		fmt.Println("No hay stock disponible")
	}

	fmt.Println("-------------------------")
}

func main() {
	producto1 := Producto{"Laptop", 500, 3}
	producto2 := Producto{"Mouse", 15, 10}
	producto3 := Producto{"Teclado", 25, 0}

	fmt.Println("LISTA DE PRODUCTOS")
	fmt.Println("=========================")

	mostrarProducto(producto1)
	mostrarProducto(producto2)
	mostrarProducto(producto3)
}
