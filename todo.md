+ [x] type client manager => lock
+ [x] **redis**/mongodb?
+ [ ] login
  + [x] **http**/websocket?
  + [ ] token=> cookie?
+ [x] primary key 
  +  self increase id
  +  given after reigster
  +  allow duplicate namme
+ [ ] **search by user name ?????**
+ [ ] message
  ``` json
  {
    from:"sender",
    to:"reveiver",
    content:"content",
    time:"",
  }
  ```
+ [ ] auth middleware
+ [ ] return object
  + resultCode guildline
+ [ ] API Doc
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
  
