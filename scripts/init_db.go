#!/bin/bash

mongo <<EOF
use blogs
db.createCollection("blogs")
EOF