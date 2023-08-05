FROM nginx:mainline-alpine AS base

FROM node:18-alpine AS build
ARG SERVER_URL
ENV VITE_SERVER_URL $SERVER_URL
WORKDIR /app/
RUN corepack enable
COPY package.json .
COPY pnpm-lock.yaml .
RUN pnpm install
COPY . .
RUN pnpm run build 

FROM base AS final
COPY --from=build /app/dist /usr/share/nginx/html