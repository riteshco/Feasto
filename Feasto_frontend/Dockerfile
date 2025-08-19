# ---- Stage 1: Build ----
FROM node:24-alpine AS builder

WORKDIR /app

# Accept the API URL as a build argument
ARG VITE_API_BASE_URL

# Create a .env file with the build argument
RUN echo "VITE_API_BASE_URL=${VITE_API_BASE_URL}" > .env

COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build

# ---- Stage 2: Serve ----
FROM nginx:stable-alpine

# Copy the built assets from the 'builder' stage
COPY --from=builder /app/dist /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]