<script>
  import { onMount } from 'svelte'
  import { getNotificationsContext } from 'svelte-notifications'

  import Input from './Input.svelte'

  const { addNotification } = getNotificationsContext()

  let settings = {}

  onMount(() => {
    fetch(`${BASE_URL}/api/v1/settings`)
      .then((res) => res.json())
      .then((data) => {
        settings = data
      })
      .catch((err) => console.log(err))
  })

  function handleDBChange(event) {
    updateSetting('dbType', event.currentTarget.value)
  }

  function toggleDebugServer() {
    if (settings.debugServer === 'true') {
      settings.debugServer = 'false'
    } else {
      settings.debugServer = 'true'
    }
    updateSetting('debugServer', settings.debugServer)
  }

  function toggleDebugTransactions() {
    if (settings.debugTransactions === 'true') {
      settings.debugTransactions = 'false'
    } else {
      settings.debugTransactions = 'true'
    }
    updateSetting('debugTransactions', settings.debugTransactions)
  }

  function updateSetting(key, value) {
    fetch(`${BASE_URL}/api/v1/settings`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        key,
        value,
      }),
    })
      .then((res) => {
        settings[key] = value

        addNotification({
          text: `Setting updated successfully`,
          position: 'bottom-left',
          type: 'success',
          removeAfter: 1000,
        })
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
  }
</script>

<form class="panel" autocomplete="off">
  <input autocomplete="false" name="hidden" type="text" style="display:none;" />
  <p class="panel-heading">Server settings</p>
  <div class="panel-body pad">
    {#if Object.keys(settings).length === 0}
      <div class="loading">Loading...</div>
    {:else}
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
                  checked={settings.debugServer === 'true'}
                  on:change={toggleDebugServer}
                />
                Server
              </label>
              <label class="check">
                <input
                  type="checkbox"
                  checked={settings.debugTransactions === 'true'}
                  on:change={toggleDebugTransactions}
                />
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
    {/if}
  </div>
</form>

<form class="panel" autocomplete="off">
  <input autocomplete="false" name="hidden" type="text" style="display:none;" />
  <p class="panel-heading">Database settings</p>
  <div class="panel-body pad">
    {#if Object.keys(settings).length === 0}
      <div class="loading">Loading...</div>
    {:else}
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
                  name="dbType"
                  checked={settings.dbType === 'sqlite'}
                  value={'sqlite'}
                  on:change={handleDBChange}
                />
                Local
              </label>
              <label class="radio">
                <input
                  type="radio"
                  name="dbType"
                  checked={settings.dbType === 'postgres'}
                  value={'postgres'}
                  on:change={handleDBChange}
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
          label="Role"
          description=""
          value={settings.dbUsername}
          updateFn={updateSetting.bind(this, 'dbUsername')}
        />

        <Input
          label="Password"
          description=""
          value={settings.dbPassword}
          updateFn={updateSetting.bind(this, 'dbPassword')}
        />
      {/if}
    {/if}
  </div>
</form>

<style>
  .pad {
    padding: 20px;
  }
</style>
