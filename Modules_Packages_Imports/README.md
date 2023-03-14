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

- There are two optional sections ```replace``` and ```exclude```.

![Alt text](Screenshot%202023-03-14%20125815.png)

## Imports and Exports

- ```import``` allows you to access exported constants, variable, functions, and types in another package.

- A package's exported identifiers cannot be accessed from another package without an import statement.

### How to export an identifier in Go?

- An identifier whose name starts with uppercase letter is exported(public) and is accessible to other package.

- An identifier whose name starts with lowercase letter is exported(public) and is only accessible within that package.

## Creating and Accessing a package

1. Check your GOPATH in environment variables and set it to the directory which contains all Go files.

2. Create a new folder with the name of the package you wish to create.
**Note:** Package names should be lowercase. Don't use snake_case or camelCase in package names.

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

## Circular Dependencies

- Two goals of Go are fast compiler and easy to undersatnd code.

- To support this, Go does not allow to have ```circular dependency``` between packages.

- This means if package A impors package B then package B cannot import package A.

- If we try to build a code which has circular dependency you'll get error.

- If two packages depend on each other, then you can merge it into a single package, if packages needs to be seperated, move the items causing circular dependency to one of the two packages or a new package.

## Renaming and Reorganizing

- If it is needed to export identifier or move them to another package within the module, instead of removing the original identifiers provide an alternate name instead.

- For functions or methods, declare a function or method that calls the original. For a constant, simply declare a new constant with same type and value, but a different name.

- To do this alias is used. Let's say we have a type Greet

```text
type Greet struct{
  x int 
  S string
}
func (g Greet) Hello(){
    return "Hello"
}
```

- If we want to access Greet by name Wish,

```text
type Wish = Greet
```

- To create an alias, we use ```type``` keyword, name of alias of original type.

- The alias will have same feild and method as original type.

**Note**: Alias is just another name for a type. If you want to add new mathods or fields in struct, it has to be done to the original type.

## Working with Modules

### Importing Third-Party Code

![Alt text](Screenshot%202023-03-14%20152814.png)

- Import github.com/sirupsen/logrus specify third party imports, they include location of package in repository.

- Once imported, we access the exported items in these packages.

![Alt text](Screenshot%202023-03-14%20152903.png)

- The require section of the go.mod file lists the module that you have imported into your module. The module name is followed by version.

- Whenever we run a go command that requires dependencies(such as go run, go build, go test, go list), any imports that aren't already in go.mod are downloaded and go.mod is updated to include the module path that contains the package and the version of the module.

### Working with Versions

- We will use ```github.com/learning-go-book/simpletax``` for understanding versions.

- The main.go file contains two third party imports:

```text
github.com/shopspring/decimal
```

- When we run ```go build```, the go.mod file will be updated

![Alt text](Screenshot%202023-03-14%20163246.png)

- Go by default will pick the latest version of a dependency when it is added to the project.

- Versioning is useful as you can specify an earlier version of a module.

- To see the versions of the module which are available use

```text
go list -m -versions github.com/learning-go-book/simpletax
```

![Alt text](Screenshot%202023-03-14%20163202.png)

- Here we have 6 versions, v1.0.0 v1.0.1 v1.1.0 v1.2.0 v1.3.0 v1.3.1. If we want to use version v1.0.0, use command

```text
go get github.com/shopspring/decimal@v1.0.0
```

![Alt text](Screenshot%202023-03-14%20163559.png)

- After execution of this command the mod file will be updated as shown in image above, the version is updated from v1.2.0 to v1.0.0 

### Minimum Version Selection

- If our project depends on two or more modules that all depends on some other module but these modules declare that they depend on different versions of that other module.

- Go resolves this by principle of **minimum version selection**.

- This means that lowest version of dependency that is declared will work in go.mod in all dependencies.

- Example: Assume our module depends on modules A, B, C which depends on module D.

- go.mod for A declares it depends on v1.1.0 of D, go.mod for B declares it depends on v1.2.0 of D, go.mod for C declares it depends on v1.2.1 of D.

- Go will import module D only once with version of 1.2.1, as it is recent specified.

### Updating to Compatible Versions

- If we have imported a module ```github.com/learning-go-book/simpletax v1.0.0``` version v1.1.0 which had a bug and now is fixed and released as v1.1.1.

- A function has also been added to the old version, it would get version v1.2.0.

- Finally, there is another version which fixed a bug found in v1.2.0, it would get version number v1.2.1

- To upgrade to bug patch release for v1.0.0, use the following command, since we are in v1.0.0 the command will take us to v1.0.1 which is fixed bug of for v1.0.0.

```text
go get -u=patch github.com/learning-go-book/simpletax
```

![Alt text](Screenshot%202023-03-14%20174113.png)

- Use following command to get the most recent version of module ```github.com/learning-go-book/simpletax```, as shown in image below after executing this command we are upgraded to v1.3.1 of module.

```text
go get -u github.com/learning-go-book/simpletax
```

![Alt text](Screenshot%202023-03-14%20174528.png)

### Vendoring

- Vendoring is the act of making your own copy of the 3rd party packages your project is using. Those copies are traditionally placed inside each project and then saved in the project repository.

- Vendoring is done to ensure that a module always builds with identical dependencies

- For vendoring use command ```go mod vendor```.

- This will create a directory called vendor in your module.

- If new dependencies are addend or updated with ```go get```, you need to run ```go mod vendor``` again to update the vendor directory.

- Vendoring will increase the size of project drastically.

### pkg.go.dev

- It is a site which automatically indexes open source Go projects.
![Alt text](Screenshot%202023-03-14%20180728.png)

- It gathers together the documentation of modules.

### Module Proxy Servers

- Every Go module is stored in a source code repository, like GitHub/GitLab. But Go doesn't fetch code directly from source code repositories.

- It sends request to proxy server run by Google. This server keeps copies of every version of all Go modules.