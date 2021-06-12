
ENDPOINT_URL=http://100.88.225.67:4566
FUNCTION=sample-lambda
ROLE=role


1. Zip the file
   zip function.zip main

2. Create Lambda function
aws lambda create-function \
		--endpoint-url http://100.88.225.67:4566 \
		--runtime go1.x \
		--role role \
		--function-name sample-lambda \
		--handler main \
		--zip-file "fileb:///home/will/lambda/function.zip"

3. Call the Endpoint
aws --endpoint-url http://localhost:4566 lambda invoke \
  --function-name sample-lambda \
  --payload '{"firstname": "foo", "lastname": "bar", "age": 41}' \
  output


  https://github.com/sebastianlacuesta/localstack-sample-go-lambda