# statuspage-api-client-go

```bash

mkdir -p ~/bin/openapitools
curl https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh > ~/bin/openapitools/openapi-generator-cli
chmod u+x ~/bin/openapitools/openapi-generator-cli
export PATH=$PATH:~/bin/openapitools/

export OPENAPI_GENERATOR_VERSION=5.0.0
openapi-generator-cli version # is 5.0.0

openapi-generator-cli generate -i api/v1/statuspage/api/openapi.yaml -o api/v1/statuspage -g go
```
