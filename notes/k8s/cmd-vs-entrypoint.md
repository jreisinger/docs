
## Docker

- `ENTRYPOINT` defines default command; gets appended
- `CMD` defines default command complete with arguments; gets replaced

```
❯ cat Dockerfile
FROM ubuntu
ENTRYPOINT [ "sleep" ]
CMD [ "5" ]
❯ docker build -t my-ubuntu .
❯ docker run my-ubuntu   # sleeps for 5 seconds
❯ docker run my-ubuntu 2 # sleeps for 2 seconds
```
