
https://stackoverflow.com/questions/48236040/exactly-which-path-do-i-curl-for-localstack-api-gateway-lambda-integration
https://github.com/vjeffrey/localstack-aws-sdk/blob/master/main.go

go bulid main

zip apiTestHandler.zip main

scp apiTestHandler.zip will@will-server:/home/will/lambda/

aws lambda create-function \
--endpoint-url http://100.88.225.67:4566
--runtime go1.x
--role apitest 
--function-name api-test-handler \
--handler main \
--zip-file "fileb:///home/will/lambda/apiTestHandler.zip" \

aws apigateway create-rest-api --name 'API Test' --endpoint-url=http://100.88.225.67:4566
// Response
"id": "3k2voifd9r",
"createdDate": 1623144567,
///

aws apigateway get-resources --rest-api-id 3k2voifd9r --endpoint-url=http://100.88.225.67:4566
// Response (Parent ID)
 "id": "qq4u0kgaie",
//

aws apigateway create-resource \
--rest-api-id 3k2voifd9r \
--parent-id qq4u0kgaie \
--path-part "{somethingId}" --endpoint-url=http://100.88.225.67:4566
// Response
    "id": "4coglefmws",
//

aws apigateway put-method \
 --rest-api-id 3k2voifd9r \
 --resource-id 4coglefmws \
 --http-method GET \
 --request-parameters "method.request.path.somethingId=true" \
 --authorization-type "NONE" \
--endpoint-url=http://100.88.225.67:4566

aws apigateway put-integration \
 --rest-api-id 3k2voifd9r \
 --resource-id 4coglefmws \
 --http-method GET \
 --type AWS_PROXY \
 --integration-http-method POST \
 --uri lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-2:000000000000:function:api-test-handler/invocations \
 --passthrough-behavior WHEN_NO_MATCH \
 --endpoint-url=http://100.88.225.67:4566



Call the Endpoint aws --endpoint-url http://localhost:4566 lambda invoke
--function-name sample-lambda
--payload '{"firstname": "foo", "lastname": "bar", "age": 41}'
output