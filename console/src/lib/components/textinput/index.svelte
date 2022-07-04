<script>
  import { createEventDispatcher, onMount } from 'svelte'
  import { getInputLabel } from '../../utils/strings'

  const dispatch = createEventDispatcher()

  //   export let type = 'text'
  let type = 'text'

  export let label = ''
  export let labelPlacement = 'top'
  export let required = false
  export let name = ''
  export let value = ''
  export let placeholder = ''
  export let disabled = false
  export let error = ''

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

  let focused = false

  let inputRef

  function onInputParentClick() {
    inputRef.focus()
  }

  function onInputChange(e) {
    dispatch('change', { name, type, value: e.originalTarget.value })
  }

  onMount(() => {
    dispatch('mount', { inputRef })
  })
</script>

<div
  class="tui-textinput"
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
      class:error={error !== ''}
      class:focused
      on:click={onInputParentClick}
    >
      <input
        bind:this={inputRef}
        {type}
        {name}
        {value}
        {placeholder}
        {disabled}
        on:input={onInputChange}
        on:focus={(e) => (focused = true)}
        on:blur={(e) => (focused = false)}
      />
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
    font-weight: 500;
    font-size: 14px;
    line-height: 16px;

    color: #232d7c;
  }

  input {
    outline: none;
    border: none;
    width: 100%;
    background: none;
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
    font-size: 16px;
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
