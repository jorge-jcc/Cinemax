FROM node:lts-alpine

# make the 'app' folder the current working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH - allow running 'vite'
ENV PATH /app/node_modules/.bin:$PATH

RUN yarn global add @vue/cli

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./

# install project dependencies
RUN yarn install

RUN chmod -R 777 node_modules/

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

EXPOSE 3000
CMD [ "yarn", "serve" ]