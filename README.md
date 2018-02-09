# hello-from

`hello-from` is a simple HTTP service that responds with its hostname and the non-loopback IPv4 addresses of its environment.  Its main use is for demoing load-balancing applications, like Kubernetes Services.

Example:

    $ docker run -ti -p 80:80 -d --name=hello-from pbitty/hello-from
    498ca7cb25e92c2c0152ea5993b91cef2b0a35de7a7dc6fbfefd21194301190b

    $ curl 192.168.99.102
    Hello from 498ca7cb25e9 (172.17.0.2)
