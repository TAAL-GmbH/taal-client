<script>
  import { createEventDispatcher } from 'svelte'
  // import icons from '../icon/svg'
  import Icon from '../icon/index.svelte'
  import { getInputLabel } from '../../utils/strings'

  const dispatch = createEventDispatcher()

  let type = 'checkbox'

  export let label = ''
  export let labelPlacement = 'left'
  export let labelAlignment = 'center' // 'start' | 'center' | 'end'
  export let required = false
  export let name = ''
  export let checked = false
  export let disabled = false
  export let error = false

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

  let height = '24px'
  let width = '24px'
  let fontSize = '16px'

  // let backgroundImage = `url('data:image/svg+xml;utf8,${icons.check}')`
  let cursor = disabled ? 'auto' : 'pointer'

  let focused = false

  let inputRef

  function onInputParentClick() {
    inputRef.focus()
    const newValue = !checked
    checked = newValue
    dispatch('change', { name, type, checked: newValue })
  }
</script>

<div
  class="tui-checkbox placement"
  style:--width-local={width}
  style:--height-local={height}
  style:--fontSize-local={fontSize}
  style:--direction-local={direction}
  style:--justify-local={justify}
  style:--label-align-local={labelAlign}
  style:--cursor-local={cursor}
  on:click={disabled ? null : onInputParentClick}
>
  <label for={name}>{getInputLabel(label, required)}</label>
  <div class="input" class:disabled class:error class:focused class:checked>
    <input
      bind:this={inputRef}
      {type}
      {name}
      {checked}
      on:focus={(e) => (focused = true)}
      on:blur={(e) => (focused = false)}
    />
    <div style="width: 20px; height: 20px;">
      <Icon name="check" size={20} />
    </div>
  </div>
</div>

<style>
  .tui-checkbox {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);
  }

  .placement {
    display: flex;
    flex-direction: var(--direction-local);
    align-items: var(--label-align-local);
    justify-content: var(--justify-local);
    gap: 8px;
  }

  label {
    font-weight: 500;
    font-size: var(--fontSize-local);
    line-height: 16px;

    color: #232d7c;
    cursor: var(--cursor-local);
  }

  input {
    outline: none;
    border: none;
    position: absolute;
    opacity: 0;
    pointer-events: none;
  }

  .input {
    display: flex;
    align-items: center;
    justify-content: center;
    width: var(--width-local);
    height: var(--height-local);

    background-color: #ffffff;
    border-style: solid;
    border-width: 1px;
    border-color: #232d7c;
    border-radius: 4px;

    color: #ffffff;
    cursor: var(--cursor-local);
    transition: color 0.2s ease-in-out, background-color 0.2s ease-in-out;
  }

  .input.focused {
    border-color: #4a3aff;
  }

  .input.error,
  .input.error.focused {
    border-color: #ff344c;
  }

  .input.checked {
    background-color: #232d7c;
  }
  .input.disabled {
    color: #8f8d94;
  }

  .disabled,
  .disabled:active,
  .input.disabled {
    background-color: #efefef;
    border-color: #232d7c;
  }
</style>
