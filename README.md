# Requirements

* docker
* docker-compose

# Usage

We just need to run `docker-compose up` and the server will be accessible on http://localhost:8000/swaggerui/

```bash
docker-compose up
```

Cause I'm not a frontend developer I've provided only SwaggerUI to test API.
I've used MongoDB cause I wanted to try it our already for a long time but didn't have tasks where I can use it :) But it fits well for
this task, cause we don't have any relations, and good that we don't need to deal with migrations.

The next step will be to write unit tests for the service, we have all interfaces and mocks in place, but I didn't have time for that, cause the task was not to spend more than a day.

And that's what I can do within a day.

# Pento tech challenge

Thanks for taking the time to do our tech challenge.

The challenge is to build a small full stack web app, that can help a freelancer track their time.

It should satisfy these user stories:

- As a user, I want to be able to start a time tracking session
- As a user, I want to be able to stop a time tracking session
- As a user, I want to be able to name my time tracking session
- As a user, I want to be able to save my time tracking session when I am done with it
- As a user, I want an overview of my sessions for the day, week and month
- As a user, I want to be able to close my browser and shut down my computer and still have my sessions visible to me when I power it up again.

## Getting started

You can fork this repo and use the fork as a basis for your project. We don't have any requirements on what stack you use to solve the task, so there is nothing set up beforehand.

## Timing

- Don't spend more than a days work on this challenge. We're not looking for perfection, rather try to show us something special and have reasons for your decisions.
- Get back to us when you have a timeline for when you are done.

## Notes

- Please focus on code quality and showcasing your skills regarding the role you are applying to.
