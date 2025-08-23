FROM node:22-slim AS builder
WORKDIR /src
COPY package.json yarn.lock ./
RUN yarn install
COPY . .
RUN yarn run lint && yarn run build

FROM node:22-slim AS runner
WORKDIR /app
COPY --from=builder /src/. /app
EXPOSE 5173
EXPOSE 4173
HEALTHCHECK --interval=10s --timeout=1s --retries=3 \
  CMD curl -f http://localhost:5173/health || exit 1
CMD ["yarn", "run", "dev"]
