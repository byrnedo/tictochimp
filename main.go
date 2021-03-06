package main

import (
	"flag"
	"fmt"
	"github.com/byrnedo/tictochimp/config"
	"github.com/byrnedo/tictochimp/models/mailchimp"
	mailchimpSpec "github.com/byrnedo/tictochimp/models/mailchimp/spec"
	"github.com/byrnedo/tictochimp/models/tictail"
	tictailSpec "github.com/byrnedo/tictochimp/models/tictail/spec"
	"os"
	"strings"
	"sync"
	"text/tabwriter"
)

var (
	configFile string
	showUsage  bool
	dryRun     bool
	w          *tabwriter.Writer
)

func init() {
	w = new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
}

func main() {

	flag.StringVar(&configFile, "conf", "", "Configuration file path")
	flag.BoolVar(&dryRun, "dry-run", false, "Only output changes, dont save to mailing list")
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
		fmt.Fprintln(os.Stderr, "Error parsing config file:"+err.Error())
		os.Exit(1)
	}
	//fmt.Println("Got config:" + cnf.GetUnderlyingData().Root.String())

	startProgram(&cnf)

}

func startProgram(cnf *config.Config) {
	var wg sync.WaitGroup
	wg.Add(2)

	existingMembers := make(chan []mailchimpSpec.Member, 1)
	listID := make(chan string, 1)
	newSubscribers := make(chan map[string]mailchimp.Subscriber, 1)

	mc, err := mailchimp.NewMailchimp(cnf.Mailchimp.AccessToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to make mailchimp client: "+err.Error())
		os.Exit(2)
	}

	go func() {
		defer wg.Done()

		_listID := getSpecifiedList(mc, cnf.Mailchimp.ListName)
		if len(_listID) == 0 {
			fmt.Fprintln(os.Stderr, "Failed to find list ID")
			os.Exit(3)
		}
		listID <- _listID
		fmt.Println("List ID = " + _listID)

		listMembers, err := mc.GetAllListMembers(_listID)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to find list members:", err.Error())
			os.Exit(4)
		}
		fmt.Printf("Found %d members in list\n", len(listMembers))
		existingMembers <- listMembers

	}()

	go func() {
		defer wg.Done()

		tt := tictail.NewTictail(cnf.Tictail.AccessToken)

		orders, err := getOrdersForProduct(tt, cnf.Tictail.StoreName, cnf.Tictail.ProductName)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to find orders:", err.Error())
			os.Exit(5)
		}

		fmt.Printf("Got %d orders for product %s\n", len(orders), cnf.Tictail.ProductName)
		if len(orders) == 0 {
			fmt.Println("Exitting..")
			os.Exit(0)
		}
		newSubscribers <- createEmailList(orders)
	}()

	wg.Wait()
	close(existingMembers)
	close(newSubscribers)
	close(listID)

	fmt.Println("Creating unique email list")
	filteredList := filterExistingSubscribers(mc, <-existingMembers, <-newSubscribers)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("###########################################")
	fmt.Printf("### Subscribers which will be added: %d ###\n", len(filteredList))
	fmt.Println("###########################################")
	fmt.Fprintln(w, "Email\tFirstName\tLastName\tAdded")
	for _, newSub := range filteredList {
		worked := "true"
		if dryRun == false {
			if err = mc.AddSubscriber(newSub, <-listID); err != nil {
				worked = "false"
				fmt.Fprintln(os.Stderr, "Error adding subscriber: "+err.Error())
			}
		} else {
			worked = "?"
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", newSub.Email, newSub.FirstName, newSub.LastName, worked)
	}
	if len(filteredList) > 0 {
		w.Flush()
	}
	//pretty.Println(newSubscribers)

}

func getSpecifiedList(mc *mailchimp.Mailchimp, listName string) (id string) {

	allLists, err := mc.GetLists()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get lists: "+err.Error())
	} else if len(allLists) == 0 {
		fmt.Fprintln(os.Stderr, "Failed to find any lists")
	}

	fmt.Println("Lists Found:")
	for _, list := range allLists {
		fmt.Println("\t" + list.Name)
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

func createEmailList(orders []tictailSpec.OrdersResponse) (uniqueSubs map[string]mailchimp.Subscriber) {
	uniqueSubs = make(map[string]mailchimp.Subscriber, 0)

	for _, order := range orders {

		nameParts := strings.SplitN(order.Customer.Name, " ", 2)
		firstName := strings.Title(nameParts[0])
		lastName := ""
		if len(nameParts) > 1 {
			lastName = strings.Title(nameParts[1])
		}
		uniqueSubs[order.Customer.Email] = mailchimp.Subscriber{
			Email:     order.Customer.Email,
			FirstName: firstName,
			LastName:  lastName,
		}

	}
	return
}

func filterExistingSubscribers(mc *mailchimp.Mailchimp, existingMembers []mailchimpSpec.Member, newSubscribers map[string]mailchimp.Subscriber) (filteredSubscribers []mailchimp.Subscriber) {
	existingMembersMap := make(map[string]mailchimpSpec.Member)
	unacTabWriter := new(tabwriter.Writer)
	unacTabWriter.Init(os.Stdout, 0, 8, 0, '\t', 0)
	numUnaccounted := 0
	for _, member := range existingMembers {
		if _, exists := newSubscribers[strings.ToLower(member.EmailAddress)]; exists == false {
			numUnaccounted++
			fmt.Fprintf(unacTabWriter, "%s\t%s\t%s\n", member.EmailAddress, member.MergeFields.FirstName, member.MergeFields.LastName)
		}
		fmt.Fprintf(w, "%s\t%s\t%s\n", member.EmailAddress, member.MergeFields.FirstName, member.MergeFields.LastName)
		existingMembersMap[strings.ToLower(member.EmailAddress)] = member
	}
	fmt.Printf("\nCurrent Members in List: %d\n\n", len(existingMembers))
	w.Flush()

	filteredSubscribers = make([]mailchimp.Subscriber, 0, 0)

	fmt.Printf("\nMembers of list which came from tictail orders: %d\n\n", len(newSubscribers))
	for _, newSubscriber := range newSubscribers {
		if _, exists := existingMembersMap[strings.ToLower(newSubscriber.Email)]; exists == true {
			fmt.Fprintf(w, "%s\t%s\t%s\n", newSubscriber.Email, newSubscriber.FirstName, newSubscriber.LastName)
			continue
		}
		filteredSubscribers = append(filteredSubscribers, newSubscriber)
	}
	w.Flush()

	fmt.Printf("\nMembers added by other means: %d\n\n", numUnaccounted)
	unacTabWriter.Flush()
	return
}
