FROM node:alpine

# set the work path
RUN mkdir /frontend
WORKDIR /frontend

# copy package.json and lock to run npm install later
COPY package*.json /frontend

RUN npm install

# copy everything to the work path
COPY . /frontend

EXPOSE 3000

CMD ["npm", "run", "dev"]
