<script>
  import { onMount } from 'svelte'
  import { getNotificationsContext } from 'svelte-notifications'

  // FontAwesome icon...
  import Fa from 'svelte-fa'
  import { faCopy, faUpload } from '@fortawesome/free-solid-svg-icons'

  const { addNotification } = getNotificationsContext()

  const stdMimeType = 'text/plain'

  let taalClientURL = BASE_URL

  let keys
  let apiKey
  let tag
  let mimeType = stdMimeType
  let data
  let filename = ''
  let mode
  let secret
  let loading = true
  let curlCommand = ''

  let files = null
  let file = null
  let fileData = null

  $: if (tag && tag.length > 0) {
    localStorage.setItem('tag', tag)
  } else if (!loading) {
    localStorage.removeItem('tag')
  }

  $: if (mode && mode.length > 0) {
    localStorage.setItem('mode', mode)
  } else if (!loading) {
    localStorage.removeItem('mode')
  }

  $: if (secret && secret.length > 0) {
    localStorage.setItem('secret', secret)
  } else if (!loading) {
    localStorage.removeItem('secret')
  }

  $: showCurl(apiKey, mimeType, tag, mode, secret, data, taalClientURL, file)

  $: submitButtonIsDisabled = data == '' || mimeType == '' || apiKey == ''
  $: inputDataDisabled = files != null

  $: if (files != null) {
    file = files[0]
    mimeType = file.type
    filename = file.name

    if (file.type.startsWith('text/')) {
      const fr = new FileReader()
      fr.onload = function () {
        data = fr.result
      }
      fr.readAsText(file)
    } else {
      data = `< ${file.name} >`
    }

    const fr = new FileReader()
    fr.onload = function () {
      fileData = fr.result
    }
    fr.readAsArrayBuffer(file)
  }

  $: if (files == null) {
    data = ''
    fileData = null
    mimeType = stdMimeType
  }

  onMount(async () => {
    const lastTag = localStorage.getItem('tag')
    tag = lastTag || ''

    const lastKey = localStorage.getItem('apiKey')

    const lastMode = localStorage.getItem('mode')
    const lastSecret = localStorage.getItem('secret')

    const r = await fetch(`${BASE_URL}/api/v1/apikeys`)
    const data = await r.json()
    keys = data.keys

    apiKey = lastKey || keys[0].api_key
    mode = lastMode || 'raw'
    secret = lastSecret || ''

    loading = false
  })

  function copyToClipboard() {
    navigator.clipboard.writeText(curlCommand)
  }

  function showCurl(key, type, tag, mode, secret, data, url, file) {
    let curl = 'curl \\\n  -X POST \\\n'
    if (key) curl += `  -H 'Authorization: Bearer ${key}' \\\n`
    if (type) curl += `  -H 'Content-Type: ${type}' \\\n`
    if (tag) curl += `  -H 'X-Tag: ${tag}' \\\n`
    switch (mode) {
      case 'hash':
        curl += `  -H 'X-Mode: ${mode}' \\\n`
        break

      case 'encrypt':
        curl += `  -H 'X-Mode: ${mode}' \\\n`
        curl += `  -H 'X-Key: ${secret}' \\\n`
        break
    }

    if (file) {
      curl += `  --data-binary @${file.name} \\\n`
    } else if (data) {
      curl += `  -d '${data}' \\\n`
    }
    curl += `${url}/api/v1/transactions`

    curlCommand = curl
  }

  function reset() {
    files = null
    filename = ''
    data = ''
    mimeType = stdMimeType
  }

  function writeData() {
    const url = `${BASE_URL}/api/v1/transactions`

    fetch(url, {
      method: 'POST',
      body: fileData ? fileData : data,
      headers: {
        Authorization: 'Bearer ' + apiKey,
        'Content-Type': mimeType,
        'X-Tag': tag,
        'X-Filename': filename,
        'X-Mode': mode,
        'X-Key': secret,
      },
    })
      .then((res) => {
        localStorage.setItem('apiKey', apiKey)

        if (!res.ok) {
          return res.text().then((text) => {
            throw new Error(text)
          })
        }

        addNotification({
          text: `Transaction submitted successfully`,
          position: 'bottom-left',
          type: 'success',
          removeAfter: 1000,
        })

        return res.json()
      })
      .catch((err) => {
        const errJson = JSON.parse(err.message)
        addNotification({
          text: `Error: ${errJson.error}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 2000,
        })

        console.log(err)
      })

    reset()
  }
</script>

<form class="panel">
  <p class="panel-heading">Submit data (Developer mode)</p>
  <div class="panel-body pad">
    <div class="columns">
      <div class="column">
        <div class="field">
          <div id="input2" class="control">
            <label class="label" for="apiKey">API Key</label>
            <div class="select">
              <select bind:value={apiKey}>
                {#if keys}
                  {#each keys as key}
                    <option value={key.api_key}>
                      {key.api_key}
                    </option>
                  {/each}
                {/if}
              </select>
            </div>
          </div>
        </div>
        <div class="field">
          <div id="input3" class="control">
            <label class="label" for="tag">Tag (optional)</label>
            <input id="tag" class="input" type="text" bind:value={tag} />
          </div>
        </div>
      </div>
      <div class="column">
        <div class="field">
          <div id="input1" class="control">
            <label class="label" for="url">TAAL Client URL</label>
            <input
              id="url"
              class="input"
              type="text"
              bind:value={taalClientURL}
            />
          </div>
        </div>
        <div class="field">
          <div id="input5" class="control">
            <label class="label" for="mimetype">MIME type</label>
            <input
              id="mimetype"
              class="input"
              type="text"
              bind:value={mimeType}
              disabled={inputDataDisabled}
            />
          </div>
        </div>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <div class="field">
          <div id="input4" class="control">
            <label class="label" for="data">Text data</label>
            <textarea
              class="textarea"
              id="data"
              type="text"
              placeholder="Enter text to send to blockchain"
              bind:value={data}
              disabled={inputDataDisabled}
              rows="3"
            />
          </div>
        </div>
      </div>
      <div class="column">
        <div class="field">
          <label class="label" for="file">File</label>
          <div class="file">
            <label class="file-label">
              <input
                class="file-input"
                type="file"
                id="file"
                name="file"
                capture
                accept="image/*, audio/*, application/json, application/pdf, video/*, text/*"
                bind:files
              />
              <span class="file-cta">
                <span class="file-icon">
                  <Fa icon={faUpload} />
                </span>
                <span class="file-label">Choose a file… </span>
              </span>
            </label>
          </div>
        </div>
      </div>

      <div class="column">
        <div class="field">
          <div id="input" class="control">
            <label class="label" for="mode">Mode</label>
            <div class="select">
              <select id="mode" bind:value={mode}>
                <option value="raw">raw</option>
                <option value="hash">hash</option>
                <option value="encrypt">encrypt</option>
              </select>
            </div>
          </div>
          {#if mode === 'encrypt'}
            <div class="field">
              <div id="input3" class="control">
                <label class="label" for="secret">Secret</label>
                <input
                  id="secret"
                  class="input"
                  type="text"
                  placeholder="This is required in encryption mode"
                  bind:value={secret}
                />
              </div>
            </div>
          {/if}
        </div>
      </div>
    </div>

    <div class="columns">
      <div class="column">
        <div class="field is-grouped is-grouped-left">
          <div class="control">
            <button
              class="button is-primary"
              on:click|preventDefault={writeData}
              disabled={submitButtonIsDisabled}>Submit transaction</button
            >
          </div>
          <div class="control">
            <button class="button is-light" on:click|preventDefault={reset}
              >Reset</button
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</form>

<form class="panel">
  <p class="panel-heading">
    cURL command
    <button class="button is-small" on:click|preventDefault={copyToClipboard}>
      <Fa icon={faCopy} />
    </button>
  </p>
  <div class="panel-body">
    <pre class="field" id="curl">{curlCommand}</pre>
  </div>
</form>

<style>
  .pad {
    padding: 20px;
  }

  pre {
    background: none;
  }
</style>
