FROM node:16

WORKDIR /usr/src/app

RUN npm i -g pnpm

COPY package*.json pnpm-lock.yaml ./

RUN pnpm install

COPY prisma/schema.prisma ./prisma/

RUN npx prisma generate

COPY . .

CMD ["yarn", "dev", "--host", "0.0.0.0"]
