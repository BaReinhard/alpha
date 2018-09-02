# Alphavantage API Lib

### Examples

```
func main() {
	opts := Options{
		Function:   "TIME_SERIES_DAILY",
		Symbol:     "MSFT",
		OutputSize: "compact",

		APIKey: os.Getenv("API_TOKEN"),
	}

	client := &http.Client{}
	resp, err := client.Get(fmt.Sprintf("%s?function=%s&symbol=%s&outputsize=%s&apikey=%s", APIURL, opts.Function, opts.Symbol, opts.OutputSize, opts.APIKey))
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print("Error: " + err.Error())
	}
	fmt.Print(string(b))
}
```
