<script>
  import { onMount } from 'svelte'
  import { getNotificationsContext } from 'svelte-notifications'

  const { addNotification } = getNotificationsContext()

  let apiKey
  let taalClientURL = 'http://localhost:9500'
  let tag
  let tagString = ''
  let mimeType = 'text/plain'
  let data

  let inputDataDisabled = false
  let inputMimeTypeDisabled = false

  let files
  let file = null
  let fileData = null

  $: tagString = tag ? tag : ''

  $: if (files != null) {
    file = files[0]
    mimeType = file.type

    if (file.type.startsWith('text/')) {
      const fr = new FileReader()
      fr.onload = function () {
        data = fr.result
      }
      fr.readAsText(file)
    } else {
      data = `< ${file.name} >`
    }

    inputDataDisabled = true
    inputMimeTypeDisabled = true

    const fr = new FileReader()
    fr.onload = function () {
      fileData = fr.result
    }
    fr.readAsArrayBuffer(file)
  }

  onMount(async () => {
    tag = localStorage.getItem('tag')
    apiKey = localStorage.getItem('apiKey')
  })

  function writeData() {
    fetch(`${taalClientURL}/api/v1/transactions`, {
      method: 'POST',
      body: fileData ? fileData : data,
      headers: {
        Authorization: 'Bearer ' + apiKey,
        'Content-Type': mimeType,
        'X-Tag': tagString,
      },
    })
      .then((res) => {
        localStorage.setItem('tag', tag)
        localStorage.setItem('apiKey', apiKey)

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
  <div id="input5" class="control">
    <label for="mimetype">MIME type</label>
    <input
      id="mimetype"
      class="input"
      type="text"
      bind:value={mimeType}
      disabled={inputMimeTypeDisabled}
    />
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
      disabled={inputDataDisabled}
    />
  </div>
</div>
<div class="field">
  <label for="file">File</label>
  <input
    type="file"
    id="file"
    name="file"
    capture
    accept="image/*, audio/*, application/json, application/pdf, video/*, text/*"
    bind:files
  />
</div>
<div class="field is-grouped">
  <div class="control">
    <button class="button is-link" on:click={writeData}
      >Submit transaction</button
    >
  </div>
</div>
