package main

import (
  "bufio"
  "fmt"
  "flag"
  "log"
  "os"
  "time"
  "strings"
)

var input string
var output string

func main() {
  start := time.Now()
  flag.StringVar(&input, "input", "text.txt", "Source file path")
  flag.StringVar(&output, "output", "result.csv", "Output file path")
  flag.Parse()
  
  fmt.Println("Parsing...")
  
  lines, err := readLines(input)
  if err != nil {
    log.Fatalf("readLines: %s", err)
  }
  
  fmt.Println("Writing...")

  if err := writeLines(lines, output); err != nil {
    log.Fatalf("writeLines: %s", err)
  }
  
  elapsed := time.Since(start)
  log.Printf("Done! Time elapsed: %s", elapsed)

}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if len(scanner.Text()) < 261 { continue }
    lines = append(lines, parseLine(scanner.Text()))
  }
  return lines, scanner.Err()
}

//WriteLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}

func parseLine(record string) string {
  dateOfBirth := strings.TrimSpace(record[135:143])
  issuedDate := strings.TrimSpace(record[251:259])
  parsedDateOfBirth := ""
  parsedIssuedDate := ""
  
  if len(dateOfBirth) > 0 {
    parsedDateOfBirth = dateOfBirth[4:]  + "-" + dateOfBirth[2:4] + "-" + dateOfBirth[:2]
  }
  if len(issuedDate) > 0 {
    parsedIssuedDate = issuedDate[4:]  + "-" + issuedDate[2:4] + "-" + issuedDate[:2]
  }
  
  fields := []string {record[:10],  // Driver Licence
  strings.TrimSpace(record[10:50]), // Last Name
  strings.TrimSpace(record[50:90]), // First Name
  strings.TrimSpace(record[90:129]), // Middle Name
  strings.TrimSpace(record[129:135]), // Suffix
  parsedDateOfBirth, // Date of Birth
  strings.TrimSpace(record[143:175]), // Address 1
  strings.TrimSpace(record[175:207]), // Address 2
  strings.TrimSpace(record[207:240]), // City
  strings.TrimSpace(record[240:242]), // State
  strings.TrimSpace(record[242:247]), // Zip Code
  strings.TrimSpace(record[247:251]), // Zip Code ext
  parsedIssuedDate, // Orig Issue Date
  strings.TrimSpace(record[259:])} // Card Type
  
  return strings.Join(fields, ",")
}