## Setup

Install the AWS Amplify CLI and configure permissions to provision services on your AWS account.

```
$ npm install -g @aws-amplify/cli
$ amplify configure
```

Add AWS Amplify to your project and initialize it. Here's what I did:

```
npm install aws-amplify
amplify init

? Choose the type of app that you're building **javascript**

? What javascript framework are you using **react** (doesn't matter here)
? Source Directory Path:  **src**
? Distribution Directory Path: **build**
? Build Command:  **npm run build**
? Start Command: **npm start**

? Do you want to use an AWS profile? **Yes**
? Please choose the profile you want to use **amplify** (whatever you named your profile)
```

Amplify will then initialize several resources: An IAM role, an S3 bucket, and a CloudFormation stack to manage those resources and any others you add.

## Write your Lambda function and test it

Run the following command and answer the prompts to generate a basic Lambda function.

The function will be created locally for you to edit before deploying. Here's what I did:

```
amplify add function

Using service: Lambda, provided by: awscloudformation
? Provide a friendly name for your resource to be used as a label for this category in the project: **quicksort**
? Provide the AWS Lambda function name: **quicksort**
? Choose the function template that you want to use: **Hello world function**
? Do you want to edit the local lambda function now? **true**
```

If your function takes an input (i.e. a request body), you'll need to update `event.json` in the same directory as your function to be a mock request body object.

When you're ready to test your function, run the following and hit enter on all of the prompts:

```
amplify function invoke **quicksort** (whatever you named your function)
```

## Deploy your Lambda function

If your function worked as expected, you can deploy it by running:

```
amplify push
```

## Setup API Gateway to serve your function

You'll need to log in to your AWS console and do a bit of setup to serve your function.

Find your function in AWS Lambda and under `Designer`, add a new trigger: `API Gateway`.

Scroll down to configure it. Choose `Create a new API`, and then for now just set security to `Open`.

Hit add, and you'll soon be given an API endpoint that you can hit with a POST request!
