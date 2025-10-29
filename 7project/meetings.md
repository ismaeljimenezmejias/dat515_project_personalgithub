# Weekly Meeting Notes

- Group 19 - Local Bike Marketplace
- Mentor: Foroozan.E

## Administrative Info

- Date: 2025-10-29
- Attendees: Ismael Jimenez
- Notetaker: Ismael Jimenez

## Progress Update (Before Meeting)


### First meeting
Design and understanding project
4h

### 2025-10-22
- Implemented the persistent database.
- Finished the main logic of the web with the search and filtering options, the user account and profile (lacking authentication).
- Defined some goals.

8-10h

### 2025-10-29
- Finished logic within the web app, the messaging between users, location in products, and the authentication (Redis Session Status, approve the log in if the password matches) and authorization (e.g. being the user owner of the product to be able to delete it), doing the app with the Redis cloud deployment stateful (before that it used cookies managed in the client side).
- Deployed the database in Railway as PostgreSQL, the Redis cache in Railway and the backend connected with my personal repo in GitHub in Railway as well.
- Also used robustness and checks as creating the logic.

14-16h

### Coding

#### 2025-10-22
- HTML files in `/templates`. Changed `Dockerfile` and `docker-compose.yaml`.

#### 2025-10-29
- messages html file, `init_db_pg.py`, `Procfile`, `redis_client.py`, `migrations.py`. Updated the rest of the HTML files.

### Documentation

- In the design

## Questions and Topics for Discussion (Before Meeting)

1. Kubernetes: what's the expected deployment for this project? I currently use managed cloud services (DB/Redis) and don't have `deployment.yaml`/`service.yaml` manifests.
   
2. GCP credits: how do I access the $50 allocation? I only see the standard $300 free credits.
   
3. Any gaps in rubric coverage you recommend prioritizing next (tests, CI, docs, metrics)?
   
4. What's advances cloud services in the checklist?
   
5. How do I have to present the project? Showing the docker-compose and also the cloud deployment or just the latest deployed final version? For instance, I had implemented the docker compose but I don't use it anymore, should I ignore it in the presentation video?

6. How is the automated deployment made? 



## Discussion Notes (During Meeting)


## Action Items for Next Week (During Meeting)

- Optimized Dockerfile
  
- Automated deployment

- Multiple replicas

- Define Kubernetes deployment approach and minimal manifests (Deployment, Service, Ingress) if required.
  
- Add a basic test suite and CI workflow (lint/tests), GitHub Actions.
 
- Start writing report once I finished that