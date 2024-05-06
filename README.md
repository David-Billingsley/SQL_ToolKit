# SQL_ToolKit

This libary is free of use, and is a helpful list of functions I have used over and over in code.

List of functions to help the end user with their projects
To use the tool add toolkit "github.com/David-Billingsley/SQL_ToolKit" to your code base

And in your function put var data DataBaseTools.Data

If a username and password is passed to the functions it will generate the user and password connection string.  If there is no user / password then the system will use the Windows Auth connection string.

## SQL_File_Read
  This function takes the following parameters fpath (filepath), server(SQL Server name), database ( database name), user, password.  It will read any file in the current directory, and attempt to pass the query into sql.  Checks to see if file is .sql suffix.  If the number of files sent is greater then 0 it will send a true message back if zero sends false back.

## Get_Coloumn_Info
  This function takes the following parameters, erver(SQL Server name), database ( database name), user, password, table ( this is the table name ).  This function will read the the coloumn names and datatypes found in the table.  If the table doesnt exisit, the funciton will return an error.  If the table does exist it will return the coloumn names and types in a map.