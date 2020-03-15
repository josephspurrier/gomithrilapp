---
id: database
title: Database
---

## Database Migrations

MySQL migrations are performed at boot by [Rove](https://github.com/josephspurrier/rove), a tool very similiar to Liquibase. The migrations are run when API starts up - they are located [here](https://github.com/josephspurrier/govueapp/blob/master/src/app/api/migration/changeset.go).