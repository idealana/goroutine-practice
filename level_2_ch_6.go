// LEVEL 2 - CHAPTER 6 - Concurrent Bank Transaction

package main

import (
	"errors"
	"fmt"
	"sync"
)

type account struct {
	ID int
	Name string
	Balance int64
	mu sync.Mutex
}

func (acc *account) Transfer(target *account, nominal int64) error {
	if target == nil {
    	return errors.New("target account is required")
    }
    
    if acc == target {
    	return errors.New("cannot transfer to the same account")
    }
    
    if nominal <= 0 {
    	return errors.New("nominal must be greater than zero")
    }
    
    // Lock accounts in a consistent order to avoid deadlock.
    first, second := acc, target
    
    if first.ID > second.ID {
    	first, second = second, first
    }
    
    first.mu.Lock()
    second.mu.Lock()
    
    defer first.mu.Unlock()
    defer second.mu.Unlock()
    
    if acc.Balance < nominal {
		return errors.New("insufficient balance")
	}
    
    acc.Balance -= nominal
    target.Balance += nominal
    
    return nil
}

func (acc *account) GetBalance() int64 {
	acc.mu.Lock()
    defer acc.mu.Unlock()
    
    return acc.Balance
}

func main() {
	accounts := []*account{
    	{ID: 1, Name: "Account 1", Balance: 10000000},
        {ID: 2, Name: "Account 2", Balance: 10000000},
        {ID: 3, Name: "Account 3", Balance: 10000000},
        {ID: 4, Name: "Account 4", Balance: 10000000},
        {ID: 5, Name: "Account 5", Balance: 10000000},
    }
    
    transfers := []struct{
    	source *account
        target *account
        amount int64
    }{
    	{accounts[0], accounts[1], 1000000},
        {accounts[0], accounts[2], 1000000},
        {accounts[1], accounts[2], 2000000},
        {accounts[2], accounts[3], 500000},
        {accounts[3], accounts[4], 5000000},
        {accounts[4], accounts[0], 7000000},
        {accounts[4], accounts[1], 10000000},
        {accounts[1], accounts[2], 1000000},
        {accounts[2], accounts[3], 1000000},
        {accounts[0], accounts[1], 1000000},
        {accounts[0], accounts[2], 1000000},
        {accounts[1], accounts[2], 2000000},
        {accounts[2], accounts[3], 500000},
        {accounts[3], accounts[4], 5000000},
    }
	
	var wg sync.WaitGroup
	wg.Add(len(transfers))
    
    for _, transfer := range transfers {
    	go func() {
          defer wg.Done()
          
          err := transfer.source.Transfer(transfer.target, transfer.amount)
          
          if err != nil {
          	fmt.Printf(
              "%s transfer failed: %v\n",
              transfer.source.Name,
              err,
            )
            return
          }
          
          fmt.Printf(
            "%s transferred %d to %s\n",
            transfer.source.Name,
            transfer.amount,
            transfer.target.Name,
          )
      }()
    }
	
	wg.Wait()
	
	fmt.Println("-- Result --")
    
    var totalBalance int64
    
    for _, acc := range accounts {
    	balance := acc.GetBalance()
        
        fmt.Println(acc.Name + ":", balance)
        
        totalBalance += balance
    }
    
	fmt.Println("Total Balance:", totalBalance)
}
