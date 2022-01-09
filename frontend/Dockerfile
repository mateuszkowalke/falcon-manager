FROM node:16

WORKDIR /usr/src/app

COPY package.json ./

RUN npm install

COPY . .

CMD ["yarn", "dev", "--host", "0.0.0.0"]

