FROM docker.io/node:24.11.0-slim AS base

WORKDIR /home
RUN --mount=type=cache,target=/root/.npm \
    --mount=type=bind,source=./package.json,target=./package.json \
    --mount=type=bind,source=./package-lock.json,target=./package-lock.json \
    npm ci

FROM base as build

WORKDIR /home
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=build /home/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
