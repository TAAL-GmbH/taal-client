<script>
  import { onMount } from 'svelte'
  import { useNavigate } from 'svelte-navigator'

  import { link } from '../../lib/utils/format'
  import Button from '../../lib/components/button/index.svelte'
  import Dropdown from '../../lib/components/dropdown/index.svelte'
  import FileUpload from '../../lib/components/file-upload/index.svelte'
  import FileTransfer from '../../lib/components/file-transfer/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Switch from '../../lib/components/switch/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import TextInput from '../../lib/components/textinput/index.svelte'
  import TextArea from '../../lib/components/textarea/index.svelte'
  import Text from '../../lib/components/text/index.svelte'

  import { spinCount } from '../../lib/stores'
  import * as api from '../../lib/api'

  import { getFileKey } from '../../lib/utils/files'
  import { success, failure } from '../../lib/utils/notifications'
  import {
    copyCurl,
    getApiKeys,
    getCorrectURL,
    getCurlCommand,
    getSettings,
    getModeItems,
    getStoreValue,
    setStoreValue,
    updateStore,
  } from './utils'
  import i18n from '../../lib/i18n'
  import { dataSize } from '../../lib/utils'

  $: t = $i18n.t
  const pageKey = 'page.send-data'

  let modeItems = []
  $: modeItems = getModeItems(t) || []

  const navigate = useNavigate()

  const stdMimeType = 'text/plain'
  let dataTransmissionInProgress = false
  let mounting = true
  // show spinner until mounted
  $spinCount++

  let devMode = getStoreValue('devmode') === 'true'
  let taalClientURL = ''
  let fileSizeLimitBytes = -1
  let mimeType = stdMimeType
  let tag = ''
  let mode = 'raw'
  let secret = ''
  let textData = ''
  let curlCommand = ''

  $: updateStore('tag', tag, mounting)
  $: updateStore('mode', mode, mounting)
  $: updateStore('secret', secret, mounting)
  $: updateStore('devmode', devMode, mounting)

  // api keys
  let keys = []
  let apiKeyItems = []
  let apiKey

  $: {
    apiKeyItems = keys.map((key) => ({
      label: key.api_key,
      value: key.api_key,
    }))
  }

  // files
  let files = []
  let fileStatusData = {}
  let fileDataMap = {}
  let hasFileOverLimit = false

  $: compactFileUpload = devMode || files.length > 0
  $: inputDataDisabled = files.length > 0

  let submitButtonText = ''
  let successFilesCount = 0
  let inProgress = false
  $: {
    successFilesCount = 0
    inProgress = false

    files.forEach((file) => {
      const key = getFileKey(file)

      if (fileStatusData[key]) {
        if (fileStatusData[key].state === 'success') {
          successFilesCount++
        }
        if (fileStatusData[key].progress < 1) {
          inProgress = true
        }
      }
    })

    submitButtonText =
      successFilesCount > 1
        ? t(`${pageKey}.submit-many-label`, { count: successFilesCount })
        : t(`${pageKey}.submit-label`)
  }

  $: submitDisabledCommon =
    !apiKey || dataTransmissionInProgress || hasFileOverLimit || inProgress

  $: submitButtonIsDisabled = devMode
    ? !textData || !mimeType || submitDisabledCommon
    : files.length === 0 || submitDisabledCommon

  $: {
    if (devMode) {
      if (files.length > 1) {
        files = [files[0]]
      }

      if (files.length > 0) {
        setFieldsFromFile(files[0])
      }

      curlCommand = getCurlCommand(
        apiKey,
        mimeType,
        tag,
        mode,
        secret,
        textData,
        taalClientURL,
        files[0]
      )
    }
  }

  function setFieldsFromFile(file) {
    mimeType = file.type || 'application/octet-stream'

    if (file.type.startsWith('text/')) {
      const fr = new FileReader()
      fr.onload = function () {
        textData = fr.result
      }
      fr.readAsText(file)
    } else {
      textData = `< ${file.name} >`
    }
  }

  function onFileSelect(e) {
    const { value } = e.detail
    const add = []
    value.forEach((file) => {
      if (!files.some((item) => getFileKey(item) === getFileKey(file))) {
        add.push(file)
      }
    })
    files = devMode ? [add[0]] : files.concat(add)
  }

  $: {
    hasFileOverLimit = false
    files.forEach((file) => {
      if (fileSizeLimitBytes !== -1 && file.size > fileSizeLimitBytes) {
        hasFileOverLimit = true
      }
    })
  }

  function onRemoveTransferFile(e) {
    const file = e.detail.value
    const key = getFileKey(file)
    files = files.filter((item) => getFileKey(item) !== key)
    if (files.length === 0) {
      textData = ''
      mimeType = stdMimeType
    }
  }

  function onSubmit() {
    writeData()
  }

  function onInputMount(e) {
    e.detail.inputRef.focus()
  }

  function writeData() {
    dataTransmissionInProgress = true

    if (files.length > 0) {
      Promise.all(
        files.map((file) => {
          writeDataItem(fileDataMap[getFileKey(file)], file.type, file.name)
        })
      ).finally(() => {
        onTransmissionEnd()
      })
    } else {
      Promise.all([writeDataItem(textData, mimeType, '')]).finally(() => {
        onTransmissionEnd()
      })
    }
  }

  let timeoutId = null

  function onTransmissionEnd() {
    timeoutId = setTimeout(() => {
      dataTransmissionInProgress = false
      reset()
      setStoreValue('tag', tag)
      setStoreValue('apiKey', apiKey)
    }, 1000)
  }

  function writeDataItem(data, mimeType, filename) {
    return api.writeTransaction(
      {
        body: data,
        headers: {
          Authorization: 'Bearer ' + apiKey,
          'Content-Type': mimeType,
          'X-Tag': tag,
          'X-Filename': filename,
          'X-Mode': mode,
          'X-Key': secret,
        },
      },
      (data) =>
        success(t(`${pageKey}.notifications.transaction-submit-success`)),
      (error) => failure(t('notifications.failure', { error }))
    )
  }

  function reset() {
    files = []
    textData = ''
    mimeType = stdMimeType
    fileStatusData = {}
    fileDataMap = {}
  }

  onMount(async () => {
    tag = getStoreValue('tag', '')
    mode = getStoreValue('mode', 'raw')
    secret = getStoreValue('secret', '')

    const settings = await getSettings(t)
    taalClientURL = getCorrectURL(settings.listenAddress)
    fileSizeLimitBytes = !isNaN(settings.fileSizeLimitBytes)
      ? parseInt(settings.fileSizeLimitBytes)
      : -1

    const keysResult = await getApiKeys(t)
    keys = keysResult.keys || []

    if (keys.length === 0) {
      failure(t(`${pageKey}.notifications.no-keys-failure`), { duration: 5000 })
      navigate('/register-key')
    }

    apiKey = getStoreValue('apiKey')
    const storedKeyExists =
      keys.filter((key) => key.api_key === apiKey).length > 0

    if ((!apiKey || !storedKeyExists) && keys.length > 0) {
      apiKey = keys[0].api_key
    }

    mounting = false
    $spinCount--

    return () => {
      if (timeoutId) {
        clearTimeout(timeoutId)
      }
    }
  })
