FROM docker.io/node:24.11.0-slim AS base

WORKDIR /home

RUN --mount=type=cache,target=/root/.npm \
    --mount=type=bind,source=./package.json,target=./package.json \
    --mount=type=bind,source=./package-lock.json,target=./package-lock.json \
    npm ci

FROM base as release

WORKDIR /home

COPY . .

ENTRYPOINT ["npm", "start"]