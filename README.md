# ManyStagings

ManyStagings is fast, lightweight and easy staging environment manager with a focus on the Developer Experience.
Orchestrators like Kubernetes, are extremely powerful for production applications, but when we use it as development environment, it becomes very complex and verbose to manage applications and services.

ManyStagings aims to integrate and abstract all the complexity of Kubernetes by bringing a friendly Web UI and a CLI that becomes the developer's best friend.

![manystagings-archtecture-padding](https://user-images.githubusercontent.com/6170412/116957360-14c88e00-ac6e-11eb-8396-944121dd8716.png)


## What problem does ManyStaging solve?

Let's imagine a scenario in which a Frontend developer needs to build an entire backend environment in order to develop a web or mobile application. He will have to worry about configuring API, Database and how the applications connect to each other. In addition, the QA process where we have only a single staging environment can become a bottleneck.

![manystagings-bottleneck-padding](https://user-images.githubusercontent.com/6170412/116957365-1a25d880-ac6e-11eb-85c8-be786e0f3cff.png)

## Why not just use docker-compose, for example?

Docker-compose solves this problem in parts, but it can still consume a lot of memory and CPU. Having a remote environment allows us to execute only what we are developing. And it is also possible to generate a public link for a QA review, for example.

Using ManyStagings the only thing we need to do is configure and manage the remote stagings environment  with a simple and intuitive CLI.


## Web UI

With the ManyStagings Web UI you can configure the entire environment, as if it were in the docker-compose: docker image, ports, environment variables and Public URL.

![image](https://user-images.githubusercontent.com/6170412/116957209-ad124300-ac6d-11eb-967f-36dd1ec19073.png)


## ManyStagings CLI

The ManyStagings Command-line is very simple to use. It looks like you are managing the applications on your local computer, but in fact everything happens remotely. It's the docker-compose experience but using Kubernetes behind the scenes.

![image](https://user-images.githubusercontent.com/6170412/116957225-bbf8f580-ac6d-11eb-8908-6c7bb383bce9.png)



