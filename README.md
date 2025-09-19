# 🌐 URL Status Checker (Go)

A simple tool written in **Go** to check whether a list of URLs are **UP** or **DOWN**.  
Reads URLs from a file (`urls.txt`) and performs HTTP GET requests to report status.  

---

## 🚀 Features
- Reads URLs from `urls.txt`  
- Checks status with HTTP GET requests  
- Reports status codes (`200 OK`, `404 Not Found`, etc.)  
- Handles timeouts for slow sites  
- Simple and lightweight  

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
PS C:\Gitbase\url-checker> go run .\main.go
[200] https://google.com ✅ UP
[200] https://github.com ✅ UP
[ERR] https://nonexistent-website-12345.com ❌ DOWN (Get "https://nonexistent-website-12345.com": dial tcp: lookup nonexistent-website-12345.com: no such host)
[200] http://example.com ✅ UP
```

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

---

## 📦 Build Executable
```bash
go build -o urlchecker
./urlchecker
```

---

## 🧑‍💻 Future Improvements
- Add concurrency (check multiple URLs in parallel)  
- Export results to JSON or CSV  
- Add colorized output for readability  
- Dockerize for easy deployment  

---

## 📜 License
MIT License – free to use and modify.  
