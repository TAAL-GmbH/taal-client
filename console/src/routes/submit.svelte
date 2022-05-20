<script>
  import { onMount } from 'svelte'
  import { getNotificationsContext } from 'svelte-notifications'

  const { addNotification } = getNotificationsContext()

  let apiKey
  let taalClientURL = 'http://localhost:9500/api/v1/write'
  let tag
  let mimeType = 'text/plain'
  let data

  let file = null
  let fileData = null

  onMount(async () => {
    tag = localStorage.getItem('tag')
    apiKey = localStorage.getItem('apiKey')
  })

  function writeData() {
    fetch(taalClientURL, {
      method: 'POST',
      body: fileData ? fileData : data,
      headers: {
        Authorization: 'Bearer ' + apiKey,
        'Content-Type': mimeType,
        'X-Tag': tag,
      },
    })
      .then(function (response) {
        localStorage.setItem('tag', tag)
        localStorage.setItem('apiKey', apiKey)

        const json = response.json()

        if (json.status != 200) {
          addNotification({
            text: `Error: ${json.error}`,
            position: 'bottom-left',
            type: 'warning',
          })
        }
      })
      .catch(function (err) {
        console.log(err)
      })
  }
</script>

<div class="field">
  <div id="input1" class="control">
    <label for="url">TAAL Client URL</label>
    <input id="url" class="input" type="text" bind:value={taalClientURL} />
  </div>
</div>
<div class="field">
  <div id="input2" class="control">
    <label for="apiKey">API Key</label>
    <input id="apiKey" class="input" type="text" bind:value={apiKey} />
  </div>
</div>
<div class="field">
  <div id="input3" class="control">
    <label for="tag">Tag (optional)</label>
    <input id="tag" class="input" type="text" bind:value={tag} />
  </div>
</div>
<div class="field">
  <div id="input4" class="control">
    <label for="data">Data</label>
    <textarea
      id="data"
      class="input"
      type="text"
      placeholder="Enter text to send to blockchain"
      bind:value={data}
    />
  </div>
</div>
<div class="field">
  <div id="input5" class="control">
    <label for="mimetype">MIME type</label>
    <input id="mimetype" class="input" type="text" bind:value={mimeType} />
  </div>
</div>
<div class="field is-grouped">
  <div class="control">
    <button class="button is-link" on:click={writeData}
      >Submit transaction</button
    >
  </div>
</div>
