import os
import hvac

client = hvac.Client(
   url = os.environ['VAULT_URL'],
   token = os.environ['VAULT_TOKEN']
)

print(client.read('secret/vaultdemo'))

