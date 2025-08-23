FROM node:22 AS builder
WORKDIR /src
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run lint && npm run build

FROM node:22 AS runner
WORKDIR /app
COPY --from=builder /src/. /app
EXPOSE 5173
EXPOSE 4173
HEALTHCHECK --interval=10s --timeout=1s --retries=3 \
  CMD curl -f http://localhost:5173/health || exit 1
CMD ["npm", "run", "dev"]
