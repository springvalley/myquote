package myquote

import (
	"fmt"

	"rsc.io/quote"
)

func QuoteHello(){
	fmt.Println(quote.Hello())
}
func QuoteGlass(){
	fmt.Println(quote.Glass())
}
func QuoteGo(){
	fmt.Println(quote.Go())
}
func QuoteOpt(){
	fmt.Println(quote.Opt())
}

func QuoteAll(){
	QuoteHello()
	QuoteGlass()
	QuoteGo()
	QuoteOpt()
}
