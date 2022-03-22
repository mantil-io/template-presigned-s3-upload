## About

This template shows how to use AWS Lambda deployed through Mantil to upload files to S3 bucket without having direct access to the AWS account or the bucket using presigned urls.

## Prerequisites

This template is created with Mantil. To download [Mantil CLI](https://docs.mantil.com/cli_install) on Mac or Linux use Homebrew

```
brew tap mantil-io/mantil
brew install mantil
```

or check [direct download links](https://docs.mantil.com/cli_install#direct-download-linux-windows-and-mac).

To deploy this application you will need an [AWS account](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/).

## Installation

To locally create a new project from this template run:

```
mantil new app --from presigned-s3-upload
cd app
```

## Configuration 

Before deploying your application you will need to create S3 bucket in the same AWS account in which your application will be deployed.
Once your bucket is created you need to add the name of the bucket to `config/environment` file as env variable for your function.

```
project:
  env:
    BUCKET: # bucket to which files will be uploaded
```

## Deploying an application

Note: If this is the first time you are using Mantil you will need to install Mantil Node on your AWS account. For detailed instructions please follow the [setup guide](https://docs.mantil.com/aws_detailed_setup/aws_credentials)

```
mantil aws install
```

After configuring the environment variable you can proceed with application deployment.

```
mantil deploy
```

This command will create a new stage for your project with default name `development` and deploy it to your node.

Now you can output the stage endpoint with `mantil env -u`. The API endpoint for your function will have the name of that function in the path, in our case that is `$(mantil env -u)/upload`.

## Generating presigned S3 upload url

To generate presigned S3 upload url for your file you will need to make a request to function API endpoint specifying name the file will have in your bucket in the request body. You can make a request either directly through endpoint URL outlined in the section above or with `mantil invoke`.

```
mantil invoke upload -d "name_of_the_file"
```

Response will contain presigned upload url which you can then use to upload file to the bucket.

```
curl (generated link) --upload-file (file you wish to upload)
```

## Cleanup

To remove the created stage with all resources from your AWS account destroy it with

```
mantil stage destroy development
```

## Final thoughts

With this template you learned how to upload files to your S3 bucket through presigned url without having direct access to the AWS account or the bucket. Check out our [documentation](https://docs.mantil.com/examples) to find more interesting templates. 

If you have any questions or comments on this concrete template or would just like to share your view on Mantil contact us at [support@mantil.com](mailto:support@mantil.com) or create an issue.
