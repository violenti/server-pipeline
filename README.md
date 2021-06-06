# Server pipeline

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://gopherize.me/gopher/2970681ff5ef6a82089f0868822f01e307825830">
<img src="images/icon.png" alt="Logo" width="100" height="100">
  </a>
</p>



This is app for sharing your .gitlab-ci.yml, for multi projects. 


# Installation

## Develoment enverioment 

```bash
export GITLAB_TOKEN=djkfjdkfjdkjd
export GITLAB=https://gitlab.com/api/v4/projects/9927708/repository/files/

```

```bash

go run main.go

```

## Docker

```bash
docker run violenti/server-pipeline:latest -e GITLAB_TOKEN=djkfjdkfjdkjd -e GITLAB=https://gitlab.com/api/v4/projects/9927708/repository/files/
```

## Kubernetes 

You can do use of the manifest that find on kubernetes folder. 

## Usage

```html
url/master/filename.yml 
```
or 

```html
url/development/filename.yml
```

Where url is the domain name of the app. And  master or development is the branch name.  


# Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

# License
[MIT](https://choosealicense.com/licenses/mit/)
