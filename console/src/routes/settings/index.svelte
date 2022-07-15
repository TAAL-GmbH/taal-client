<script>
  import { onMount } from 'svelte'
  import { getNotificationsContext } from 'svelte-notifications'

  import Checkbox from '../../lib/components/checkbox/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Radio from '../../lib/components/radio/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Text from '../../lib/components/text/index.svelte'
  import TextInput from '../../lib/components/textinput/index.svelte'
  import Spinner from '../../lib/components/spinner/index.svelte'

  import { spinCount } from '../../lib/stores'
  import * as api from '../../lib/api'

  const { addNotification } = getNotificationsContext()

  let settings = {}

  onMount(() => {
    getSettings()
  })

  function getSettings() {
    api.getSettings(
      (data) => {
        addNotification({
          text: `Settings retrieved successfully`,
          position: 'bottom-left',
          type: 'success',
          removeAfter: 1000,
        })

        settings = data
      },
      (error) => {
        addNotification({
          text: `Error: ${error}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 5000,
        })
      }
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

        addNotification({
          text: `Setting updated successfully`,
          position: 'bottom-left',
          type: 'success',
          removeAfter: 1000,
        })
      },
      (error) => {
        addNotification({
          text: `Error: ${error}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 2000,
        })
      }
    )
  }

  function onChange(e) {
    const { name, group, type, value, checked } = e.detail
    // console.log(
    //   'name =',
    //   name,
    //   ' group =',
    //   group,
    //   ' type =',
    //   type,
    //   ' checked =',
    //   checked
    // )
    switch (type) {
      case 'search':
      case 'text':
        updateSetting(name, value || '')
        break
      case 'checkbox':
        updateSetting(name, (checked || false).toString())
        break
      case 'radio':
        updateSetting(group, checked ? name : '')
        break
    }
  }

  function onInputMount(e) {
    e.detail.inputRef.focus()
  }
</script>

<PageWithMenu>
  <div class="island">
    <Heading value="Settings" />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading value="Server settings" size={2} />
    </div>
    <Spacer h={24} />
    <!-- type="search" is an undesired workaround for 1Password: https://1password.community/discussion/117501/as-a-web-developer-how-can-i-disable-1password-filling-for-a-specific-field -->
    <TextInput
      type="search"
      name="listenAddress"
      label="Listen address"
      placeholder="localhost:9500"
      value={settings.listenAddress}
      autocomplete="off"
      confirm
      on:change={onChange}
      on:mount={onInputMount}
    />
    <Spacer h={8} />
    <Text
      size={5}
      color="#8F8D94"
      value={'TaalClient will listen on the specified address. The default is "localhost:9500" To bind to all interfaces on your machine, omit the host part and only specify the port, for example ":9500"'}
    />
    <Spacer h={24} />
    <TextInput
      name="taalUrl"
      label="TAAL URL"
      placeholder="https://api.taal.com"
      value={settings.taalUrl}
      confirm
      on:change={onChange}
    />
    <Spacer h={8} />
    <Text
      size={5}
      color="#8F8D94"
      value={'TaalClient communicates with Taal API servers which are hosted at https://api.taal.com.'}
    />
    <Spacer h={24} />
    <TextInput
      name="taalTimeout"
      label="TAAL timeout"
      placeholder="10s"
      value={settings.taalTimeout}
      confirm
      on:change={onChange}
    />
    <Spacer h={8} />
    <Text
      size={5}
      color="#8F8D94"
      value={'The default timeout is 10 seconds. This settings can be provided in milliseconds (ms), seconds (s) or minutes (m).'}
    />
    <Spacer h={24} />
    <Text size={4} value={'Debug'} />
    <Spacer h={8} />
    <Text
      size={5}
      color="#8F8D94"
      value={'Debug modes will add extra output to the console output of this service. Server debug will output all access to server endpoints. Transactions debug will output funding and data transaction raw hex.'}
    />
    <Spacer h={16} />
    <Checkbox
      name="debugServer"
      label="Server"
      checked={settings.debugServer === 'true'}
      labelPlacement="right"
      on:change={onChange}
    />
    <Spacer h={16} />
    <Checkbox
      name="debugTransactions"
      label="Transactions"
      checked={settings.debugTransactions === 'true'}
      labelPlacement="right"
      on:change={onChange}
    />
    <Spacer h={32} />
    <div class="sub-row">
      <Heading value="Transaction data" size={2} />
    </div>
    <Spacer h={24} />
    <Text size={4} value={'Database mode'} />
    <Spacer h={16} />
    <div class="radio-row">
      <Radio
        name="sqlite"
        group="dbType"
        label="Local"
        checked={settings.dbType === 'sqlite'}
        labelPlacement="right"
        on:change={onChange}
      />
      <Radio
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
        name="dbFilename"
        label="File name"
        placeholder="./taal_client.db"
        value={settings.dbFilename}
        confirm
        on:change={onChange}
      />
    {:else}
      <Spacer h={24} />
      <TextInput
        name="dbHost"
        label="Host"
        placeholder="localhost"
        value={settings.dbHost || ''}
        confirm
        on:change={onChange}
      />
      <Spacer h={24} />
      <TextInput
        name="dbPort"
        label="Port"
        placeholder="5432"
        value={settings.dbPort || ''}
        confirm
        on:change={onChange}
      />
      <Spacer h={24} />
      <TextInput
        name="dbName"
        label="Database name"
        placeholder="taal_client"
        value={settings.dbName || ''}
        confirm
        on:change={onChange}
      />
      <Spacer h={24} />
      <TextInput
        name="dbUsername"
        label="Role"
        value={settings.dbUsername || ''}
        confirm
        on:change={onChange}
      />
      <Spacer h={24} />
      <TextInput
        name="dbPassword"
        label="Password"
        value={settings.dbPassword || ''}
        confirm
        on:change={onChange}
      />
    {/if}
  </div>
</PageWithMenu>

{#if $spinCount > 0}
  <Spinner />
{/if}

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
