package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)


func main() {
	// fmt.Println("Hello, World!")

	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	// c.OnHTML("title", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	type Book struct {
		Title string
		Cover string
		Price string
	}

	// basic scraper

	// c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
	// 	var book Book
	// 	book.Cover = e.ChildAttr(".image_container img", "src")
	// 	book.Title = e.ChildAttr(".image_container img", "alt")
	// 	book.Price = e.ChildText(".price_color")
	// 	fmt.Println(book.Title, "https://books.toscrape.com/" + book.Cover, book.Price)
	// })

	// pagination

	// c.OnHTML(".pager li.next a", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	fmt.Println(link)
	// 	c.Visit(e.Request.AbsoluteURL(link))
	// })

	// c.OnHTML(".next > a", func(e *colly.HTMLElement) {
	// 	nextPage := e.Request.AbsoluteURL(e.Attr("href"))
	// 	c.Visit(nextPage)
	// })

	file, err := os.Create("export.csv")
		if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush() // flush everything into the file

	headers := []string{"Title", "Cover", "Price"}
	writer.Write(headers)

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		var book Book
		book.Cover = e.ChildAttr(".image_container img", "src")
		book.Title = e.ChildAttr(".image_container img", "alt")
		book.Price = e.ChildText(".price_color")
		row := []string{book.Title, "https://books.toscrape.com/" + book.Cover, book.Price}
		writer.Write(row)
	})

	c.Visit("https://books.toscrape.com/")

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	// })

	
}