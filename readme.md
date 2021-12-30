# Durin



## Error Code
  + success:0
  + resource
    + 1: user
    + 2: message
  + type
    + 1. not authorized
    + 2: missing parameter
    + 3: parameter error
    + 4: not found
  + serial

## TODO
+ [x] type client manager => lock
+ [x] **redis**/mongodb?
+ [ ] login
  + [x] **http**/websocket?
  + [x] **token**=> cookie?
+ [x] primary key 
  +  self increase id
  +  given after reigster
  +  allow duplicate namme
+ [x] ~~**search by user name ?????**~~
+ [x] message
  ``` json
  {
    from:"sender",
    to:"reveiver",
    content:"content",
    time:"",
  }
  ```
+ [x] ~~auth middleware~~
+ [x] return object
  + resultCode guildline
+ [ ] avatar
+ [ ] message type
  + [ ] plain
  + [ ] img
+ [ ] account setting
  + [ ] change name/password
+ [x] API Doc
  + [ ] login
  + [ ] logout
  + [ ] register
    +  /user => post
  + [ ] search user
    + /user?name="nameToSearch" => get
    + /user?id="idToSearch" => get
  + [ ] modify user info
    + /user => put
  + [ ] delete user
    + /user => delete
  + [ ] messageWS
    + /message?id="userId" => get
  
