# api.one.adapttive.com
Serverless API for link preview

# Installation

- Install Go Lang

    `sudo apt-get install golang`

- Place the project inside `\home\<user>\go\src\<project>`

- Build Go Lambda functions

    `make`

- Install Serverless Framework

  `npm install -g serverless`

- Login (one time)

    `serverless login`

- Deploy

    `serverless deploy --verbose`

# Endpoints

- Ping

    `https://api.one.adapttive.com/ping/`

- Preview

    `https://api.one.adapttive.com/preview/`
