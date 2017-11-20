s3env
===

## Description

Load environment variables from AWS S3.

## Usage

### Setup

#### 1. Upload environment variable

Create a bucket and, put a dotenv format file.  
Optionally, you can use base64 encoded file.

```
$ cat dotenv
# This is a comment
ENABLE_DEBUG_MODE=y # comment
ENABLE_TEST_MODE=${ENABLE_DEBUG_MODE}"

$ cat dotenv | base64 | aws s3 cp - s3://your-bucket-name/dotenv # Upload to s3
```

#### 2. Set AWS credential and s3env environment variables

Set below environment variables, and [aws credential environment variables](http://docs.aws.amazon.com/cli/latest/userguide/cli-environment.html).

```
S3ENV_BASE64ENCODE=y
S3ENV_BUCKET_NAME=your-bucket-name
S3ENV_KEY_NAME=dotenv
```

### Use environment variable from S3

#### Load to current shell

```
$ source <(s3env show --export)
$ echo $ENABLE_DEBUG_MODE
y
```

#### Run with some command

```
$ s3env run irb
irb(main):001:0> ENV["ENABLE_DEBUG_MODE"]
=> "y"
```

## Inspired by

[the40san/dotenv-s3: load your dotenv from s3](https://github.com/the40san/dotenv-s3)
[timakin/ssm2env: Environments injection tool for AWS EC2, with SSM (EC2 Parameter Store).](https://github.com/timakin/ssm2env)
[okzk/env-injector](https://github.com/okzk/env-injector)
