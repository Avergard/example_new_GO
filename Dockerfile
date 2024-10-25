version: '3.1'

services:
  db:
    image: postgres:14
    container_name: postgres_container
    environment:
      POSTGRES_USER: your_username
      POSTGRES_PASSWORD: your_password
      POSTGRES_DB: your_database
    ports:
      - "5432:5432"
    volumes:
      - ./postgres-data:/var/lib/postgresql/data