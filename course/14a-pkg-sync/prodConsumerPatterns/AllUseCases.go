package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

// These use cases from the course:  https://app.pluralsight.com/library/courses/go-programming-concurrent/table-of-contents

// Concurrency patterns
// Convers 4 patterns
// 1) 1-1 : Single Producer and Single Consumer
// 2) 1-* : Single Producer and Multiple Consumers
// 3) *-1 : Multiple Producers and Single Consumers
// 4) *-* : Multiple Producers and Multiple Consumers
func main() {
    // channelDemo()

    // without encapsulation
    //OrderProcess()
    // with encapsulation
    //OrderProcessEncapsulation()

    // 1) 1-1 : Single Producer and Single Consumer
    ConcurrencyPattern_OneProducer2OneConsumer()


    // 2) 1-* : Single Producer and Multiple Consumers
    ConcurrencyPattern_singleProducer2MultipleConsumbers()

    // 3) *-1 : Multiple Producers and Single Consumers
    ConcurrencyPattern_MultipleProducer2SingleConsumbers()

    // 4) *-* : Multiple Producers and Multiple Consumers
    ConcurrencyPattern_MultipleProducer2MultipleConsumbers()

}

func ConcurrencyPattern_MultipleProducer2MultipleConsumbers(){
    var wg sync.WaitGroup

    receivedOrderCh := receiveOrdersEncapsulateChannelObj()
    validOrderCh, invalidOrderCh := validateOrdersEncapsulationChannelObj(receivedOrderCh)

    // Here we have Multiple Produces in this 
    reservedInventoryCh := reserveInventory(validOrderCh)

    // Here we have multiple Consumbers, we need to pass wg as main need to wait until all consumbers are done
    fillOrdersWithMultipleConsumers(reservedInventoryCh, &wg)
   
    wg.Add(1)
    go func(invalidOrderCh <- chan invalidOrder){
        for or := range invalidOrderCh{
            fmt.Printf("Invalid order received: %v. Issue: %v\n", or.order, or.err)
        }
        wg.Done()
    }(invalidOrderCh)
   
    wg.Wait()
}

func fillOrdersWithMultipleConsumers(in <- chan order, wg *sync.WaitGroup)  {
    const workers = 3
    wg.Add(workers)
    for i:=0;i<workers;i++{
        go func(){
            for or := range in {
                or.Status = filled
                fmt.Printf("Order has been completed: %v \n", or)
            }
            wg.Done()
        }()
    }
}

func ConcurrencyPattern_MultipleProducer2SingleConsumbers(){
    var wg sync.WaitGroup

    receivedOrderCh := receiveOrdersEncapsulateChannelObj()
    validOrderCh, invalidOrderCh := validateOrdersEncapsulationChannelObj(receivedOrderCh)

    // Here we have Multiple Produces in this 
    reservedInventoryCh := reserveInventory(validOrderCh)
    fillOrdersCh := fillOrders(reservedInventoryCh)
   
    wg.Add(2)
    go func(invalidOrderCh <- chan invalidOrder){
        for or := range invalidOrderCh{
            fmt.Printf("Invalid order received: %v. Issue: %v\n", or.order, or.err)
        }
        wg.Done()
    }(invalidOrderCh)
    go func(fillOrdersCh <- chan order){
        for or := range fillOrdersCh{
            fmt.Printf("Order has been completed: %v \n", or)
        }
        wg.Done()
    }(fillOrdersCh)
    wg.Wait()
}



func fillOrders(in <- chan order) <- chan order{
    out := make(chan order)

    go func(){
        for or := range in {
            or.Status = filled
            out <- or
        }
        close(out)
    }()
    return out
}



