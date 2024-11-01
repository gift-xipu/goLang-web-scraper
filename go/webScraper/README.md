```markdown
# Web Scraper

This is a simple web scraper built using Go and the [Colly](https://github.com/gocolly/colly) library. It scrapes product data from an e-commerce website and saves the information to a CSV file.

## How It Works

The scraper performs the following steps:

1. **Initialization**: It initializes a Colly collector with the allowed domain to ensure that only requests to that domain are permitted.
2. **Scraping Logic**: The collector listens for HTML elements that match the specified selectors (in this case, `li.product`) and extracts relevant product data (URL, image source, name, and price) from each product item.
3. **Data Storage**: After scraping is completed, it writes the collected product data into a CSV file named `products.csv` with appropriate headers.

## Prerequisites

Make sure you have the following installed:

- Go (version 1.15 or higher)
- Git (for cloning the repository)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/web-scraper.git
   cd web-scraper
   ```

2. Install the required dependencies:
   ```bash
   go mod tidy
   ```

## How to Build and Run

To build and run the scraper, follow these steps:

1. Ensure you are in the project directory where your Go files are located.
2. Run the following command to execute the scraper:
   ```bash
   go run .
   ```

## Output

Upon successful execution, the scraper will create a file named `products.csv` in the project directory. The CSV file will contain the following columns:

- **Url**: The product URL
- **Image**: The product image URL
- **Name**: The product name
- **Price**: The product price

### Example of `products.csv` Output

```csv
Url,Image,Name,Price
https://www.scrapingcourse.com/product1.jpg,https://www.scrapingcourse.com/product1.jpg,Product 1,$19.99
https://www.scrapingcourse.com/product2.jpg,https://www.scrapingcourse.com/product2.jpg,Product 2,$29.99
...
```

## Notes

- Make sure that the target website allows web scraping and complies with its `robots.txt` file.
- Modify the HTML selectors in the code if the structure of the target website changes.
- You can customize the output CSV filename and format as needed.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

### Instructions
1. **Replace** `https://github.com/your-username/web-scraper.git` with the actual URL of your repository.
2. **Add any additional information** as needed, such as license details or specific usage instructions based on your implementation. 

This `README.md` provides a clear overview of your project, making it easier for others (or yourself in the future) to understand how to use it.
