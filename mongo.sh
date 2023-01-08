#!/bin/bash

mongo <<EOF
use cliente-oculto-auth
db.createUser({
  user: "admin",
  pwd: "admin",
  roles: [
    {
      role: "readWrite",
      db: "cliente-oculto-auth"
    }
  ]
})
EOF