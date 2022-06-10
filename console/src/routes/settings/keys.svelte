<script>
  import { onMount } from 'svelte'
  import Fa from 'svelte-fa'
  import { faPlus } from '@fortawesome/free-solid-svg-icons'

  import Key from './key.svelte'

  import { getNotificationsContext } from 'svelte-notifications'

  const { addNotification } = getNotificationsContext()

  let keys
  let isActive
  let loading = true

  function showModal() {
    isActive = true
  }

  function closeModal() {
    isActive = false
  }

  function getApiKeys() {
    fetch(`${BASE_URL}/api/v1/apikeys`)
      .then((res) => {
        if (!res.ok) {
          return res.text().then((text) => {
            throw new Error(text)
          })
        }

        return res.json()
      })
      .then((data) => {
        keys = data.keys
      })
      .catch((err) => {
        var errMessage = ''
        try {
          const errJson = JSON.parse(err.message)
          errMessage = errJson.error
        } catch (error) {
          errMessage = err
        }

        addNotification({
          text: `Failed to load api keys: ${errMessage}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 2000,
        })

        console.log(errMessage)
      })
  }

  onMount(() => {
    getApiKeys()
  })

  let apiKey = ''

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

          closeModal()
          getApiKeys()

          return res.json()
        }
      })
      .catch((err) => {
        const errJson = JSON.parse(err.message)
        addNotification({
          text: `Error: ${errJson.error}`,
          position: 'bottom-left',
          type: 'danger',
        })

        console.log(err)
      })
  }
</script>

<div class="panel">
  <p class="panel-heading">
    <span class="wrapper">
      <span class="left"> Registered API keys </span>
      <span class="icon" on:click={showModal}>
        <Fa icon={faPlus} />
      </span>
    </span>
  </p>
  <div class="panel-body pad">
    {#if keys}
      <table class="table">
        <tr>
          <th>Created</th>
          <th>API Key</th>
          <th>Address</th>
        </tr>
        {#each keys as key}
          <Key {key} />
        {/each}
      </table>
    {/if}
  </div>
</div>

<div class="modal {isActive ? 'is-active' : ''}">
  <div class="modal-background" />
  <div class="modal-card">
    <header class="modal-card-head">
      <p class="modal-card-title">Register API key</p>
      <button class="delete" aria-label="close" on:click={closeModal} />
    </header>
    <section class="modal-card-body">
      <p class="content is-small">
        API keys can be obtained from <a href="https://console.taal.com"
          >https://console.taal.com</a
        >
        Registering an API key with this TaalClient will allow you to read and write
        transactions to the BSV blockchain.
      </p>
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
          <button class="button is-link" on:click={register}>Register</button>
        </div>
      </div>
    </section>
  </div>
</div>

<style>
  .pad {
    padding: 20px;
  }

  .wrapper {
    display: flex;
    flex-direction: row;
    align-items: stretch;
    width: 100%;
  }

  .wrapper > .left {
    flex: 1;
  }
</style>
