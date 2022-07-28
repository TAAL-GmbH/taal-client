<script>
  import { createEventDispatcher } from 'svelte'

  import Icon from '../icon/index.svelte'
  import Progress from '../progress/index.svelte'
  import Text from '../text/index.svelte'
  import { getInputLabel } from '../../utils/strings'
  import { getFileKey as getKey } from '../../utils/files'
  import { dataSize } from '../../utils/format'

  const dispatch = createEventDispatcher()

  export let testId = null

  export let label = ''
  export let labelPlacement = 'top'
  export let labelAlignment = 'start' // 'start' | 'center' | 'end'
  export let name = ''
  export let idField = 'name'
  export let required = false
  export let disabled = false
  export let valid = true
  export let error = ''

  // upload state
  export let onError = (error) => {}
  export let getFileKey = getKey
  export let files = []
  export let fileStatusData = {} // exposed mostly to allow reactive binding for reading values
  export let fileDataMap = {} // exposed mostly to allow reactive binding for reading values
  export let imageSrcData = {} // exposed mostly to allow reactive binding for reading values
  export let supportedImageSrcDataFileTypes = ['image/png', 'image/jpeg']

  let readerMap = {}

  $: {
    if (files.length === 0) {
      Object.values(readerMap).forEach((reader) => reader.abort())
      imageSrcData = {}
      fileStatusData = {}
      fileDataMap = {}
      readerMap = {}
      onData()
      onStatus()
    } else {
      files.forEach((file) => {
        const key = getFileKey(file)
        // process image preview
        if (
          !imageSrcData[key] &&
          supportedImageSrcDataFileTypes.includes(file.type)
        ) {
          const reader = new FileReader()

          reader.onloadend = () => (imageSrcData[key] = reader.result)
          reader.onerror = () => onError(reader.error)
          reader.readAsDataURL(file)
        }
        // read file data
        if (!fileDataMap[key] && !readerMap[key]) {
          const reader = new FileReader()
          readerMap[key] = reader
          fileStatusData[key] = { state: 'progress', progress: 0 }
          onStatus()

          reader.onload = () => {
            fileStatusData[key] = { state: 'success', progress: 1 }
            fileDataMap[key] = reader.result
            onStatus()
            onData()
          }
          reader.onprogress = (data) => {
            if (data.lengthComputable) {
              fileStatusData[key] = {
                state: 'progress',
                progress: data.loaded / data.total,
              }
            }
          }
          reader.onabort = () => {
            fileStatusData[key] = { state: 'cancelled', progress: 1 }
            onStatus()
          }
          reader.onerror = () => {
            fileStatusData[key] = { state: 'failure', progress: 1 }
            onStatus()
            onError(reader.error)
          }
          reader.readAsArrayBuffer(file)
        }
      })
    }
  }

  let direction = 'row'
  let justify = 'flex-start'

  $: {
    switch (labelPlacement) {
      case 'top':
        direction = 'column'
        break
      case 'bottom':
        direction = 'column-reverse'
        justify = 'flex-end'
        break
      case 'left':
        direction = 'row'
        break
      case 'right':
        direction = 'row-reverse'
        justify = 'flex-end'
        break
    }
  }

  let labelAlign = 'center'

  $: {
    switch (labelAlignment) {
      case 'start':
        labelAlign = 'flex-start'
        break
      case 'center':
        labelAlign = 'center'
        break
      case 'end':
        labelAlign = 'flex-end'
        break
    }
  }

  let fontSize = '14px'
  let gap = '8px'

  $: {
    gap = label ? '8px' : '0'
  }

  let focused = false

  function onCancel(file) {
    const key = getFileKey(file)
    readerMap[key].abort()
    dispatch('cancel', { name, value: file })
  }

  function onRemove(file) {
    const key = getFileKey(file)
    files = files.filter((item) => getFileKey(item) !== key)
    if (imageSrcData[key]) {
      delete imageSrcData[key]
    }
    if (fileDataMap[key]) {
      delete fileDataMap[key]
      onData()
    }
    if (fileStatusData[key]) {
      delete fileStatusData[key]
      onStatus()
    }
    if (readerMap[key]) {
      delete readerMap[key]
    }
    dispatch('remove', { name, value: file })
  }

  function onData() {
    dispatch('data', { name, value: fileDataMap })
  }

  function onStatus() {
    dispatch('status', { name, value: fileStatusData })
  }

  function getLabelColor(file) {
    return fileStatusData[getFileKey(file)]?.state === 'failure'
      ? '#FF344C'
      : '#282933'
  }
