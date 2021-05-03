# ManyStagings

ManyStagings is fast, lightweight and easy staging environment manager with a great focus on the Developer Experience.
Orchestrators like Kubernetes, are extremely powerful for production applications. But when we use it as development environment, it becomes very complex and verbose to manage applications and services.
ManyStagings aims to integrate and abstract all the complexity of Kubernetes by bringing a friendly Web UI and a CLI that becomes the developer's best friend.

![image](https://gblobscdn.gitbook.com/assets%2F-MYc4vBTATLslDGIBVhf%2F-MZjU9twlH1fjuxVibVq%2F-MZjXc_LQUyAUpQJPBJo%2Fimage.png?alt=media&token=0534786e-94d4-47aa-9395-86b77315e865)

## What problem does ManyStaging solve?

The main idea is that each developer can have their own remote staging environment. Let's imagine a scenario in which a Frontend developer needs to build an entire backend environment in order to develop a web or mobile application. He will have to worry about configuring API, Database, how the applications connect to each other, etc.

## Why not just use docker-compose, for example?

Docker-compose solves this problem in parts, but it can still consume a lot of memory and CPU. Having a remote environment allows us to execute only what we are developing. And it is also possible to generate a public link for a QA review, for example.

With ManyStaging the only thing we need to do is configure the CLI and run:

```
$ ms up

$ ms status

Environment: Carlos's Environment
+-------------+---------+--------+------+------------------------------------------------+
| APPLICATION | STATUS  |  AGE   | PORT |                   PUBLIC URL                   |
+-------------+---------+--------+------+------------------------------------------------+
| postgres    | RUNNING | 13m37s | 5432 |                                                |
| hello-world | PAUSED  |    -   |   80 | https://carlos-hello-world.ms.myproject.com    |
| nginx       | RUNNING | 13m35s |   80 | https://carlos-nginx.ms.myproject.com          |
+-------------+---------+--------+------+------------------------------------------------+
```

