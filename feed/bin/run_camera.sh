#!/bin/bash

sudo venv/bin/gunicorn --timeout 3600 -w 2 -b 0.0.0.0:9090 --error-logfile - --access-logfile - feed:app
