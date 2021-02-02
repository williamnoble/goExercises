aws lambda create-function --function-name goLambda --runtime go1.x --zip-file fileb://function.zip --handler main --role arn:aws:iam::196410176949:role/execution_role

aws lambda invoke --function-name goLambda --payload '{ "name": "WILLIAM" }' response.json

aws lambda invoke --function-name goLambda --payload '{ "name": "WILLIAM" }' /tmp/output.json
