---
title: Happy Handlers and Swagger Codegen
author: Joseph Spurrier
authorURL: http://github.com/josephspurrier
authorFBID: 1225770017
---

Let's improve those HTTP handlers! In 2018, I [released code](https://github.com/josephspurrier/h) that forces an HTTP status code and an error to be returned
by all of the handlers so it's very clear what happens in each of the code paths.
[View commit](https://github.com/josephspurrier/govueapp/commit/4f77cb041d71ae06b07d93bb22a8aba40441503f).

There is a new bind package (thanks to the [go-playground](https://github.com/go-playground/validator)) that handles mapping data from a request to a struct
and also validates it. [View commit](https://github.com/josephspurrier/govueapp/commit/06457c34511871d720e1aa84da8fba88a725c79d). 

### Swagger Code Generation

No one likes to write an API spec by hand so I added in Swagger annotations which
allows us to use [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) tool
to generate a Swagger spec from it!

<!--truncate-->

I [worked with the Swagger team](https://github.com/go-swagger/go-swagger/issues/1545)
to get code added in that allows us to store the structs for each request inside
of the func body to help clean up the cody.
[View commit](https://github.com/josephspurrier/govueapp/commit/34286e553ae2ede8a03abd9422774f04a9389783).

You can view the latest Swagger spec [here](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/josephspurrier/govueapp/master/src/app/api/static/swagger/swagger.json).

