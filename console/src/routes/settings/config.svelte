<script>
  import { onMount } from 'svelte'
  import Fa from 'svelte-fa'
  import { faLock, faLockOpen } from '@fortawesome/free-solid-svg-icons'
  import Input from './Input.svelte'

  import { listen } from 'svelte/internal'

  let settings = {}

  onMount(async () => {
    const res = await fetch(`${BASE_URL}/api/v1/settings`)
    settings = await res.json()

    console.log('SETTINGS', JSON.stringify(settings, null, 2))
  })

  let showPassword = false

  function toggleShowPassword() {
    showPassword = !showPassword
  }

  function toggleDebugServer() {
    settings.debugServer = !settings.debugServer
    saveSetting('debugServer', settings.debugServer)
  }

  function updateSetting(key, value) {
    saveSetting(key, value)
      .then(() => {
        console.log('Saved', key, value)
        settings[key] = value
      })
      .catch((err) => {
        console.log('ERROR', err)
      })
  }

  async function saveSetting(key, value) {
    try {
      await fetch(`${BASE_URL}/api/v1/settings/${key}/${value}`, {
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
    <Input
      label="Listen address"
      description="TaalClient will listen on the specified address. The default is localhost:9500. To bind to all interfaces on your machine, omit the host part and only specify the port, for example &quot;:9500&quot;."
      value={settings.listenAddress}
      placeholder="localhost:9500"
      updateFn={updateSetting.bind(this, 'listenAddress')}
    />

    <Input
      label="Taal URL"
      description="TaalClient communicates with Taal API servers which are hosted at https://api.taal.com."
      value={settings.taalUrl}
      placeholder="https://api.taal.com"
      updateFn={updateSetting.bind(this, 'taalUrl')}
    />

    <Input
      label="Taal timeout"
      description="The default timeout is 10s (10 seconds). This settings can be provided a milliseconds (ms), seconds (s) or minutes (m)."
      value={settings.taalTimeout}
      placeholder="10s"
      updateFn={updateSetting.bind(this, 'taalTimeout')}
    />

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
      <Input
        label="Filename"
        description=""
        value={settings.dbFilename}
        placeholder="./taal_client.db"
        updateFn={updateSetting.bind(this, 'dbFilename')}
      />
    {:else}
      <Input
        label="Host"
        description=""
        value={settings.dbHost}
        placeholder="localhost"
        updateFn={updateSetting.bind(this, 'dbHost')}
      />

      <Input
        label="Port"
        description=""
        value={settings.dbPort}
        placeholder="5432"
        updateFn={updateSetting.bind(this, 'dbPort')}
      />

      <Input
        label="Database name"
        description=""
        value={settings.dbName}
        placeholder="taal_client"
        updateFn={updateSetting.bind(this, 'dbName')}
      />

      <Input
        label="Username"
        description=""
        value={settings.dbUsername}
        placeholder="database username"
        updateFn={updateSetting.bind(this, 'dbUsername')}
      />

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
                on:input={updateSetting.bind(this, 'dbPassword')}
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
