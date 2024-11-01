package main

import (
    "encoding/csv"
    "log"
    "os"
    "github.com/gocolly/colly"
)

// Product struct to store scraped data
type Product struct {
    Url   string
    Image string
    Name  string
    Price string
}

func collector(link string) *colly.Collector {
    c := colly.NewCollector(
        colly.AllowedDomains("www.scrapingcourse.com"), // Specify the allowed domain
    )

    // Create products slice in the correct scope
    var products []Product

    c.OnHTML("li.product", func(e *colly.HTMLElement) {
        // Initialize a new Product instance
        product := Product{
            Url:   e.ChildAttr("a", "href"),
            Image: e.ChildAttr("img", "src"),
            Name:  e.ChildText(".product-name"),
            Price: e.ChildText(".price"),
        }
        // Add the product instance to the slice
        products = append(products, product)
    })

    // Store the data to a CSV after extraction
    c.OnScraped(func(r *colly.Response) {
        // Open the CSV file
        file, err := os.Create("products.csv")
        if err != nil {
            log.Fatalln("Failed to create output CSV file", err)
        }
        defer file.Close()

        // Initialize a file writer
        writer := csv.NewWriter(file)
        defer writer.Flush()

        // Write the CSV headers
        headers := []string{
            "Url",
            "Image",
            "Name",
            "Price",
        }
        if err := writer.Write(headers); err != nil {
            log.Fatalln("Failed to write headers to CSV:", err)
        }

        // Write each product as a CSV row
        for _, product := range products {
            record := []string{
                product.Url,
                product.Image,
                product.Name,
                product.Price,
            }
            if err := writer.Write(record); err != nil {
                log.Printf("Failed to write record to CSV: %v", err)
                continue
            }
        }
    })

    return c
}

func main() {
    url := "www.scrapingcourse.com/ecommerce/"
    col := collector(url)
    err := col.Visit("https://" + url) // Use the full URL
    if err != nil {
        log.Fatalf("Failed to visit site: %v", err)
    }
}

