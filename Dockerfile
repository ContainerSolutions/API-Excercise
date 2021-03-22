FROM node:14
COPY /swagger/package*.json ./
RUN npm install --quiet
COPY /swagger ./
CMD ["npm", "start"]