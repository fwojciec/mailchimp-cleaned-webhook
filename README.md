# mailchimp-cleaned-webhook

This is a target for Mailchimp webhook request with data about emails cleaned from the mailing list.

## Motivation

Mailchimp doesn't send cleaned email notification, in fact it's quite difficult to find out when an email is cleaned. Sometimes it's important to monitor which emails get cleaned, in case this is an unwanted outcome.

## How to use

1. Configure AWS SES service as required: https://docs.aws.amazon.com/sdk-for-go/api/service/ses/#SES.SendEmail
2. Install serverless globally:

```
npm install -g serverless
```

3. Add `.env` file which defines the following variables:

```
TO_EMAIL=to@email.com
FROM_EMAIL=from@email.com
```

4. Install dependencies:

```
go get -u github.com/fwojciec/email
```

```
npm install -D serverless-dotenv-plugin
```

5. Deploy:

```
sls deploy
```

6. Configure your Mailchimp list webhook to use the url of the lambda to send cleaned notifications to.
