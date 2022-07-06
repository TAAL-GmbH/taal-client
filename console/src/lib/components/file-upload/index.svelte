<script>
  import { createEventDispatcher } from 'svelte'

  import Button from '../button/index.svelte'
  import Icon from '../icon/index.svelte'
  import Text from '../text/index.svelte'
  import { getInputLabel } from '../../utils/strings'

  const dispatch = createEventDispatcher()

  let type = 'file'

  export let label = ''
  export let labelPlacement = 'top'
  export let labelAlignment = 'start' // 'start' | 'center' | 'end'
  export let required = false
  export let name = ''
  export let disabled = false
  export let error = ''
  export let accept = '*'
  export let multiple = false

  export let compact = false

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

  let inputRef
  let btnRef

  function onInputParentClick() {
    // btnRef.focus()
  }

  function toArray(fileList) {
    return Array.from(fileList)
  }

  function onSelect() {
    inputRef.click()
  }

  function onInputChange(e) {
    dispatch('change', { name, type, value: toArray(inputRef.files) })
  }

  let dragOver = false

  function onDrop(e) {
    dragOver = false
    dispatch('change', { name, type, value: toArray(e.dataTransfer.files) })
  }
  function onDragOver(e) {
    dragOver = true
  }
  function onDragLeave(e) {
    dragOver = false
  }
</script>

<div
  class="tui-file-upload"
  class:compact
  style:--direction-input-local={compact ? 'row' : 'column'}
  style:--prompt-gap-local={compact ? '15px' : '26px'}
  style:--fontSize-local={fontSize}
  style:--direction-local={direction}
  style:--justify-local={justify}
  style:--label-align-local={labelAlign}
  style:--gap-local={gap}
  on:click={disabled ? null : onInputParentClick}
>
  <div class="placement">
    {#if label}
      <label for={name}>{getInputLabel(label, required)}</label>
    {/if}
    <div
      class="input"
      class:disabled
      class:error={error !== ''}
      class:focused
      class:dragOver
      on:drop|preventDefault={onDrop}
      on:dragover|preventDefault={onDragOver}
      on:dragleave|preventDefault={onDragLeave}
    >
      <input
        bind:this={inputRef}
        {type}
        {name}
        {accept}
        {multiple}
        {disabled}
        on:input={onInputChange}
        on:focus={(e) => (focused = true)}
        on:blur={(e) => (focused = false)}
      />
      <div class="prompt">
        <div class="icon">
          <Icon name="feather_upload-cloud" size={44} />
        </div>
        <div class="text">
          <Text
            value="Drag and drop or select a file"
            size={4}
            color={disabled ? '#8F8D94' : '#232D7C'}
          />
          <Text value="File size no more than 10MB" size={5} color="#8F8D94" />
        </div>
      </div>
      <div class="action">
        <Button
          bind:this={btnRef}
          variant="secondary"
          size="small"
          {disabled}
          forceDisabledBorderDark
          on:click={onSelect}
        >
          Select file
        </Button>
      </div>
    </div>
  </div>
  {#if error}
    <div class="error-msg">{error}</div>
  {/if}
</div>

<style>
  .tui-file-upload {
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

  input {
    position: absolute;
    outline: none;
    border: none;
    opacity: 0;
    pointer-events: none;
  }

  .input {
    box-sizing: var(--box-sizing);

    display: flex;
    flex-direction: var(--direction-input-local);
    align-items: center;
    justify-content: space-between;

    width: 100%;
    height: 260px;
    padding: 42px 25px;

    background-color: #ffffff;
    border: 1px dashed #232d7c;
    border-radius: 4px;
  }
  .tui-file-upload.compact .input {
    height: 120px;
    padding: 0 25px;
  }
  .input.dragOver {
    background-color: #eff8ff;
  }

  .input.error,
  .input.error.focused {
    border: 1px dashed #ff344c;
  }

  .error-msg {
    margin-top: 8px;

    font-weight: 500;
    font-size: 12px;
    line-height: 16px;
    letter-spacing: -0.01em;

    color: #ff344c;
  }

  .prompt {
    display: flex;
    flex-direction: var(--direction-input-local);
    align-items: center;
    gap: var(--prompt-gap-local);
  }
  .icon {
    width: 44px;
    height: 44px;
    color: #232d7c;
  }
  .disabled .icon {
    color: #8f8d94;
  }
  .text {
    text-align: center;
  }
  .tui-file-upload.compact .text {
    text-align: left;
  }

  .disabled,
  .disabled:active {
    background-color: #efefef;
    border-color: #8f8d94;
    color: #8f8d94;
  }
</style>
