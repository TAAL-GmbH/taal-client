<script>
  import { onMount } from 'svelte'
  import Fa from 'svelte-fa'
  import { faLock, faLockOpen } from '@fortawesome/free-solid-svg-icons'

  import Register from './register.svelte'

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
</script>

<form class="panel">
  <p class="panel-heading">Server settings</p>
  <div class="panel-body pad">
    <div class="field is-horizontal">
      <div class="field-label is-normal has-text-left">
        <label class="label">Listen address:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
              class="input"
              type="text"
              placeholder="localhost:9500"
              value={settings.listenAddress}
            />
          </p>
        </div>
      </div>
    </div>

    <div class="field is-horizontal">
      <div class="field-label is-normal has-text-left">
        <label class="label">Taal URL:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
              class="input"
              type="text"
              placeholder="https://api.taal.com"
              value={settings.taalUrl}
            />
          </p>
        </div>
      </div>
    </div>

    <div class="field is-horizontal">
      <div class="field-label is-normal has-text-left">
        <label class="label">Taal timeout:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
              class="input"
              type="text"
              placeholder="10s"
              value={settings.taalTimeout}
            />
          </p>
        </div>
      </div>
    </div>
  </div>
</form>

<p class="panel-heading">Database settings</p>
<div class="panel-body pad">
  <div class="field is-horizontal">
    <div class="field-label is-normal has-text-left">
      <label class="label">Database mode:</label>
    </div>
    <div class="field-body">
      <div class="field">
        <p class="control">
          <label class="radio">
            <input type="radio" bind:group={settings.dbType} value={'sqlite'} />
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
        <label class="label">Filename:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
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
        <label class="label">Host:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
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
        <label class="label">Port:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
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
        <label class="label">Database name:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
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
        <label class="label">Username:</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control">
            <input
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
        <label class="label"
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

<style>
  .pad {
    padding: 20px;
  }
</style>