</script>

<PageWithMenu>
  <div class="island">
    <Heading value={t(`${pageKey}.send-data-label`)} />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading value={t(`${pageKey}.parameters-label`)} size={2} />
      <Switch
        name="devMode"
        label={t(`${pageKey}.developer-mode-label`)}
        bind:checked={devMode}
      />
    </div>
    <Spacer h={24} />
    <Dropdown
      name="apiKey"
      label={t(`${pageKey}.api-key-label`)}
      required
      value={apiKey}
      items={apiKeyItems}
      disabled={!apiKeyItems || apiKeyItems.length <= 1}
      on:change={(e) => {
        apiKey = e.detail.value
        setStoreValue('apiKey', apiKey)
      }}
      on:mount={onInputMount}
    />
    {#if devMode}
      <Spacer h={24} />
      <TextInput
        name="taalClientUrl"
        label={t(`${pageKey}.taal-client-url-label`)}
        bind:value={taalClientURL}
      />
      <Spacer h={24} />
      <TextInput
        name="mimeType"
        label={t(`${pageKey}.mime-type-label`)}
        bind:value={mimeType}
        disabled={inputDataDisabled}
      />
    {/if}
    <Spacer h={24} />
    <TextInput name="tag" label={t(`${pageKey}.tag-label`)} bind:value={tag} />
    <Spacer h={32} />
    <div class="sub-row">
      <Heading value={t(`${pageKey}.transaction-data-label`)} size={2} />
    </div>
    <Spacer h={24} />
    <div class="mode-row">
      <Dropdown
        name="mode"
        label={t(`${pageKey}.mode-label`)}
        required
        bind:value={mode}
        items={modeItems}
      />
      {#if mode === 'encrypt'}
        <TextInput
          name="secret"
          label={t(`${pageKey}.secret-label`)}
          bind:value={secret}
        />
      {/if}
    </div>
    {#if devMode}
      <Spacer h={24} />
      <TextArea
        name="textData"
        label={t(`${pageKey}.text-data-label`)}
        required
        bind:value={textData}
        disabled={inputDataDisabled}
      />
    {/if}
    <Spacer h={24} />
    <FileUpload
      name="fileUpload"
      label={t(`${pageKey}.file-upload.label`)}
      titleText={t(`${pageKey}.file-upload.title`)}
      hintText={fileSizeLimitBytes === -1
        ? t(`${pageKey}.file-upload.hint-no-limit`)
        : t(`${pageKey}.file-upload.hint`, {
            size: dataSize(fileSizeLimitBytes),
          })}
      selectText={t(`${pageKey}.file-upload.select`)}
      multiple={!devMode}
      required
      compact={compactFileUpload}
      on:change={onFileSelect}
    />
    {#if files.length > 0}
      <Spacer h={24} />
      <FileTransfer
        name="fileTransfer"
        label={t(`${pageKey}.file-transfer-label`)}
        {files}
        error={hasFileOverLimit
          ? t(`${pageKey}.max-filesize-limit-error`, {
              size: dataSize(fileSizeLimitBytes || 0),
            })
          : ''}
        onError={(error) => failure(t('notifications.failure', { error }))}
        on:remove={onRemoveTransferFile}
        bind:fileDataMap
        bind:fileStatusData
      />
    {/if}
    {#if devMode}
      <Spacer h={32} />
      <div class="sub-row">
        <Heading value={t(`${pageKey}.curl-command-label`)} size={2} />
        <Button
          variant="ghost"
          icon="document-duplicate"
          size="small"
          on:click={() => copyCurl(curlCommand)}
        >
          {t(`${pageKey}.curl-copy-label`)}
        </Button>
      </div>
      <Spacer h={24} />
      <TextArea
        name="curlCommand"
        readonly
        minH={160}
        value={curlCommand}
        disabled={hasFileOverLimit}
      />
    {/if}
    <Spacer h={64} />
    <div class="buttons">
      <Button
        variant="ghost"
        size="large"
        disabled={submitButtonIsDisabled}
        on:click={reset}
      >
        {t(`${pageKey}.reset-label`)}
      </Button>
      <div>
        <Button
          size="large"
          iconAfter="arrow-narrow-right"
          disabled={submitButtonIsDisabled}
          on:click={onSubmit}
        >
          {submitButtonText}
        </Button>
        <Spacer h={12} />
        <div>
          <Text
            size={5}
            html={true}
            color="#8F8D94"
            value={t(`${pageKey}.disclaimer`, {
              link: link('https://beta-console.taal.com/terms-and-conditions', 'client agreement'),
            })}
          />
        </div>
      </div>
    </div>
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

  .mode-row {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(calc(50% - 16px), 1fr));
    column-gap: 16px;
    row-gap: 24px;
  }

  .buttons {
    display: flex;
    justify-content: flex-end;
    flex-wrap: wrap;
    align-items: top;
    gap: 24px;
  }
</style>
