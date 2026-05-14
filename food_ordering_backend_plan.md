# Food Ordering Platform Backend Plan

## Overview

Build a multi-tenant food ordering backend using:

- Go
- Gin
- GORM
- MySQL
- Redis
- RabbitMQ
- Docker

Architecture style:

- Modular Monolith
- Shared Database
- Shared Schema
- Subdomain-based multi-tenancy

---

# High-Level Architecture

```text
Clients
  ├── Customer App
  ├── Shop Admin Dashboard
  └── Super Admin Dashboard

        ↓

Go Backend API

        ↓

Modules
  ├── Auth
  ├── Tenant
  ├── Product
  ├── Category
  ├── Cart
  ├── Orders
  ├── Payments
  ├── Notifications
  └── Analytics

        ↓

Infrastructure
  ├── MySQL
  ├── Redis
  ├── RabbitMQ
  └── Object Storage
```

---

# Phase 1 — Foundation

## Goals

- Initialize Go project
- Set up infrastructure
- Create project structure
- Configure Docker
- Create health check endpoint

---

## Tech Stack

### Backend

- Go
- Gin
- GORM

### Database

- MySQL
- Redis

### Messaging

- RabbitMQ

### Infrastructure

- Docker
- Docker Compose

---

## Initial Dependencies

```bash
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/redis/go-redis/v9
go get github.com/rabbitmq/amqp091-go
```

---

## Folder Structure

```text
backend/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── auth/
│   ├── tenant/
│   ├── product/
│   ├── category/
│   ├── order/
│   ├── cart/
│   ├── middleware/
│   ├── database/
│   └── shared/
│
├── configs/
├── migrations/
├── scripts/
├── pkg/
├── go.mod
└── docker-compose.yml
```

---

# Phase 2 — Multi-Tenancy

## Multi-Tenancy Strategy

- Shared Database
- Shared Schema

Each tenant-owned table contains:

```sql
tenant_id
```

---

## Tenant Flow

```text
pizza.example.com
```

↓

Extract subdomain:

```text
pizza
```

↓

Find tenant:

```sql
SELECT * FROM tenants
WHERE subdomain = 'pizza';
```

↓

Attach `tenant_id` to request context.

---

## Tenant Middleware Responsibilities

- Extract subdomain
- Resolve tenant
- Validate tenant
- Attach tenant_id to context

---

## Important Rule

Every tenant-owned query MUST filter by tenant_id.

### Bad

```sql
SELECT * FROM products WHERE id = 1;
```

### Good

```sql
SELECT * FROM products
WHERE id = 1
AND tenant_id = 2;
```

---

# Phase 3 — Authentication & Authorization

## User Types

### Customer

Can:
- Browse products
- Place orders

### Shop Admin

Can:
- Manage products
- Manage categories
- Manage orders

### Super Admin

Can:
- Manage tenants
- Monitor platform

---

## Auth Strategy

Use:

- JWT Access Tokens
- Refresh Tokens

---

## JWT Payload

```json
{
  "user_id": 1,
  "tenant_id": 2,
  "role": "SHOP_ADMIN"
}
```

---

## Middleware Stack

```text
TenantMiddleware
AuthMiddleware
RoleMiddleware
```

---

# Phase 4 — Product Module

## Features

### Admin Features

- Create products
- Update products
- Delete products
- Upload product images
- Manage availability

### Customer Features

- Browse menu
- Search products
- Filter categories

---

## Database Tables

```text
products
categories
product_variants
product_images
```

All tables contain:

```sql
tenant_id
```

---

# Phase 5 — Cart System

## Recommended Approach

Use Redis-backed carts.

Benefits:
- Fast
- Scalable
- Temporary storage

---

## Cart Example

```json
{
  "user_id": 1,
  "items": [
    {
      "product_id": 10,
      "quantity": 2
    }
  ]
}
```

---

# Phase 6 — Order System

## Order Lifecycle

```text
PENDING
→ CONFIRMED
→ PREPARING
→ READY
→ DELIVERED
```

---

## Order Tables

```text
orders
order_items
payments
```

---

## Important Rule

Copy product data into order items during order creation.

Do NOT rely on live product data later.

---

# Phase 7 — Async Processing

Use RabbitMQ for:

- Notifications
- Email sending
- Analytics
- Invoice generation
- Background jobs

---

## Event Flow

```text
API
  ↓
Publish Event
  ↓
Worker
  ↓
Process
```

---

# Phase 8 — Payments

## Payment Flow

```text
Create Order
↓
Create Payment Intent
↓
Payment Gateway
↓
Webhook Verification
↓
Mark Order Paid
```

---

## Important Rule

Never trust frontend payment success.

Always verify payments using webhooks.

---

# Phase 9 — Notifications

## Channels

- Email
- SMS
- Push notifications

Use event-driven architecture.

---

# Phase 10 — Observability

## Logging

Use:
- zap
- slog

---

## Metrics

Use:
- Prometheus

---

## Tracing

Use:
- OpenTelemetry

---

# Security Checklist

## Passwords

Use:
- bcrypt
- argon2

Never store plain text passwords.

---

## Tenant Validation

Always validate tenant ownership.

---

## Rate Limiting

Protect:
- Login
- Register
- OTP endpoints

---

# API Structure

## Public APIs

```text
GET /menu
GET /products
```

---

## Auth APIs

```text
POST /auth/login
POST /auth/register
```

---

## Admin APIs

```text
POST /admin/products
PUT /admin/products/:id
DELETE /admin/products/:id
```

---

# Database Best Practices

## Standard Columns

Every table should contain:

```sql
created_at
updated_at
deleted_at
```

---

## IDs

Internal:
- BIGINT

External:
- UUID

---

# Recommended Development Order

## Stage 1

- Project setup
- Docker setup
- Database setup
- Tenant middleware
- Health check API

---

## Stage 2

- Authentication
- Authorization
- User management

---

## Stage 3

- Categories
- Product CRUD
- Image uploads

---

## Stage 4

- Cart
- Orders
- Order status flow

---

## Stage 5

- RabbitMQ workers
- Notifications
- Payments

---

## Stage 6

- Analytics
- Monitoring
- Performance optimization

---

# Recommended MVP Scope

Build this first:

- Tenant system
- Authentication
- Product CRUD
- Customer menu
- Cart
- Order placement

---

# Key Focus Areas

Most important backend challenges:

1. Multi-tenancy isolation
2. Order workflow consistency
3. Authorization
4. Async processing

---

# Recommended Architecture Style

Use:

- Clean modular architecture
- Practical abstractions
- Modular monolith

Avoid:

- Premature microservices
- Over-engineering
- Excessive abstraction

---

# Final Recommendation

Focus on:
- Shipping MVP quickly
- Keeping architecture clean
- Strong tenant isolation
- Async processing patterns

Scale complexity only when necessary.
