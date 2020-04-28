[![Go Report Card](https://goreportcard.com/badge/github.com/jjmengze/heroku-go)]

# heroku-go
This is very simple web server in golang ,it will show that heroku open port for this application. 
## how to use heroku

1. login your heroku
```bash
$heroku login
...
```

2. create a project for this application
```bash
$heroku create <YOUR_APP_NAME>
```

3. clone this repository
```bash
$git clone https://github.com/jjmengze/heroku-go.git
```
4. add a remote heroko url at this project
```bash
$git:remote -a <YOUR_APP_NAME>
```

5. deploy to heroku
```bash
$git push heroku master
```
