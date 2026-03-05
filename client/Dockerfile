# install node and build Angular
FROM node:18-alpine AS build
WORKDIR /app

COPY package*.json ./

ENV CYPRESS_INSTALL_BINARY=0

RUN npm install

COPY . .

RUN npm run build --prod

# serve built app using nginx
FROM nginx:alpine
COPY --from=build /app/dist/client /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]]
