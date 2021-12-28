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
  
