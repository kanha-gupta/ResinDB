# How to Run using Docker ?

This guide explains how to run the `resin` CLI tool using Docker.

## Step 1: Pull the Docker Image

First, pull the latest version of the `resin` image from Docker Hub:

```bash
docker pull kanhag4163/resin:latest
```

## Step 2: Run the Container

Start a Docker container from the image with an interactive terminal:

```bash
docker run -it kanhag4163/resin:latest
```

## Step 3: Interact with the CLI Tool

Once inside the container's shell session, you can interact with the `resin` tool. Type `./resin` followed by any commands or options to use the tool:

```bash
./resin [command] [options]
```

Replace `[command]` and `[options]` with the specific commands and options for your tool.

## Conclusion

Following these steps, you can easily use the `resinDB` CLI tool in a Docker environment. For any more information, please use `-h` tag 
