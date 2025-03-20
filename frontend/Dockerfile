FROM node:22-alpine AS builder
WORKDIR /src
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

FROM node:22-alpine AS runner
WORKDIR /app
COPY --from=builder /src/. /app
EXPOSE 5173
# TODO: /health endpoint
# RUN apk add curl
# HEALTHCHECK --interval=10s --timeout=1s --retries=3 \
#   CMD curl -f http://localhost:5173/health || exit 1
CMD ["npm", "run", "dev"]
