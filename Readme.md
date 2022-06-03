# Readme for Techpro.Club

## Note:

Rename the `src/.env.sample` to `src/.env` and replace with actual credentials.

## Documentation for AWS SES

For emails, we are using Amazon SES. Please go through the documentation https://docs.aws.amazon.com/ses/latest/dg/smtp-credentials.html
Apart from that, generated SMTP Credentials will not work in place of `SES_ACCESS_ID` & `SES_ACCESS_SECRET`. You will need to go to **IAM > Users > Security Credentials > Create Access Key**. Note the credentials and replace the values in .env variables.

Also, remember to apply `AmazonSesSendingAccess` from **IAM > Users > Permissions > Add Permissions**

The above steps are before you unlock the sending limits in SES. Read https://docs.aws.amazon.com/ses/latest/dg/manage-sending-quotas.html
