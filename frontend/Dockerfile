FROM node:18.10-alpine3.15 AS build

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]


# FROM node:20 AS build

# WORKDIR /app

# COPY . .

# # Pass environment variables explicitly (needed for Vite)
# ARG VITE_BACKEND_URL
# ARG VITE_STATISTICS_URL

# ENV VITE_BACKEND_URL=$VITE_BACKEND_URL
# ENV VITE_STATISTICS_URL=$VITE_STATISTICS_URL

# RUN npm ci
# RUN npm run build

# # Serve with nginx (optional step)
# FROM nginx:alpine
# COPY --from=build /app/dist /usr/share/nginx/html
# CMD ["nginx", "-g", "daemon off;"]
