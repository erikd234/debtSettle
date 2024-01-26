package main

import (
	"fmt"
	"math"
	"sort"
)

type player struct {
	name   string
	amount float32 // +/- amount owed or owing
}

type payment struct {
	from   string
	to     string
	amount float32
}

func isCloseToZero(x float32) bool {
	epsilon := 0.1
	return math.Abs(float64(x)) <= epsilon
}

func settleDebt(players []player) {
	payments := []payment{}
	var sendAmount float32
	sort.Slice(players, func(i, j int) bool {
		return players[i].amount > players[j].amount
	})
	fmt.Println(players)
	hIdx := 0                //histhest not paid index
	lIdx := len(players) - 1 // Lowest current playing idx
	var pAmount float32      // post transaxtion amoutn
	for hIdx <= lIdx {
		// loser try to pay winner
		winner := &(players[hIdx])
		loser := &(players[lIdx])
		pAmount = winner.amount + loser.amount
		var p payment
		if isCloseToZero(pAmount) {
			// This means both winner and loser perfectly paid off
			sendAmount = loser.amount
			winner.amount = 0
			loser.amount = 0
			// Adjust amount to be 0 in the original array
			p = payment{from: loser.name, to: winner.name, amount: sendAmount}
			hIdx++
			lIdx--
		} else if pAmount > 0 {
			// Loser is now fully paid, winner
			// needs more from another loser
			sendAmount = loser.amount
			// Here we will modify the amount the winner left to be paid
			winner.amount = winner.amount + sendAmount
			p = payment{from: loser.name, to: winner.name, amount: sendAmount}
			lIdx--
		} else if pAmount < 0 {
			// Loser has overpaid winner,
			// give the loser pack
			// amount amount below 0
			// At this point we can increment
			// the winner, keep the loser the same
			sendAmount = loser.amount - pAmount // mental check like -7.34 - - 2.34 = 5 . this means 5 should be sent loser amount would be 2.34.
			loser.amount = pAmount
			p = payment{from: loser.name, to: winner.name, amount: sendAmount}
			hIdx++
		}
		payments = append(payments, p)
	}
	fmt.Println(players)
	fmt.Println(payments)
	printPayments(payments)
}

func printPayments(ps []payment) {
        fmt.Println()
	for _, p := range ps {
		amount := fmt.Sprintf("%.2f", math.Abs(float64(p.amount)))
		fmt.Println(p.from, " pays $", amount, " to ", p.to)
	}
        fmt.Println()
}

func main() {
	test_players := []player{
		{"Alice", 10.0},
		{"Bob", -5.0},
		{"Charlie", -5.0},
	}
	fmt.Println("Test1")
	settleDebt(test_players)
	// Test Case 1: Simple case with two players
	test_players1 := []player{
		{"Dave", 20.0},
		{"Eve", -20.0},
	}
	settleDebt(test_players1)

	fmt.Println("Test2")
	// Test Case 2: Multiple players with no debt
	test_players2 := []player{
		{"Fred", 0.0},
		{"Gina", 0.0},
		{"Hank", 0.0},
	}
	settleDebt(test_players2)

	fmt.Println("Test3")
	// Test Case 3: Complex case with multiple debts
	test_players3 := []player{
		{"Ivy", 15.0},
		{"Hellmuth", 3.33},
		{"dad", 3.33},
		{"Mom", 3.33},
		{"John", -10.0},
		{"Marry", -10.0},
		{"Karen", 5.0},
		{"Leo", -10.0},
	}
	settleDebt(test_players3)

	fmt.Println("Test3")
	// Test Case 4: Case with large debts and credits
	test_players4 := []player{
		{"Mia", 100.0},
		{"Nick", -30.0},
		{"Oscar", -70.0},
	}
	settleDebt(test_players4)
	fmt.Println("Test5")
	// Test Case 5: Case with floating point precision
	test_players5 := []player{
		{"TrevDon'tFold", 60.05},
		{"Jordan J", 38.58},
		{"Matt Frost", 34.69},
		{"a-dawg", -4.96},
		{"jordan", -10},
		{"Young", -10},
		{"matoosh", -20},
		{"Densi", -40},
		{"trev", -48.10},
	}
	settleDebt(test_players5)
}
