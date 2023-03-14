# Modules, Packages and Imports

## Repositories, modules and packages

- Library management in Go is based on three concepts.

  - Repository : It is a palce in version control system where the source code for a project is stored.

  - Module : A module is a collection of Go packages stored in a file tree with a ```go.mod``` file at its root.

  - Packages : The purpose of a package is to design and maintain a large number of programs by grouping related features together into single units so that they can be easy to maintain and understand and independent of the other package programs.

## go.mod

- Collection of Go source code becomes a module when there's a ```go.mod``` file in root directory.

- Command ```go mod init <MODULE_PATH>``` makes the current directory the root of module.

- MODULE_PATH is unique name that identifies the module, it is case-sensitive(avoid using upercase letters within it).

### Structure of go.mod file

- go.mod file starts with declaration that consist of keyword module and module's path.

- Next, it specifies minimum compatible version of Go

- Next, the ```require``` section lists the modules that your module depends on and version required for each one.

```text
module learning

go 1.17

require (
 github.com/sirupsen/logrus v1.9.0
 go.mongodb.org/mongo-driver v1.11.1
)
```

## Imports and Exports

- ```import``` allows you to access exported constants, variable, functions, and types in another package.

- A package's exported identifiers cannot be accessed from another package without an import statement.

### How to export an identifier in Go?

- An identifier whose name starts with uppercase letter is exported(public) and is accessible to other package.

- An identifier whose name starts with lowercase letter is exported(public) and is only accessible within that package.

## Creating and Accessing a package

1. Check your GOPATH in environment variables and set it to the directory which contains all Go files.

2. Create a new folder with the name of the package you wish to create.

3. In the folder created in step 2, create your go file that holds the Go package code you wish to create.

4. In the image below a package named objects is created which contains some Go files. First line of all the files in the objects package starts with ```package objects```.
![Alt text](Screenshot%202023-03-14%20115742.png)

5. The following contents are in the file main.go in root directory.
![Alt text](Screenshot%202023-03-14%20115730.png)

## Package comments and godoc

- godoc is Go's format for writing comments that are automatically converted into documentation.

- Conventions to write godoc comment:

```text
1. Place the comment directly before the item being documented with no blanck lines between the comment and declaration of the item.

2. Start the comment with two forward slashes (//) followed by name of the item.

3. Use a blank comment to break the comment into multiple paragraph.

4. Insert preformatted comments by indenting lines.
```

- The image below shows a well-commented file representing package-level comment, struct comment and function comment.
![Alt text](Screenshot%202023-03-14%20122945.png)

