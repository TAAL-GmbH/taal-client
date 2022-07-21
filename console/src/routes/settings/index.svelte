<script>
  import { onMount } from 'svelte'

  import Checkbox from '../../lib/components/checkbox/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Radio from '../../lib/components/radio/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Text from '../../lib/components/text/index.svelte'
  import TextInput from '../../lib/components/textinput/index.svelte'

  import * as api from '../../lib/api'
  import { success, failure } from '../../lib/utils/notifications'

  // injected by svelte-navigator
  export let location = null
  export let navigate = null

  let settings = {}

  onMount(() => {
    getSettings()
  })

  function getSettings() {
    api.getSettings(
      (data) => {
        success('Settings retrieved successfully', { duration: 1000 })
        settings = data
      },
      (error) => failure(`Error: ${error}`, { duration: 5000 })
    )
  }

  function updateSetting(key, value) {
    api.updateSetting(
      {
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          key,
          value,
        }),
      },
      (data) => {
        settings[key] = value
        success('Setting updated successfully', { duration: 1000 })
      },
      (error) => failure(`Error: ${error}`, { duration: 2000 })
    )
  }

  function onChange(e) {
    const { name, group, type, value, checked } = e.detail
    switch (type) {
      case 'text':
        // updateSetting(name, value || '')
        break
      case 'checkbox':
        updateSetting(name, (checked || false).toString())
        break
      case 'radio':
        updateSetting(group, checked ? name : '')
        break
    }
  }
  function onConfirm(e) {
    const { name, type, value } = e.detail
    switch (type) {
      case 'text':
        // we avoid listenAddress textinput name to prevent 1password icon from showing up
        const useName = name === 'listenAddr' ? 'listenAddress' : name
        updateSetting(useName, value || '')
        break
    }
  }

  function onInputMount(e) {
    e.detail.inputRef.focus()
  }
</script>

<PageWithMenu>
  <div class="island">
    <Heading testId={'settings-label'} value="Settings" />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading
        testId={'server-settings-label'}
        value="Server settings"
        size={2}
      />
    </div>
    <Spacer h={24} />
    <TextInput
      testId={'listen-address-text-input'}
      name="listenAddr"
      label="Listen address"
      placeholder="localhost:9500"
      value={settings.listenAddress}
      autocomplete="off"
      confirm
      on:change={onChange}
      on:confirm={onConfirm}
      on:mount={onInputMount}
    />
    <Spacer h={8} />
    <Text
      testId={'listen-address-hint-text'}
      size={5}
      color="#8F8D94"
      value={'TaalClient will listen on the specified address. The default is "localhost:9500" To bind to all interfaces on your machine, omit the host part and only specify the port, for example ":9500"'}
    />
    <Spacer h={24} />
    <TextInput
      testId={'taal-url-input-text'}
      name="taalUrl"
      label="TAAL URL"
      placeholder="https://api.taal.com"
      value={settings.taalUrl}
      confirm
      on:change={onChange}
      on:confirm={onConfirm}
    />
    <Spacer h={8} />
    <Text
      testId={'taal-url-hint-text'}
      size={5}
      color="#8F8D94"
      value={'TaalClient communicates with Taal API servers which are hosted at https://api.taal.com.'}
    />
    <Spacer h={24} />
    <TextInput
      testId={'taal-timeout-input-text'}
      name="taalTimeout"
      label="TAAL timeout"
      placeholder="10s"
      value={settings.taalTimeout}
      confirm
      on:change={onChange}
      on:confirm={onConfirm}
    />
    <Spacer h={8} />
    <Text
      testId={'taal-timeout-hint-text'}
      size={5}
      color="#8F8D94"
      value={'The default timeout is 10 seconds. This settings can be provided in milliseconds (ms), seconds (s) or minutes (m).'}
    />
    <Spacer h={24} />
    <Text testId={'debug-text'} size={4} value={'Debug'} />
    <Spacer h={8} />
    <Text
      testId={'debug-hint-text'}
      size={5}
      color="#8F8D94"
      value={'Debug modes will add extra output to the console output of this service. Server debug will output all access to server endpoints. Transactions debug will output funding and data transaction raw hex.'}
    />
    <Spacer h={16} />
    <Checkbox
      testId={'server-checkbox'}
      name="debugServer"
      label="Server"
      checked={settings.debugServer === 'true'}
      labelPlacement="right"
      on:change={onChange}
    />
    <Spacer h={16} />
    <Checkbox
      testId={'transactions-checkbox'}
      name="debugTransactions"
      label="Transactions"
      checked={settings.debugTransactions === 'true'}
      labelPlacement="right"
      on:change={onChange}
    />
    <Spacer h={32} />
    <div class="sub-row">
      <Heading
        testId={'transaction-data-label'}
        value="Transaction data"
        size={2}
      />
    </div>
    <Spacer h={24} />
    <Text testId={'database-mode-label'} size={4} value={'Database mode'} />
    <Spacer h={16} />
    <div class="radio-row">
      <Radio
        testId={'local-radio-btn'}
        name="sqlite"
        group="dbType"
        label="Local"
        checked={settings.dbType === 'sqlite'}
        labelPlacement="right"
        on:change={onChange}
      />
      <Radio
        testId={'remote-radio-btn'}
        name="postgres"
        group="dbType"
        label="Remote"
        checked={settings.dbType === 'postgres'}
        labelPlacement="right"
        on:change={onChange}
      />
    </div>
    {#if settings.dbType === 'sqlite'}
      <Spacer h={24} />
      <TextInput
        testId={'filename-text-input'}
        name="dbFilename"
        label="File name"
        placeholder="./taal_client.db"
        value={settings.dbFilename}
        confirm
        on:change={onChange}
        on:confirm={onConfirm}
      />
    {:else}
      <Spacer h={24} />
      <TextInput
        testId={'host-text-input'}
        name="dbHost"
        label="Host"
        placeholder="localhost"
        value={settings.dbHost || ''}
        confirm
        on:change={onChange}
        on:confirm={onConfirm}
      />
      <Spacer h={24} />
      <TextInput
        testId={'port-text-input'}
        name="dbPort"
        label="Port"
        placeholder="5432"
        value={settings.dbPort || ''}
        confirm
        on:change={onChange}
        on:confirm={onConfirm}
      />
      <Spacer h={24} />
      <TextInput
        testId={'database-name-text-input'}
        name="dbName"
        label="Database name"
        placeholder="taal_client"
        value={settings.dbName || ''}
        confirm
        on:change={onChange}
        on:confirm={onConfirm}
      />
      <Spacer h={24} />
      <TextInput
        testId={'role-text-input'}
        name="dbUsername"
        label="Role"
        value={settings.dbUsername || ''}
        confirm
        on:change={onChange}
        on:confirm={onConfirm}
      />
      <Spacer h={24} />
      <TextInput
        testId={'password-text-input'}
        name="dbPassword"
        label="Password"
        value={settings.dbPassword || ''}
        confirm
        on:change={onChange}
        on:confirm={onConfirm}
      />
    {/if}
  </div>
</PageWithMenu>

<style>
  .island {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 920px;
  }

  .sub-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .radio-row {
    display: flex;
    align-items: center;
    gap: 24px;
  }
</style>
