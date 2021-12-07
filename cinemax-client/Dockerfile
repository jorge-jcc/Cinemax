# etapa de compilación
#FROM node:9.11.1-alpine as build-stage
#WORKDIR /app
#COPY package*.json ./
#RUN yarn install
#COPY . .
#RUN yarn build

# etapa de producción
FROM nginx:1.13.12-alpine as production-stage
COPY ./dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]