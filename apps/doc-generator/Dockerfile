FROM node:22 as builder
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

# Use NodeJS server for the app.
FROM node:18 
# install chrome for puppeteer
RUN apt-get update \
    && apt-get install -y wget gnupg \
    && wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list' \
    && apt-get update \
    && apt-get install -y google-chrome-stable fonts-ipafont-gothic fonts-wqy-zenhei fonts-thai-tlwg fonts-kacst fonts-freefont-ttf libxss1 \
      --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

RUN npm install -g npm@latest
# Copy files as a non-root user. The `node` user is built in the Node image.
WORKDIR /app
RUN chown node:node ./
USER node
# Defaults to production, docker-compose overrides this to development on build and run.
ARG NODE_ENV=production
ENV NODE_ENV $NODE_ENV
# Install dependencies first, as they change less often than code.
COPY --chown=node:node package*.json ./
RUN npm ci
COPY --chown=node:node . .
COPY --chown=node:node --from=builder /app/dist /app/dist
CMD ["node", "dist/index.js"]
