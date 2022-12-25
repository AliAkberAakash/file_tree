## File Naming convention
To define your own structure, create a `config.json` file in the root of your project directory. The name of the file must be <b>config.json</b> and it needs to be a valid json file.

## Properties
Each entry must folow the following structure:
```
{
  "name":"feature",
  "type":"FD",
  "children":[]
}
```

This unit is called a `node` and will be refered to as `node` in the following section.

### Property: name

The name of first node on the json file must always be feature. It will be replaced by the feature name you provide while running the file_tree command.

Th name property represents what kind of file or folder it is. For example, you might have repositories in all of your features. So a child will have the name property set to "repository". And it will be added to end of the feature. 
For feature "Login" and name "repository" the created folder will be `login_repository`.

Example config file:

```
{
  "name":"feature",
  "type":"FD",
  "children":[
      {
          "name":"repository",
          "type":"FD",
          "children":[
              {
                  "name":"repository",
                  "type":"FL"
              }
          ]
      }
  ]
}
```

Run the command: 
```
file_tree login js
```

generated files:
```
 login/repository/login_repository.js
```

### Property: type

The type denotes weather the node represents a file or a folder. Use `FD` for folders and `FL` for files. Any other value will be ignored and throw exception. 

### Property: children

The cihldren property takes an array of nodes. Each node must be valid and can be of either type file or folder. It represents the nested folders and files. If the type of a node is `FL` the children property can be ommited.
 
