# Frontend Dockerfile
FROM node:16 AS frontend-build

WORKDIR /app/frontend
COPY ./package.json ./package-lock.json ./
RUN npm install
COPY . .
RUN npm run build

# Use a lightweight web server to serve the static files
FROM nginx:alpine
COPY --from=frontend-build /app/frontend/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