</script>

<div
  class="tui-file-transfer"
  data-test-id={testId}
  style:--fontSize-local={fontSize}
  style:--direction-local={direction}
  style:--justify-local={justify}
  style:--label-align-local={labelAlign}
  style:--gap-local={gap}
>
  <div class="placement">
    {#if label}
      <label for={name}>{getInputLabel(label, required)}</label>
    {/if}
    <div
      class="content"
      class:disabled
      class:error={!valid || error !== ''}
      class:focused
    >
      {#each files as file (file[idField])}
        <div
          class="item"
          class:success={fileStatusData[getFileKey(file)] &&
            fileStatusData[getFileKey(file)].state === 'success'}
        >
          <div class="info">
            <div class="icon">
              {#if imageSrcData[getFileKey(file)]}
                <img
                  src={imageSrcData[getFileKey(file)]}
                  style="width: 24px; height: 24px;"
                  alt="thumb-preview"
                />
              {:else}
                <Icon name="bi_file-earmark-image" size={24} />
              {/if}
            </div>
            <div class="text">
              <div class="labels">
                <Text value={file.name} size={5} color={getLabelColor(file)} />
                <Text
                  value={dataSize(file.size)}
                  size={5}
                  color={getLabelColor(file)}
                />
              </div>
              {#if fileStatusData[getFileKey(file)] && fileStatusData[getFileKey(file)].progress < 1}
                <Progress
                  ratio={fileStatusData[getFileKey(file)]
                    ? fileStatusData[getFileKey(file)].progress
                    : 0}
                  size="small"
                  normalColor="#232D7C"
                />
              {/if}
            </div>
          </div>
          {#if fileStatusData[getFileKey(file)]}
            {#if fileStatusData[getFileKey(file)].progress < 1}
              <div class="action" on:click={() => onCancel(file)}>
                <Icon name="close" size={18} />
              </div>
            {:else}
              <div class="action" on:click={() => onRemove(file)}>
                <Icon name="trash" size={18} />
              </div>
            {/if}
          {/if}
        </div>
      {/each}
    </div>
  </div>
  {#if error}
    <div class="error-msg">{error}</div>
  {/if}
</div>

<style>
  .tui-file-transfer {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);

    display: flex;
    flex-direction: column;
    flex-grow: 1;
  }

  .placement {
    display: flex;
    flex-direction: var(--direction-local);
    align-items: var(--label-align-local);
    justify-content: var(--justify-local);
    gap: var(--gap-local);
  }

  label {
    font-weight: 400;
    font-size: var(--fontSize-local);
    line-height: 18px;
    letter-spacing: -0.02em;

    color: #232d7c;
  }

  .content {
    box-sizing: var(--box-sizing);

    width: 100%;
    min-height: 120px;
    padding: 25px;

    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 20px;

    background-color: #ffffff;
    border: 1px solid #232d7c;
    border-radius: 4px;
  }

  .content.error,
  .content.error.focused {
    border: 1px solid #ff344c;
  }

  .error-msg {
    margin-top: 8px;

    font-weight: 500;
    font-size: 12px;
    line-height: 16px;
    letter-spacing: -0.01em;

    color: #ff344c;
  }

  .disabled,
  .disabled:active {
    background-color: #efefef;
    border-color: #8f8d94;
    color: #8f8d94;
  }

  .item {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  .info {
    width: 100%;
    display: flex;
    align-items: center;
    gap: 10px;
  }
  .icon {
    width: 24px;
    height: 24px;
    color: rgba(0, 148, 255, 0.5);
    opacity: 0.5;
  }
  .text {
    width: 80%;
    text-align: left;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    gap: 8px;
  }
  .labels {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    opacity: 0.5;
  }
  .item.success .icon {
    opacity: 1;
  }
  .item.success .labels {
    opacity: 1;
  }
  .action {
    width: 18px;
    height: 18px;
    color: #232d7c;
  }
  .action:hover {
    cursor: pointer;
  }
  .disabled .action,
  .disabled .action:hover {
    cursor: auto;
    color: #8f8d94;
  }
</style>
