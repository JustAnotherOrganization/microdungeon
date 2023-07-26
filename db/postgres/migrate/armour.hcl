table "armour" {
  schema = schema.microdungeon

  column "id" {
    type = serial
    null = false
  }

  column "name" {
    type = varchar(255)
    null = false
  }

  column "aliases" {
    type = sql("text[]")
    null = true
  }

  column "description" {
    type = varchar(255)
    null = true
  }

  column "weight" {
    type = bigint
    null = false
  }

  column "type" {
    type = bigint
    null = false
  }

  column "ac" {
    type = bigint
    null = false
  }

  column "worn" {
    type = bool
    null = false
  }

  column "attributes" {
    type = jsonb
    null = true
  }

  primary_key {
    columns = [
      column.id
    ]
  }

  index "armour_name" {
    columns = [
      column.name
    ]
    unique = true
  }
}