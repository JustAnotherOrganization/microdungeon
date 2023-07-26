table "characters" {
  schema = schema.microdungeon

  column "id" {
    type = serial
    null = false
  }

  column "name" {
    type = varchar(255)
    null = false
  }

  column "player_name" {
    type = varchar(255)
    null = false
  }

  column "race_name" {
    type = varchar(255)
    null = false
  }
  foreign_key "race_name" {
    columns = [column.race_name]
    ref_columns = [table.races.column.name]
  }

  column "class_name" {
    type = varchar(255)
    null = false
  }
  foreign_key "class_name" {
    columns = [column.class_name]
    ref_columns = [table.classes.column.name]
  }

  column "level" {
    type = bigint
    null = false
  }

  column "alignment_ethical" {
    type = bigint
    null = false
  }

  column "alignment_moral" {
    type = bigint
    null = false
  }

  column "experience" {
    type = bigint
    null = false
  }

  column "abilities" {
    type = integer
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

  column "inventory" {
    type = sql("text[]")
    null = true
  }

  primary_key {
    columns = [
      column.id
    ]
  }

  index "character_name" {
    columns = [
      column.name
    ]
    unique = true
  }
}
