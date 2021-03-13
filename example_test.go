package vclutil_test

import (
	"fmt"
	"math"

	vclutil "github.com/John-Tonny/vclsuite_vclutil"
)

func ExampleAmount() {

	a := vclutil.Amount(0)
	fmt.Println("Zero Satoshi:", a)

	a = vclutil.Amount(1e8)
	fmt.Println("100,000,000 Satoshis:", a)

	a = vclutil.Amount(1e5)
	fmt.Println("100,000 Satoshis:", a)
	// Output:
	// Zero Satoshi: 0 BTC
	// 100,000,000 Satoshis: 1 BTC
	// 100,000 Satoshis: 0.001 BTC
}

func ExampleNewAmount() {
	amountOne, err := vclutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := vclutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := vclutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := vclutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 BTC
	// 0.01234567 BTC
	// 0 BTC
	// invalid bitcoin amount
}

func ExampleAmount_unitConversions() {
	amount := vclutil.Amount(44433322211100)

	fmt.Println("Satoshi to kBTC:", amount.Format(vclutil.AmountKiloBTC))
	fmt.Println("Satoshi to BTC:", amount)
	fmt.Println("Satoshi to MilliBTC:", amount.Format(vclutil.AmountMilliBTC))
	fmt.Println("Satoshi to MicroBTC:", amount.Format(vclutil.AmountMicroBTC))
	fmt.Println("Satoshi to Satoshi:", amount.Format(vclutil.AmountSatoshi))

	// Output:
	// Satoshi to kBTC: 444.333222111 kBTC
	// Satoshi to BTC: 444333.222111 BTC
	// Satoshi to MilliBTC: 444333222.111 mBTC
	// Satoshi to MicroBTC: 444333222111 μBTC
	// Satoshi to Satoshi: 44433322211100 Satoshi
}
