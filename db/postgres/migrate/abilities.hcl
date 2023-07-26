table "abilities" {
  schema = schema.microdungeon

  column "id" {
    type = serial
    null = false
  }

  column "strength" {
    type = bigint
    null = false
  }

  column "dexterity" {
    type = bigint
    null = false
  }

  column "constitution" {
    type = bigint
    null = false
  }

  column "intelligence" {
    type = bigint
    null = false
  }

  column "wisdom" {
    type = bigint
    null = false
  }

  column "charisma" {
    type = bigint
    null = false
  }

  primary_key {
    columns = [
      column.id
    ]
  }
}