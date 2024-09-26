---
slug: create-lambda-aws-cli
title: How to Create Lambda from AWS CLI 
authors: [kbbgl]
tags: [cloud, lambda, aws]
---

### Create Function Code

```python
# function.py
import requests
def main(event, context):   
    response = requests.get("https://docs.aws.amazon.com")
    print(response.text)
    return response.text
if __name__ == "__main__":   
    main('', '')
```

### Create Dependency Zip

```bash
pwd
python_project 

tree .
└── function.py

# Install all Python dependencies into `package` folder
pip3  install --target ./package requests

tree .
└── package
    └── ...
└── function.py

# Create zip of dependency
cd package
zip -r ../python_project.zip .

# Add file used for lambda function to root of the zip
cd ..
zip -g python_project.zip function.py
```

### Create Lambda function from Source

```bash
cd python_project

aws lambda create-function --function-name python-project --zip-file fileb://python_project.zip --handler function.main --runtime python3.8 --role arn:aws:iam::$account_id:role/$role_name
```

### Inboke Lambda Function

```bash
aws lambda invoke --function-name requests-function --payload '{"key1": "value1", "key2": "value2", "key3": "value3"}' output.txt
```
