<script>
  import { createEventDispatcher, onMount } from 'svelte'
  import Icon from '../icon/index.svelte'
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
  export let placeholder = null
  export let disabled = false
  export let valid = true
  export let error = ''
  export let autocomplete = null

  // in confirm mode, changes are local until confirm is clicked,
  // or alternatively changes can be reset to previous non-local value
  export let confirm = false
  let localValue = value

  $: {
    localValue = value
  }

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

  let height = '60px'
  let padding = '10px 32px'
  let fontSize = '16px'

  $: {
    switch (size) {
      case 'large':
        height = '60px'
        padding = '10px 32px'
        fontSize = '16px'
        break
      case 'medium':
        height = '48px'
        padding = '10px 32px'
        fontSize = '16px'
        break
      case 'small':
        height = '36px'
        padding = '4px 16px'
        fontSize = '14px'
        break
    }
  }

  let inputOpts = {}

  $: {
    inputOpts = {}

    if (autocomplete) {
      inputOpts.autocomplete = autocomplete
    }
    if (placeholder) {
      inputOpts.placeholder = placeholder
    }
  }

  let focused = false

  let inputRef

  function onInputParentClick() {
    inputRef.focus()
  }

  function onInputChange(e) {
    if (confirm) {
      localValue = e.srcElement.value
    } else {
      value = e.srcElement.value
    }
    // we always pass through the raw value here to aid validators, etc
    dispatch('change', { name, type, value: e.srcElement.value })
  }

  onMount(() => {
    dispatch('mount', { inputRef })
  })

  function doConfirm() {
    value = localValue
    dispatch('confirm', { name, type, value })
  }

  function doReset() {
    localValue = value
  }

  function onKeyDown(e) {
    if (!e) e = window.event
    const keyCode = e.code || e.key
    if (keyCode === 'Enter') {
      doConfirm()
      return false
    } else if (keyCode === 'Escape') {
      doReset()
      return false
    }
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
</script>

<div
  class="tui-textinput"
  data-test-id={testId}
  style:--height-local={height}
  style:--padding-local={padding}
  style:--fontSize-local={fontSize}
  style:--direction-local={direction}
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
      on:click={onInputParentClick}
    >
      <input
        bind:this={inputRef}
        {type}
        {name}
        value={confirm ? localValue : value}
        {disabled}
        {...inputOpts}
        on:input={onInputChange}
        on:focus={() => onFocusAction('focus')}
        on:blur={() => onFocusAction('blur')}
        on:keydown={onKeyDown}
      />
      {#if confirm && localValue !== value}
        <div class="confirm-row">
          <div
            class="confirm-icon"
            style="color: #6EC492; width: 20px; height: 20px;"
            on:click={doConfirm}
          >
            <Icon name="check" size={20} />
          </div>
          <div
            class="confirm-icon"
            style="color: #FF344C; width: 17px; height: 17px;"
            on:click={doReset}
          >
            <Icon name="close" size={17} />
          </div>
        </div>
      {/if}
    </div>
  </div>
  {#if error}
    <div class="error-msg">{error}</div>
  {/if}
</div>

<style>
  .tui-textinput {
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

  input {
    outline: none;
    border: none;
    width: 100%;
    background-color: inherit;

    font-weight: 400;
    font-size: var(--fontSize-local);
    font-family: var(--font-family);
    line-height: 24px;
    letter-spacing: -0.01em;
  }

  .input {
    display: flex;
    align-items: center;
    padding: 1px 16px;
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
    padding: 0px 15px;
  }

  .input.error,
  .input.error.focused {
    border-width: 1px;
    border-color: #ff344c;
    padding: 1px 16px;
  }

  .confirm-row {
    display: flex;
    align-items: center;
    height: var(--height-local);
    gap: 4px;

    background-color: rgba(255, 255, 255, 0.8);
    z-index: 2;
    margin-left: -40px;
  }
  .confirm-icon {
    width: 18px;
    height: 18px;
  }
  .confirm-icon:hover {
    cursor: pointer;
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
