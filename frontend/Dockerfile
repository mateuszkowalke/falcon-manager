FROM node:16

WORKDIR /app
COPY package.json package-lock.json ./

RUN npm config set registry "http://registry.npmjs.org/"
RUN npm install

COPY . .

RUN npm run dev

