FROM node:17-alpine as build
WORKDIR /app
COPY package.json .
RUN npm i
COPY . .
EXPOSE 5173

# Start service
CMD ["npm", "run","dev"]