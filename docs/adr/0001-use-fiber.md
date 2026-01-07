# ADR-0001: Use Fiber as the HTTP framework
**Status:** Accepted  
**Date:** 2026-01-07  
**Owner:** Team

## Context
We need a Go HTTP framework to implement REST APIs with routing, middleware, and consistent error handling.
The project allows Fiber/Gin/Echo. Team familiarity and speed matter due to the 2-week timeline.

## Decision
We will use **Fiber** as the HTTP framework.

## Alternatives considered
- **Gin**: very common and stable, good ecosystem
- **Echo**: solid framework with clean APIs

## Consequences
- All HTTP routing/middleware will be implemented using Fiber.
- Handlers will remain thin and delegate business rules to service layer (clean layering).
- Fiber-specific code must be limited to the HTTP layer (`internal/http/...`) to keep business logic testable.

## Notes
- We will implement a centralized Fiber error handler to enforce a consistent error response format.

