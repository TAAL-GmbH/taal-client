<script>
  import { getNotificationsContext } from 'svelte-notifications'

  const { addNotification } = getNotificationsContext()

  export let key

  function revoke() {
    fetch(`${BASE_URL}/api/v1/apikeys/${key}`, {
      method: 'DELETE',
    })
      .then((res) => {
        if (!res.ok) {
          return res.text().then((text) => {
            throw new Error(text)
          })
        } else {
          return res.json()
        }
      })
      .catch((err) => {
        const errJson = JSON.parse(err.message)
        addNotification({
          text: `Error: ${errJson.error}`,
          position: 'bottom-left',
          type: 'warning',
        })

        console.log(err)
      })
  }
</script>

<tr>
  <td class="is-vcentered">{key.createdAt.replace('T', ' ').slice(0, 16)}</td>
  <td class="is-vcentered">{key.api_key}</td>
  <td class="is-vcentered" title={key.publicKey}>{key.address}</td>
  <td><button class="button is-primary" on:click={revoke}>Revoke</button></td>
</tr>
