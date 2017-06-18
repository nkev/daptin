package datastore

import (
  "github.com/artpar/api2go"
  "github.com/artpar/goms/server/resource"
  "github.com/artpar/goms/server/fsm_manager"
)

var StandardColumns = []api2go.ColumnInfo{
  {
    Name:            "id",
    ColumnName:      "id",
    DataType:        "INTEGER",
    IsPrimaryKey:    true,
    IsAutoIncrement: true,
    ExcludeFromApi:  true,
    ColumnType:      "id",
  },
  {
    Name:         "created_at",
    ColumnName:   "created_at",
    DataType:     "timestamp",
    DefaultValue: "current_timestamp",
    ColumnType:   "datetime",
    IsIndexed:    true,
  },
  {
    Name:       "updated_at",
    ColumnName: "updated_at",
    DataType:   "timestamp",
    IsIndexed:  true,
    IsNullable: true,
    ColumnType: "datetime",
  },
  {
    Name:           "deleted_at",
    ColumnName:     "deleted_at",
    DataType:       "timestamp",
    ExcludeFromApi: true,
    IsIndexed:      true,
    IsNullable:     true,
    ColumnType:     "datetime",
  },
  {
    Name:       "reference_id",
    ColumnName: "reference_id",
    DataType:   "varchar(40)",
    IsIndexed:  true,
    ColumnType: "alias",
  },
  {
    Name:       "permission",
    ColumnName: "permission",
    DataType:   "int(11)",
    IsIndexed:  false,
    ColumnType: "value",
  },
  {
    Name:         "status",
    ColumnName:   "status",
    DataType:     "varchar(20)",
    DefaultValue: "'pending'",
    IsIndexed:    true,
    ColumnType:   "state",
  },
}

var StandardRelations = []api2go.TableRelation{
  api2go.NewTableRelation("world_column", "belongs_to", "world"),
  api2go.NewTableRelation("action", "belongs_to", "world"),
  api2go.NewTableRelation("world", "has_many", "smd"),
}

var SystemSmds = []fsm_manager.LoopbookFsmDescription{

}
var SystemActions = []resource.Action{
  {
    Name:   "upload_system_schema",
    Label:  "Upload features",
    OnType: "world",
    InFields: []api2go.ColumnInfo{
      {
        Name:       "Schema JSON file",
        ColumnName: "schema_json_file",
        ColumnType: "file.json",
        IsNullable: false,
      },
    },
    OutFields: []resource.Outcome{
      {
        Type:   "system_json_schema_update",
        Method: "EXECUTE",
        Attributes: map[string]string{
          "json_schema": "$schema_json_file",
        },
      },
    },
  },
  {
    Name:     "download_system_schema",
    Label:    "Download system schema",
    OnType:   "world",
    InFields: []api2go.ColumnInfo{},
    OutFields: []resource.Outcome{
      {
        Type:       "system_json_schema_download",
        Method:     "EXECUTE",
        Attributes: map[string]string{},
      },
    },
  },
  {
    Name:     "invoke_become_admin",
    Label:    "Become GoMS admin",
    OnType:   "world",
    InFields: []api2go.ColumnInfo{},
    OutFields: []resource.Outcome{
      {
        Type:   "become_admin",
        Method: "EXECUTE",
        Attributes: map[string]string{
          "user_id": "$user.id",
        },
      },
    },
  },
  {
    Name:   "signup",
    Label:  "Sign up on Goms",
    OnType: "user",
    InFields: []api2go.ColumnInfo{
      {
        Name:       "name",
        ColumnName: "name",
        ColumnType: "label",
        IsNullable: false,
      },
      {
        Name:       "email",
        ColumnName: "email",
        ColumnType: "email",
        IsNullable: false,
      },
      {
        Name:       "password",
        ColumnName: "password",
        ColumnType: "password",
        IsNullable: false,
      },
      {
        Name:       "Password Confirm",
        ColumnName: "passwordConfirm",
        ColumnType: "password",
        IsNullable: false,
      },
    },
    OutFields: []resource.Outcome{
      {
        Type:      "user",
        Method:    "POST",
        Reference: "user",
        Attributes: map[string]string{
          "name":     "$name",
          "email":    "$email",
          "password": "$password",
        },
      },
      {
        Type:      "usergroup",
        Method:    "POST",
        Reference: "usergroup",
        Attributes: map[string]string{
          "name": "!'Home group for ' + user.name",
        },
      },
      {
        Type:      "user_user_id_has_usergroup_usergroup_id",
        Method:    "POST",
        Reference: "user_usergroup",
        Attributes: map[string]string{
          "user_id":      "$user.reference_id",
          "usergroup_id": "$usergroup.reference_id",
        },
      },
    },
  },
  {
    Name:   "signin",
    Label:  "Sign in to Goms",
    OnType: "user",
    InFields: []api2go.ColumnInfo{
      {
        Name:       "email",
        ColumnName: "email",
        ColumnType: "email",
        IsNullable: false,
      },
      {
        Name:       "password",
        ColumnName: "password",
        ColumnType: "password",
        IsNullable: false,
      },
    },
    OutFields: []resource.Outcome{
      {
        Type:   "jwt.token",
        Method: "EXECUTE",
        Attributes: map[string]string{
          "email":    "$email",
          "password": "$password",
        },
      },
    },
  },
}

