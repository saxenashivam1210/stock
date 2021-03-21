package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Type int32

const (
	TypeBuy  Type = 0
	TypeSell Type = 1
)

type TradeOrders struct {
	BuyOrders []Order
	SellOrder []Order
}
type Order struct{
	Id string
	Type Type
	Quantity int64
	Rate float64
	Traded bool
}

func main() {
	in, err := ioutil.ReadFile("/Users/shivamsaxena/test/main/input.txt")
	if err != nil {
		log.Println("File reading error", err)
		return
	}

	buyOrders := make([]Order, 0)
	sellOrders := make([]Order, 0)
	input := TradeOrders{
		BuyOrders: buyOrders,
		SellOrder: sellOrders,
	}

	for _, d := range strings.Split(string(in), "\n"){
		log.Println("data: ",d)
		s := strings.Split(d, " ")
		var t Type
		if s[3] == "buy"{
			t = TypeBuy
		}else{
			t = TypeSell
		}
		var r string
		if s[len(s)-2] == ""{
			r = s[len(s)-3]
		}else{
			r = s[len(s)-2]
		}
		rate, _ := strconv.ParseFloat(r , 64)
		quantity, _ := strconv.ParseInt(s[len(s)-1], 10, 64)
		order := Order{
			Id:       s[0],
			Type:     t,
			Rate:     rate,
			Quantity:  quantity,
			Traded:   false,
		}
		if t == TypeSell {
			input.SellOrder = append(input.SellOrder, order)
		} else {
			input.BuyOrders = append(input.BuyOrders, order)
		}

	}
	input.trade()

}

func (t *TradeOrders) trade(){
	tradeOpen := true
	for tradeOpen{
		for j, sell := range t.SellOrder{
			for i, buy :=range t.BuyOrders{
				var quantity int64
				if buy.Rate >= sell.Rate && !t.BuyOrders[i].Traded && !t.SellOrder[j].Traded{
					if buy.Quantity > sell.Quantity{
						quantity = t.SellOrder[j].Quantity
						t.BuyOrders[i].Quantity = t.BuyOrders[i].Quantity - t.SellOrder[j].Quantity
						t.SellOrder[j].Quantity = 0
						t.SellOrder[j].Traded = true
					}else if buy.Quantity < sell.Quantity{
						quantity = t.BuyOrders[i].Quantity
						t.SellOrder[j].Quantity = t.SellOrder[j].Quantity - t.BuyOrders[i].Quantity
						t.BuyOrders[i].Quantity = 0
						t.BuyOrders[i].Traded = true
					}else{
						quantity = t.SellOrder[j].Quantity
						t.SellOrder[j].Quantity = 0
						t.BuyOrders[i].Quantity = 0
						t.BuyOrders[i].Traded = true
						t.SellOrder[j].Traded = true
					}
					log.Println(buy.Id, sell.Rate, quantity, sell.Id)
					break
				}else{

					if i == len(t.BuyOrders) - 1 && j == len(t.SellOrder) - 1{
						tradeOpen = false
					}
				}
			}
		}
	}
}