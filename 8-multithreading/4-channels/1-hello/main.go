package main

// Thread 1
func main() {
	canal := make(chan string) // cria um canal (vazio)

	// Thread 2
	go func() {
		canal <- "Hello World!" // a mensagem é enviada para o canal (canal bloqueado)
	}()

	// Thread 1
	msg := <-canal // a mensagem é recebida do canal (esvaziando o canal)
	println(msg)
}
