<script>

    import { getNotificationsContext } from 'svelte-notifications';

    const { addNotification } = getNotificationsContext();
    
    let apiKey
    async function doPost () {
        const res = await fetch(`${BASE_URL}/api/v1/register/${apiKey}`, {
            method: 'POST'
	    })

	    const json = await res.json()

        if (json.status != 200) {
            addNotification({text: `Error: ${json.error}`, position: 'bottom-left'});
        };
	}
</script>
<div class="field">
    <div id="text1" class="control">
      <input class="input" type="text" placeholder="API Key" bind:value={apiKey} />
    </div>
  </div>
  <div class="field is-grouped">
    <div class="control">
      <button class="button is-link" on:click="{doPost}">Register new API key</button>
    </div>
</div>
