# Antibot Backend
## Requirements
 - Go (tested on 1.16.5)

## Usage
Run:  
```bash
make run
```
Test:  
```bash
curl -X POST http://localhost:8081/api/v1/analyze -d '{"data":"test"}'
```
