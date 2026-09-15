# statuspage-api-client-go

Go client library for the [Atlassian Statuspage API (v1)](https://developer.statuspage.io/).

This library is automatically generated from the Statuspage OpenAPI specification using `openapi-generator` and custom scripts to patch upstream spec inconsistencies.

## Installation

```bash
go get github.com/sbecker59/statuspage-api-client-go
```

## Usage Example

```go
package main

import (
	"context"
	"fmt"
	"log"

	sp "github.com/sbecker59/statuspage-api-client-go/api/v1/statuspage"
)

func main() {
	cfg := sp.NewConfiguration()
	client := sp.NewAPIClient(cfg)

	// Configure API Key authentication
	ctx := context.WithValue(
		context.Background(),
		sp.ContextAPIKeys,
		map[string]sp.APIKey{
			"api_key": {
				Key:    "YOUR_STATUSPAGE_API_KEY",
				Prefix: "oauth",
			},
		},
	)

	// List pages
	pages, _, err := client.PagesAPI.GetPages(ctx).Execute()
	if err != nil {
		log.Fatalf("Failed to fetch pages: %v", err)
	}

	for _, page := range pages {
		fmt.Printf("Page Name: %s | ID: %s\n", page.GetName(), page.GetId())
	}
}
```

## Development & Regenerating the Client

The client is generated from the Statuspage OpenAPI specification using scripts located in the `scripts/` directory.

### Requirements

- Node.js (v18+)
- Java (required by `openapi-generator-cli`)
- Go (v1.20+)

### Steps to Regenerate

1. Install script dependencies:
   ```bash
   cd scripts
   npm install
   npx puppeteer browsers install chrome
   ```

2. Extract the latest OpenAPI specification:
   ```bash
   node extract-spec.js
   ```

3. Patch known upstream spec issues:
   ```bash
   node fix-spec.js
   ```

4. Regenerate the Go client:
   ```bash
   cd ..
   rm -rf api/v1/statuspage
   npx @openapitools/openapi-generator-cli generate \
     -i scripts/developer_statuspage_io.json \
     -g go \
     -o api/v1/statuspage \
     --git-user-id sbecker59 \
     --git-repo-id statuspage-api-client-go \
     --skip-validate-spec \
     --enable-post-process-file
   rm -f api/v1/statuspage/go.mod api/v1/statuspage/go.sum
   gofmt -w api/v1/statuspage
   ```

## License

MIT License. See [LICENSE](LICENSE) for details.

