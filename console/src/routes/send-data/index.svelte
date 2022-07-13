<script>
  import { onMount } from 'svelte'
  import { getNotificationsContext } from 'svelte-notifications'

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
  import Spinner from '../../lib/components/spinner/index.svelte'

  import { spinCount } from '../../lib/stores'
  import * as api from '../../lib/api'

  import { getFileKey } from '../../lib/utils/files'
  import {
    copyCurl,
    getApiKeys,
    getCorrectURL,
    getCurlCommand,
    getSettings,
    modeItems,
    getStoreValue,
    setStoreValue,
    getTimeoutMillis,
    updateStore,
  } from './utils'

  const { addNotification } = getNotificationsContext()

  const stdMimeType = 'text/plain'
  let disableButtonsCount = 0
  let loading = true
  let timeout = 0

  let devMode = getStoreValue('devmode') === 'true'
  let taalClientURL = ''
  let mimeType = stdMimeType
  let tag = ''
  let mode = 'raw'
  let secret = ''
  let textData = ''
  let curlCommand = ''

  $: updateStore('tag', tag, loading)
  $: updateStore('mode', mode, loading)
  $: updateStore('secret', secret, loading)
  $: updateStore('devmode', devMode, loading)

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

  $: compactFileUpload = devMode || files.length > 0

  // let files = null
  let file = null
  let fileData = null
  let filename = ''

  let files = []
  let imageSrcData = {}
  let supportedImageSrcDataFileTypes = ['image/png', 'image/jpeg']
  let fileProgressData = {}
  // let fileProgressData = {
  //   'Logged out.png_1643790959194': { state: 'progress', progress: 0.5 },
  // }

  let inputDataDisabled = files.length > 0
  let submitButtonIsDisabled = true
  let submitButtonText = 'Submit transaction'

  $: {
    if (devMode) {
      // dev mode
      submitButtonText = 'Submit transaction'
      submitButtonIsDisabled =
        !textData || !mimeType || !apiKey || disableButtonsCount > 0

      if (files.length > 1) {
        files = [files[0]]
      }

      inputDataDisabled = false

      if (files.length > 0) {
        setFieldsFromFile(files[0])
        inputDataDisabled = true
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
    } else {
      // standard mode
      submitButtonText =
        files.length > 1
          ? `Submit ${files.length} transactions`
          : 'Submit transaction'
      submitButtonIsDisabled =
        files.length === 0 || !apiKey || disableButtonsCount > 0
    }
  }

  function setFieldsFromFile(file) {
    mimeType = file.type
    filename = file.name

    if (file.type.startsWith('text/')) {
      const fr = new FileReader()
      fr.onload = function () {
        textData = fr.result
      }
      fr.readAsText(file)
    } else {
      textData = `< ${file.name} >`
    }

    const fr = new FileReader()
    fr.onload = function () {
      fileData = fr.result
    }
    fr.readAsArrayBuffer(file)
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
    files.forEach((file) => {
      const key = getFileKey(file)
      if (
        !imageSrcData[key] &&
        supportedImageSrcDataFileTypes.includes(file.type)
      ) {
        const reader = new FileReader()
        reader.readAsDataURL(file)
        reader.onloadend = () => (imageSrcData[key] = reader.result)
      }
    })
  }

  function onCancelTransfer(e) {
    const file = e.detail.value
    const fileKey = getFileKey(file)
    console.log('onCancelTransfer: file = ', file.name)
    fileProgressData[fileKey] = { state: 'cancelled', progress: 1 }
  }

  function onRemoveTransferFile(e) {
    const file = e.detail.value
    const fileKey = getFileKey(file)
    console.log('onRemoveTransferFile: file = ', file.name)
    files = files.filter((item) => getFileKey(item) !== fileKey)
    if (imageSrcData[fileKey]) {
      delete imageSrcData[fileKey]
    }
    if (files.length === 0) {
      textData = ''
      fileData = null
      filename = ''
      mimeType = stdMimeType
    }
  }

  $: {
    console.log('taalClientURL =', taalClientURL)
  }

  let checkUploadsCompleteIntervalId = null

  function checkUploadsComplete() {
    if (checkUploadsCompleteIntervalId) {
      clearInterval(checkUploadsCompleteIntervalId)
    }
    checkUploadsCompleteIntervalId = setInterval(() => {
      if (disableButtonsCount === 0) {
        clearInterval(checkUploadsCompleteIntervalId)
        checkUploadsCompleteIntervalId = null
        reset()
      }
    }, 2000)
  }

  function onSubmit() {
    if (files.length > 0) {
      for (const file of files) {
        uploadFile(file)
      }
      checkUploadsComplete()
    } else {
      writeData()
    }
  }

  function uploadFile(file) {
    const url = `${BASE_URL}/api/v1/transactions`
    const xhr = new XMLHttpRequest()
    const formData = new FormData()
    const fileKey = getFileKey(file)
    xhr.open('POST', url, true)

    disableButtonsCount++

    xhr.upload.addEventListener('progress', function (e) {
      console.log('progress = ', e.loaded / e.total)
      // TODO: right now the backend does not give proper progress updates
      //       basically (e.loaded / e.total) evaluates to 1 right at the start already
      //       we don't want to indicate completed progress when upload is still in progress, so
      //       not doing an update here for now
      // fileProgressData[fileKey] = {
      //   state: 'progress',
      //   progress: e.loaded / e.total,
      // }
    })

    xhr.addEventListener('readystatechange', function (e) {
      console.log(
        'readystatechange: xhr.readyState = ',
        xhr.readyState,
        ' xhr.status = ',
        xhr.status
      )

      // we reached a final state
      if (xhr.readyState === 4) {
        if (xhr.status >= 200 && xhr.status < 300) {
          fileProgressData[fileKey] = { state: 'success', progress: 1 }

          addNotification({
            text: `Successfully uploaded file: ${file.name}`,
            position: 'bottom-left',
            type: 'success',
            removeAfter: 2000,
          })
        } else {
          fileProgressData[fileKey] = { state: 'failure', progress: 1 }

          addNotification({
            text: `Upload failed for: ${file.name}`,
            position: 'bottom-left',
            type: 'danger',
            removeAfter: 2000,
          })
        }
        disableButtonsCount--
      }
    })

    // initialise start state
    fileProgressData[fileKey] = { state: 'progress', progress: 0 }

    formData.append('file', file)

    // headers
    xhr.setRequestHeader('Authorization', 'Bearer ' + apiKey)
    xhr.setRequestHeader('Content-Type', file.type)
    xhr.setRequestHeader('X-Tag', tag)
    xhr.setRequestHeader('X-Filename', file.name)
    xhr.setRequestHeader('X-Mode', mode)
    xhr.setRequestHeader('X-Key', secret)

    // set timeout
    xhr.timeout = getTimeoutMillis(timeout)

    xhr.send(formData)
  }

  function onInputMount(e) {
    e.detail.inputRef.focus()
  }

  async function writeData() {
    disableButtonsCount++

    await api.writeData(
      {
        body: fileData ? fileData : data,
        headers: {
          Authorization: 'Bearer ' + apiKey,
          'Content-Type': mimeType,
          'X-Tag': tag,
          'X-Filename': filename,
          'X-Mode': mode,
          'X-Key': secret,
        },
      },
      (data) => {
        addNotification({
          text: `Transaction submitted successfully`,
          position: 'bottom-left',
          type: 'success',
          removeAfter: 1000,
        })

        setStoreValue('apiKey', apiKey)
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

    reset()
    disableButtonsCount--
  }

  function reset() {
    files = []
    textData = ''
    mimeType = stdMimeType
    imageSrcData = {}
    fileProgressData = {}
  }

  onMount(async () => {
    tag = getStoreValue('tag', '')
    mode = getStoreValue('mode', 'raw')
    secret = getStoreValue('secret', '')

    const settings = await getSettings(addNotification)
    taalClientURL = getCorrectURL(settings.listenAddress)
    timeout = settings.taalTimeout

    const keysResult = await getApiKeys(addNotification)
    // TODO handle the situation when no api keys are registered yet
    //      suggestion is to show a message with a redirect button to register key page

    console.log('keysResult = ', keysResult)
    keys = keysResult.keys || []

    apiKey = getStoreValue('apiKey')
    if (!apiKey && keys.length > 0) {
      apiKey = keys[0].api_key
    }

    loading = false
  })
</script>

<PageWithMenu>
  <div class="island">
    <Heading value="Send data" />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading value="Parameters" size={2} />
      <Switch
        name="devMode"
        label="Developer mode"
        checked={devMode}
        on:change={(e) => (devMode = e.detail.checked)}
      />
    </div>
    <Spacer h={24} />
    <Dropdown
      name="apiKey"
      label="API key"
      required
      value={apiKey}
      items={apiKeyItems}
      disabled={!apiKeyItems || apiKeyItems.length <= 1}
      on:change={(e) => (apiKey = e.detail.value)}
      on:mount={onInputMount}
    />
    {#if devMode}
      <Spacer h={24} />
      <TextInput
        name="taalClientUrl"
        label="TAAL Client URL"
        value={taalClientURL}
        on:change={(e) => (taalClientURL = e.detail.value)}
      />
      <Spacer h={24} />
      <TextInput
        name="mimeType"
        label="MIME type"
        value={mimeType}
        disabled={inputDataDisabled}
        on:change={(e) => (mimeType = e.detail.value)}
      />
    {/if}
    <Spacer h={24} />
    <TextInput
      name="tag"
      label="Tag"
      value={tag}
      on:change={(e) => (tag = e.detail.value)}
    />
    <Spacer h={32} />
    <div class="sub-row">
      <Heading value="Transaction data" size={2} />
    </div>
    <Spacer h={24} />
    <div class="mode-row">
      <Dropdown
        name="mode"
        label="Mode"
        required
        value={mode}
        items={modeItems}
        on:change={(e) => (mode = e.detail.value)}
      />
      {#if mode === 'encrypt'}
        <TextInput
          name="secret"
          label="Secret"
          value={secret}
          on:change={(e) => (secret = e.detail.value)}
        />
      {/if}
    </div>
    {#if devMode}
      <Spacer h={24} />
      <TextArea
        name="textData"
        label="Text data"
        required
        value={textData}
        disabled={inputDataDisabled}
        on:change={(e) => (textData = e.detail.value)}
      />
    {/if}
    <Spacer h={24} />
    <FileUpload
      name="fileUpload"
      label="File"
      required
      compact={compactFileUpload}
      on:change={onFileSelect}
    />
    {#if files.length > 0}
      <Spacer h={24} />
      <FileTransfer
        name="fileTransfer"
        label="Files added"
        {files}
        {imageSrcData}
        {fileProgressData}
        on:cancel={onCancelTransfer}
        on:remove={onRemoveTransferFile}
      />
    {/if}
    {#if devMode}
      <Spacer h={32} />
      <div class="sub-row">
        <Heading value="cURL command" size={2} />
        <Button
          variant="ghost"
          icon="document-duplicate"
          size="small"
          on:click={() => copyCurl(curlCommand, addNotification)}
        >
          Copy
        </Button>
      </div>
      <Spacer h={24} />
      <TextArea name="curlCommand" readonly minH={160} value={curlCommand} />
    {/if}
    <Spacer h={64} />
    <div class="buttons">
      <Button
        variant="ghost"
        size="large"
        disabled={submitButtonIsDisabled}
        on:click={reset}>Reset</Button
      >
      <Button
        size="large"
        iconAfter="arrow-narrow-right"
        disabled={submitButtonIsDisabled}
        on:click={onSubmit}
      >
        {submitButtonText}
      </Button>
    </div>
  </div>
</PageWithMenu>

{#if $spinCount > 0 || loading}
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
    align-items: center;
    gap: 24px;
  }
</style>
