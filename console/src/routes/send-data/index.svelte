<script>
  import Button from '../../lib/components/button/index.svelte'
  import Dropdown from '../../lib/components/dropdown/index.svelte'
  import FileUpload from '../../lib/components/file-upload/index.svelte'
  import FileTransfer from '../../lib/components/file-transfer/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Switch from '../../lib/components/switch/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Text from '../../lib/components/text/index.svelte'
  import TextInput from '../../lib/components/textinput/index.svelte'
  import TextArea from '../../lib/components/textarea/index.svelte'
  import { valueSet } from '../../lib/utils/types'
  import { getFileKey } from '../../lib/utils/files'

  let textInputs = {
    apiKey: '',
    taalClientUrl: '',
    mimeType: '',
    tag: '',
    textData: '',
    copyCurlText:
      'curl \\ \n\t -X POST \\ \n\t -H ‘Authorization: Bearer mainnet_...',
  }

  let dropdowns = {
    mode: '',
  }

  let checks = {
    devMode: true,
  }

  let files = []
  let imageSrcData = {}
  let supportedImageSrcDataFileTypes = ['image/png', 'image/jpeg']
  let fileProgressData = {}
  // let fileProgressData = {
  //   'profile.jpg_1643963589484': { state: 'progress', progress: 0.5 },
  // }

  let compactFileUpload = false

  $: {
    compactFileUpload = checks.devMode || files.length > 0

    // localStorage.setItem('devmode', checks.devMode)

    files.forEach((file) => {
      if (
        !imageSrcData[getFileKey(file)] &&
        supportedImageSrcDataFileTypes.includes(file.type)
      ) {
        let reader = new FileReader()
        reader.readAsDataURL(file)
        reader.onloadend = function () {
          imageSrcData[getFileKey(file)] = reader.result
        }
      }
    })
  }

  function uploadFile(file) {
    const url = 'TODO-GET-URL'
    const xhr = new XMLHttpRequest()
    const formData = new FormData()
    const fileKey = getFileKey(file)
    xhr.open('POST', url, true)

    xhr.upload.addEventListener('progress', function (e) {
      console.log('progress = ', e.loaded / e.total)
      fileProgressData[fileKey] = {
        state: 'progress',
        progress: e.loaded / e.total,
      }
    })

    xhr.addEventListener('readystatechange', function (e) {
      if (xhr.readyState === 4 && xhr.status === 200) {
        fileProgressData[fileKey] = { state: 'success', progress: 1 }

        addNotification({
          text: `Successfully uploaded file: ${file.name}`,
          position: 'bottom-left',
          type: 'success',
          removeAfter: 2000,
        })
      } else if (xhr.readyState === 4 && xhr.status !== 200) {
        fileProgressData[fileKey] = { state: 'failure', progress: 1 }

        addNotification({
          text: `Upload failed for: ${file.name}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 2000,
        })
      }
    })

    // initialise start state
    fileProgressData[fileKey] = { state: 'progress', progress: 0 }

    formData.append('file', file)
    xhr.send(formData)
  }

  function onChange(e) {
    const { name, group, type, value, checked } = e.detail
    console.log('onChange: name =', name, ' type =', type, ' value =', value)
    switch (type) {
      case 'text':
        textInputs[name] = value || ''
        break
      case 'dropdown':
        dropdowns[name] = value || ''
        break
      case 'checkbox':
        checks[name] = checked || false
        break
      case 'file':
        files = files.concat(value)
        break
    }
  }

  function onInputMount(e) {
    e.detail.inputRef.focus()
  }

  function onCopyCurl() {
    console.log('onCopyCurl')
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
  }
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
        checked={checks['devMode']}
        on:change={onChange}
      />
    </div>
    <Spacer h={24} />
    <TextInput
      name="apiKey"
      label="API key"
      required
      value={textInputs['apiKey']}
      on:change={onChange}
      on:mount={onInputMount}
    />
    {#if checks.devMode}
      <Spacer h={24} />
      <TextInput
        name="taalClientUrl"
        label="TAAL Client URL"
        value={textInputs['taalClientUrl']}
        on:change={onChange}
      />
      <Spacer h={24} />
      <TextInput
        name="mimeType"
        label="MIME type"
        value={textInputs['mimeType']}
        on:change={onChange}
      />
    {/if}
    <Spacer h={24} />
    <TextInput
      name="tag"
      label="Tag"
      value={textInputs['tag']}
      on:change={onChange}
    />
    <Spacer h={32} />
    <div class="sub-row">
      <Heading value="Transaction data" size={2} />
    </div>
    <Spacer h={24} />
    <div class="drop">
      <Dropdown
        class="drop"
        name="mode"
        label="Mode"
        required
        value={dropdowns['mode']}
        items={[
          { value: '1', label: 'Mode 1' },
          { value: '2', label: 'Mode 2' },
        ]}
        on:change={onChange}
      />
    </div>
    {#if checks.devMode}
      <Spacer h={24} />
      <TextArea
        name="textData"
        label="Text data"
        required
        value={textInputs['textData']}
        on:change={onChange}
      />
    {/if}
    <Spacer h={24} />
    <FileUpload
      name="fileUpload"
      label="File"
      required
      compact={compactFileUpload}
      on:change={onChange}
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
    {#if checks.devMode}
      <Spacer h={32} />
      <div class="sub-row">
        <Heading value="cURL command" size={2} />
        <Button
          variant="ghost"
          icon="document-duplicate"
          size="small"
          on:click={onCopyCurl}
        >
          Copy
        </Button>
      </div>
      <Spacer h={24} />
      <TextArea
        name="copyCurlText"
        readonly
        value={textInputs['copyCurlText']}
        on:change={onChange}
      />
    {/if}
    <Spacer h={64} />
    <div class="buttons">
      <Button variant="ghost" size="large">Cancel</Button>
      <Button size="large" iconAfter="arrow-narrow-right"
        >Submit transaction</Button
      >
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

  .drop {
    max-width: 445px;
  }

  .buttons {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    gap: 24px;
  }
</style>
