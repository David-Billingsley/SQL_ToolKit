# SQL_ToolKit

This libary is free of use, and is a helpful list of functions I have used over and over in code.

List of functions to help the end user with their projects
To use the tool add toolkit "github.com/David-Billingsley/SQL_ToolKit" to your code base

And in your function put var data DataBaseTools.Data

## SQL_File_Read
  This function takes the following parameters fpath (filepath), server(SQL Server name), database ( database name).  It will read any file in the current directory, and attempt to pass the query into sql.  Checks to see if file is .sql suffix.  If the number of files sent is greater then 0 it will send a true message back if zero sends false back.
