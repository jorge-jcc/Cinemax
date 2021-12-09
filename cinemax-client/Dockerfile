# etapa de compilación
#FROM node:alpine as build-stage
#WORKDIR /app
#COPY package*.json ./
#RUN yarn install
#COPY . .
#RUN yarn build

# etapa de producción

FROM nginx:alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY ./dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]