func ConcurrencyPattern_singleProducer2MultipleConsumbers(){
    var wg sync.WaitGroup

    receivedOrderCh := receiveOrdersEncapsulateChannelObj()
    validOrderCh, invalidOrderCh := validateOrdersEncapsulationChannelObj(receivedOrderCh)
    reservedInventoryCh := reserveInventory(validOrderCh)

    // add multiple consumers instead of just one
    const workers = 3
    wg.Add(workers)
    for i:=0 ;i <workers; i++{
        go func (reservedInventoryCh <- chan order){
            for or:= range reservedInventoryCh{
                fmt.Printf("Inventory reserved for: %v\n", or)
            }
            wg.Done()
        }(reservedInventoryCh)
    }
    
    wg.Add(1)
    go func(invalidOrderCh <- chan invalidOrder){
        for or := range invalidOrderCh{
            fmt.Printf("Invalid order received: %v. Issue: %v\n", or.order, or.err)
        }
        wg.Done()
    }(invalidOrderCh)
    wg.Wait()
}

func ConcurrencyPattern_OneProducer2OneConsumer(){
    var wg sync.WaitGroup

    receivedOrderCh := receiveOrdersEncapsulateChannelObj()
    validOrderCh, invalidOrderCh := validateOrdersEncapsulationChannelObj(receivedOrderCh)
    reservedInventoryCh := reserveInventory(validOrderCh)
    wg.Add(2)
    go func (reservedInventoryCh <- chan order){
        for or:= range reservedInventoryCh{
            fmt.Printf("Inventory reserved for: %v\n", or)
        }
        wg.Done()
    }(reservedInventoryCh)
    go func(invalidOrderCh <- chan invalidOrder){
        for or := range invalidOrderCh{
            fmt.Printf("Invalid order received: %v. Issue: %v\n", or.order, or.err)
        }
        wg.Done()
    }(invalidOrderCh)
    wg.Wait()
}
 
func receiveOrdersEncapsulateChannelObj() (chan order) {
    out := make(chan order)
    go func(){
        for _, rawOrder := range rawOrders{
            var newOrder order
            fmt.Println("processing order:",rawOrder)
            err := json.Unmarshal([]byte(rawOrder),&newOrder)
            if err != nil{
                log.Print("Got err while unmarshaling order json::",err)
                continue
            }
            out <- newOrder
            //orders = append(orders,newOrder)
        }
        close(out)
    }()
  return out
}

func reserveInventory(in <- chan order) <- chan order{
    out := make (chan order)
    var wg sync.WaitGroup
    const workers = 3
    wg.Add(workers)
    for i:=0;i < workers;i++{
        go func(){
            for or := range in {
                or.Status = reserved
                out <- or
            }
            wg.Done() // runs for every worker
           // close(out) we can not close here
        }()
    }
    go func(){
        wg.Wait() // wait until all workers done
        close(out) // now close common out channel
    }()
    return out
}

 

func validateOrdersEncapsulationChannelObj(in <- chan order) ( chan order, chan invalidOrder){
    // order := <- in
    out := make(chan order)
    errCh := make(chan invalidOrder, 1)
    go func(){
        for order := range in{
            if order.Quantity <=0 {
                s := fmt.Sprintf("err: sending this order to invalidOrder Channel: quantity should not be negative: ",order.Quantity)
                //fmt.Println(s)
                errCh <- invalidOrder{order: order, err: errors.New(s)}
            } else {
                out <- order
            }
        }
        close(out)
        close(errCh)
    }()
    return out, errCh
 }

func (o order) String() string{
return fmt.Sprintf("Product code: %v, Quantity: %v, Status: %v \n", o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

type orderStatus int
type order struct{
    ProductCode int
    Quantity float64
    Status orderStatus
    }
    type invalidOrder struct{
        order order
        err error
}
    
var rawOrders = []string{
    `{"productCode": 1111, "quantity":5,"status":1}`,
    `{"productCode": 2222, "quantity":42.3,"status":1}`,
    `{"productCode": 3333, "quantity":19,"status":1}`,
    `{"productCode": 4444, "quantity":8,"status":1}`,
}
    
const(
    none orderStatus = iota
    new
    received
    reserved
    filled
)
func orderStatusToText(o orderStatus) string{
    switch o {
    case none:
        return "none"
    case new:
        return "new"
    case received:
        return "received"
    case reserved:
        return "reserved"
    case filled:
        return "filled"
    default:
        return "unknow status"
    }
}
