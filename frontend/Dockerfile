FROM node:22-alpine AS builder
WORKDIR /src
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

FROM node:22-alpine AS runner
WORKDIR /bin
COPY --from=builder /src .
EXPOSE 5173
CMD ["npm", "run", "dev"]
