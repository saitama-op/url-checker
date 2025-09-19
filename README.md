# 🌐 URL Status Checker (Go) with Concurrency, CSV & JSON Export

A robust tool written in **Go** to check whether a list of URLs are **UP** or **DOWN**, using a configurable worker pool for concurrency and exporting results to CSV and JSON.  

---

## 🚀 Features
- Reads URLs from a file (`urls.txt`)  
- Checks status with HTTP GET requests  
- Worker pool for safe concurrency  
- Reports status codes (`200 OK`, `404 Not Found`, etc.)  
- Failed URLs appear first in the tabular output  
- Colorized console output (Red=DOWN, Green=UP)  
- Exports results to CSV and optional JSON  

---

## 📂 Project Structure
```
url-checker/
├── main.go        # Main program
├── urls.txt       # List of URLs to check
├── go.mod
└── README.md
```

---

## 📝 Example `urls.txt`
```
https://google.com
https://github.com
https://nonexistent-website-12345.com
http://example.com
```

---

## ⚡ Example Output
```
PS C:\Gitbase\url-checker> go run .\main.go -json="result.json"
URL                                                Status   Code   Error
------------------------------------------------------------------------------------------
https://nonexistent-website-12345.com              DOWN     0      Get "https://nonexistent-website-12345.com": dial tcp: lookup nonexistent-website-12345.com: no such host
http://example.com                                 UP       200
https://github.com                                 UP       200
https://google.com                                 UP       200
Results saved to results.csv
Results saved to result.json 
```
<img width="1892" height="252" alt="image" src="https://github.com/user-attachments/assets/187fd404-f990-4e46-9c8a-7a89d6213487" />

---

## 🛠 Installation & Usage

### Clone the repo
```bash
git clone https://github.com/your-username/url-checker.git
cd url-checker
```

### Install dependencies
```bash
go mod tidy
```

### Run the program
```bash
go run main.go
```

### Command-line flags
- `-workers` : Number of concurrent workers (default 20)  
- `-file`    : File containing URLs (default `urls.txt`)  
- `-csv`     : CSV output file (default `results.csv`)  
- `-json`    : Optional JSON output file  

#### Examples
```bash
# Default workers and CSV
go run main.go

# Custom workers and CSV
go run main.go -workers=50 -csv=myresults.csv

# Export to JSON
go run main.go -json=results.json

# Custom URL file + JSON + CSV
go run main.go -file=myurls.txt -workers=30 -csv=output.csv -json=output.json
```

---

## 🧑‍💻 Future Improvements
- Colorized console output (already implemented)  
- Dockerize for deployment  
- Integrate with monitoring systems for automated alerts  

---

## 📜 License
MIT License – free to use and modify.  
