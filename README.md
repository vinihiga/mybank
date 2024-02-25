# MyBank - Backend
**A self-study project to learn Go Lang, Microservices Architecture and Relational Databases.**

This project aims to develop two basic backend services based on the "Rinha de Backend 2024 Q1" challenge (https://github.com/zanfranceschi/rinha-de-backend-2024-q1/issues), focusing on:

* Retrieving bank statements: Fetching account statements for specific users.

* Simple transactions: Enabling basic transactions within a Brazilian banking system context, taking into account "limite de saldo" (credit limit).

While the "Rinha de Backend" competition itself wasn't my intended participation due to my focus on building front-end mobile apps for the Apple ecosystem, this project serves as a learning experience for backend development. Future participation in such competitions may be considered.

## Architectural Design
### Components:

* Nginx: Acts as both a reverse proxy and load balancer, distributing incoming requests across microservices.

* 2 instances of microservices. Each one having both:
    * Transactions Service: Handles user actions like "withdraw" and "deposit" via the [/clientes/{id}/transacoes]() (POST) endpoint.
    * Statements Service: Retrieves full user statements via the [/clientes/{id}/extrato]() (GET) endpoint.

* Database: A simple relational database using PostgreSQL. See "sql" folder for the scripts.

## Code Complexity

Oh yes, in this code you are going to see some Big-O notations, because I want to also practice asymptotic analysis.

## Author

Vinícius Hiroshi Higa - @vinihiga
