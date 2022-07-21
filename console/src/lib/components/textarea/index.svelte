<script>
  import { createEventDispatcher, onMount, afterUpdate } from 'svelte'
  import Icon from '../icon/index.svelte'
  import { drag } from '../../actions/drag'
  import { clamp } from '../../utils/numbers'
  import { getInputLabel } from '../../utils/strings'

  const dispatch = createEventDispatcher()

  export let testId = null

  //   export let type = 'text'
  let type = 'text'

  export let label = ''
  export let labelPlacement = 'top'
  export let required = false
  export let name = ''
  export let value = ''
  export let placeholder = ''
  export let disabled = false
  export let readonly = false
  export let valid = true
  export let error = ''

  export let resize = true
  export let minH = 80
  export let maxH = 400

  // direction
  let direction = 'row'

  $: {
    switch (labelPlacement) {
      case 'top':
        direction = 'column'
        break
      case 'bottom':
        direction = 'column-reverse'
        break
      case 'left':
        direction = 'row'
        break
      case 'right':
        direction = 'row-reverse'
        break
    }
  }

  // size
  export let size = 'medium'

  let height = 120
  let nativeHeight = 60
  let padding = '10px 32px'
  let paddingFocused = '9px 31px'
  let fontSize = '16px'

  $: {
    switch (size) {
      case 'large':
        height = clamp(120, minH, maxH)
        padding = '10px 32px'
        paddingFocused = '9px 31px'
        fontSize = '16px'
        break
      case 'medium':
        height = clamp(108, minH, maxH)
        padding = '10px 17px'
        paddingFocused = '9px 16px'
        fontSize = '16px'
        break
      case 'small':
        height = clamp(84, minH, maxH)
        padding = '4px 16px'
        paddingFocused = '3px 15px'
        fontSize = '14px'
        break
    }
  }

  let focused = false

  let inputRef

  function onInputParentClick() {
    inputRef.focus()
  }

  function onInputChange(e) {
    dispatch('change', { name, type, value: e.srcElement.value })
  }

  afterUpdate(() => {
    updateSize()
  })

  onMount(() => {
    dispatch('mount', { inputRef })
    updateSize()
  })

  function updateSize() {
    let newNativeHeight = inputRef.scrollHeight - 4
    nativeHeight = newNativeHeight

    // let newHeight =
    //   size === 'small' ? newNativeHeight + 8 : newNativeHeight + 20
    // newHeight = clamp(newHeight, minH, maxH)

    // height = newHeight
  }

  function onKeyUp(e) {
    updateSize()
  }

  let dragStartH = null

  function dragStart(e) {
    dragStartH = height
  }

  function dragMove(e) {
    if (!dragStartH) {
      return
    }
    const { x, y } = e.detail
    let newHeight = dragStartH + y
    newHeight = clamp(newHeight, minH, maxH)

    height = newHeight
  }

  function dragStop() {
    dragStartH = null
  }

  function onFocusAction(eventName) {
    switch (eventName) {
      case 'blur':
        focused = false
        break
      case 'focus':
        focused = true
        break
    }
    dispatch(eventName)
  }

  let optProps = {}
  $: {
    if (testId) {
      optProps['data-test-id'] = testId
    }
  }
</script>

<div
  class="tui-textarea"
  style:--height-local={height + 'px'}
  style:--padding-local={padding}
  style:--padding-focused-local={paddingFocused}
  style:--fontSize-local={fontSize}
  style:--direction-local={direction}
  style:--native-height-local={nativeHeight + 'px'}
  {...optProps}
>
  <div class="placement">
    {#if label}
      <label for={name}>{getInputLabel(label, required)}</label>
    {/if}
    <div
      class="input"
      class:disabled
      class:error={!valid || error !== ''}
      class:focused
      class:resize
      class:dragging={dragStartH === null}
      on:click={disabled ? null : onInputParentClick}
    >
      <div class="control-container" tabindex={-1}>
        <textarea
          bind:this={inputRef}
          {type}
          {name}
          {value}
          {placeholder}
          {disabled}
          {readonly}
          on:input={onInputChange}
          on:focus={() => onFocusAction('focus')}
          on:blur={() => onFocusAction('blur')}
          on:keyup={onKeyUp}
        />
      </div>
      <div
        class="resize-icon"
        use:drag
        on:dragstart={disabled ? null : dragStart}
        on:dragmove={dragMove}
        on:dragstop={dragStop}
      >
        <Icon name="drag-corner" size={14} />
      </div>
    </div>
  </div>
  {#if error}
    <div class="error-msg">{error}</div>
  {/if}
</div>

<style>
  .tui-textarea {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);

    display: flex;
    flex-direction: column;
    flex-grow: 1;
  }

  .placement {
    display: flex;
    flex-direction: var(--direction-local);
    gap: 8px;
  }

  label {
    font-weight: 400;
    font-size: 14px;
    line-height: 18px;
    letter-spacing: -0.02em;

    color: #232d7c;
  }

  textarea {
    outline: none;
    border: none;
    height: var(--native-height-local);
    width: 100%;
    resize: none;
    overflow: hidden;
    background-color: inherit;

    font-weight: 400;
    font-size: var(--fontSize-local);
    font-family: var(--font-family);
    line-height: 24px;
    letter-spacing: -0.01em;
  }

  .control-container {
    width: 100%;
    overflow-x: hidden;
    overflow-y: auto;
    height: var(--height-local);
  }

  .resize-icon {
    width: 14;
    height: 14;
    margin-right: -16px;
    margin-top: calc(var(--height-local) - 6px);
    color: #eee;
    cursor: auto;
  }
  .input.resize .resize-icon {
    color: #232d7c;
    cursor: ns-resize;
  }
  .input.disabled,
  .input.disabled.resize .resize-icon {
    color: #8f8d94;
    cursor: auto;
  }

  .input {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    padding: var(--padding-local);
    height: var(--height-local);

    background-color: #ffffff;
    border-style: solid;
    border-width: 1px;
    border-color: #232d7c;
    border-radius: 4px;

    font-weight: 400;
    font-size: var(--fontSize-local);
    line-height: 24px;
    letter-spacing: -0.01em;

    color: #282933;
  }

  .input.focused {
    border-width: 2px;
    border-color: #4a3aff;
    padding: var(--padding-focused-local);
  }

  .input.error,
  .input.error.focused {
    border-width: 1px;
    border-color: #ff344c;
    padding: var(--padding-local);
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
  }
</style>
