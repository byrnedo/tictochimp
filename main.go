package main

import (
	"flag"
	"fmt"
	"github.com/byrnedo/tictochimp/config"
	"github.com/byrnedo/tictochimp/models/mailchimp"
	"github.com/byrnedo/tictochimp/models/tictail"
	tictailSpec "github.com/byrnedo/tictochimp/models/tictail/spec"
	"os"
)

var (
	configFile string
	showUsage  bool
)

func main() {

	flag.StringVar(&configFile, "conf", "", "Configuration file path")
	flag.BoolVar(&showUsage, "help", false, "Show usage information")
	flag.Parse()

	if len(configFile) == 0 {
		showUsage = true
	}

	if showUsage {
		flag.Usage()
		os.Exit(1)
	}

	var cnf = config.Config{}
	err := cnf.ParseFile(configFile)
	if err != nil {
		fmt.Println("Error parsing config file:" + err.Error())
		os.Exit(1)
	}
	fmt.Println("Got config:" + cnf.GetUnderlyingData().Root.String())

	startProgram(&cnf)

}

func startProgram(cnf *config.Config) {

	mc := mailchimp.NewMailchimp(cnf.Mailchimp.AccessToken)
	listID := getSpecifiedList(mc, cnf.Mailchimp.ListName)
	if len(listID) == 0 {
		fmt.Println("Failed to find list ID")
		os.Exit(1)
	}
	fmt.Println("List ID = " + listID)

	tt := tictail.NewTictail(cnf.Tictail.AccessToken)

	orders, err := getOrdersForProduct(tt, cnf.Tictail.StoreName, cnf.Tictail.ProductName)

	if err != nil {
		fmt.Println("Failed to find orders:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Got %d orders for product %s\n", len(orders), cnf.Tictail.ProductName)
	if len(orders) == 0 {
		fmt.Println("Exitting..")
		os.Exit(0)
	}
	fmt.Println("Creating unique email list")

}

func getSpecifiedList(mc *mailchimp.Mailchimp, listName string) (id string) {

	allLists, err := mc.GetLists()

	if err != nil {
		fmt.Println("Failed to get lists: " + err.Error())
	} else if len(allLists) == 0 {
		fmt.Println("Failed to find any lists")
	}

	for _, list := range allLists {
		if list.Name == listName {
			id = list.Id
		}
	}
	return
}

func getOrdersForProduct(tt *tictail.Tictail, storeName string, productName string) ([]tictailSpec.OrdersResponse, error) {
	me, err := tt.GetMe()
	if err != nil {
		return nil, fmt.Errorf("Error getting 'me' response from tictail: %s", err.Error())
	}

	var storeID string
	if me.Subdomain == storeName {
		storeID = me.ID
	}
	fmt.Printf("%#v\n", me)
	if len(storeID) == 0 {
		return nil, fmt.Errorf("Failed to find store ID for %s", storeName)
	}

	allOrders, err := tt.GetAllOrders(storeID)
	if err != nil {
		return nil, fmt.Errorf("Error getting orders for store (name=%s, id=%s): %s", storeName, storeID, err.Error())
	}

	ordersForProduct := []tictailSpec.OrdersResponse{}
	for _, order := range allOrders {

		for _, item := range order.Items {
			if item.Product.Title == productName {
				ordersForProduct = append(ordersForProduct, order)
			}
		}

	}

	return ordersForProduct, nil
}
