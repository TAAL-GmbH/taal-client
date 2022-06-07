<script>
  // svelte-ignore unused-export-let
  export let location

  import { onMount } from 'svelte'
  import { getNotificationsContext } from 'svelte-notifications'

  // FontAwesome icon...
  import Fa from 'svelte-fa'
  import { faFile, faUpload } from '@fortawesome/free-solid-svg-icons'

  let settings = {}

  let dropZone

  let files

  $: if (files) {
    const promises = [...files].map((f) => {
      return readFile(f)
    })

    Promise.all(promises).then((values) => {
      arr = Object.assign([], values)
    })
  }

  let arr = []

  $: console.log('ARRAY OF FILES:', arr)

  let over = false

  function handleDragEnter(e) {
    e.preventDefault()
    over = true
  }

  function handleDragLeave(e) {
    e.preventDefault()
    over = false
  }

  async function handleDragDrop(e) {
    e.preventDefault()

    over = false

    const data = e.dataTransfer
    const fileList = await data.files

    const a = []

    for (let i = 0; i < fileList.length; i++) {
      a.push(await readFile(fileList[i]))
    }

    arr = Object.assign([], a)
  }

  function readFile(file) {
    return new Promise((resolve, reject) => {
      let reader = new FileReader()

      reader.onload = () => {
        const data = reader.result
        // let view = new Uint8Array(data)
        resolve({
          file,
          data,
        })
      }

      reader.onerror = () => {
        resolve({
          file,
          error: reader.error,
        })
      }

      reader.readAsArrayBuffer(file)
    })
  }

  let encrypt
  let hashOnly
  let autoSubmit

  let keys
  const { addNotification } = getNotificationsContext()

  let apiKey
  let tag = ''
  let tagString = ''

  $: tagString = tag ? tag : ''
  $: submitButtonIsDisabled = files == [] && apiKey == ''

  onMount(async () => {
    let res = await fetch(`${BASE_URL}/api/v1/settings`)
    settings = await res.json()

    if (localStorage.getItem('tag') !== 'null') {
      tag = localStorage.getItem('tag')
    }

    const lastKey = localStorage.getItem('apiKey')

    res = await fetch(`${BASE_URL}/api/v1/apikeys`)
    const data = await res.json()
    keys = data.keys

    apiKey = lastKey || keys[0].api_key
  })

  function reset() {
    files = null
    arr = []
  }

  function writeData() {
    const url = `${BASE_URL}/api/v1/transactions`

    arr.forEach((a) => {
      const headers = {
        Authorization: 'Bearer ' + apiKey,
        'Content-Type': a.file.type,
      }

      if (tagString) {
        headers['X-Tag'] = tagString
      }

      if (mode !== 'raw') {
        headers['X-Mode'] = mode
      }

      fetch(url, {
        method: 'POST',
        body: a.data,
        headers,
      })
        .then((res) => {
          localStorage.setItem('tag', tag)
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
            removeAfter: 2000,
          })

          return res.json()
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
    })
  }
</script>

<form class="panel">
  <p class="panel-heading">Submit transactions</p>
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
      </div>

      <div class="column">
        <div class="field">
          <div id="input3" class="control">
            <label class="label" for="tag">Tag (optional)</label>
            <input id="tag" class="input" type="text" bind:value={tag} />
          </div>
        </div>
      </div>
    </div>

    <div class="field is-horizontal">
      <div class="field-label is-normal has-text-left">
        <label for="debug" class="label">Options:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p id="debug" class="control">
            <label class="check">
              <input type="checkbox" checked={encrypt} on:change={encrypt} />
              Encrypt data
            </label>
            <label class="check">
              <input type="checkbox" checked={hashOnly} />
              Hash data (do not store original media)
            </label>
            <label class="check">
              <input type="checkbox" checked={autoSubmit} />
              Auto submit
            </label>
          </p>
        </div>
      </div>
    </div>
  </div>

  <section class="columns is-mobile is-centered">
    <div
      class="column is-8 droparea {over ? 'over' : ''}"
      on:dragenter={handleDragEnter}
      on:dragleave={handleDragLeave}
      on:drop={handleDragDrop}
      bind:this={dropZone}
      id="drop_zone"
      ondragover="return false"
    >
      <Fa icon={faFile} size="3x" />
      <p class="top">
        Upload something with the file dialog or by dragging and dropping files
        onto the dashed region
      </p>

      <label class="file-label top">
        <input
          class="file-input"
          type="file"
          id="file"
          name="file"
          capture
          multiple
          accept="image/*, audio/*, application/json, application/pdf, video/*, text/*"
          bind:files
        />
        <span class="file-cta">
          <span class="file-icon">
            <Fa icon={faUpload} />
          </span>
          <span class="file-label"> Choose a file… </span>
        </span>
      </label>
    </div>
  </section>
</form>

{#if arr.length > 0}
  <form class="panel">
    <p class="panel-heading">Queue</p>
    <div class="panel-body pad">
      <div class="columns">
        <div class="column">
          {#each arr as a}
            <p>{a.file.name}</p>
          {/each}
        </div>
      </div>
    </div>

    <div class="field is-grouped is-grouped-left">
      <div class="control">
        <button
          class="button is-primary"
          on:click={writeData}
          disabled={submitButtonIsDisabled}
          >{arr.length === 0
            ? 'Submit'
            : `Submit ${arr.length} transaction(s)`}</button
        >
      </div>
      <div class="control">
        <button class="button is-light" on:click={reset}>Reset</button>
      </div>
    </div>
  </form>
{/if}

<style>
  .pad {
    padding: 20px;
  }

  .top {
    padding-top: 1rem;
  }

  .over {
    border-color: red;
  }
  .droparea {
    margin: 1rem auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    /* width: 384px; */
    /* max-width: 100%; */
    /* height: 160px; */
    border: 2px dashed grey;
    border-radius: 15px;
  }
</style>
