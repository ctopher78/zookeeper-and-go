## Zookeeper and Go

Example code and Vagrant configurations for getting started with
Zookeeper and Go.

Please see the [Getting Started with Zookeeper and Go](https://mmcgrana.github.io/2014/05/getting-started-with-zookeeper-and-go.html)
blog post for details on this code and using Zookeeper with Go.

NOTE (ctopher78): I've edited this forked repo to use docker instead of a vagrant vm to spin up the zookeeper environement.

```console
$ docker-compose -f stack.yml up
```

More information about the docker zookeeper image can be found here:
https://docs.docker.com/samples/library/zookeeper/#start-a-zookeeper-server-instance

`ex-*.go` files are example programs running with `go run` as above.

`sim-*.txt` are failure simulation notes, showing how to run basic
failure simulations in the environment described above and the
results they should produce.
