// go run main.go repo:golang/go is:open

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"example.com/m/Chapter04/Exercise4.10/github"
)

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	onedayago := now.AddDate(0, 0, -1)
	onemonthago := now.AddDate(0, -1, 0)
	oneyearago := now.AddDate(-1, 0, 0)

	// only store pointers
	var in1dayissue []*github.Issue
	var in1monissue []*github.Issue
	var in1yearissue []*github.Issue

	// Within one day/month/year
	for _, item := range result.Items {
		if item.CreatedAt.After(onedayago) {
			in1dayissue = append(in1dayissue, item)
		}
		if item.CreatedAt.After(onemonthago) {
			in1monissue = append(in1monissue, item)
		}
		if item.CreatedAt.After(oneyearago) {
			in1yearissue = append(in1yearissue, item)
		}
	}

	// print issues within one d/m/y
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range in1dayissue {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("---------------------------------------")

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range in1monissue {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("---------------------------------------")
	for _, item := range in1yearissue {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
