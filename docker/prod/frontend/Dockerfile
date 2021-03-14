# ref. https://cli.vuejs.org/guide/deployment.html#docker-nginx

# build-stage

FROM node:14-alpine as builder

RUN mkdir /frontend

WORKDIR /frontend

COPY frontend .

RUN yarn && yarn build


# deploy-stage

FROM nginx:1-alpine

COPY --from=builder /frontend/dist /frontend

COPY docker/prod/frontend/nginx.conf /etc/nginx/nginx.conf