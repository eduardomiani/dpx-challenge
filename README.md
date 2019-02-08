# dpx-challenge

This is a sample zip code client, that connects into Digipix homologation servers. 

## Running with Docker

There is already exists a stable docker image. So, first, pull the image from repository:

```
docker pull eduardomiani/dpx-challenge
```

Then, run a new container from this image:

```
docker run --rm --name dpx-challenge -p {{PORT}}:8080 eduardomiani/dpx-challenge
```
Replace the {{PORT}} with the host port of to your liking.
Access http://localhost:{{PORT}}

