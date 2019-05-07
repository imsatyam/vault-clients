var options = {
    apiVersion: 'v1',
    endpoint: process.env.VAULT_ADDR, 
    token: process.env.VAULT_TOKEN
};
  
var vault = require("node-vault")(options);

/** For us vault is already initialized. We do not need this.
 
vault.init({ secret_shares: 1, secret_threshold: 1 })
  .then( (result) => {
    var keys = result.keys;
    vault.token = result.root_token;
    return vault.unseal({ secret_shares: 1, key: keys[0] })
})
.catch(console.error);
 */

console.log("getting secret...");
vault.read('secret/vaultdemo')
.then((result) => {
    console.log('secret retrieved...');
    console.log(result);
})
.catch(console.error);


/** All  operations
vault.write('secret/vaultdemonode', { username: "nodeuser", password: "nodepass" })
.then( () => vault.read('secret/vaultdemonode'))
.then( () => print())
.then( () => vault.delete('secret/vaultdemonode'))
.catch(console.error);
 */