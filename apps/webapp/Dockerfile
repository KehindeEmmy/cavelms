#############################################
# App Builder: From developrnt to production
#############################################
FROM node:lts-alpine as builder
WORKDIR /app

# BUILD WEBAPP
COPY ./apps/webapp/package*.json ./webapp/
RUN cd webapp && npm install

COPY ./apps/webapp/prisma ./webapp/prisma
RUN cd webapp && npx prisma generate

COPY ./apps/webapp ./webapp
RUN cd webapp && npm run build
RUN cd webapp && npm prune --production

#############################################
# App Runner: Production image runner
#############################################
FROM node:lts-alpine as webapp

WORKDIR /app

COPY --from=builder /app/webapp/build ./build/
COPY --from=builder /app/webapp/prisma ./prisma/
COPY --from=builder /app/webapp/package.json ./package.json
COPY --from=builder /app/webapp/node_modules ./node_modules/

ENV NODE_ENV=production
CMD [  "npm", "run", "prod:migrate" ]