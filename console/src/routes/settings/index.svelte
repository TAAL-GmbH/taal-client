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
  import i18n from '../../lib/i18n'

  $: t = $i18n.t
  const pageKey = 'page.settings'

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
        success(t(`${pageKey}.notifications.get-settings-success`), {
          duration: 1000,
        })
        settings = data
      },
      (error) =>
        failure(t('notifications.failure', { error }), { duration: 5000 })
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
        success(t(`${pageKey}.notifications.update-settings-success`), {
          duration: 1000,
        })
      },
      (error) =>
        failure(t('notifications.failure', { error }), { duration: 2000 })
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
    <Heading testId={'settings-label'} value={t(`${pageKey}.settings-label`)} />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading
        testId={'server-settings-label'}
        value={t(`${pageKey}.server-settings-label`)}
        size={2}
      />
    </div>
    <Spacer h={24} />
    <TextInput
      testId={'listen-address-text-input'}
      name="listenAddr"
      label={t(`${pageKey}.listen-address-label`)}
      placeholder={t(`${pageKey}.listen-address-placeholder`)}
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
      value={t(`${pageKey}.listen-address-hint`)}
    />
    <Spacer h={24} />
    <TextInput
      testId={'taal-url-input-text'}
      name="taalUrl"
      label={t(`${pageKey}.taal-url-label`)}
      placeholder={t(`${pageKey}.taal-url-placeholder`)}
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
      value={t(`${pageKey}.taal-url-hint`)}
    />
    <Spacer h={24} />
    <TextInput
      testId={'taal-timeout-input-text'}
      name="taalTimeout"
      label={t(`${pageKey}.taal-timeout-label`)}
      placeholder={t(`${pageKey}.taal-timeout-placeholder`)}
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
      value={t(`${pageKey}.taal-timeout-hint`)}
    />
    <Spacer h={24} />
    <TextInput
      testId={'filesize-limit-input-text'}
      name="fileSizeLimitBytes"
      label={t(`${pageKey}.filesize-limit-label`)}
      placeholder={t(`${pageKey}.filesize-limit-placeholder`)}
      value={settings.fileSizeLimitBytes}
      confirm
      on:change={onChange}
      on:confirm={onConfirm}
    />
    <Spacer h={8} />
    <Text
      testId={'filesize-limit-hint-text'}
      size={5}
      color="#8F8D94"
      value={t(`${pageKey}.filesize-limit-hint`)}
    />
    <Spacer h={24} />
    <Text testId={'debug-text'} size={4} value={t(`${pageKey}.debug-label`)} />
    <Spacer h={8} />
    <Text
      testId={'debug-hint-text'}
      size={5}
      color="#8F8D94"
      value={t(`${pageKey}.debug-hint`)}
    />
    <Spacer h={16} />
    <Checkbox
      testId={'server-checkbox'}
      name="debugServer"
      label={t(`${pageKey}.server-checkbox-label`)}
      checked={settings.debugServer === 'true'}
      labelPlacement="right"
      on:change={onChange}
    />
    <Spacer h={16} />
    <Checkbox
      testId={'transactions-checkbox'}
      name="debugTransactions"
      label={t(`${pageKey}.transactions-checkbox-label`)}
      checked={settings.debugTransactions === 'true'}
      labelPlacement="right"
      on:change={onChange}
    />
    <Spacer h={32} />
    <div class="sub-row">
      <Heading
        testId={'transaction-data-label'}
        value={t(`${pageKey}.transaction-data-label`)}
        size={2}
      />
    </div>
    <Spacer h={24} />
    <Text
      testId={'database-mode-label'}
      size={4}
      value={t(`${pageKey}.database-mode-label`)}
    />
    <Spacer h={16} />
    <div class="radio-row">
      <Radio
        testId={'local-radio-btn'}
        name="sqlite"
        group="dbType"
        label={t(`${pageKey}.local-radio-label`)}
        checked={settings.dbType === 'sqlite'}
        labelPlacement="right"
        on:change={onChange}
      />
      <Radio
        testId={'remote-radio-btn'}
        name="postgres"
        group="dbType"
        label={t(`${pageKey}.remote-radio-label`)}
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
        label={t(`${pageKey}.filename-label`)}
        placeholder={t(`${pageKey}.filename-placeholder`)}
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
        label={t(`${pageKey}.host-label`)}
        placeholder={t(`${pageKey}.host-placeholder`)}
        value={settings.dbHost || ''}
        confirm
        on:change={onChange}
        on:confirm={onConfirm}
      />
      <Spacer h={24} />
      <TextInput
        testId={'port-text-input'}
        name="dbPort"
        label={t(`${pageKey}.port-label`)}
        placeholder={t(`${pageKey}.port-placeholder`)}
        value={settings.dbPort || ''}
        confirm
        on:change={onChange}
        on:confirm={onConfirm}
      />
      <Spacer h={24} />
      <TextInput
        testId={'database-name-text-input'}
        name="dbName"
        label={t(`${pageKey}.db-name-label`)}
        placeholder={t(`${pageKey}.db-name-placeholder`)}
        value={settings.dbName || ''}
        confirm
        on:change={onChange}
        on:confirm={onConfirm}
      />
      <Spacer h={24} />
      <TextInput
        testId={'role-text-input'}
        name="dbUsername"
        label={t(`${pageKey}.role-label`)}
        value={settings.dbUsername || ''}
        confirm
        on:change={onChange}
        on:confirm={onConfirm}
      />
      <Spacer h={24} />
      <TextInput
        testId={'password-text-input'}
        name="dbPassword"
        label={t(`${pageKey}.password-label`)}
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
