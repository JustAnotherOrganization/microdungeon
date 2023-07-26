table "classes" {
  schema = schema.microdungeon

  column "id" {
    type = serial
    null = false
  }

  column "name" {
    type = varchar(255)
    null = false
  }

  column "hit_die" {
    type = bigint
    null = false
  }

  column "primary_abilities" {
    type = sql("bigint[]")
    null = false
  }

  column "saves" {
    type = sql("bigint[]")
    null = true
  }

  primary_key {
    columns = [
      column.id
    ]
  }

  index "class_name" {
    columns = [
      column.name
    ]
    unique = true
  }
}
