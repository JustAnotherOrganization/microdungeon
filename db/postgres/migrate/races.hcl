table "races" {
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
    null = false
  }

  column "abilities" {
    type = bigint
    null = false
  }
  foreign_key "abilities" {
    columns = [column.abilities]
    ref_columns = [table.abilities.column.id]
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

  index "race_name" {
    columns = [
      column.name
    ]
    unique = true
  }

  # FIXME: should index attributes...
}