FROM node:20-slim

WORKDIR /app

# Install requirements
COPY package.json package-lock.json ./
RUN npm install

# Copy files
COPY . ./

# For Docker healthchecking
RUN apt-get update && apt-get install wget -y
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.19/grpc_health_probe-linux-amd64 && chmod +x /bin/grpc_health_probe
HEALTHCHECK CMD ["/bin/grpc_health_probe", "-addr=127.0.0.1:9090"]

EXPOSE 9090
CMD ["npm", "run", "start"]