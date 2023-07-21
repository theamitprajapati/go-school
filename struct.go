package main

import "fmt"

type config struct {
	host string
	port int
}

func setting(ht string) *config {

	p := config{host: ht}
	p.port = 8956
	return &p
}
func main() {

	fmt.Println(config{"localhost", 25052})
	fmt.Println(config{"k", 001000})
	fmt.Println(&config{"data", 001000})

	x := config{"local", 8080}

	fmt.Println(x.port)

	y := &x

	y.host = "Google.com"

	fmt.Println(y.host)

	fmt.Println(x.host)

	direct := struct {
		name string
		age  int
	}{
		"amit",
		33,
	}

	fmt.Println(direct)

}
