package main

import (
	"colectivo"
	"fmt"
)

func main() {
	pub := colectivo.Publisher{
		Subs: []colectivo.Subscriber{colectivo.WithStdOut()},
	}

	if err := pub.Publish(colectivo.Event{
		Bytes: []byte("hello\n"),
	}); err != nil {
		fmt.Println(err)
	}

	test(colectivo.Optional{V: nil})
}

func test(o colectivo.Optional) {
	fmt.Println(o.OrElse("."))
}
