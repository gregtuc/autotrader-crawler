# autotrader-crawler

## Quick Start

- Execute `docker-compose up -d` to start the program.

## Overview

- On execution, the container will collect all car data from autotrader.ca
- When complete, the data will be stored in a MongoDB collection named after the UNIX time in which the program first began running. The container will stop when complete.
- The MongoDB container is backed up in a Volume.
