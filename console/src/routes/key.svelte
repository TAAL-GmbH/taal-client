<script>
  export let key
</script>

<tr>
  <td>{key.createdAt.replace('T', ' ').slice(0, 16)}</td>
  <td>{key.api_key}</td>
  <td title={key.publicKey}>{key.address}</td>
  <td><button class="button is-primary">Revoke</button></td>
</tr>
