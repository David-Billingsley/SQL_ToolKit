# SQL_ToolKit

This libary is free of use, and is a helpful list of functions I have used over and over in code.

List of functions to help the end user with their projects
To use the tool add toolkit "github.com/David-Billingsley/SQL_ToolKit" to your code base

And in your function put var data DataBaseTools.Data

## SQLFileRead
  This function takes the following parameters fpath (filepath), server(SQL Server name), database ( database name).  It will read any file in the current directory, and attempt to pass the query into sql.  Currently it is set for all file types need to refine to just .sql