var StandardTables = []TableInfo{
  {
    TableName: "world",
    IsHidden:  true,
    Columns: []api2go.ColumnInfo{
      {
        Name:       "table_name",
        ColumnName: "table_name",
        IsNullable: false,
        IsUnique:   true,
        DataType:   "varchar(200)",
        ColumnType: "name",
      },
      {
        Name:           "schema_json",
        ColumnName:     "schema_json",
        DataType:       "text",
        IsNullable:     false,
        ExcludeFromApi: true,
        ColumnType:     "json",
      },
      {
        Name:         "default_permission",
        ColumnName:   "default_permission",
        DataType:     "int(4)",
        IsNullable:   false,
        DefaultValue: "644",
        ColumnType:   "value",
      },

      {
        Name:         "is_top_level",
        ColumnName:   "is_top_level",
        DataType:     "bool",
        IsNullable:   false,
        DefaultValue: "true",
        ColumnType:   "truefalse",
      },
      {
        Name:         "is_hidden",
        ColumnName:   "is_hidden",
        DataType:     "bool",
        IsNullable:   false,
        DefaultValue: "false",
        ColumnType:   "truefalse",
      },
    },
  },
  {
    TableName: "world_column",
    IsHidden:  true,
    Columns: []api2go.ColumnInfo{
      {
        Name:       "name",
        ColumnName: "name",
        DataType:   "varchar(100)",
        IsIndexed:  true,
        IsNullable: false,
        ColumnType: "name",
      },
      {
        Name:       "column_name",
        ColumnName: "column_name",
        DataType:   "varchar(100)",
        IsIndexed:  true,
        IsNullable: false,
        ColumnType: "name",
      },
      {
        Name:       "column_type",
        ColumnName: "column_type",
        DataType:   "varchar(100)",
        IsNullable: false,
        ColumnType: "label",
      },
      {
        Name:         "is_primary_key",
        ColumnName:   "is_primary_key",
        DataType:     "bool",
        IsNullable:   false,
        DefaultValue: "false",
        ColumnType:   "truefalse",
      },
      {
        Name:         "is_auto_increment",
        ColumnName:   "is_auto_increment",
        DataType:     "bool",
        IsNullable:   false,
        DefaultValue: "false",
        ColumnType:   "truefalse",
      },
      {
        Name:         "is_indexed",
        ColumnName:   "is_indexed",
        DataType:     "bool",
        IsNullable:   false,
        DefaultValue: "false",
        ColumnType:   "truefalse",
      },
      {
        Name:         "is_unique",
        ColumnName:   "is_unique",
        DataType:     "bool",
        IsNullable:   false,
        DefaultValue: "false",
        ColumnType:   "truefalse",
      },
      {
        Name:         "is_nullable",
        ColumnName:   "is_nullable",
        DataType:     "bool",
        IsNullable:   false,
        DefaultValue: "false",
        ColumnType:   "truefalse",
      },
      {
        Name:         "is_foreign_key",
        ColumnName:   "is_foreign_key",
        DataType:     "bool",
        IsNullable:   false,
        DefaultValue: "false",
        ColumnType:   "truefalse",
      },
      {
        Name:         "include_in_api",
        ColumnName:   "include_in_api",
        DataType:     "bool",
        IsNullable:   false,
        DefaultValue: "true",
        ColumnType:   "truefalse",
      },
      {
        Name:       "foreign_key_data",
        ColumnName: "foreign_key_data",
        DataType:   "varchar(100)",
        IsNullable: true,
        ColumnType: "content",
      },
      {
        Name:       "default_value",
        ColumnName: "default_value",
        DataType:   "varchar(100)",
        IsNullable: true,
        ColumnType: "content",
      },
      {
        Name:       "data_type",
        ColumnName: "data_type",
        DataType:   "varchar(50)",
        IsNullable: true,
        ColumnType: "label",
      },
    },
  },
  {
    TableName: "user",
    Columns: []api2go.ColumnInfo{
      {
        Name:       "name",
        ColumnName: "name",
        IsIndexed:  true,
        DataType:   "varchar(80)",
        ColumnType: "name",
      },
      {
        Name:       "email",
        ColumnName: "email",
        DataType:   "varchar(80)",
        IsIndexed:  true,
        IsUnique:   true,
        ColumnType: "email",
      },

      {
        Name:           "password",
        ColumnName:     "password",
        DataType:       "varchar(100)",
        ExcludeFromApi: true,
        ColumnType:     "password",
        IsNullable:     true,
      },
      {
        Name:         "confirmed",
        ColumnName:   "confirmed",
        DataType:     "boolean",
        ColumnType:   "truefalse",
        IsNullable:   false,
        DefaultValue: "false",
      },
    },
  },
  {
    TableName: "usergroup",
    Columns: []api2go.ColumnInfo{
      {
        Name:       "name",
        ColumnName: "name",
        IsIndexed:  true,
        DataType:   "varchar(80)",
        ColumnType: "name",
      },
    },
  },
  {
    TableName: "action",
    Columns: []api2go.ColumnInfo{
      {
        Name:       "action_name",
        IsIndexed:  true,
        ColumnName: "action_name",
        DataType:   "varchar(100)",
        ColumnType: "name",
      },
      {
        Name:       "label",
        ColumnName: "label",
        IsIndexed:  true,
        DataType:   "varchar(100)",
        ColumnType: "label",
      },
      {
        Name:       "in_fields",
        ColumnName: "in_fields",
        DataType:   "text",
        ColumnType: "json",
      },
      {
        Name:       "out_fields",
        ColumnName: "out_fields",
        DataType:   "text",
        ColumnType: "json",
      },
    },
  },
  {
    TableName: "smd",
    IsHidden:  true,
    Columns: []api2go.ColumnInfo{
      {
        Name:       "name",
        ColumnName: "name",
        IsIndexed:  true,
        DataType:   "varchar(100)",
        ColumnType: "label",
        IsNullable: false,
      },
      {
        Name:       "label",
        ColumnName: "label",
        DataType:   "varchar(100)",
        ColumnType: "label",
        IsNullable: false,
      },
      {
        Name:       "initial_state",
        ColumnName: "initial_state",
        DataType:   "varchar(100)",
        ColumnType: "label",
        IsNullable: false,
      },
      {
        Name:       "events",
        ColumnName: "events",
        DataType:   "text",
        ColumnType: "json",
        IsNullable: false,
      },
    },
  },
}

type TableInfo struct {
  TableName         string `db:"table_name"`
  TableId           int
  DefaultPermission int64 `db:"default_permission"`
  Columns           []api2go.ColumnInfo
  StateMachines     []fsm_manager.LoopbookFsmDescription
  Relations         []api2go.TableRelation
  IsTopLevel        bool `db:"is_top_level"`
  Permission        int64
  UserId            uint64 `db:"user_id"`
  IsHidden          bool   `db:"is_hidden"`
}
