<script>
  export let key

</script>

<tr>
  <td class="is-vcentered">{key.createdAt.replace('T', ' ').slice(0, 16)}</td>
  <td class="is-vcentered">{key.api_key}</td>
  <td class="is-vcentered" title={key.publicKey}>{key.address}</td>
</tr>
