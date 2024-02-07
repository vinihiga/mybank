# MyBank - Backend
**A self-study project to learn Go Lang, Microservices Architecture and Relational Databases.**

The idea of the project is to create 2 simple services based on **zanfranceschi/rinha-de-backend-2024-q1** where we must get the statement of certain bank user and also to allow simple transactions based on a brazilian banking system where the user can have "credits" (in portuguese "limite de saldo").

I didn't want to participate in the competition "rinha-de-backend" because my idea is to learn how to develop backend systems and my main focus is to develop front-end mobile Apps for Apple ecosystem. Perhaps in the future I can participate. Who knows. 😁

## The Architectural Design

We are using Nginx as a proxy that will also work as load balancer. Then we have 2 microservices where are going to handle 2 endpoints:

*- [POST] /clientes/{id}/transacoes*

*- [GET] /clientes/{id}/extrato*

Where the first one is to handle the actions the user can do like "withdraw" or "deposit" and the second one to get the full "statement".

In the future we are going to have a simple relation database but it's still in WIP.

## Code Complexity

Oh yes, in this code you are going to see some Big-O notations, because I want to also practice asymptotic analysis.

## Author

Vinícius Hiroshi Higa - @vinihiga
