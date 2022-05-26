<script>
  import { onMount } from 'svelte'
  import Fa from 'svelte-fa'
  import { faLock, faLockOpen } from '@fortawesome/free-solid-svg-icons'

  import { listen } from 'svelte/internal'

  let settings = {}
  let showPassword = false

  function toggleShowPassword() {
    showPassword = !showPassword
  }

  onMount(async () => {
    const res = await fetch(`${BASE_URL}/api/v1/settings`)
    settings = await res.json()
  })

  function toggleDebugServer() {
    settings.debugServer = !settings.debugServer
    saveSetting('debugServer')
  }

  async function saveSetting(key, val) {
    try {
      await fetch(`${BASE_URL}/api/v1/settings/${key}/${settings[key]}`, {
        method: 'PUT',
      })

      addNotification({
        text: `Setting updated successfully`,
        position: 'bottom-left',
        type: 'success',
      })
    } catch (err) {
      const errJson = JSON.parse(err.message)
      addNotification({
        text: `Error: ${errJson.error}`,
        position: 'bottom-left',
        type: 'warning',
      })

      console.log(err)
    }
  }
</script>

<form class="panel">
  <p class="panel-heading">Server settings</p>
  <div class="panel-body pad">
    <div class="field is-horizontal">
      <div class="field-label is-normal has-text-left">
        <label for="listenAddress" class="label">Listen address:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
              id="listenAddress"
              class="input"
              type="text"
              placeholder="localhost:9500"
              value={settings.listenAddress}
            />
          </p>
        </div>
      </div>
    </div>
    <p class="content is-small">
      TaalClient will listen on the specified address. The default is
      "localhost:9500" To bind to all interfaces on your machine, omit the host
      part and only specify the port, for example ":9500".
    </p>

    <div class="field is-horizontal">
      <div class="field-label is-normal has-text-left">
        <label for="taalUrl" class="label">Taal URL:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
              id="taalUrl"
              class="input"
              type="text"
              placeholder="https://api.taal.com"
              value={settings.taalUrl}
            />
          </p>
        </div>
      </div>
    </div>
    <p class="content is-small">
      TaalClient communicates with Taal API servers which are hosted at
      https://api.taal.com.
    </p>

    <div class="field is-horizontal">
      <div class="field-label is-normal has-text-left">
        <label for="timeout" class="label">Taal timeout:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
              id="timeout"
              class="input"
              type="text"
              placeholder="10s"
              value={settings.taalTimeout}
            />
          </p>
        </div>
      </div>
    </div>
    <p class="content is-small">
      The default timeout is 10s (10 seconds). This settings can be provided a
      milliseconds (ms), seconds (s) or minutes (m).
    </p>

    <div class="field is-horizontal">
      <div class="field-label is-normal has-text-left">
        <label for="debug" class="label">Debug:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p id="debug" class="control">
            <label class="check">
              <input
                type="checkbox"
                checked={settings.debugServer}
                on:change={toggleDebugServer}
              />
              Server
            </label>
            <label class="check">
              <input type="checkbox" checked={settings.debugTransactions} />
              Transactions
            </label>
          </p>
        </div>
      </div>
    </div>
    <p class="content is-small">
      Debug modes will add extra output to the console output of this service.
      Server debug will output all access to server endpoints. Transactions
      debug will output funding and data transaction raw hex.
    </p>
  </div>
</form>

<form class="panel">
  <p class="panel-heading">Database settings</p>
  <div class="panel-body pad">
    <div class="field is-horizontal">
      <div class="field-label is-normal has-text-left">
        <label for="mode" class="label">Database mode:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p id="mode" class="control">
            <label class="radio">
              <input
                type="radio"
                bind:group={settings.dbType}
                value={'sqlite'}
              />
              Local
            </label>
            <label class="radio">
              <input
                type="radio"
                bind:group={settings.dbType}
                value={'postges'}
              />
              Remote
            </label>
          </p>
        </div>
      </div>
    </div>

    {#if settings.dbType === 'sqlite'}
      <div class="field is-horizontal">
        <div class="field-label is-normal has-text-left">
          <label for="filename" class="label">Filename:</label>
        </div>
        <div class="field-body">
          <div class="field">
            <p class="control">
              <input
                id="filename"
                class="input"
                type="text"
                placeholder="./taal_client.db"
                value={settings.dbFilename}
              />
            </p>
          </div>
        </div>
      </div>
    {:else}
      <div class="field is-horizontal">
        <div class="field-label is-normal has-text-left">
          <label for="host" class="label">Host:</label>
        </div>
        <div class="field-body">
          <div class="field">
            <p class="control">
              <input
                id="host"
                class="input"
                type="text"
                placeholder="localhost"
                value={settings.dbHost}
              />
            </p>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label is-normal has-text-left">
          <label for="port" class="label">Port:</label>
        </div>
        <div class="field-body">
          <div class="field">
            <p class="control">
              <input
                id="port"
                class="input"
                type="text"
                placeholder="5432"
                value={settings.dbPort}
              />
            </p>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label is-normal has-text-left">
          <label for="name" class="label">Database name:</label>
        </div>
        <div class="field-body">
          <div class="field">
            <p class="control">
              <input
                id="name"
                class="input"
                type="text"
                placeholder="taal_client"
                value={settings.dbName}
              />
            </p>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label is-normal has-text-left">
          <label for="username" class="label">Username:</label>
        </div>
        <div class="field-body">
          <div class="field">
            <p class="control">
              <input
                id="username"
                class="input"
                type="text"
                placeholder="database username"
                value={settings.dbUsername}
              />
            </p>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label is-normal has-text-left">
          <label for="password" class="label"
            >Password:
            <span class="icon is-small is-right" on:click={toggleShowPassword}>
              <Fa icon={showPassword ? faLockOpen : faLock} />
            </span>
          </label>
        </div>
        <div class="field-body">
          <div class="field">
            <p class="control">
              <input
                id="password"
                class="input"
                type={showPassword ? 'text' : 'password'}
                placeholder="database password"
                value={settings.dbPassword}
              />
            </p>
          </div>
        </div>
      </div>
    {/if}
  </div>
</form>

<style>
  .pad {
    padding: 20px;
  }
</style>
