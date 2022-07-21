<script>
  import { createEventDispatcher } from 'svelte'

  import Icon from '../icon/index.svelte'
  import Progress from '../progress/index.svelte'
  import Text from '../text/index.svelte'
  import { getInputLabel } from '../../utils/strings'
  import { getFileKey } from '../../utils/files'
  import { dataSize } from '../../utils/format'

  const dispatch = createEventDispatcher()

  export let testId = null

  export let label = ''
  export let labelPlacement = 'top'
  export let labelAlignment = 'start' // 'start' | 'center' | 'end'
  export let name = ''
  export let idField = 'name'
  export let files = []
  export let required = false
  export let disabled = false
  export let valid = true
  export let error = ''
  export let imageSrcData = {}
  export let fileProgressData = {}

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
    dispatch('cancel', { name, value: file })
  }

  function onRemove(file) {
    dispatch('remove', { name, value: file })
  }

  let optProps = {}
  $: {
    if (testId) {
      optProps['data-test-id'] = testId
    }
  }
</script>

<div
  class="tui-file-transfer"
  style:--fontSize-local={fontSize}
  style:--direction-local={direction}
  style:--justify-local={justify}
  style:--label-align-local={labelAlign}
  style:--gap-local={gap}
  {...optProps}
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
        <div class="item">
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
                <Text
                  value={file.name}
                  size={5}
                  color="rgba(40, 41, 51, 0.5)"
                />
                <Text
                  value={dataSize(file.size)}
                  size={5}
                  color="rgba(40, 41, 51, 0.5)"
                />
              </div>
              <Progress
                ratio={fileProgressData[getFileKey(file)]
                  ? fileProgressData[getFileKey(file)].progress
                  : 0}
                size="small"
                normalColor="#232D7C"
              />
            </div>
          </div>
          {#if fileProgressData[getFileKey(file)]}
            {#if fileProgressData[getFileKey(file)].progress < 1}
              <div class="action" on:click={() => onCancel(file)}>
                <Icon name="close" size={18} />
              </div>
            {/if}
          {:else}
            <div class="action" on:click={() => onRemove(file)}>
              <Icon name="trash" size={18} />
            </div>
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
  }
  .text {
    width: 80%;
    text-align: left;
  }
  .labels {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 8px;
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
