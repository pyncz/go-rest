set -e

mongo << EOF
use $MONGO_INITDB_DATABASE

db.createUser({
  user: '$MONGO_INITDB_USERNAME',
  pwd: '$MONGO_INITDB_PASSWORD',
  roles: [{
    role: 'readWrite',
    db: '$MONGO_INITDB_DATABASE'
  }]
})

// Create required collections
db.createCollection('tasks', { capped: false })
db.createCollection('tags', { capped: false })
db.createCollection('items', { capped: false })

// Set default indexes
db.tags.createIndex({ slug: 1 }, { unique: true })
EOF
