FROM node:22

WORKDIR /app

#if changes in any file does not run again as uses cached resource

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host"]