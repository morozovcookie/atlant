Atlant Server
=================

This project represent my own variant of realization test from Tensigma LTD. Full description yoy can find [here (RU)](DESCRIPTION_RU.md).

<!-- place badges here -->


# Table of Contents

- [Requirements](#requirements)
- [Architecture](#architecture)
    - [Full picture](#full-picture)
    - [Fetch process](#fetch-process)
    - [List process](#list-process)
- [Usage](#usage)
- [TODO](#todo)


# Requirements

Below you can find list of minimal requirements of items which I used for building this project:

- GoLang >=1.15.3
- GoLangCI-Lint >=1.31.0
- Docker >= 19.03.13
- Docker-Compose >= 1.27.4


# Architecture

In this section you can find a visualisation of how it should work (and maybe some comments about why I choose exactly that method or another).

## Full picture

Let is start from the full picture. This picture represents all components involved in this project.

<!-- place for full picture -->

For the basic two process I decided make a two diagrams: component and sequence, and these diagrams could describe each process very accurate.

## Fetch process

Fetch process is a process when you pass csv file url (only http(s) available) to the server which download, parse and save files data into MongoDB.  
I decided to use the Apache Kafka and [the second service](cmd/processor) for guaranteed and asynchronous file data processing. 
In this case we have more freedom to extend file processing logic because all this stuff will be in the background, hide from the client.

<!-- place for on fetch process -->

<!-- place for UML diagram -->

## List process

This process is very simple: client make a call, atlantserver receive this call, retrieve data from MongoDB and return it back to the client.
In this case we have situation when two services share one database and one collection and this is normal, 
because first service (processor) using database only for "write" operations, and the second (atlantserver) - only for "read" operations.

<!-- place for on list process -->

<!-- place for UML diagram -->


# Usage

This project contains for executable files:

- [atlantclient](cmd/atlantclient/README.md) - console client for interaction with atlantserver
- [atlantserver](cmd/atlantserver/README.md) - main but not only one part of this project
- [processor](cmd/processor/README.md) - this little service responsible for saving line from csv-file into MongoDB

More about each project component you can find in specified README files.

For using all project components you can use next commands:

```bash
# Build and run containers, create topic and apply migrations
$ make run


# Stop project
$ make stop
```

# TODO

Here is list of some features that could be implemented in the future:

- [x] Finish docker-compose deploy
- [ ] Metrics
- [ ] Liveness/Readiness probes
- [ ] Producer and consumer compression
- [ ] CI/CD
- [ ] Deploy to k8s
- [ ] Build Docker images using werf
- [ ] More tests:
    - [ ] Unit tests
    - [ ] Integration tests
    - [ ] Functional tests
- [ ] GitHub Actions
- [ ] Code docs
- [x] Use earliest offset in Kafka and changes history for indempotence
- [x] Write changes history on products fetching
- [ ] Move Kafka from wurstmeister to confluentc
- [ ] Support more protocols, not only http(s)
- [ ] Use high-availability MongoDB cluster
- [ ] Move from Nginx ingres to Envoy mesh network
- [ ] Replace JSON with Avro
- [ ] Configurable encoders and decoders
- [ ] Verify Kafka Cluster high availability
- [ ] Circuit breaker for file fetcher
