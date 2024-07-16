# Project Management / Organisation Tool

## Idea

todo: naming/branding

A tool to simplify all your organisation needs.

Cater to wide audience (individuals, freelancers/contractors, startups, small/medium businesses, large organisations) - focus on customisation to help meet the best solution to meet needs in each places.

* Organisation could have specific needs, then teams in the organisation can customise to fit them best

plan: offer as hosted service + in future allow to be self-hosted / open source 



## Features/Functionality

### Homepage / User Dashboard

* Customisable by the user for their needs
  * User can select their role(s) "developer", "manager", "designer", etc
    * Help customise the experience based on the role
  * User can add widgets such as:
    * Main goal for today / main tasks (did you achieve that goal? - at the end of the day)

* Display of the user's teams / groups
* Display of the user's "active" / "current" / "upcoming" projects
* User's calendar / overview
  * Display of deadlines from projects
  * Display of synced calendar events (i.e. user's meetings, etc)
  * Display of calendar events in all their teams / org calendars
  * Display of any clashes
  * Events which want an RSVP response



### Syncing from Other Sources

* User can sync their calendars (sync through CalDAV)
* User can sync Wakatime / Wakapi data
* GitHub data can be synced
  * Commits brought in, etc
* Each group / team / organisation have their own calendars
  * Events can be added that can be seen by all members (i.e. company wide meeting, etc)
  * Collaborative organisation boards, etc
* With project boards, etc - syncing from places such as Trello, Linear, Asana, etc (/ migrating from)
* Wakatime compatible API for recieving from wakatime-cli directly
  *  https://github.com/wakatime/wakatime-cli



### Groupings

* Users can create their own accounts.
  * -> user dashboard

* Users can be invited to organisations / added to teams
  * Added TO the team / added as an outsider, etc (permissioning system)
  * Can share / attach info to/with certain orgs / teams
  * Summaries / overviews can be generated at a team / organisation level
* Organisation managers can receive email summaries, etc



### Teams / Products / Projects

* Product / Project Roadmaps
* Ticket boards



### Time Records / Time Sheets

* Ability to create records from timer
  * Timer functionality can be expanded (running in background, automatically stopping with inactivity, etc)

* Use data from Wakapi, Google Calendar + more as basis to build up
  * Summaries of the day from Wakapi, etc
* Users can reference info from other places within their time records (i.e. referencing branches & commits - from GitHub pull requests and commits, and wakatime / wakapi)
* Time records can be labelled / focused on specific projects (i.e. time record only for the user - time record from the user added to organisation, referencing project in organisation)
  * Organisation enable time records as a feature
  * User can create their own plan / "organisation" (better naming to encapsulate the two use cases), which simplifies sustainability approach



### Sustainability / marketability / pricing

* Organisations could be billed per user / based on features being used
* Modularise some features and not include in OS release



## Tech Brainstorm

### Core

* https://kit.svelte.dev/ (combined with Go)

* Postgres / NoSQL (leaning towards NoSQL for some elements)

* gRPC + Protobufs
  * with grpc-gateway
  * strong schema-based models (+ grpc-gateway's generated openapi schema)
    * openapi schema -> svelete
    * https://dev.to/fbjorn/a-typed-http-client-for-sveltekit-88b
    * https://github.com/openapi-ts/openapi-typescript
  * wrap calls to other services internally (all to `localhost` for now - then the services can be split apart / scaled separately if/when needed)
* OTel for tracing
  * and error logging: https://opentelemetry.io/docs/specs/otel/trace/exceptions/

### Svelte Graph / Data Visualisation

* LayerCake: https://layercake.graphics/ (https://github.com/mhkeller/layercake)
* **LayerChart (uses LayerCake):** https://www.layerchart.com/ (https://github.com/techniq/layerchart)
  * Works without Svelte UX

* D3 (also used by LayerCake): https://d3js.org/

* Svelte Flow: https://svelteflow.dev/

### UI / Design

* Skeleton: https://www.skeleton.dev/docs/introduction
  * Seems too basic
* Flowbite: https://flowbite-svelte.com/
  * Seems okay, good candidate
* MaterialUI for Svelte: https://sveltematerialui.com/INSTALL.md
  * Not the nicest looking
* **Shadcn for Svelte: https://www.shadcn-svelte.com/**
  * Very nice. Like normal shad
* AgnosticUI: https://www.agnosticui.com/
  * Seems nice, but like a skeleton framework
* STWUI (svelte-tailwindcss UI): https://github.com/N00nDay/stwui
  * Seems okay, but seems early days



# Research

* Wakapi
* Wakatime
* Trello
* Toggl
* Asana
* Linear
* Productive
* Notion
* YouTrack
  * https://www.jetbrains.com/youtrack
* ProofHub
  * https://www.proofhub.com/solutions/marketing-project-management-software
  * https://www.proofhub.com/pricing