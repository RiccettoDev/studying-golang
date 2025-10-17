package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // o defer espera tudo ser executado e depois executa algo por último, no caso aqui o cancel()

	go func() {
		//time.Sleep(10 * time.Second) quando roda mais de 5 segundos a reserva funciona
		time.Sleep(3 * time.Second) // quando leva menos de 5 segundos a reserva é cancelada
		cancel()
	}()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Tempo excedido para reservar o quarto")
	case <-time.After(5 * time.Second):
		fmt.Println("Quarto reservado com sucesso!")
	}
}
