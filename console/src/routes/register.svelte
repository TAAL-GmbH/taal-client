<script>
  import { getNotificationsContext } from 'svelte-notifications'

  const { addNotification } = getNotificationsContext()

  let apiKey

  function register() {
    fetch(`${BASE_URL}/api/v1/apikeys/${apiKey}`, {
      method: 'POST',
    })
      .then((res) => {
        if (!res.ok) {
          return res.text().then((text) => {
            throw new Error(text)
          })
        } else {
          addNotification({
            text: `API key registered successfully`,
            position: 'bottom-left',
            type: 'success',
          })
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

<div class="field">
  <div id="text1" class="control">
    <input
      class="input"
      type="text"
      placeholder="API Key"
      bind:value={apiKey}
    />
  </div>
</div>
<div class="field is-grouped">
  <div class="control">
    <button class="button is-link" on:click={register}>Register API key</button>
  </div>
</div>
