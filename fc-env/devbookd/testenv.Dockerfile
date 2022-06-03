# This Dockerfile is for creating a testing environment for devbookd.
# docker build -t devbookd-testenv . -f testenv.Dockerfile


FROM node:16

COPY bin/devbookd /usr/bin/devbookd

WORKDIR code
RUN npm init -y

# Set env vars for devbook-daemon
RUN echo RUN_CMD=node >> /.dbkenv
# Format: RUN_ARGS=arg1 arg2 arg3
RUN echo RUN_ARGS=index.js >> /.dbkenv
RUN echo WORKDIR=/code >> /.dbkenv
# Relative to the WORKDIR env.
RUN echo ENTRYPOINT=index.js >> /.dbkenv

WORKDIR /
