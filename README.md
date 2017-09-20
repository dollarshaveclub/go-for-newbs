<p align="center">
  <a href="https://blog.golang.org/gopher">
    <img height="150" src="https://user-images.githubusercontent.com/1074042/30151230-7714164c-9363-11e7-86c4-caca0b76e6b4.png" alt="Go Gopher, borrowed from https://blog.golang.org/gopher" />
  </a>
</p>
<h1 align="center">Go For Newbs</h1>
<h5 align="center">Real project build for learning Go.</h5>
---
<p align="center">
  <a href="setup">Setup</a>
  <a href="user">Setup User</a>
</p>
---

<h2 id="setup">Setup</h2>

Clone
```shell
   git clone https://github.com/dollarshaveclub/go-for-newbs.git
   ```

<h2 id="user">Project 1: Example user microservice</h1>
In this project, a user microservice will be setup and tested using [go](https://golang.org/).

1. to get up and running
   ```shell
   git clone https://github.com/dollarshaveclub/go-for-newbs.git
   ```
2. Then setup
   ```shell
   cd go-for-newbs/user-microservice
   ```
3. Then build
   ```shell
   .user-moicroservice
   ```
4. Create a user
   ```shell
  curl -X POST -d '{ "id" : 1, "name" : "foo" }' "http://0.0.0.0:8088/set"
   ```
5. View in the [browser](http://0.0.0.0:8088/?id=1)
   -  You should see a user
      ```json
      {
        {
          "id": 1,
          "name": "foo",
          "email": "",
          "address": {
          "street_line_1": "",
          "street_line_2": "",
          "city": "",
          "state": "",
          "zipcode": ""
        }
      }
      ```

