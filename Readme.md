Demo: https://rayyildiz-go.herokuapp.com

Deploy
===
You can deploy it to heroku with the following commands:


  $ git clone https://github.com/rayyildiz/rayyildiz-go-heroku-example.git
  $ cd rayyildiz-go-heroku-example
  $ heroku create -b https://github.com/heroku/heroku-buildpack-go.git [NAME]
  $ git push heroku master
  $ heroku open
